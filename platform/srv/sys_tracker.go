// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tracker.go

package srv

import (
	"context"
	"errors"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
)

type SysTracker struct {
	svc.Svc
}

func NewSysTracker() *SysTracker {
	return &SysTracker{}
}

// PageFormatter defined TODO
func (srv *SysTracker) PageFormatter(db1 *xorm.Engine) func(*xorm.Engine, []map[string]interface{}) ([]map[string]interface{}, error) {
	return func(db2 *xorm.Engine, items []map[string]interface{}) (data []map[string]interface{}, err error) {
		if uids, ok := slice.GetFieldSliceByName(items, "user_id").([]string); ok {
			users := []types.SysUser{}
			err := db1.Cols("id", "name").In("id", uids).Find(&users)
			if err != nil {
				return data, err
			}
			data, err = slice.PatchSliceByField(items, users, "user_id", "id", "user_name#name")
			if err != nil {
				return data, err
			}
		}
		return data, err
	}
}

// TODO defined srv
func (srv *SysTracker) TODO(ctx context.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
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
