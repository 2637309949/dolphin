// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/log"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

// XLogger defined
type XLogger struct {
	*logrus.Logger
	showSQL bool
	level   log.LogLevel
}

// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {
	gin.LogFormatterParams
	Domain  string
	Token   string
	UserID  int64
	Header  []byte
	ReqBody []byte
	ResBody []byte
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// NopCloser defined
type NopCloser struct {
	rc io.ReadCloser
	w  io.Writer
}

// Close defined
func (rc *NopCloser) Close() error {
	return rc.rc.Close()
}

// Read defined
func (rc *NopCloser) Read(p []byte) (n int, err error) {
	n, err = rc.rc.Read(p)
	if n > 0 {
		if n, err := rc.w.Write(p[:n]); err != nil {
			return n, err
		}
	}
	return n, err
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
}

// TrackerStore store logs
var TrackerStore func(domain string, beans *[]types.SysTracker) error
var logWorkerPool chan chan *LogFormatterParams
var logChannel chan *LogFormatterParams

// Formatter log formatter params
func Formatter(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("%3d | %13v | %15s |%-7s %s%s",
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}

// DumpRecv defined tracker recorder
func DumpRecv() func(ctx *Context, p *LogFormatterParams) {
	return func(ctx *Context, p *LogFormatterParams) {
		token, _ := App.Identity.OAuth2.BearerAuth(ctx.Request())
		p.Token = token
		if tokenInfo, err := App.Identity.OAuth2.Manager.LoadAccessToken(token); err == nil {
			p.Domain = tokenInfo.GetDomain()
			it, _ := strconv.ParseInt(tokenInfo.GetUserID(), 10, 64)
			p.UserID = it
		}

		jobChannel := <-logWorkerPool
		go func(p *LogFormatterParams) {
			jobChannel <- p
		}(p)
	}
}

func InitTracker() {
	logWorkerPool = make(chan chan *LogFormatterParams, 20)
	logChannel = make(chan *LogFormatterParams, 150)
	TrackerStore = func(domain string, beans *[]types.SysTracker) error {
		if db := App.Manager.GetBusinessDB(domain); db != nil {
			db, err := db.Clone()
			if err != nil {
				logrus.Error(err)
				return err
			}
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
		var timer = time.NewTicker(5 * time.Second)
		for {
			logWorkerPool <- logChannel
			select {
			case <-timer.C:
				logs := bucket.Reset()
				bmp := map[string][]types.SysTracker{}
				beans := funk.Map(logs, func(entity interface{}) types.TrackerInfo {
					item := entity.(*LogFormatterParams)
					st := types.TrackerInfo{Domain: null.StringFrom(item.Domain)}
					st.SysTracker = &types.SysTracker{
						Token:      null.StringFrom(item.Token),
						UserId:     null.IntFrom(item.UserID),
						StatusCode: null.IntFrom(int64(item.StatusCode)),
						Latency:    null.FloatFrom(item.Latency.Seconds()),
						ClientIp:   null.StringFrom(item.ClientIP),
						Method:     null.StringFrom(item.Method),
						Path:       null.StringFrom(item.Path),
						Header:     item.Header,
						ReqBody:    item.ReqBody,
						ResBody:    item.ResBody,
						CreateTime: null.TimeFromNow(),
						Creater:    null.IntFrom(1),
						UpdateTime: null.TimeFromNow(),
						Updater:    null.IntFrom(1),
						IsDelete:   null.IntFrom(0),
					}
					return st
				}).([]types.TrackerInfo)
				beans = funk.Filter(beans, func(entity types.TrackerInfo) bool {
					return entity.Domain.String != ""
				}).([]types.TrackerInfo)
				funk.ForEach(beans, func(entity types.TrackerInfo) {
					if _, ok := bmp[entity.Domain.String]; !ok {
						bmp[entity.Domain.String] = []types.SysTracker{}
					}
					bmp[entity.Domain.String] = append(bmp[entity.Domain.String], *entity.SysTracker)
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
	xlogger := &XLogger{Logger: logrus.StandardLogger()}
	xlogger.SetLevel(log.DEFAULT_LOG_LEVEL)
	xlogger.ShowSQL(IsDebugging())
	return xlogger
}

// InitLogger defined TODO
func InitLogger() {
	var writer io.Writer = os.Stdout
	util.SetFormatter(term.IsTerminal(unix.Stdout))
	if viper.GetString("dir.log") != "" {
		logf, err := rotatelogs.New(
			viper.GetString("dir.log")+"/%Y%m%d%H",
			rotatelogs.WithMaxAge(24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			logrus.Printf("failed to create rotatelogs: %v", err)
			return
		}
		writer = io.MultiWriter(writer, logf)
	}

	logrus.SetOutput(writer)
	InitTracker()
}
