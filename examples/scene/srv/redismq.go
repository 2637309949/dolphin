// Code generated by dol build. Only Generate by tools if not existed.
// source: redismq.go

package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"scene/svc"
	"scene/types"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/go-errors/errors"
	"github.com/go-redis/redis/v7"
	"github.com/kak-tus/ami"
	"github.com/spf13/viper"
)

type errorLogger struct{}

func (l *errorLogger) AmiError(err error) {
	println("Got error from Ami:", err.Error())
}

// RedisProducer defined TODO
var RedisProducer *ami.Producer

// RedisConsumer defined TODO
var RedisConsumer *ami.Consumer

func init() {
	var err error
	RedisProducer, err = ami.NewProducer(ami.ProducerOptions{
		Name:              "ami",
		ErrorNotifier:     &errorLogger{},
		PendingBufferSize: 10000000,
		PipeBufferSize:    50000,
		PipePeriod:        time.Microsecond * 1000,
		ShardsCount:       10,
	}, &redis.ClusterOptions{
		Addrs:        viper.GetStringSlice("redis.addr"),
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	})
	if err != nil {
		logrus.Error(context.TODO(), err)
	}
	RedisConsumer, err = ami.NewConsumer(ami.ConsumerOptions{
		Name:              "ami",
		Consumer:          "ami",
		ErrorNotifier:     &errorLogger{},
		PendingBufferSize: 10000000,
		PipeBufferSize:    50000,
		PipePeriod:        time.Microsecond * 1000,
		PrefetchCount:     100,
		ShardsCount:       10,
	}, &redis.ClusterOptions{
		Addrs:        viper.GetStringSlice("redis.addr"),
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	})
	if err != nil {
		logrus.Error(context.TODO(), err)
	}
}

type RedisMq struct {
	*svc.ServiceContext
}

func NewRedisMq() *RedisMq {
	return &RedisMq{}
}

// SetServiceContext defined TODO
func (srv *RedisMq) SetServiceContext(svc *svc.ServiceContext) {
	srv.ServiceContext = svc
}

// Producer defined TODO
func (srv *RedisMq) Producer(ctx context.Context, db *xorm.Engine, params types.AmiInfo) (interface{}, error) {
	aiStr, err := json.Marshal(params)
	if err != nil {
		logrus.Error(ctx, "failed to marshal:", err)
		return nil, err
	}
	RedisProducer.Send(string(aiStr))
	return nil, nil
}

// Consumer defined TODO
func (srv *RedisMq) Consumer(ctx context.Context, db *xorm.Engine, params map[string]interface{}) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err.(error), 3)
			fmt.Print(string(goErr.Stack()))
		}
	}()
	var items []types.AmiInfo
	c := RedisConsumer.Start()
	cwt, cancel := context.WithCancel(context.TODO())
	go func(context.Context) {
		for {
			select {
			case m, ok := <-c:
				if !ok {
					cancel()
					return
				}
				RedisConsumer.Ack(m)
				value := types.AmiInfo{}
				if err := json.Unmarshal([]byte(m.Body), &value); err != nil {
					logrus.Error(ctx, "failed to unmarshal:", err)
				}
				items = append(items, value)
			case <-time.After(3 * time.Second):
				cancel()
				return
			}
		}
	}(cwt)
	<-cwt.Done()
	// RedisConsumer.Stop()
	return items, nil
}

// TODO defined TODO
func (srv *RedisMq) TODO(ctx context.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
	cwt, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	chi := func(context.Context) chan interface{} {
		chi := make(chan interface{}, 1)
		go func() {
			time.Sleep(1 * time.Second)
			chi <- 100
			cancel()
			close(chi)
		}()
		return chi
	}(cwt)
	for range ticker.C {
		select {
		case <-cwt.Done():
			logrus.Infoln(ctx, "child process interrupt...")
			return <-chi, cwt.Err()
		default:
			logrus.Infoln(ctx, "awaiting job...")
		}
	}
	return nil, errors.New("no implementation found")
}
