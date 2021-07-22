// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// RootRelativePath defined TODO
const RootRelativePath = "/"

type (
	// HttpHandler defined TODO
	HttpHandler interface {
		lifeHook
		ServeHTTP(http.ResponseWriter, *http.Request)
		Handle(string, string, ...HandlerFunc)
	}
	// Restful defined TODO
	Restful struct {
		*gin.Engine
		dol *Dolphin
	}
)

// OnStart defined TODO
func (gh *Restful) OnStart(ctx context.Context) error {
	go func() {
		gh.rebuildGlobalRoutes()
		logrus.Infof("http listen on port:%v", viper.GetString("http.port"))
		if err := gh.Engine.Run(fmt.Sprintf(":%v", viper.GetString("http.port"))); err != nil {
			logrus.Fatal(err)
		}
	}()
	return nil
}

// rebuildGlobalRoutes gin global handlers
func (gh *Restful) rebuildGlobalRoutes() *Restful {
	hf := funk.Map(gh.dol.Handlers, func(hf HandlerFunc) gin.HandlerFunc { return gh.unWrapHandler(hf) }).([]gin.HandlerFunc)
	gh.NoRoute(hf...)
	gh.NoMethod(hf...)
	return gh
}

// OnStop defined TODO
func (gh *Restful) OnStop(ctx context.Context) error {
	return nil
}

// unWrapHandler defined TODO
func (gh *Restful) unWrapHandler(h HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := gh.dol.pool.Get().(*Context)
		c.reset()
		c.Context = ctx
		// uwrap Context
		for k := range ctx.Keys {
			switch t := ctx.Keys[k].(type) {
			case *xorm.Engine:
				c.DB = t
			case AuthInfo:
				c.AuthInfo = t
			}
		}
		h(c)
		gh.dol.pool.Put(c)
	}
}

// Handle defined TODO
func (gh *Restful) Handle(httpMethod, absolutePath string, handlerFuncs ...HandlerFunc) {
	hls := funk.Map(handlerFuncs, func(hf HandlerFunc) gin.HandlerFunc { return gh.unWrapHandler(hf) }).([]gin.HandlerFunc)
	nuHandlers := len(handlerFuncs)
	handlerName := nameOfFunction(handlerFuncs[nuHandlers-1])
	DebugPrintRoute(httpMethod, absolutePath, handlerName, nuHandlers)
	gh.Engine.Handle(httpMethod, absolutePath, hls...)
}

// DebugPrintRoute defined TODO
func DebugPrintRoute(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	logrus.Infof("%-6s %-25s ", httpMethod, absolutePath)
}

// nameOfFunction defined TODO
func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// NewGinHandler defined TODO
func NewGinHandler(dol *Dolphin) HttpHandler {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = logrus.StandardLogger().Out
	h := &Restful{Engine: gin.New()}
	h.dol = dol
	return h
}
