// Code generated by dol build. Only Generate by tools if not existed.
// source: kafka.go

package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"scene/model"
	"time"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/util/worker"
	"github.com/go-errors/errors"
	kafka "github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer
var Reader *kafka.Reader
var KafkaDispatcher *worker.Dispatcher
var WokerCntChannel = make(chan bool, 2)

func init() {
	broker1Address, topic := "172.16.10.191:9092", "score-topic"
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	Writer = kafka.NewWriter(kafka.WriterConfig{
		Dialer:       dialer,
		Brokers:      []string{broker1Address},
		Topic:        topic,
		BatchSize:    2,
		BatchTimeout: 2 * time.Second,
		RequiredAcks: 1,
		Balancer:     &kafka.Hash{},
	})
	Reader = kafka.NewReader(kafka.ReaderConfig{
		Dialer:      dialer,
		Brokers:     []string{broker1Address},
		Topic:       topic,
		GroupID:     "kafka-group",
		MinBytes:    5,
		MaxBytes:    1e6,
		MaxWait:     3 * time.Second,
		StartOffset: kafka.LastOffset,
		// Logger:      log.New(os.Stdout, "kafka reader: ", 0),
	})
	go KafkaConsumer(context.Background())
}

// KafkaProducer defined srv
func KafkaProducer(ctx context.Context, params model.KafkaInfo) error {
	kpStr, err := json.Marshal(&params)
	if err != nil {
		return err
	}
	err = Writer.WriteMessages(ctx, kafka.Message{Value: kpStr})
	if err != nil {
		return err
	}
	return nil
}

// KafkaConsumer defined srv
func KafkaConsumer(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err.(error), 3)
			fmt.Print(string(goErr.Stack()))
		}
	}()
	for {
		msg, err := Reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(err)
			break
		}
		value := model.KafkaInfo{}
		if err := json.Unmarshal(msg.Value, &value); err != nil {
			logrus.Error("failed to unmarshal:", err)
		}
		WokerCntChannel <- true
		go func() {
			Woker(msg.Value)
			<-WokerCntChannel
		}()
	}
}

func Woker(value []byte) {
	fmt.Println("------start woker: ", string(value))
	time.Sleep(10 * time.Second)
}
