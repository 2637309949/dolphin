// Code generated by dol build. Only Generate by tools if not existed.
// source: article.go

package srv

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/app"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ArticleTODO defined srv
func ArticleTODO(ctx *gin.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
	cwt, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	chi := func(cwt context.Context) chan interface{} {
		chi := make(chan interface{}, 1)
		go func() {
			time.Sleep(1 * time.Second)
			chi <- 100
			close(chi)
			cancel()
		}()
		return chi
	}(cwt)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
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

func init() {
	id, _ := app.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
		fmt.Println("hello")
	})
	entry, _ := app.App.Manager.Cron().RefreshFunc(id, "*/3 * * * * *")
	app.App.Manager.Cron().TryFunc(entry)
	app.App.Manager.Worker().AddJobHandler("hello", func(args model.Worker) (interface{}, error) {
		fmt.Printf("topic=%v, payload=%v", "hello", args.Payload)
		return map[string]interface{}{
			"score": 99,
		}, nil
	})
}
