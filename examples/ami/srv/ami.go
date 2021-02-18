// Code generated by dol build. Only Generate by tools if not existed.
// source: ami.go

package srv

import (
	"time"

	"github.com/2637309949/dolphin/packages/ami"
	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/redis"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

type errorLogger struct{}

func (l *errorLogger) AmiError(err error) {
	println("Got error from Ami:", err.Error())
}

// Producer defined
var Producer *ami.Producer

// Consumer defined
var Consumer *ami.Consumer

func init() {
	var err error
	Producer, err = ami.NewProducer(
		ami.ProducerOptions{
			Name:              "ami",
			ErrorNotifier:     &errorLogger{},
			PendingBufferSize: 10000000,
			PipeBufferSize:    50000,
			PipePeriod:        time.Microsecond * 1000,
			ShardsCount:       10,
		},
		&redis.ClusterOptions{
			Addrs:        []string{":6379"},
			ReadTimeout:  time.Second * 60,
			WriteTimeout: time.Second * 60,
		},
	)
	if err != nil {
		panic(err)
	}

	Consumer, err = ami.NewConsumer(
		ami.ConsumerOptions{
			Name:              "ami",
			Consumer:          "ami",
			ErrorNotifier:     &errorLogger{},
			PendingBufferSize: 10000000,
			PipeBufferSize:    50000,
			PipePeriod:        time.Microsecond * 1000,
			PrefetchCount:     100,
			ShardsCount:       10,
		},
		&redis.ClusterOptions{
			Addrs:        []string{":6379"},
			ReadTimeout:  time.Second * 60,
			WriteTimeout: time.Second * 60,
		},
	)
	if err != nil {
		panic(err)
	}

	c := Consumer.Start()

	go func() {
		for {
			m, _ := <-c
			println("Got", m.Body, "ID", m.ID)
			Consumer.Ack(m)
		}
	}()
}

// AmiAction defined srv
func AmiAction(ctx *gin.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
	for i := 0; i < 10000; i++ {
		Producer.Send("{}")
	}
	return nil, nil
}
