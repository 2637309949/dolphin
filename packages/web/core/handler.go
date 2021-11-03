package core

import (
	"net/http"
)

// Handler defined TODO
type Handler interface {
	Handler() http.Handler
	Handle(string, string, ...HandlerFunc)
}

var defaultHandler Handler

func GetHandler() Handler {
	return defaultHandler
}

func SetHandler(e Handler) {
	defaultHandler = e
}
