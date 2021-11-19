package web

import (
	"net/http"

	"github.com/2637309949/dolphin/packages/web/core"
)

var defaultApp *core.Web

func init() {
	defaultApp = core.New()
}

// Web defined TODO
func Web() *core.Web {
	return defaultApp
}

// Group defined TODO
func Group(relativePath string, handlers ...core.HandlerFunc) *core.RouterGroup {
	return defaultApp.Group(relativePath, handlers...)
}

// ServeHTTP defined TODO
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defaultApp.ServeHTTP(w, r)
}

// Handle defined TODO
func Handle(m string, r string, h ...core.HandlerFunc) {
	defaultApp.Handle(m, r, h...)
}

// Static defined TODO
func Static(relativePath, root string) {
	defaultApp.Static(relativePath, root)
}

// StaticFS defined TODO
func StaticFS(relativePath string, fs http.FileSystem) {
	defaultApp.StaticFS(relativePath, fs)
}

// Use defined TODO
func Use(middleware ...core.HandlerFunc) {
	defaultApp.Use(middleware...)
}

// Run defined TODO
func Run(addr ...string) error {
	return defaultApp.Run(addr...)
}
