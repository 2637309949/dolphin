package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	// "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"

	//  "github.com/2637309949/dolphin/platform/conf"
	_ "github.com/2637309949/dolphin/platform/conf"
	"github.com/spf13/viper"

	"github.com/2637309949/dolphin/platform/app"
)

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
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("origin", "http://localhost:"+viper.GetString("http.port"))
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	app.App.Http.ServeHTTP(w, req)
	h(w)
}

// Get defined
func Get(url string, h func(w *httptest.ResponseRecorder)) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	handler(req, h)
}

// Post defined
func Post(url string, payform M, h func(w *httptest.ResponseRecorder)) {
	jm, err := json.Marshal(&payform)
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

// HttpTest defined
func HttpTest(reqPath string, funk func(ctx *Context), t *testing.T, p ...M) {
	switch len(p) {
	case 0:
		Get(reqPath, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	default:
		Post(reqPath, p[0], func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	}
}

// TestMain defined
func TestMain(m *testing.M) {
	app.App.Init()
	TestSysUserLogin(nil)
	os.Exit(m.Run())
}

// SetToken defined
func SetToken(token string) {
	AccessToken = token
}

var AccessToken string

// TestSysMenuPage defined
func TestSysUserLogin(t *testing.T) {
	HttpTest("/api/sys/user/login", XTestSysUserLogin, t, M{"domain": "localhost", "name": "admin", "password": "admin"})
}

// TestSysMenuPage defined
func TestSysMenuPage(t *testing.T) {
	HttpTest("/api/sys/menu/page", XTestSysMenuPage, t)
}
