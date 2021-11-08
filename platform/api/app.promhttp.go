package api

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PromHttp defined
type PromHttp struct {
}

// NewPromHttp defined
func NewPromHttp() *PromHttp {
	ctr := &PromHttp{}
	return ctr
}

// PromHttpRoutes defined TODO
func PromHttpRoutes(dol *Dolphin) {
	g, i := dol.Group("/"), PromHttpInstance
	g.Handle("GET", "/metrics", i.Metrics)
}

// Domain defined TODO
func (ctr *PromHttp) Metrics(ctx *Context) {
	promhttp.Handler().ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

// PromHttpInstance defined TODO
var PromHttpInstance = NewPromHttp()
