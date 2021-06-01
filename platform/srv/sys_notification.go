// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_notification.go

package srv

import (
	"context"
	"errors"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// SysNotificationTODO defined srv
func SysNotificationTODO(ctx *gin.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
	cwt, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	go func(ctx context.Context) {
	}(cwt)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			logrus.Infoln("child job...")
		}
	}
	return nil, errors.New("no implementation found")
}
