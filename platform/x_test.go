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

	"github.com/2637309949/dolphin/platform/app"
)

var Token string

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
	Response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
	Context struct {
		*httptest.ResponseRecorder
		T testingT
	}
	M map[string]interface{}
)

// handler defined
func handler(req *http.Request, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+Token)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("origin", "http://localhost:"+viper.GetString("http.port"))
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	app.App.Http.ServeHTTP(w, req)
	h(w)
}

// Get defined TODO
func Get(url string, params []M, h func(w *httptest.ResponseRecorder)) {
	if len(params) > 0 {
		url += "?"
		for k, v := range params[0] {
			url += fmt.Sprintf("%v=%v&", k, v)
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	handler(req, h)
}

// Options defined TODO
func Options(url string, params []M, h func(w *httptest.ResponseRecorder)) {
	if len(params) > 0 {
		url += "?"
		for k, v := range params[0] {
			url += fmt.Sprintf("%v=%v&", k, v)
		}
	}
	req, err := http.NewRequest("OPTIONS", url, nil)
	if err != nil {
		panic(err)
	}
	handler(req, h)
}

// Post defined TODO
func Post(url string, params []M, h func(w *httptest.ResponseRecorder)) {
	jm, err := json.Marshal(&params[0])
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(jm))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	handler(req, h)
}

// Put defined TODO
func Put(url string, params []M, h func(w *httptest.ResponseRecorder)) {
	jm, err := json.Marshal(&params[0])
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(jm))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	handler(req, h)
}

// Head defined TODO
func Head(url string, params []M, h func(w *httptest.ResponseRecorder)) {
	if len(params) > 0 {
		url += "?"
		for k, v := range params[0] {
			url += fmt.Sprintf("%v=%v&", k, v)
		}
	}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		panic(err)
	}
	handler(req, h)
}

// HttpHandle defined
func HttpHandle(method, reqPath string, funk func(ctx *Context), t *testing.T, params ...M) {
	switch method {
	case "GET":
		Get(reqPath, params, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "HEAD":
		Head(reqPath, params, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "OPTIONS":
		Options(reqPath, params, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "POST":
		Post(reqPath, params, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	case "PUT":
		Put(reqPath, params, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	}
}

// TestMain defined
func TestMain(m *testing.M) {
	util.Ensure(app.App.Reflesh())
	TestSysUserLogin(nil)
	os.Exit(m.Run())
}

// SetToken defined
func SetToken(token string) {
	Token = token
}

// TestSysMenuPage defined
func TestSysUserLogin(t *testing.T) {
	HttpHandle("POST", "/api/sys/user/login", XTestSysUserLogin, t, M{"domain": "localhost", "name": "admin", "password": "admin"})
}

// TestSysMenuPage defined
func TestSysMenuPage(t *testing.T) {
	HttpHandle("GET", "/api/sys/menu/page", XTestSysMenuPage, t)
}
