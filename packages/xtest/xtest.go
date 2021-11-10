package xtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/spf13/viper"
)

type XTest struct {
	hdl   http.Handler
	token string
}

// SetToken defined TODO
func (x *XTest) SetToken(tk string) {
	x.token = tk
}

// serveHttp defined TODO
func (x *XTest) serveHttp(w http.ResponseWriter, req *http.Request) {
	x.hdl.ServeHTTP(w, req)
}

// handler defined TODO
func (x *XTest) handler(req *http.Request, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+x.token)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("origin", "http://localhost"+viper.GetString("http.port"))
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	x.serveHttp(w, req)
	h(w)
}

// Get defined TODO
func (x *XTest) Get(url string, args interface{}, h func(w *httptest.ResponseRecorder)) {
	var payload = map[string]interface{}{}
	br, _ := json.Marshal(args)
	json.Unmarshal(br, &payload)
	url += "?"
	for k, v := range payload {
		url += fmt.Sprintf("%v=%v&", k, v)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	x.handler(req, h)
}

// Options defined TODO
func (x *XTest) Options(url string, args interface{}, h func(w *httptest.ResponseRecorder)) {
	var payload = map[string]interface{}{}
	br, _ := json.Marshal(args)
	json.Unmarshal(br, &payload)
	url += "?"
	for k, v := range payload {
		url += fmt.Sprintf("%v=%v&", k, v)
	}
	req, err := http.NewRequest("OPTIONS", url, nil)
	if err != nil {
		panic(err)
	}
	x.handler(req, h)
}

// Post defined TODO
func (x *XTest) Post(url string, args interface{}, h func(w *httptest.ResponseRecorder)) {
	jm, err := json.Marshal(&args)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(jm))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	x.handler(req, h)
}

// Put defined TODO
func (x *XTest) Put(url string, args interface{}, h func(w *httptest.ResponseRecorder)) {
	jm, err := json.Marshal(&args)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(jm))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	x.handler(req, h)
}

// Head defined TODO
func (x *XTest) Head(url string, args interface{}, h func(w *httptest.ResponseRecorder)) {
	var payload = map[string]interface{}{}
	br, _ := json.Marshal(args)
	json.Unmarshal(br, &payload)
	url += "?"
	for k, v := range payload {
		url += fmt.Sprintf("%v=%v&", k, v)
	}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		panic(err)
	}
	x.handler(req, h)
}

// HttpHandle defined TODO
func (x *XTest) Handle(method, reqPath string, funk func(ctx *Context), t testing, args interface{}) {
	switch method {
	case "GET":
		x.Get(reqPath, args, func(w *httptest.ResponseRecorder) { funk(NewContext(w, t)) })
	case "HEAD":
		x.Head(reqPath, args, func(w *httptest.ResponseRecorder) { funk(NewContext(w, t)) })
	case "OPTIONS":
		x.Options(reqPath, args, func(w *httptest.ResponseRecorder) { funk(NewContext(w, t)) })
	case "POST":
		x.Post(reqPath, args, func(w *httptest.ResponseRecorder) { funk(NewContext(w, t)) })
	case "PUT":
		x.Put(reqPath, args, func(w *httptest.ResponseRecorder) { funk(NewContext(w, t)) })
	}
}

// New defined TODO
func New(hdl http.Handler) *XTest {
	return &XTest{hdl: hdl}
}
