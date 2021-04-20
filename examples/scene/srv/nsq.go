// Code generated by dol build. Only Generate by tools if not existed.
// source: nsq.go

package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"scene/model"
	"time"

	"github.com/go-errors/errors"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

type messageHandler struct{}

// NsqProducer defined
var NsqProducer *nsq.Producer

// NsqConsumer defined
var NsqConsumer *nsq.Consumer

var tcpNsqdAddrr = "127.0.0.1:4161"

func init() {
	var err error
	cfg := nsq.NewConfig()
	// cfg.LookupdPollInterval = 10 * time.Second
	NsqConsumer, err = nsq.NewConsumer("nsq-test", "test-channel", cfg)
	if err != nil {
		panic(err)
	}
	NsqProducer, err = nsq.NewProducer(tcpNsqdAddrr, cfg)
	if err != nil {
		panic(err)
	}
	go NConsumer()
}

// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}
	logrus.Infoln("NSQL HandleMessage: ", string(m.Body))
	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	m.Finish()
	return nil
}

// NProducer defined srv
func NProducer(ctx *gin.Context, db *xorm.Engine, params model.NsqInfo) (interface{}, error) {
	aiStr, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	err = NsqProducer.Publish("nsq-test", aiStr)
	// if err != nil {
	// 	return nil, err
	// }
	// NsqProducer.Stop()
	return nil, err
}

// NConsumer defined srv
func NConsumer() (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err.(error), 3)
			fmt.Print(string(goErr.Stack()))
		}
	}()

	NsqConsumer.AddHandler(&messageHandler{})
	err := NsqConsumer.ConnectToNSQLookupd(tcpNsqdAddrr)
	if err != nil {
		logrus.Error("failed to ConnectToNSQLookupd:", err)
		return nil, err
	}
	// NsqConsumer.Stop()
	return nil, err
}

// NsqTODO defined srv
func NsqTODO(ginCtx *gin.Context, db *xorm.Engine, actCtx context.Context, params struct{}) (interface{}, error) {
	actCtx, cancel := context.WithTimeout(actCtx, 5*time.Second)
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-actCtx.Done():
				logrus.Infoln("child process interrupt...")
				return
			default:
				logrus.Infoln("child job...")
			}
		}
	}(actCtx)
	defer cancel()
	<-actCtx.Done()
	time.Sleep(1 * time.Second)
	logrus.Infoln("main process exit!")
	return nil, errors.New("no implementation found")
}
