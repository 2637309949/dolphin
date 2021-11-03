package core

import (
	"path"
)

type (
	HandlerFunc   interface{}
	HandlersChain []HandlerFunc
)

// RouterGroup defined TODO
type RouterGroup struct {
	web      *Web
	Handlers []HandlerFunc
	basePath string
}

// Handle defined TODO
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlerFuncs ...HandlerFunc) {
	hls := group.combineHandlers(handlerFuncs)
	relativePath = group.calculateAbsolutePath(relativePath)
	GetHandler().Handle(httpMethod, relativePath, hls...)
}

// Use defined TODO
func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
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
