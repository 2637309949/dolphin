package svc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

// Kafka defined TODO
type Kafka interface {
	WriteMessages(ctx context.Context, topic string, message interface{}) error
	ReadMessage(ctx context.Context, topic string, cb func([]byte), cnt ...int) error
}

type XKafka struct {
	dialer *kafka.Dialer
}

// WriteMessages defined TODO
func (x *XKafka) WriteMessages(ctx context.Context, topic string, message interface{}) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Dialer:       x.dialer,
		Brokers:      viper.GetStringSlice("kafka.broker"),
		Topic:        topic,
		BatchSize:    2,
		BatchTimeout: 2 * time.Second,
		RequiredAcks: 1,
		Balancer:     &kafka.Hash{},
	})
	kpStr, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = writer.WriteMessages(ctx, kafka.Message{Value: kpStr})
	if err != nil {
		return err
	}
	return nil
}

// ReadMessage defined TODO
func (x *XKafka) ReadMessage(ctx context.Context, topic string, cb func([]byte), cnt ...int) error {
	cch := util.SomeOne(cnt, int(1)).(int)
	var WokerCntChannel = make(chan bool, cch)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Dialer:      x.dialer,
		Brokers:     viper.GetStringSlice("kafka.broker"),
		Topic:       topic,
		GroupID:     "group",
		MinBytes:    5,
		MaxBytes:    1e6,
		MaxWait:     3 * time.Second,
		StartOffset: kafka.LastOffset,
		// Logger:      log.New(os.Stdout, "kafka reader: ", 0),
	})
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			logrus.Error(ctx, err)
			return err
		}
		WokerCntChannel <- true
		go func() {
			cb(msg.Value)
			<-WokerCntChannel
		}()
	}
}

func NewXKafka() *XKafka {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	return &XKafka{dialer: dialer}
}
