package types

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/logrus"
)

// With defined TODO
func (m *PageList) With(s interface{}) *PageList {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		panic(errors.New("not ptr found"))
	}
	mbt, err := json.Marshal(m.Data)
	if err != nil {
		logrus.Error(context.TODO(), err)
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		logrus.Error(context.TODO(), err)
	}

	// unmarshal back
	mbt, err = json.Marshal(s)
	if err != nil {
		logrus.Error(context.TODO(), err)
	}
	if err := json.Unmarshal(mbt, &m.Data); err != nil {
		logrus.Error(context.TODO(), err)
	}
	return m
}
