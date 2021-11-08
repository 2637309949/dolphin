package api

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics defined TODO
type Metrics struct {
}

// NewMetrics defined TODO
func NewMetrics() *Metrics {
	ctr := &Metrics{}
	return ctr
}

// MetricsRoutes defined TODO
func MetricsRoutes(dol *Dolphin) {
	g, i := dol.Group("/"), MetricsInstance
	g.Handle("GET", "/metrics", i.Metrics)
}

// Metrics defined TODO
func (ctr *Metrics) Metrics(ctx *Context) {
	promhttp.Handler().ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

// MetricsInstance defined TODO
var MetricsInstance = NewMetrics()
