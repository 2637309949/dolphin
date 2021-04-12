// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_table_column.go

package srv

import (
	"context"
	"errors"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// SysTableColumnTODO defined srv
func SysTableColumnTODO(ginCtx *gin.Context, db *xorm.Engine, actCtx context.Context, params struct{}) (interface{}, error) {
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
	logrus.Infoln("main process exit!")
	return nil, errors.New("no implementation found")
}
