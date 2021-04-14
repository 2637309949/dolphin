package redismq

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/imdario/mergo"
	"github.com/ssgreg/repeat"
)

// NewConsumer creates new consumer client for Ami
//
// Note, that you MUST set in ClusterOptions ReadTimeout and WriteTimeout to at
// least 30-60 seconds, if you set big PipeBufferSize (50000 in examples) and
// if you do producing big messages.
// Reason: one full buffer is send in one big-sized query to Redis.
// So this big query may don't be completed in time and will be retransmitted,
// may be forewer retransmitted.
// This is especially strictly for Producer, because it send big queries with
// big messages. And not so strictly for Consumer, because it send big queries
// with not so big messages (only ids of ACKed messages).
func NewConsumer(opt ConsumerOptions, ropt *redis.Options) (*Consumer, error) {
	if err := mergo.Merge(&opt, defaultConsumerOptions); err != nil {
		return nil, err
	}

	client, err := newClient(clientOptions{
		name:        opt.Name,
		ropt:        ropt,
		shardsCount: opt.ShardsCount,
	})
	if err != nil {
		return nil, err
	}

	cCons := make(chan Message, opt.PrefetchCount)
	cAck := make(chan Message, opt.PendingBufferSize)

	cn := &Consumer{
		cAck:   cAck,
		cCons:  cCons,
		cl:     client,
		notif:  opt.ErrorNotifier,
		opt:    opt,
		wgAck:  &sync.WaitGroup{},
		wgCons: &sync.WaitGroup{},
	}

	cn.wgAck.Add(1)

	go cn.ack()

	return cn, nil
}

// Start consume from queue.
//
// Start read messages from Redis streams and return channel.
func (c *Consumer) Start() chan Message {
	for i := 0; i < int(c.opt.ShardsCount); i++ {
		c.wgCons.Add(1)

		go c.consume(i)
	}

	return c.cCons
}

// Stop queue client.
//
// Stop reading messages from Redis streams and lock until all
// read messages being processed.
func (c *Consumer) Stop() {
	c.needStop = true

	c.wgCons.Wait()
	close(c.cCons)
	c.stopped = true
}

// Close queue client
//
// Lock until all ACK messages will be sent to Redis.
func (c *Consumer) Close() {
	close(c.cAck)
	c.wgAck.Wait()
}

func (c *Consumer) consume(shard int) {
	group := fmt.Sprintf("qu_%s_group", c.opt.Name)
	stream := fmt.Sprintf("qu{%d}_%s", shard, c.opt.Name)

	lastID := "0-0"
	checkBacklog := true

	// Millisecond is minimal for Redis
	block := time.Second * 1
	if c.opt.Block != 0 {
		block = c.opt.Block
	}

	for {
		if c.needStop {
			break
		}

		var id string
		if checkBacklog {
			id = lastID
		} else {
			id = ">"
		}

		var res []redis.XStream

		err := repeat.Repeat(
			repeat.Fn(func() error {
				var err error

				res, err = c.cl.rDB.XReadGroup(context.Background(), &redis.XReadGroupArgs{
					Block:    block,
					Consumer: c.opt.Consumer,
					Count:    c.opt.PrefetchCount,
					Group:    group,
					Streams:  []string{stream, id},
				}).Result()

				if err != nil && err != redis.Nil {
					if c.notif != nil {
						c.notif.AmiError(err)
					}

					return repeat.HintTemporary(err)
				}

				return nil
			}),
			repeat.StopOnSuccess(),
			repeat.WithDelay(repeat.FullJitterBackoff(500*time.Millisecond).Set()),
		)

		if err != nil {
			if c.notif != nil {
				c.notif.AmiError(err)
			}

			continue
		}

		if checkBacklog && (len(res) == 0 || len(res[0].Messages) == 0) {
			checkBacklog = false
			continue
		}

		for _, s := range res {
			for _, m := range s.Messages {
				lastID = m.ID

				msg := Message{
					Group:  group,
					ID:     m.ID,
					Stream: stream,
				}

				v, ok := m.Values["m"]
				if !ok {
					if c.notif != nil {
						c.notif.AmiError(errors.New("Incorrect message format: no \"m\" field in message with id " + m.ID))
					}

					c.Ack(msg)

					continue
				}

				msg.Body = v.(string)
				c.cCons <- msg
			}
		}
	}

	c.wgCons.Done()
}

// Ack acknowledges message
//
// Function not only do XACK call, but additionally it deletes message
// from stream with XDELETE.
// Ack do not do immediately, but pushed to send buffer and sended to Redis
// in other goroutine.
func (c *Consumer) Ack(m Message) {
	c.cAck <- m
}

func (c *Consumer) ack() {
	started := time.Now()
	tick := time.NewTicker(c.opt.PipePeriod)

	toAck := make(map[string][]Message)
	cnt := make(map[string]int)

	for {
		var (
			doStop bool
			stream string
		)

		select {
		case m, more := <-c.cAck:
			if !more {
				doStop = true
			} else {
				stream = m.Stream

				if toAck[stream] == nil {
					toAck[stream] = make([]Message, c.opt.PipeBufferSize)
					cnt[stream] = 0
				}

				toAck[stream][cnt[stream]] = m
				cnt[stream]++
			}
		case <-tick.C:
		}

		if doStop {
			c.sendAckAllStreams(toAck, cnt)
			break
		}

		if cnt[stream] >= int(c.opt.PipeBufferSize) {
			c.sendAckStreamWithLock(toAck[stream][0:cnt[stream]])
			cnt[stream] = 0
		} else if time.Since(started) >= c.opt.PipePeriod && len(c.cAck) == 0 {
			// Don't send by time if there are more messages in channel
			// Prefer to collect them in batch to speedup producing
			c.sendAckAllStreams(toAck, cnt)
			started = time.Now()
		}
	}

	c.wgAck.Done()
}

func (c *Consumer) sendAckAllStreams(toAck map[string][]Message, cnt map[string]int) {
	for stream := range toAck {
		c.sendAckStreamWithLock(toAck[stream][0:cnt[stream]])
		cnt[stream] = 0
	}
}

func (c *Consumer) sendAckStreamWithLock(lst []Message) {
	if len(lst) == 0 {
		return
	}

	ids := make([]string, len(lst))
	for i, m := range lst {
		ids[i] = m.ID
	}

	c.wgAck.Add(1)

	go func() {
		c.sendAckStream(lst[0].Stream, lst[0].Group, ids)
		c.wgAck.Done()
	}()
}

func (c *Consumer) sendAckStream(stream string, group string, ids []string) {
	err := repeat.Repeat(
		repeat.Fn(func() error {
			pipe := c.cl.rDB.TxPipeline()

			pipe.XAck(context.Background(), stream, group, ids...)
			pipe.XDel(context.Background(), stream, ids...)

			_, err := pipe.Exec(context.Background())

			if err != nil {
				if c.notif != nil {
					c.notif.AmiError(err)
				}

				return repeat.HintTemporary(err)
			}

			return nil
		}),
		repeat.StopOnSuccess(),
		repeat.WithDelay(repeat.FullJitterBackoff(500*time.Millisecond).Set()),
	)

	if err != nil && c.notif != nil {
		c.notif.AmiError(err)
	}
}
