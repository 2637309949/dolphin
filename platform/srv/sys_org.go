// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_org.go

package srv

import (
	"context"
	"errors"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/sirupsen/logrus"
)

type SysOrg struct {
	svc.Svc
}

func NewSysOrg() *SysOrg {
	return &SysOrg{}
}

// TODO defined srv
func (srv *SysOrg) TODO(ctx context.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
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
