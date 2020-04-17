// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: app.go

package app

import (
	"sync"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/viper"
	pApp "github.com/2637309949/dolphin/platform/app"
)

type (
	// Engine defined parse app engine
	Engine struct {
		*pApp.Engine
		Handler Handler
		pool    sync.Pool
	}
	// Context defined http handle hook context
	Context struct {
		*pApp.Context
		engine *Engine
	}
	// RouterGroup defines struct that extend from gin.RouterGroup
	RouterGroup struct {
		*pApp.RouterGroup
		engine *Engine
	}
	// HandlerFunc defines the handler used by gin middleware as return value.
	HandlerFunc func(*Context)
	// Handler defined hooks
	Handler interface {
		OnHandler(name string, h func(ctx *Context)) func(*Context)
	}
	// IHander defined handler
	IHander struct{}
)

// OnHandler defined event
func (i *IHander) OnHandler(name string, h func(ctx *Context)) func(*Context) {
	return h
}

func (e *Engine) allocateContext() *Context {
	return &Context{engine: e}
}

// Auth middles
func (e *Engine) Auth() func(*Context) {
	return func(ctx *Context) {
		e.Engine.Auth()(ctx.Context)
	}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{engine: e, RouterGroup: e.Engine.Group(relativePath, handlers...)}
}

// HandlerFunc convert to pApp.HandlerFunc
func (e *Engine) HandlerFunc(h HandlerFunc) (phf pApp.HandlerFunc) {
	return pApp.HandlerFunc(func(ctx *pApp.Context) {
		c := e.pool.Get().(*Context)
		c.Context = ctx
		h(c)
		e.pool.Put(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) pApp.HandlerFunc {
		return rg.engine.HandlerFunc(h)
	}).([]pApp.HandlerFunc)...)
}

// InvokeEngine build engine
func InvokeEngine(build func(*Engine)) func(*pApp.Engine) {
	return func(*pApp.Engine) {
		build(App)
	}
}

// InvokeContext build context
func InvokeContext(httpMethod string, relativePath string, handlers ...HandlerFunc) func(*pApp.Engine) {
	return func(base *pApp.Engine) {
		group := App.Group(viper.GetString("http.prefix"))
		group.Handle(httpMethod, relativePath, handlers...)
	}
}

func buildEngine() *Engine {
	e := &Engine{Engine: pApp.App, Handler: &IHander{}}
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	return e
}

// App instance
var App = buildEngine()
