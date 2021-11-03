// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"net/http/httputil"
	"os"
	"time"

	"github.com/2637309949/dolphin/packages/trace"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

// Hostname returns the name of the host, if no hostname, a random id is returned.
func Hostname() string {
	var err error
	var hostname string
	hostname, err = os.Hostname()
	if err != nil {
		hostname = trace.RandId()
	}
	return hostname
}

// HttpTrace defined TODO
func HttpTrace() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		carrier, err := trace.Extract(trace.HttpFormat, ctx.Request.Header)
		if err != nil && err != trace.ErrInvalidCarrier {
			logrus.Error(err)
		}
		c, span := trace.StartServerSpan(ctx.Request.Context(), carrier, Hostname(), ctx.Request.RequestURI)
		defer span.Finish()
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}

// Cors defined TODO
func Cors(origin string, headers string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", headers)
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Disposition")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

// Recovery defined TODO
func Recovery() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(ctx.Request, false)
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				logrus.Errorf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
				ctx.JSON(200, map[string]interface{}{
					"code":   "500",
					"detail": goErr.Error(),
					"status": "",
				})
			}
		}()
		ctx.Next()
	}
}

// DumpBody instance a Logger middleware with config.
func DumpBody(dumpCall func(*gin.Context, *LogFormatterParams)) func(ctx *gin.Context) {
	formatter := Formatter
	notlogged := []string{}
	skip := map[string]struct{}{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		var bufWriter bytes.Buffer
		var bufReader bytes.Buffer
		var buf []byte
		if cb, ok := c.Get(gin.BodyBytesKey); ok {
			if cbb, ok := cb.([]byte); ok {
				buf = cbb
			}
		}
		if buf == nil {
			c.Request.Body = &NopCloser{c.Request.Body, &bufWriter}
		}

		c.Writer = &bodyLogWriter{body: &bufReader, ResponseWriter: c.Writer}

		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// resove header
		hr := bytes.Buffer{}
		c.Request.Header.Write(&hr)

		// Process request
		c.Next()

		// Bytes set
		if buf == nil {
			buf = bufWriter.Bytes()
			c.Set(gin.BodyBytesKey, buf)
		}

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				LogFormatterParams: gin.LogFormatterParams{
					Request: c.Request,
					// IsTerm:  isTerm,
					Keys: c.Keys,
				},
				Header:  hr.Bytes(),
				ReqBody: buf,
				ResBody: bufReader.Bytes(),
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			logrus.Info(formatter(param.LogFormatterParams))

			// LogFormatterParams defined recorder
			dumpCall(c, &param)
		}
	}
}
