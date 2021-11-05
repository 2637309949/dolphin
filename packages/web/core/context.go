package core

import (
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

// Context defined TODO
type Context interface {
	// Set/Get is used to store a new key/value pair exclusively for this context.
	Set(string, interface{})
	Get(string) interface{}

	// Next
	Next()
	Abort()

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
	String(string, ...interface{})

	// net/http
	Request() *http.Request
	SetRequest(r *http.Request)

	ResponseWriter() http.ResponseWriter

	// context
	Deadline() (time.Time, bool)
	Done() <-chan struct{}
	Err() error
	Value(interface{}) interface{}

	// std reponse json
	JSON(code int, i interface{})
	Success(interface{})
	Fail(error)
}
