package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	//  "github.com/2637309949/dolphin/platform/conf"
	_ "github.com/2637309949/dolphin/platform/conf"

	"github.com/2637309949/dolphin/platform/util"
	"github.com/spf13/viper"

	"hello/api"
)

var token string
var x *XTest

type (
	testingT interface {
		Deadline() (time.Time, bool)
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
		Parallel()
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
	M map[string]interface{}
)

// ParseBody defined TODO
func (ctx *Context) ParseBody(arg interface{}) error {
	return json.Unmarshal(ctx.Body.Bytes(), arg)
}

// serveHttp defined TODO
func (x *XTest) serveHttp(w http.ResponseWriter, req *http.Request) {
	api.App.Http.ServeHTTP(w, req)
}

// handler defined TODO
func (x *XTest) handler(req *http.Request, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("origin", "http://localhost:"+viper.GetString("http.port"))
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	x.serveHttp(w, req)
	h(w)
}

// Get defined TODO
func (x *XTest) Get(url string, args interface{}, h func(w *httptest.ResponseRecorder)) {
	brM := M{}
	br, _ := json.Marshal(args)
	json.Unmarshal(br, &brM)
	url += "?"
	for k, v := range brM {
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
	brM := M{}
	br, _ := json.Marshal(args)
	json.Unmarshal(br, &brM)
	url += "?"
	for k, v := range brM {
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
	brM := M{}
	br, _ := json.Marshal(args)
	json.Unmarshal(br, &brM)
	url += "?"
	for k, v := range brM {
		url += fmt.Sprintf("%v=%v&", k, v)
	}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		panic(err)
	}
	x.handler(req, h)
}

// HttpHandle defined TODO
func (x *XTest) Handle(method, reqPath string, funk func(ctx *Context), t *testing.T, args interface{}) {
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
	x.Handle("POST", "/api/sys/user/login", XTestSysUserLogin, t, XTestSysUserLoginM)
}

// TestArticleAdd defined TODO
func TestArticleAdd(t *testing.T) {
	x.Handle("POST", "/api/article/add", XTestArticleAdd, t, XTestArticleAddM)
}

// TestArticleBatchAdd defined TODO
func TestArticleBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/article/batch_add", XTestArticleBatchAdd, t, XTestArticleBatchAddM)
}

// TestArticleDel defined TODO
func TestArticleDel(t *testing.T) {
	x.Handle("DELETE", "/api/article/del", XTestArticleDel, t, XTestArticleDelM)
}

// TestArticleBatchDel defined TODO
func TestArticleBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/article/batch_del", XTestArticleBatchDel, t, XTestArticleBatchDelM)
}

// TestArticleUpdate defined TODO
func TestArticleUpdate(t *testing.T) {
	x.Handle("PUT", "/api/article/update", XTestArticleUpdate, t, XTestArticleUpdateM)
}

// TestArticleBatchUpdate defined TODO
func TestArticleBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/article/batch_update", XTestArticleBatchUpdate, t, XTestArticleBatchUpdateM)
}

// TestArticlePage defined TODO
func TestArticlePage(t *testing.T) {
	x.Handle("GET", "/api/article/page", XTestArticlePage, t, XTestArticlePageM)
}

// TestArticleGet defined TODO
func TestArticleGet(t *testing.T) {
	x.Handle("GET", "/api/article/get", XTestArticleGet, t, XTestArticleGetM)
}

// TestArticleTodo defined TODO
func TestArticleTodo(t *testing.T) {
	x.Handle("GET", "/api/article/todo", XTestArticleTodo, t, XTestArticleTodoM)
}