package xtest

import (
	"encoding/json"
	"net/http/httptest"
)

// Context defined TODO
type Context struct {
	testing
	*httptest.ResponseRecorder
}

// Bind defined TODO
func (ctx *Context) Bind(v interface{}) error {
	return json.Unmarshal(ctx.ResponseRecorder.Body.Bytes(), &v)
}

// Body defined TODO
func (ctx *Context) Body() (*struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Detail string      `json:"detail"`
	Status string      `json:"status"`
}, error) {
	ret := struct {
		Code   int         `json:"code"`
		Data   interface{} `json:"data"`
		Detail string      `json:"detail"`
		Status string      `json:"status"`
	}{}
	err := json.Unmarshal(ctx.ResponseRecorder.Body.Bytes(), &ret)
	return &ret, err
}

// NewContext defined TODO
func NewContext(w *httptest.ResponseRecorder, t testing) *Context {
	return &Context{t, w}
}
