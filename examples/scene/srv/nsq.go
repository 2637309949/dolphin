// Code generated by dol build. Only Generate by tools if not existed.
// source: nsq.go

package srv

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-errors/errors"

	"scene/svc"
	"scene/types"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

var (
	// NsqProducer defined
	NsqProducer *nsq.Producer
	// NsqConsumer defined
	NsqConsumer *nsq.Consumer
	// tcpNsqdAddrr defined
	tcpNsqdAddrr = "172.20.10.3:4161"
)

type messageHandler struct{}

type Nsq struct {
	*svc.ServiceContext
}

func NewNsq() *Nsq {
	return &Nsq{}
}

// SetServiceContext defined TODO
func (h *messageHandler) SetServiceContext(svc *svc.ServiceContext) {
	srv.ServiceContext = svc
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

// Producer defined srv
func (srv *Nsq) Producer(ctx context.Context, db *xorm.Engine, params types.NsqInfo) (interface{}, error) {
	aiStr, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	err = NsqProducer.Publish("nsq-test", aiStr)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// TODO defined srv
func (srv *Nsq) TODO(ctx context.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
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
			logrus.Infoln("child process interrupt...")
			return <-chi, cwt.Err()
		default:
			logrus.Infoln("awaiting job...")
		}
	}
	return nil, errors.New("no implementation found")
}

// NConsumer defined srv
func NConsumer() (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err.(error), 3)
			logrus.Error(string(goErr.Stack()))
		}
	}()

	NsqConsumer.AddHandler(&messageHandler{})
	err := NsqConsumer.ConnectToNSQLookupd(tcpNsqdAddrr)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func init() {
	var err error
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 10 * time.Second
	NsqConsumer, err = nsq.NewConsumer("nsq-test", "test-channel", cfg)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	NsqProducer, err = nsq.NewProducer(tcpNsqdAddrr, cfg)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	go NConsumer()
}
