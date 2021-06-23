// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: app.go

package app

import (
	"path"
	"strings"
	"sync"
	"time"

	"regexp"

	"github.com/2637309949/dolphin/platform/app"
	"github.com/2637309949/dolphin/platform/util"
)

type (
	// HandlerFunc defined
	HandlerFunc func(ctx *Context)
	// HandlersChain defined
	HandlersChain []HandlerFunc
	// Dolphin defined parse app engine
	Dolphin struct {
		RouterGroup
		*app.Dolphin
		pool sync.Pool
	}
	// Context defined http handle hook context
	Context struct {
		*app.Context
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

// HandlerFunc convert to app.HandlerFunc
func (dol *Dolphin) HandlerFunc(h HandlerFunc) (phf app.HandlerFunc) {
	return app.HandlerFunc(func(ctx *app.Context) {
		c := dol.pool.Get().(*Context)
		c.Context = ctx
		h(c)
		dol.pool.Put(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		group.handle(methods[i], relativePath, handlers...)
	}
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	re, err := regexp.Compile("^[A-Z]+$")
	if matches := re.MatchString(httpMethod); !matches || err != nil {
		panic("http method " + httpMethod + " is not valid")
	}
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	hls := []app.HandlerFunc{}
	for j := 0; j < len(handlers); j++ {
		hls = append(hls, group.dol.HandlerFunc(handlers[j]))
	}
	group.dol.Http.Handle(httpMethod, absolutePath, hls...)
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
		app.Auth(auth...)(ctx.Context)
	}
}

// Roles middles
func Roles(roles ...string) HandlerFunc {
	return func(ctx *Context) {
		app.Roles(roles...)(ctx.Context)
	}
}

// Cache middles
func Cache(time time.Duration) HandlerFunc {
	return func(ctx *Context) {
		app.Cache(time)(ctx.Context)
	}
}

// NewDolphin defined init dol you can custom engine
func NewDolphin() *Dolphin {
	dol := &Dolphin{
		Dolphin: app.App,
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
		},
	}
	dol.pool.New = func() interface{} { return dol.allocateContext() }
	dol.RouterGroup.dol = dol
	return dol
}

var (
	// App instance
	App = NewDolphin()
	// Run defined
	Run = App.Run
)

func init() {
	SyncModel(App)
	SyncController(App)
	SyncService(App)
}
