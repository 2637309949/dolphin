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
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/slice"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

// TrackerStore store logs
var TrackerStore func(beans *[]model.SysTracker) error
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
		logChannel <- p
	}
}

func initTracker() {
	logChannel = make(chan *plugin.LogFormatterParams, 600)
	TrackerStore = func(beans *[]model.SysTracker) error {
		App.PlatformDB.ShowSQL(false)
		_, err := App.PlatformDB.Insert(*beans)
		App.PlatformDB.ShowSQL(viper.GetString("app.mode") == "debug")
		return err
	}
	go func() {
		var bucket slice.SyncSlice
		for {
			select {
			case <-time.Tick(6 * time.Second):
				logs := bucket.Reset()
				beans := funk.Map(logs, func(entity interface{}) model.SysTracker {
					item := entity.(*plugin.LogFormatterParams)
					return model.SysTracker{
						ID:         null.StringFromUUID(),
						AppName:    null.StringFrom(viper.GetString("app.name")),
						Token:      null.StringFrom(item.Token),
						Domain:     null.StringFrom(item.Domain),
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
					}
				}).([]model.SysTracker)
				if len(beans) > 0 {
					err := TrackerStore(&beans)
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
