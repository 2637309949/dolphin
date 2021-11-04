package core

import (
	"errors"
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
	if defaultHandler != nil {
		panic(errors.New("handler has been registered"))
	}
	defaultHandler = e
}
