// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: api.go

package api

import (
	"path"
	"strings"
	"sync"
	"time"

	"aisle/svc"

	"regexp"

	"github.com/2637309949/dolphin/platform/api"
	appSvc "github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/util"
)

var (
	App *Dolphin
	Run func()
)

type (
	// HandlerFunc defined
	HandlerFunc func(ctx *Context)
	// HandlersChain defined
	HandlersChain []HandlerFunc
	// Dolphin defined parse app engine
	Dolphin struct {
		RouterGroup
		*api.Dolphin
		pool sync.Pool
	}
	// Context defined http handle hook context
	Context struct {
		*api.Context
	}
	// RouterGroup defines struct that extend from gin.RouterGroup
	RouterGroup struct {
		dol      *Dolphin
		Handlers []HandlerFunc
		basePath string
	}
)

func (dol *Dolphin) allocateContext() *Context {
	return &Context{}
}

// HandlerFunc convert to api.HandlerFunc
func (dol *Dolphin) HandlerFunc(h HandlerFunc) (phf api.HandlerFunc) {
	return api.HandlerFunc(func(ctx *api.Context) {
		c := dol.pool.Get().(*Context)
		c.Context = ctx
		h(c)
		dol.pool.Put(c)
	})
}

// Handle defined TODO
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		re, err := regexp.Compile("^[A-Z]+$")
		if matches := re.MatchString(methods[i]); !matches || err != nil {
			panic("http method " + methods[i] + " is not valid")
		}
		relativePath := group.calculateAbsolutePath(relativePath)
		handlers = group.combineHandlers(handlers)
		hls := []api.HandlerFunc{}
		for j := 0; j < len(handlers); j++ {
			hls = append(hls, group.dol.HandlerFunc(handlers[j]))
		}
		group.handle(methods[i], relativePath, hls...)
	}
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers ...api.HandlerFunc) {
	group.dol.Http.Handle(httpMethod, relativePath, handlers...)
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(63) {
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	if relativePath == "" {
		return group.basePath
	}

	finalPath := path.Join(group.basePath, relativePath)
	if util.LastChar(relativePath) == '/' && util.LastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

// Use defined TODO
func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

// Group defined TODO
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		dol:      group.dol,
	}
}

// Auth middles
func Auth(auth ...string) HandlerFunc {
	return func(ctx *Context) {
		api.Auth(auth...)(ctx.Context)
	}
}

// Roles middles
func Roles(roles ...string) HandlerFunc {
	return func(ctx *Context) {
		api.Roles(roles...)(ctx.Context)
	}
}

// Cache middles
func Cache(time time.Duration) HandlerFunc {
	return func(ctx *Context) {
		api.Cache(time)(ctx.Context)
	}
}

// NewDolphin defined init dol you can custom engine
func NewDolphin() *Dolphin {
	rg := RouterGroup{Handlers: nil, basePath: "/"}
	dol := Dolphin{Dolphin: api.App, RouterGroup: rg}
	dol.pool.New = func() interface{} { return dol.allocateContext() }
	dol.RouterGroup.dol = &dol
	return &dol
}

func init() {
	dol, svcHelper := NewDolphin(), svc.NewSvcHepler(appSvc.NewSvcHepler(api.RedisClient))
	dol.SyncModel()
	dol.SyncController()
	dol.SyncService()
	dol.SyncSrv(svcHelper)
	App, Run = dol, dol.Run
}