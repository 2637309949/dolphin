package app

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/util/file"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type HttpHandler interface {
	Handle(httpMethod, relativePath string, handlers ...HandlerFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	SetAllocateContextFunk(func(func(*Context)))
	OnStart(context.Context) error
	OnStop(context.Context) error
}

type ginHandler struct {
	gin      *gin.Engine
	httpSrv  *http.Server
	allocCtx func(func(*Context))
}

func (gh *ginHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gh.gin.ServeHTTP(w, req)
}

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
			h.Handler(c)
		})
	})
}

func (gh *ginHandler) SetAllocateContextFunk(alloc func(func(*Context))) {
	gh.allocCtx = alloc
}

func (gh *ginHandler) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		method := methods[i]
		group := gh.gin.Group("/")
		hls := []gin.HandlerFunc{}
		for j := 0; j < len(handlers); j++ {
			hl := handlers[j]
			for k := 0; k < len(hl.Interceptor); k++ {
				ir := hl.Interceptor[k]
				hls = append(hls, gh.handlerFunc(ir))
			}
			hls = append(hls, gh.handlerFunc(hl))
		}
		group.Handle(method, relativePath, hls...)
	}
}

// NewGinHandler defined
func NewGinHandler(e *Engine) HttpHandler {
	gin.DefaultWriter = logrus.StandardLogger().Out
	gin.SetMode(viper.GetString("app.mode"))

	gn := gin.New()
	gn.Use(plugin.CORS())
	gn.Static(viper.GetString("http.static"), path.Join(file.Getwd(), viper.GetString("http.static")))
	gn.Use(plugin.Recovery())
	gn.Use(plugin.Override(gn.HandleContext))
	gn.Use(plugin.Tracker(Tracker(e)))

	h := &ginHandler{gin: gn}
	h.httpSrv = &http.Server{Addr: fmt.Sprintf(":%v", viper.GetString("http.port"))}
	h.SetAllocateContextFunk(func(f func(*Context)) {
		c := e.pool.Get().(*Context)
		c.reset()
		f(c)
		e.pool.Put(c)
	})
	return h
}
