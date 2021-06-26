package svc

import (
	"github.com/2637309949/dolphin/platform/util/http"
)

type Http interface {
	// Get returns *HttpRequest with GET method.
	Get(url string) *http.HttpRequest
	// Post returns *HttpRequest with POST method.
	Post(url string) *http.HttpRequest
	// Put returns *HttpRequest with PUT method.
	Put(url string) *http.HttpRequest
	// Delete returns *HttpRequest DELETE method.
	Delete(url string) *http.HttpRequest
	// Head returns *HttpRequest with HEAD method.
	Head(url string) *http.HttpRequest
}

// Get returns *HttpRequest with GET method.
func (svc *SvcHepler) Get(url string) *http.HttpRequest {
	return http.NewRequest(url, "GET")
}

// Post returns *HttpRequest with POST method.
func (svc *SvcHepler) Post(url string) *http.HttpRequest {
	return http.NewRequest(url, "POST")
}

// Put returns *HttpRequest with PUT method.
func (svc *SvcHepler) Put(url string) *http.HttpRequest {
	return http.NewRequest(url, "PUT")
}

// Delete returns *HttpRequest DELETE method.
func (svc *SvcHepler) Delete(url string) *http.HttpRequest {
	return http.NewRequest(url, "DELETE")
}

// Head returns *HttpRequest with HEAD method.
func (svc *SvcHepler) Head(url string) *http.HttpRequest {
	return http.NewRequest(url, "HEAD")
}
