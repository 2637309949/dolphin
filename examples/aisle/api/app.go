// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: api.go

package api

import (
	"fmt"
	"path"
	"strings"
	"sync"
	"time"

	"aisle/svc"

	"regexp"

	"github.com/2637309949/dolphin/platform/api"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/thoas/go-funk"
)

// vars defined
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
		*api.Dolphin
		RouterGroup
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

// Reflesh defined TODO
func (dol *Dolphin) Reflesh() error {
	dol.SyncModel()
	dol.SyncSrv(svc.NewServiceContext(api.CacheStore))
	dol.SyncController()
	return nil
}

// Run defined TODO
func (dol *Dolphin) Run() {
	if err := dol.Reflesh(); err != nil {
		panic(fmt.Errorf("%v\n%v", err, string(errors.Wrap(err, 2).Stack())))
	}
	if err := dol.Dolphin.Run(); err != nil {
		panic(fmt.Errorf("%v\n%v", err, string(errors.Wrap(err, 2).Stack())))
	}
}

// Use defined TODO
func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
	hls := funk.Map(middleware, group.unWrapContext).([]api.HandlerFunc)
	group.dol.Dolphin.Handlers = append(group.dol.Dolphin.Handlers, hls...)
}

// Group defined TODO
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	handlers = append(funk.Map(group.dol.Dolphin.Handlers, group.wrapContext).([]HandlerFunc), handlers...)
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		dol:      group.dol,
	}
}

// Handle defined TODO
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		re, err := regexp.Compile("^[A-Z]+$")
		if matches := re.MatchString(methods[i]); !matches || err != nil {
			panic("http method " + methods[i] + " is not valid")
		}
		relativePath := group.calculateAbsolutePath(relativePath)
		combineHandlers := group.combineHandlers(handlers)
		hls := funk.Map(combineHandlers, func(h HandlerFunc) api.HandlerFunc { return group.unWrapContext(h) }).([]api.HandlerFunc)
		group.dol.Http.Handle(methods[i], relativePath, hls...)
	}
}

// unWrapContext defined TODO
func (group *RouterGroup) unWrapContext(h HandlerFunc) api.HandlerFunc {
	return func(ctx *api.Context) {
		c := group.dol.pool.Get().(*Context)
		c.Context = ctx
		h(c)
		group.dol.pool.Put(c)
	}
}

// wrapContext defined TODO
func (group *RouterGroup) wrapContext(h api.HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		h(ctx.Context)
	}
}

// combineHandlers defined TODO
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

// calculateAbsolutePath defined TODO
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

// NewContext defined TODO
func NewContext(dol *Dolphin) *Context {
	return &Context{}
}

// New defined init dol you can custom engine
func New() *Dolphin {
	rg := RouterGroup{Handlers: nil, basePath: "/"}
	dol := Dolphin{Dolphin: api.App, RouterGroup: rg}
	dol.pool.New = func() interface{} { return NewContext(&dol) }
	dol.RouterGroup.dol = &dol
	return &dol
}

func init() {
	app := New()
	App, Run = app, app.Run
}
