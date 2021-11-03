package core

import (
	"io"
	"mime/multipart"
)

// Context defined TODO
type Context interface {
	// Set/Get is used to store a new key/value pair exclusively for this context.
	Set(string, interface{})
	Get(string) (interface{}, bool)

	// /user/:id
	Param(string) string

	// /path?id=1234&name=Manu&value=
	Query(key string) string

	// parsed multipart form, including file uploads
	MultipartForm() (*multipart.Form, error)

	// http header
	Header(key, value string)
	GetHeader(key string) string

	// return stream data.
	ShouldBindWith(interface{}) error
	GetRawData() ([]byte, error)

	// Set-Cookie header to the ResponseWriter's headers
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	Cookie(name string) (string, error)

	// render reponse
	HTML(io.Reader, ...interface{})
	XML(io.Reader, ...interface{})
	File(io.Reader, string, ...interface{})

	// next
	Next()

	// std reponse json
	Success(interface{})
	Fail(error)
}

var defaultContext Context

func GetContext() Context {
	return defaultContext
}
