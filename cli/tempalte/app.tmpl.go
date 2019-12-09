// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tempalte

// TmplGin defined template
var TmplGin = `// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: app.go

package app

import (
	"sync"

	pApp "github.com/2637309949/dolphin/cli/platform/app"

	"github.com/2637309949/dolphin/cli/gin"
	"github.com/thoas/go-funk"
)

type (
	// Engine defined parse app engine
	Engine struct {
		*pApp.Engine
		pool sync.Pool
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
)

func (e *Engine) allocateContext() *Context {
	return &Context{engine: e}
}

// BuildEngine build engine
func BuildEngine(build func(*Engine)) func(*pApp.Engine) {
	return func(base *pApp.Engine) {
		engine.Engine = base
		build(engine)
	}
}

// Auth middles
func (e *Engine) Auth(mode ...pApp.AuthType) func(*Context) {
	return func(ctx *Context) {
		e.Engine.Auth(mode...)(ctx.Context)
	}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{engine: e, RouterGroup: e.Engine.Group(relativePath, handlers...)}
}

// HandlerFunc convert to pApp.HandlerFunc
func (hf HandlerFunc) HandlerFunc(e *Engine) (phf pApp.HandlerFunc) {
	phf = pApp.HandlerFunc(func(ctx *pApp.Context) {
		c := e.pool.Get().(*Context)
		c.Context = ctx
		hf(c)
		e.pool.Put(c)
	})
	phf.HandlerFunc(e.Engine)
	return
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) pApp.HandlerFunc {
		return h.HandlerFunc(rg.engine)
	}).([]pApp.HandlerFunc)...)
}

// NewEngine init Engine
func NewEngine() *Engine {
	e := &Engine{}
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	return e
}

// Engine instance
var engine = NewEngine()
`
