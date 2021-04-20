package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	//  "github.com/2637309949/dolphin/platform/conf"
	_ "github.com/2637309949/dolphin/platform/conf"
	// "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"

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
	engine *gin.Engine
}

func (t HTTPTools) handler(m, p string, body io.Reader, h func(w *httptest.ResponseRecorder)) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(m, p, body)
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	req.Host = "local"
	t.engine.ServeHTTP(w, req)
	h(w)
}

func (t HTTPTools) Get(p string, h func(w *httptest.ResponseRecorder)) {
	t.handler("GET", p, nil, h)
}

func (t HTTPTools) Post(p string, body io.Reader, h func(w *httptest.ResponseRecorder)) {
	t.handler("POST", p, body, h)
}

var httpTools *HTTPTools

func TestMain(m *testing.M) {
	app.App.Init()
	httpTools = &HTTPTools{}
	httpTools.engine = app.App.Gin
	XTestSysUserLogin()
	os.Exit(m.Run())
}

func TestSysMenuPage(t *testing.T) { XTestSysMenuPage(t) }
