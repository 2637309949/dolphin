package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	//  "github.com/2637309949/dolphin/platform/conf"
	_ "github.com/2637309949/dolphin/platform/conf"

	// "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"

	"github.com/2637309949/dolphin/platform/app"
	"github.com/gin-gonic/gin"
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
	HTTPTools struct {
		*gin.Engine
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
)

// handler defined
func (t HTTPTools) handler(req *http.Request, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	t.ServeHTTP(w, req)
	h(w)
}

// Get defined
func (t HTTPTools) Get(url string, h func(w *httptest.ResponseRecorder)) {
	req, _ := http.NewRequest("GET", url, nil)
	t.handler(req, h)
}

// Post defined
func (t HTTPTools) Post(url string, payform map[string]interface{}, h func(w *httptest.ResponseRecorder)) {
	jm, _ := json.Marshal(&payform)
	req, _ := http.NewRequest("POST", url, bufio.NewReader(strings.NewReader(string(jm))))
	req.Header.Add("Content-Type", "application/json")
	t.handler(req, h)
}

// HttpTest defined
func HttpTest(reqPath string, funk func(ctx *Context), t *testing.T, p ...map[string]interface{}) {
	switch len(p) {
	case 0:
		httpTools.Get(reqPath, func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	default:
		httpTools.Post(reqPath, p[0], func(w *httptest.ResponseRecorder) { funk(&Context{w, t}) })
	}
}

// TestMain defined
func TestMain(m *testing.M) {
	app.App.Init()
	httpTools = &HTTPTools{}
	httpTools.Engine = app.App.Gin
	TestSysUserLogin(nil)
	os.Exit(m.Run())
}

// SetToken defined
func SetToken(token string) {
	AccessToken = token
}

var httpTools *HTTPTools
var AccessToken string

// TestSysMenuPage defined
func TestSysUserLogin(t *testing.T) {
	HttpTest("/api/sys/user/login", XTestSysUserLogin, t, map[string]interface{}{
		"domain":   "localhost",
		"name":     "admin",
		"password": "admin",
	})
}

// TestSysMenuPage defined
func TestSysMenuPage(t *testing.T) {
	HttpTest("/api/sys/menu/page", XTestSysMenuPage, t)
}
