// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"net/http/httputil"
	"os"
	"time"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/trace"
	"github.com/2637309949/dolphin/packages/trace/tracespec"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/gin-gonic/gin"
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

// Roles middles TODO
func Roles(roles ...string) func(ctx *Context) {
	return func(ctx *Context) {
		db := ctx.MustDB()
		svc, userId := svc.NewXDB(), ctx.MustToken().GetUserID()
		if !svc.InRole(db, userId, roles...) {
			ctx.Fail(errors.ErrAccessDenied)
			return
		}
		ctx.Next()
	}
}

// Auth middles TODO
func Auth(auth ...string) func(ctx *Context) {
	return func(ctx *Context) {
		for _, ah := range auth {
			provider := App.Identity.GetProvider(ah)
			if provider == nil {
				ctx.Fail(errors.ErrNotFoundProvider)
				return
			}

			tokenInfo, ok := provider.Verify(ctx)
			if !ok {
				ctx.Fail(errors.ErrAuthenticationFailed)
				return
			}

			if tokenInfo.GetDomain() != "" {
				db := App.Manager.GetBusinessDB(tokenInfo.GetDomain())
				if db == nil {
					ctx.Fail(errors.ErrInvalidDomain)
					return
				}
				ctx.Set("DB", db)
			}

			ctx.Set("AuthInfo", tokenInfo)
			ctx.Next()
		}
	}
}

// HttpTrace defined TODO
func HttpTrace() func(ctx *Context) {
	return func(ctx *Context) {
		req := ctx.Request()
		carrier, err := trace.Extract(trace.HttpFormat, req.Header)
		if err != nil && err != trace.ErrInvalidCarrier {
			logrus.Error(ctx, err)
		}
		c, span := trace.StartServerSpan(req.Context(), carrier, Hostname(), req.RequestURI)
		defer span.Finish()

		ctx.Set(string(tracespec.TracingKey), span)
		ctx.SetRequest(req.WithContext(c))
		ctx.Next()
	}
}

// Cors defined TODO
func Cors(origin string, headers string) func(ctx *Context) {
	return func(ctx *Context) {
		ctx.ResponseWriter().Header().Set("Access-Control-Allow-Origin", origin)
		ctx.ResponseWriter().Header().Set("Access-Control-Max-Age", "86400")
		ctx.ResponseWriter().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.ResponseWriter().Header().Set("Access-Control-Allow-Headers", headers)
		ctx.ResponseWriter().Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Disposition")
		ctx.ResponseWriter().Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request().Method == "OPTIONS" {
			ctx.Status(200)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}

// Recovery defined TODO
func Recovery() func(ctx *Context) {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(ctx.Request(), false)
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				logrus.Errorf(ctx, "[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
				ctx.Fail(goErr)
			}
		}()
		ctx.Next()
	}
}

// DumpBody instance a Logger middleware with config.
func DumpBody(dumpCall func(*Context, *LogFormatterParams)) func(ctx *Context) {
	formatter := Formatter
	notlogged := []string{}
	skip := map[string]struct{}{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *Context) {
		var bufWriter bytes.Buffer
		var bufReader bytes.Buffer
		var buf []byte
		if cb := c.Get(gin.BodyBytesKey); cb != nil {
			if cbb, ok := cb.([]byte); ok {
				buf = cbb
			}
		}
		if buf == nil {
			c.Request().Body = &NopCloser{c.Request().Body, &bufWriter}
		}

		c.Writer = &bodyLogWriter{body: &bufReader, ResponseWriter: c.Writer}

		// Start timer
		start := time.Now()
		path := c.Request().URL.Path
		raw := c.Request().URL.RawQuery

		// resove header
		hr := bytes.Buffer{}
		c.Request().Header.Write(&hr)

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
					Request: c.Request(),
					Keys:    c.Keys,
				},
				Header:  hr.Bytes(),
				ReqBody: buf,
				ResBody: bufReader.Bytes(),
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = c.ClientIP()
			param.Method = c.Request().Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			logrus.Info(c, formatter(param.LogFormatterParams))

			// LogFormatterParams defined recorder
			dumpCall(c, &param)
		}
	}
}
