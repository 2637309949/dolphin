// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"context"
	"encoding/json"

	"github.com/2637309949/dolphin/packages/logrus"
)

const (
	WorkerStatusInitial = iota + 100
	WorkerStatusProccessing
	WorkerStatusInterrupt
	WorkerStatusFinish
)

// TableName table name of defined SysUser
func (m *Worker) Unmarshal(v interface{}) error {
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(m.Payload); err != nil {
		logrus.Error(context.TODO(), err)
		return err
	}
	if err := json.Unmarshal(bs, v); err != nil {
		logrus.Error(context.TODO(), err)
		return err
	}
	return nil
}
