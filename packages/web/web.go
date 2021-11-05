package web

import (
	"net/http"

	"github.com/2637309949/dolphin/packages/web/core"
	_ "github.com/2637309949/dolphin/packages/web/gin"
	// _ "github.com/2637309949/dolphin/packages/web/echo"
)

var defaultApp *core.Web

func init() {
	app := core.New()
	defaultApp = app
}

func Web() *core.Web {
	return defaultApp
}

func Group(relativePath string, handlers ...core.HandlerFunc) *core.RouterGroup {
	return defaultApp.Group(relativePath, handlers...)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defaultApp.ServeHTTP(w, r)
}

func Handle(m string, r string, h ...core.HandlerFunc) {
	defaultApp.Handle(m, r, h...)
}

func Static(relativePath, root string) {
	defaultApp.Static(relativePath, root)
}

func Use(middleware ...core.HandlerFunc) {
	defaultApp.Use(middleware...)
}

func Run(addr ...string) error {
	return defaultApp.Run(addr...)
}
