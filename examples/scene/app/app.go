// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: app.go

package app

import (
	"path"
	"strings"
	"sync"
	"time"

	"github.com/2637309949/dolphin/platform/app"
)

type (
	// HandlerFunc defined
	HandlerFunc func(ctx *Context)
	// Engine defined parse app engine
	Engine struct {
		*app.Engine
		pool sync.Pool
	}
	// Context defined http handle hook context
	Context struct {
		*app.Context
		engine *Engine
	}
	// RouterGroup defines struct that extend from gin.RouterGroup
	RouterGroup struct {
		Routes   []Route
		basePath string
		engine   *Engine
	}
	// Route defines the handler used by gin middleware as return value
	Route struct {
		Method, RelativePath string
		Interceptor          []Route
		Handler              HandlerFunc
	}
)

// HF2Handler defined
func HF2Handler(h func(ctx *Context)) Route {
	return Route{Handler: h}
}

func (e *Engine) allocateContext() *Context {
	return &Context{engine: e}
}

// Group handlers
func (e *Engine) Group(relativePath string, routes ...Route) *RouterGroup {
	return &RouterGroup{Routes: routes, basePath: relativePath, engine: e}
}

// HandlerFunc convert to app.HandlerFunc
func (e *Engine) HandlerFunc(h HandlerFunc) (phf app.HandlerFunc) {
	return app.HandlerFunc(func(ctx *app.Context) {
		c := e.pool.Get().(*Context)
		c.Context = ctx
		h(c)
		e.pool.Put(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, routes ...Route) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		method := methods[i]
		absPath := path.Join(rg.basePath, relativePath)
		rs := append(rg.Routes, routes...)
		hls := []app.HandlerFunc{}
		for j := 0; j < len(rs); j++ {
			route := rs[j]
			for k := 0; k < len(route.Interceptor); k++ {
				ir := route.Interceptor[k]
				hls = append(hls, rg.engine.HandlerFunc(ir.Handler))
			}
			hls = append(hls, rg.engine.HandlerFunc(route.Handler))
		}
		rg.engine.Http.Handle(method, absPath, hls...)
	}
}

// Auth middles
func Auth(auth ...string) Route {
	return HF2Handler(func(ctx *Context) {
		app.Auth(auth...).Handler(ctx.Context)
	})
}

// Roles middles
func Roles(roles ...string) Route {
	return HF2Handler(func(ctx *Context) {
		app.Roles(roles...).Handler(ctx.Context)
	})
}

// Cache middles
func Cache(time time.Duration) Route {
	return HF2Handler(func(ctx *Context) {
		app.Cache(time).Handler(ctx.Context)
	})
}

// NewEngine defined init engine you can custom engine
func NewEngine() *Engine {
	e := &Engine{Engine: app.App}
	e.pool.New = func() interface{} { return e.allocateContext() }
	return e
}

var (
	// App instance
	App = NewEngine()
	// Run defined
	Run = App.Run
)

func init() {
	SyncModel(App)
	SyncController(App)
	SyncService(App)
}
