// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"fmt"
	"net/http"

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
	gin      *gin.Engine
	httpSrv  *http.Server
	allocCtx func(func(*Context))
}

// ServeHTTP defined TODO
func (gh *ginHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gh.gin.ServeHTTP(w, req)
}

// OnStart defined TODO
func (gh *ginHandler) OnStart(ctx context.Context) error {
	logrus.Infof("http listen on port:%v", viper.GetString("http.port"))
	gh.httpSrv.Handler = gh.gin
	go func() {
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
		gh.allocCtx(func(c *Context) {
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

func (gh *ginHandler) Handle(httpMethod, relativePath string, handlerFuncs ...HandlerFunc) {
	hls := funk.Map(handlerFuncs, func(hf HandlerFunc) gin.HandlerFunc { return gh.handlerFunc(hf) }).([]gin.HandlerFunc)
	gh.gin.Handle(httpMethod, relativePath, hls...)
}

// NewGinHandler defined TODO
func NewGinHandler(dol *Dolphin) HttpHandler {
	gin.DefaultWriter = logrus.StandardLogger().Out
	gin.SetMode(viper.GetString("app.mode"))
	h := &ginHandler{gin: gin.New()}
	h.httpSrv = &http.Server{Addr: fmt.Sprintf(":%v", viper.GetString("http.port"))}
	h.allocCtx = func(f func(*Context)) {
		c := dol.pool.Get().(*Context)
		c.reset()
		f(c)
		dol.pool.Put(c)
	}
	return h
}
