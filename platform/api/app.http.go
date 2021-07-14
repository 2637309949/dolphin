// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// RootRelativePath defined TODO
const RootRelativePath = "/"

// HttpHandler defined TODO
type HttpHandler interface {
	http.Handler
	Handle(string, string, ...HandlerFunc)
	OnStart(context.Context) error
	OnStop(context.Context) error
}

type ginHandler struct {
	*gin.Engine
	handlers        *HandlersChain
	httpSrv         *http.Server
	allocateContext func(func(*Context))
}

// OnStart defined TODO
func (gh *ginHandler) OnStart(ctx context.Context) error {
	go func() {
		// register gin global handlers
		hf := funk.Map(*gh.handlers, func(hf HandlerFunc) gin.HandlerFunc { return gh.handlerFunc(hf) }).([]gin.HandlerFunc)
		gh.NoRoute(hf...)
		gh.NoMethod(hf...)
		gh.httpSrv.Handler = gh
		logrus.Infof("http listen on port:%v", viper.GetString("http.port"))
		if err := gh.httpSrv.ListenAndServe(); err != nil {
			logrus.Fatal(err)
		}
	}()
	return nil
}

// OnStop defined TODO
func (gh *ginHandler) OnStop(ctx context.Context) error {
	if err := gh.httpSrv.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
		return err
	}
	return nil
}

func (gh *ginHandler) handlerFunc(h HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		gh.allocateContext(func(c *Context) {
			c.Context = ctx
			for k := range ctx.Keys {
				switch t := ctx.Keys[k].(type) {
				case *xorm.Engine:
					c.DB = t
				case AuthInfo:
					c.AuthInfo = t
				}
			}
			h(c)
		})
	})
}

// Handle defined TODO
func (gh *ginHandler) Handle(httpMethod, absolutePath string, handlerFuncs ...HandlerFunc) {
	hls := funk.Map(handlerFuncs, func(hf HandlerFunc) gin.HandlerFunc { return gh.handlerFunc(hf) }).([]gin.HandlerFunc)
	gh.Engine.Handle(httpMethod, absolutePath, hls...)
}

// DebugPrintRouteFunc defined TODO
func DebugPrintRouteFunc(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	if gin.IsDebugging() {
		logrus.Infof("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, strings.Split(handlerName, ".(")[0], nuHandlers)
	}
}

// NewGinHandler defined TODO
func NewGinHandler(dol *Dolphin) HttpHandler {
	gin.SetMode(viper.GetString("app.mode"))
	gin.DefaultWriter = logrus.StandardLogger().Out
	gin.DebugPrintRouteFunc = DebugPrintRouteFunc
	h := &ginHandler{Engine: gin.New()}
	h.handlers, h.allocateContext = &dol.Handlers, dol.allocateContext
	h.httpSrv = &http.Server{Addr: fmt.Sprintf(":%v", viper.GetString("http.port"))}
	return h
}
