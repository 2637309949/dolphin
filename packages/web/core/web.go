package core

import (
	"net/http"
)

type Web struct {
	RouterGroup
}

// New defined TODO
func New() *Web {
	web := Web{RouterGroup: RouterGroup{basePath: "/"}}
	web.RouterGroup.web = &web
	return &web
}

// resolveAddress defined TODO
func (w *Web) resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

func (w *Web) Run(addr ...string) error {
	address := w.resolveAddress(addr)
	srv := http.Server{Addr: address, Handler: GetHandler().Handler()}
	return srv.ListenAndServe()
}
