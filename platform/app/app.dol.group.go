// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/2637309949/dolphin/platform/util"
)

type (
	// HandlerFunc defined TODO
	HandlerFunc func(ctx *Context)
	// HandlersChain defined TODO
	HandlersChain []HandlerFunc
	// RouterGroup defined TODO
	RouterGroup struct {
		dol      *Dolphin
		Handlers HandlersChain
		basePath string
	}
	onlyFilesFS struct {
		fs http.FileSystem
	}
	neuteredReaddirFile struct {
		http.File
	}
)

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
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		re, err := regexp.Compile("^[A-Z]+$")
		if matches := re.MatchString(methods[i]); !matches || err != nil {
			panic("http method " + methods[i] + " is not valid")
		}
		handlerFuncs = group.combineHandlers(handlerFuncs)
		relativePath := group.calculateAbsolutePath(relativePath)
		group.handle(methods[i], relativePath, handlerFuncs...)
	}
}

// Static defined TODO
func (group *RouterGroup) Static(relativePath, root string) {
	group.StaticFS(relativePath, Dir(root, false))
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

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := group.calculateAbsolutePath(relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		if _, noListing := fs.(*onlyFilesFS); noListing {
			c.Writer.WriteHeader(http.StatusNotFound)
		}
		file := c.Param("filepath")
		f, err := fs.Open(file)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		f.Close()
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
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
		dol:      group.dol,
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
	}
}
