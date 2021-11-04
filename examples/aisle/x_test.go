// Code generated by dol build. DO NOT EDIT.
// source: x_test.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/2637309949/dolphin/platform/util"
	"github.com/spf13/viper"

	"aisle/api"
)

var token string
var x *XTest

type (
	testingT interface {
		// Deadline() (time.Time, bool)
		Error(args ...interface{})
		Errorf(format string, args ...interface{})
		Fail()
		FailNow()
		Failed() bool
		Fatal(args ...interface{})
		Fatalf(format string, args ...interface{})
		Helper()
		Log(args ...interface{})
		Logf(format string, args ...interface{})
		Name() string
		// Parallel()
		Skip(args ...interface{})
		SkipNow()
		Skipf(format string, args ...interface{})
		Skipped() bool
	}
	XTest struct {
	}
	Response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
	Context struct {
		*httptest.ResponseRecorder
		testingT testingT
	}
	Payload map[string]interface{}
)

// ParseBody defined TODO
func (ctx *Context) ParseBody(arg interface{}) error {
	return json.Unmarshal(ctx.Body.Bytes(), arg)
}

// serveHttp defined TODO
func (x *XTest) serveHttp(w http.ResponseWriter, req *http.Request) {
	api.App.ServeHTTP(w, req)
}

// handler defined TODO
func (x *XTest) handler(req *http.Request, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+token)
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
	payload := Payload{}
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
	payload := Payload{}
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
	payload := Payload{}
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
func (x *XTest) Handle(method, reqPath string, funk func(ctx *Context), t testingT, args interface{}) {
	switch method {
	case "GET":
		x.Get(reqPath, args, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "HEAD":
		x.Head(reqPath, args, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "OPTIONS":
		x.Options(reqPath, args, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "POST":
		x.Post(reqPath, args, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "PUT":
		x.Put(reqPath, args, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	}
}

// SetToken defined
func SetToken(tk string) {
	token = tk
}

// TestMain defined
func TestMain(m *testing.M) {
	x = &XTest{}
	util.Ensure(api.App.Reflesh())
	TestSysUserLogin(nil)
	os.Exit(m.Run())
}

// TestSysUserLogin defined TODO
func TestSysUserLogin(t *testing.T) {
	x.Handle("POST", "/api/sys/user/login", XTestSysUserLogin, t, XTestSysUserLoginRequest)
}

// TestOrganAdd defined TODO
// go test -v -test.run TestOrganAdd
func TestOrganAdd(t *testing.T) {
	x.Handle("POST", "/api/organ/add", XTestOrganAdd, t, XTestOrganAddRequest)
}

// BenchmarkOrganAdd defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganAdd -test.benchmem=true
func BenchmarkOrganAdd(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("POST", "/api/organ/add", XTestOrganAdd, t, XTestOrganAddRequest)
		}
	})
}

// TestOrganBatchAdd defined TODO
// go test -v -test.run TestOrganBatchAdd
func TestOrganBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/organ/batch_add", XTestOrganBatchAdd, t, XTestOrganBatchAddRequest)
}

// BenchmarkOrganBatchAdd defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganBatchAdd -test.benchmem=true
func BenchmarkOrganBatchAdd(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("POST", "/api/organ/batch_add", XTestOrganBatchAdd, t, XTestOrganBatchAddRequest)
		}
	})
}

// TestOrganDel defined TODO
// go test -v -test.run TestOrganDel
func TestOrganDel(t *testing.T) {
	x.Handle("DELETE", "/api/organ/del", XTestOrganDel, t, XTestOrganDelRequest)
}

// BenchmarkOrganDel defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganDel -test.benchmem=true
func BenchmarkOrganDel(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("DELETE", "/api/organ/del", XTestOrganDel, t, XTestOrganDelRequest)
		}
	})
}

// TestOrganBatchDel defined TODO
// go test -v -test.run TestOrganBatchDel
func TestOrganBatchDel(t *testing.T) {
	x.Handle("PUT", "/api/organ/batch_del", XTestOrganBatchDel, t, XTestOrganBatchDelRequest)
}

// BenchmarkOrganBatchDel defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganBatchDel -test.benchmem=true
func BenchmarkOrganBatchDel(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("PUT", "/api/organ/batch_del", XTestOrganBatchDel, t, XTestOrganBatchDelRequest)
		}
	})
}

// TestOrganUpdate defined TODO
// go test -v -test.run TestOrganUpdate
func TestOrganUpdate(t *testing.T) {
	x.Handle("PUT", "/api/organ/update", XTestOrganUpdate, t, XTestOrganUpdateRequest)
}

// BenchmarkOrganUpdate defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganUpdate -test.benchmem=true
func BenchmarkOrganUpdate(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("PUT", "/api/organ/update", XTestOrganUpdate, t, XTestOrganUpdateRequest)
		}
	})
}

// TestOrganBatchUpdate defined TODO
// go test -v -test.run TestOrganBatchUpdate
func TestOrganBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/organ/batch_update", XTestOrganBatchUpdate, t, XTestOrganBatchUpdateRequest)
}

// BenchmarkOrganBatchUpdate defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganBatchUpdate -test.benchmem=true
func BenchmarkOrganBatchUpdate(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("PUT", "/api/organ/batch_update", XTestOrganBatchUpdate, t, XTestOrganBatchUpdateRequest)
		}
	})
}

// TestOrganPage defined TODO
// go test -v -test.run TestOrganPage
func TestOrganPage(t *testing.T) {
	x.Handle("GET", "/api/organ/page", XTestOrganPage, t, XTestOrganPageRequest)
}

// BenchmarkOrganPage defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganPage -test.benchmem=true
func BenchmarkOrganPage(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("GET", "/api/organ/page", XTestOrganPage, t, XTestOrganPageRequest)
		}
	})
}

// TestOrganGet defined TODO
// go test -v -test.run TestOrganGet
func TestOrganGet(t *testing.T) {
	x.Handle("GET", "/api/organ/get", XTestOrganGet, t, XTestOrganGetRequest)
}

// BenchmarkOrganGet defined TODO
// go test -v -test.run=none -test.bench=^BenchmarkOrganGet -test.benchmem=true
func BenchmarkOrganGet(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Handle("GET", "/api/organ/get", XTestOrganGet, t, XTestOrganGetRequest)
		}
	})
}
