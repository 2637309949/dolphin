// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrotate"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/log"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

// XLogger defined
type XLogger struct {
	*logrus.Logger
	showSQL bool
	level   log.LogLevel
}

// ShowSQL implement ILogger
func (s *XLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement ILogger
func (s *XLogger) IsShowSQL() bool {
	return s.showSQL
}

// Level implement ILogger
func (s *XLogger) Level() log.LogLevel {
	return s.level
}

// SetLevel implement ILogger
func (s *XLogger) SetLevel(l log.LogLevel) {
	s.level = l
	return
}

// TrackerStore store logs
var TrackerStore func(domain string, beans *[]model.SysTracker) error
var logWorkerPool chan chan *plugin.LogFormatterParams
var logChannel chan *plugin.LogFormatterParams

// Tracker defined tracker recorder
func Tracker(e *Engine) func(ctx *gin.Context, p *plugin.LogFormatterParams) {
	return func(ctx *gin.Context, p *plugin.LogFormatterParams) {
		token, _ := e.OAuth2.BearerAuth(ctx.Request)
		p.Token = token
		if tokenInfo, err := e.OAuth2.Manager.LoadAccessToken(token); err == nil {
			p.Domain = tokenInfo.GetDomain()
			p.UserID = tokenInfo.GetUserID()
		}

		// would not be block <-logWorkerPool
		// but <- p will
		jobChannel := <-logWorkerPool
		go func(p *plugin.LogFormatterParams) {
			jobChannel <- p
		}(p)
	}
}

func initTracker() {
	logWorkerPool = make(chan chan *plugin.LogFormatterParams, 20)
	logChannel = make(chan *plugin.LogFormatterParams, 150)
	TrackerStore = func(domain string, beans *[]model.SysTracker) error {
		if db := App.Manager.GetBusinessDB(domain); db != nil {
			db, err := db.Clone()
			db.ShowSQL(false)
			_, err = db.Insert(*beans)
			if err != nil {
				logrus.Error(err)
				return err
			}
			db.Close()
			return nil
		}
		return errors.Errorf("no datasource found, %v", domain)
	}
	go func() {
		var bucket slice.SyncSlice
		type SysTracker struct {
			model.SysTracker
			Domain null.String `xorm:"varchar(36) comment('域名') 'domain'" json:"domain" xml:"domain"`
		}
		for {
			logWorkerPool <- logChannel
			select {
			case <-time.Tick(5 * time.Second):
				logs := bucket.Reset()
				bmp := map[string][]model.SysTracker{}
				beans := funk.Map(logs, func(entity interface{}) SysTracker {
					item := entity.(*plugin.LogFormatterParams)
					return SysTracker{
						SysTracker: model.SysTracker{
							ID:         null.StringFromUUID(),
							Token:      null.StringFrom(item.Token),
							UserId:     null.StringFrom(item.UserID),
							StatusCode: null.IntFrom(int64(item.StatusCode)),
							Latency:    null.FloatFrom(item.Latency.Seconds()),
							ClientIp:   null.StringFrom(item.ClientIP),
							Method:     null.StringFrom(item.Method),
							Path:       null.StringFrom(item.Path),
							Header:     item.Header,
							ReqBody:    item.ReqBody,
							ResBody:    item.ResBody,
							CreateTime: null.TimeFrom(time.Now()),
							CreateBy:   model.DefaultAdmin.ID,
							UpdateTime: null.TimeFrom(time.Now()),
							UpdateBy:   model.DefaultAdmin.ID,
							DelFlag:    null.IntFrom(0),
						},
						Domain: null.StringFrom(item.Domain),
					}
				}).([]SysTracker)
				beans = funk.Filter(beans, func(entity SysTracker) bool {
					return entity.Domain.String != ""
				}).([]SysTracker)
				funk.ForEach(beans, func(entity SysTracker) {
					if _, ok := bmp[entity.Domain.String]; !ok {
						bmp[entity.Domain.String] = []model.SysTracker{}
					}
					bmp[entity.Domain.String] = append(bmp[entity.Domain.String], entity.SysTracker)
				})
				for k, v := range bmp {
					err := TrackerStore(k, &v)
					if err != nil {
						logrus.Error(err)
					}
				}
			case info := <-logChannel:
				bucket.Append(info)
			}
		}
	}()
}

func createXLogger() interface{} {
	xlogger := &XLogger{
		Logger: logrus.StandardLogger(),
	}
	xlogger.SetLevel(log.DEFAULT_LOG_LEVEL)
	xlogger.ShowSQL(viper.GetString("app.mode") == "debug")
	return xlogger
}

func init() {
	var writer io.Writer
	util.SetFormatter(terminal.IsTerminal(unix.Stdout))
	if terminal.IsTerminal(unix.Stdout) {
		writer = os.Stdout
	}
	dir := path.Join(viper.GetString("dir.log"), viper.GetString("app.name"))
	logf, err := logrotate.New(
		dir+".%Y%m%d%H",
		logrotate.WithMaxAge(24*time.Hour),
		logrotate.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Printf("failed to create rotatelogs: %s", err)
		return
	}
	if writer != nil {
		writer = io.MultiWriter(writer, logf)
	} else {
		writer = logf
	}
	logrus.SetOutput(writer)
	initTracker()
}
