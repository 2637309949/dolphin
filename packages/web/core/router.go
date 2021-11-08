package core

import (
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	HandlerFunc   interface{}
	HandlersChain []HandlerFunc
	onlyFilesFS   struct {
		fs http.FileSystem
	}
	neuteredReaddirFile struct {
		http.File
	}
)

// nameOfFunction defined TODO
func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// DebugPrintRoute defined TODO
func DebugPrintRoute(httpMethod, absolutePath, _ string, _ int) {
	logrus.Infof("%-6s %-25s ", httpMethod, absolutePath)
}

// routerGroup defined TODO
type RouterGroup struct {
	Handlers []HandlerFunc
	basePath string
	web      *Web
}

// Dir defined TODO
func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyFilesFS{fs}
}

// Open defined TODO
func (fs onlyFilesFS) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

// Readdir defined TODO
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

// Handle defined TODO
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlerFuncs ...HandlerFunc) {
	methods := strings.Split(httpMethod, ",")
	for i := 0; i < len(methods); i++ {
		hls := group.combineHandlers(handlerFuncs)
		relativePath = group.calculateAbsolutePath(relativePath)

		nuHandlers := len(handlerFuncs)
		handlerName := nameOfFunction(handlerFuncs[nuHandlers-1])
		DebugPrintRoute(httpMethod, relativePath, handlerName, nuHandlers)

		// handler request
		handler := MustHandler()
		handler.Handle(methods[i], relativePath, hls...)
	}
}

// Static defined TODO
func (group *RouterGroup) Static(relativePath, root string) {
	group.StaticFS(relativePath, Dir(root, false))
}

// Use defined TODO
func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

// StaticFS defined TODO
func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := group.createStaticHandler(relativePath, fs)
	urlPattern := path.Join(relativePath, "/*filepath")
	group.Handle("GET", urlPattern, handler)
	group.Handle("HEAD", urlPattern, handler)
}

// createStaticHandler defined TODO
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := group.calculateAbsolutePath(relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c Context) {
		if _, noListing := fs.(*onlyFilesFS); noListing {
			c.ResponseWriter().WriteHeader(http.StatusNotFound)
		}
		file := c.Param("filepath")
		f, err := fs.Open(file)
		if err != nil {
			c.ResponseWriter().WriteHeader(http.StatusNotFound)
			return
		}
		f.Close()
		fileServer.ServeHTTP(c.ResponseWriter(), c.Request())
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
	if LastChar(relativePath) == '/' && LastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

// Group defined TODO
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
	}
}
