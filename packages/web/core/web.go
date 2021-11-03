package core

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

// Done returns a channel of signals to block on after starting the
func (w *Web) done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

func (w *Web) Run(addr ...string) error {
	address := w.resolveAddress(addr)
	srv := http.Server{Addr: address, Handler: GetHandler().Handler()}
	err := srv.ListenAndServe()
	if err != nil {
		log.Print(err)
		return err
	}
	signal := w.done()
	<-signal
	return nil
}
