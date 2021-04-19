// Code generated by dol build. Only Generate by tools if not existed.
// source: redismq.go

package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"scene/model"
	"time"

	"github.com/2637309949/dolphin/packages/redismq"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type errorLogger struct{}

func (l *errorLogger) AmiError(err error) {
	println("Got error from Ami:", err.Error())
}

// RedisProducer defined
var RedisProducer *redismq.Producer

// RedisConsumer defined
var RedisConsumer *redismq.Consumer

func init() {
	var err error
	RedisProducer, err = redismq.NewProducer(
		redismq.ProducerOptions{
			Name:              "ami",
			ErrorNotifier:     &errorLogger{},
			PendingBufferSize: 10000000,
			PipeBufferSize:    50000,
			PipePeriod:        time.Microsecond * 1000,
			ShardsCount:       10,
		},
		&redis.Options{
			Addr:         ":6379",
			ReadTimeout:  time.Second * 60,
			WriteTimeout: time.Second * 60,
		},
	)
	if err != nil {
		panic(err)
	}
	RedisConsumer, err = redismq.NewConsumer(
		redismq.ConsumerOptions{
			Name:              "ami",
			Consumer:          "ami",
			ErrorNotifier:     &errorLogger{},
			PendingBufferSize: 10000000,
			PipeBufferSize:    50000,
			PipePeriod:        time.Microsecond * 1000,
			PrefetchCount:     100,
			ShardsCount:       10,
		},
		&redis.Options{
			Addr:         ":6379",
			ReadTimeout:  time.Second * 60,
			WriteTimeout: time.Second * 60,
		},
	)
	if err != nil {
		panic(err)
	}
}

// Producer defined srv
func Producer(ctx *gin.Context, db *xorm.Engine, params model.AmiInfo) (interface{}, error) {
	aiStr, err := json.Marshal(params)
	if err != nil {
		logrus.Error("failed to marshal:", err)
		return nil, err
	}
	RedisProducer.Send(string(aiStr))
	return nil, nil
}

// Consumer defined srv
func Consumer(ctx *gin.Context, db *xorm.Engine, params map[string]interface{}) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err.(error), 3)
			fmt.Print(string(goErr.Stack()))
		}
	}()
	var items []model.AmiInfo
	c := RedisConsumer.Start()
	cwt, cancel := context.WithCancel(context.TODO())
	go func(cwt context.Context) {
		for {
			select {
			case m, ok := <-c:
				if !ok {
					cancel()
					return
				}
				RedisConsumer.Ack(m)
				value := model.AmiInfo{}
				if err := json.Unmarshal([]byte(m.Body), &value); err != nil {
					logrus.Error("failed to unmarshal:", err)
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