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

type testingT interface {
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

// HTTPTools defined  tools
type HTTPTools struct {
	*gin.Engine
}

// Response defined
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (t HTTPTools) handler(req *http.Request, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	t.ServeHTTP(w, req)
	h(w)
}

func (t HTTPTools) Get(url string, h func(w *httptest.ResponseRecorder)) {
	req, _ := http.NewRequest("GET", url, nil)
	t.handler(req, h)
}

func (t HTTPTools) Post(url string, payform map[string]interface{}, h func(w *httptest.ResponseRecorder)) {
	jm, _ := json.Marshal(&payform)
	req, _ := http.NewRequest("POST", url, bufio.NewReader(strings.NewReader(string(jm))))
	req.Header.Add("Content-Type", "application/json")
	t.handler(req, h)
}

var httpTools *HTTPTools

func TestMain(m *testing.M) {
	app.App.Init()
	httpTools = &HTTPTools{}
	httpTools.Engine = app.App.Gin
	XTestSysUserLogin()
	os.Exit(m.Run())
}

// SetToken defined
func SetToken(token string) {
	AccessToken = token
}

var AccessToken string

// XTestSysUserLogin defined
func XTestSysUserLogin() {
	type (
		Data struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}
		AutoGenerated struct {
			Code int  `json:"code"`
			Data Data `json:"data"`
		}
	)
	httpTools.Post("/api/sys/user/login",
		map[string]interface{}{
			"domain":   "localhost",
			"name":     "admin",
			"password": "admin",
		},
		func(w *httptest.ResponseRecorder) {
			ret := AutoGenerated{}
			json.Unmarshal(w.Body.Bytes(), &ret)
			SetToken(ret.Data.AccessToken)
		},
	)
}

func TestSysMenuPage(t *testing.T) { XTestSysMenuPage(t) }
