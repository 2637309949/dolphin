package core

import (
	"errors"
	"net/http"
)

var (
	defaultHandler Handler
)

// Handler defined TODO
type Handler interface {
	Handler() http.Handler
	Handle(string, string, ...HandlerFunc)
}

// GetHandler defined TODO
func GetHandler() Handler {
	return defaultHandler
}

// MustHandler defined TODO
func MustHandler() Handler {
	if defaultHandler == nil {
		panic(errors.New("handler has been registered"))
	}
	return defaultHandler
}

// SetHandler defined TODO
func SetHandler(e Handler) {
	if defaultHandler != nil {
		panic(errors.New("handler has been registered"))
	}
	defaultHandler = e
}
