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

	"github.com/2637309949/dolphin/platform/api"
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

// TestSysAppFunAdd defined TODO
func TestSysAppFunAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/app/fun/add", XTestSysAppFunAdd, t, XTestSysAppFunAddM)
}

// TestSysAppFunBatchAdd defined TODO
func TestSysAppFunBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/app/fun/batch_add", XTestSysAppFunBatchAdd, t, XTestSysAppFunBatchAddM)
}

// TestSysAppFunDel defined TODO
func TestSysAppFunDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/app/fun/del", XTestSysAppFunDel, t, XTestSysAppFunDelM)
}

// TestSysAppFunBatchDel defined TODO
func TestSysAppFunBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/app/fun/batch_del", XTestSysAppFunBatchDel, t, XTestSysAppFunBatchDelM)
}

// TestSysAppFunUpdate defined TODO
func TestSysAppFunUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/app/fun/update", XTestSysAppFunUpdate, t, XTestSysAppFunUpdateM)
}

// TestSysAppFunBatchUpdate defined TODO
func TestSysAppFunBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/app/fun/batch_update", XTestSysAppFunBatchUpdate, t, XTestSysAppFunBatchUpdateM)
}

// TestSysAppFunPage defined TODO
func TestSysAppFunPage(t *testing.T) {
	x.Handle("GET", "/api/sys/app/fun/page", XTestSysAppFunPage, t, XTestSysAppFunPageM)
}

// TestSysAppFunTree defined TODO
func TestSysAppFunTree(t *testing.T) {
	x.Handle("GET", "/api/sys/app/fun/tree", XTestSysAppFunTree, t, XTestSysAppFunTreeM)
}

// TestSysAppFunGet defined TODO
func TestSysAppFunGet(t *testing.T) {
	x.Handle("GET", "/api/sys/app/fun/get", XTestSysAppFunGet, t, XTestSysAppFunGetM)
}

// TestSysAreaAdd defined TODO
func TestSysAreaAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/area/add", XTestSysAreaAdd, t, XTestSysAreaAddM)
}

// TestSysAreaBatchAdd defined TODO
func TestSysAreaBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/area/batch_add", XTestSysAreaBatchAdd, t, XTestSysAreaBatchAddM)
}

// TestSysAreaDel defined TODO
func TestSysAreaDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/area/del", XTestSysAreaDel, t, XTestSysAreaDelM)
}

// TestSysAreaBatchDel defined TODO
func TestSysAreaBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/area/batch_del", XTestSysAreaBatchDel, t, XTestSysAreaBatchDelM)
}

// TestSysAreaUpdate defined TODO
func TestSysAreaUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/area/update", XTestSysAreaUpdate, t, XTestSysAreaUpdateM)
}

// TestSysAreaBatchUpdate defined TODO
func TestSysAreaBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/area/batch_update", XTestSysAreaBatchUpdate, t, XTestSysAreaBatchUpdateM)
}

// TestSysAreaPage defined TODO
func TestSysAreaPage(t *testing.T) {
	x.Handle("GET", "/api/sys/area/page", XTestSysAreaPage, t, XTestSysAreaPageM)
}

// TestSysAreaGet defined TODO
func TestSysAreaGet(t *testing.T) {
	x.Handle("GET", "/api/sys/area/get", XTestSysAreaGet, t, XTestSysAreaGetM)
}

// TestSysAttachmentAdd defined TODO
func TestSysAttachmentAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/attachment/add", XTestSysAttachmentAdd, t, XTestSysAttachmentAddM)
}

// TestSysAttachmentBatchAdd defined TODO
func TestSysAttachmentBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/attachment/batch_add", XTestSysAttachmentBatchAdd, t, XTestSysAttachmentBatchAddM)
}

// TestSysAttachmentUpload defined TODO
func TestSysAttachmentUpload(t *testing.T) {
	x.Handle("POST", "/api/sys/attachment/upload", XTestSysAttachmentUpload, t, XTestSysAttachmentUploadM)
}

// TestSysAttachmentExport defined TODO
func TestSysAttachmentExport(t *testing.T) {
	x.Handle("GET", "/api/sys/attachment/export", XTestSysAttachmentExport, t, XTestSysAttachmentExportM)
}

// TestSysAttachmentDel defined TODO
func TestSysAttachmentDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/attachment/del", XTestSysAttachmentDel, t, XTestSysAttachmentDelM)
}

// TestSysAttachmentBatchDel defined TODO
func TestSysAttachmentBatchDel(t *testing.T) {
	x.Handle("POST", "/api/sys/attachment/batch_del", XTestSysAttachmentBatchDel, t, XTestSysAttachmentBatchDelM)
}

// TestSysAttachmentUpdate defined TODO
func TestSysAttachmentUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/attachment/update", XTestSysAttachmentUpdate, t, XTestSysAttachmentUpdateM)
}

// TestSysAttachmentBatchUpdate defined TODO
func TestSysAttachmentBatchUpdate(t *testing.T) {
	x.Handle("POST", "/api/sys/attachment/batch_update", XTestSysAttachmentBatchUpdate, t, XTestSysAttachmentBatchUpdateM)
}

// TestSysAttachmentPage defined TODO
func TestSysAttachmentPage(t *testing.T) {
	x.Handle("GET", "/api/sys/attachment/page", XTestSysAttachmentPage, t, XTestSysAttachmentPageM)
}

// TestSysAttachmentGet defined TODO
func TestSysAttachmentGet(t *testing.T) {
	x.Handle("GET", "/api/sys/attachment/get", XTestSysAttachmentGet, t, XTestSysAttachmentGetM)
}

// TestSysCasLogout defined TODO
func TestSysCasLogout(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/logout", XTestSysCasLogout, t, XTestSysCasLogoutM)
}

// TestSysCasAffirm defined TODO
func TestSysCasAffirm(t *testing.T) {
	x.Handle("POST", "/api/sys/cas/affirm", XTestSysCasAffirm, t, XTestSysCasAffirmM)
}

// TestSysCasAuthorize defined TODO
func TestSysCasAuthorize(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/authorize", XTestSysCasAuthorize, t, XTestSysCasAuthorizeM)
}

// TestSysCasToken defined TODO
func TestSysCasToken(t *testing.T) {
	x.Handle("POST", "/api/sys/cas/token", XTestSysCasToken, t, XTestSysCasTokenM)
}

// TestSysCasURL defined TODO
func TestSysCasURL(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/url", XTestSysCasURL, t, XTestSysCasURLM)
}

// TestSysCasOauth2 defined TODO
func TestSysCasOauth2(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/oauth2", XTestSysCasOauth2, t, XTestSysCasOauth2M)
}

// TestSysCasRefresh defined TODO
func TestSysCasRefresh(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/refresh", XTestSysCasRefresh, t, XTestSysCasRefreshM)
}

// TestSysCasCheck defined TODO
func TestSysCasCheck(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/check", XTestSysCasCheck, t, XTestSysCasCheckM)
}

// TestSysCasProfile defined TODO
func TestSysCasProfile(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/profile", XTestSysCasProfile, t, XTestSysCasProfileM)
}

// TestSysCasQrcode defined TODO
func TestSysCasQrcode(t *testing.T) {
	x.Handle("GET", "/api/sys/cas/qrcode", XTestSysCasQrcode, t, XTestSysCasQrcodeM)
}

// TestSysClientAdd defined TODO
func TestSysClientAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/client/add", XTestSysClientAdd, t, XTestSysClientAddM)
}

// TestSysClientBatchAdd defined TODO
func TestSysClientBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/client/batch_add", XTestSysClientBatchAdd, t, XTestSysClientBatchAddM)
}

// TestSysClientDel defined TODO
func TestSysClientDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/client/del", XTestSysClientDel, t, XTestSysClientDelM)
}

// TestSysClientBatchDel defined TODO
func TestSysClientBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/client/batch_del", XTestSysClientBatchDel, t, XTestSysClientBatchDelM)
}

// TestSysClientUpdate defined TODO
func TestSysClientUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/client/update", XTestSysClientUpdate, t, XTestSysClientUpdateM)
}

// TestSysClientBatchUpdate defined TODO
func TestSysClientBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/client/batch_update", XTestSysClientBatchUpdate, t, XTestSysClientBatchUpdateM)
}

// TestSysClientPage defined TODO
func TestSysClientPage(t *testing.T) {
	x.Handle("GET", "/api/sys/client/page", XTestSysClientPage, t, XTestSysClientPageM)
}

// TestSysClientGet defined TODO
func TestSysClientGet(t *testing.T) {
	x.Handle("GET", "/api/sys/client/get", XTestSysClientGet, t, XTestSysClientGetM)
}

// TestSysCommentAdd defined TODO
func TestSysCommentAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/comment/add", XTestSysCommentAdd, t, XTestSysCommentAddM)
}

// TestSysCommentBatchAdd defined TODO
func TestSysCommentBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/comment/batch_add", XTestSysCommentBatchAdd, t, XTestSysCommentBatchAddM)
}

// TestSysCommentDel defined TODO
func TestSysCommentDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/comment/del", XTestSysCommentDel, t, XTestSysCommentDelM)
}

// TestSysCommentBatchDel defined TODO
func TestSysCommentBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/comment/batch_del", XTestSysCommentBatchDel, t, XTestSysCommentBatchDelM)
}

// TestSysCommentUpdate defined TODO
func TestSysCommentUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/comment/update", XTestSysCommentUpdate, t, XTestSysCommentUpdateM)
}

// TestSysCommentBatchUpdate defined TODO
func TestSysCommentBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/comment/batch_update", XTestSysCommentBatchUpdate, t, XTestSysCommentBatchUpdateM)
}

// TestSysCommentPage defined TODO
func TestSysCommentPage(t *testing.T) {
	x.Handle("GET", "/api/sys/comment/page", XTestSysCommentPage, t, XTestSysCommentPageM)
}

// TestSysCommentGet defined TODO
func TestSysCommentGet(t *testing.T) {
	x.Handle("GET", "/api/sys/comment/get", XTestSysCommentGet, t, XTestSysCommentGetM)
}

// TestSysDataPermissionAdd defined TODO
func TestSysDataPermissionAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/data/permission/add", XTestSysDataPermissionAdd, t, XTestSysDataPermissionAddM)
}

// TestSysDataPermissionBatchAdd defined TODO
func TestSysDataPermissionBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/data/permission/batch_add", XTestSysDataPermissionBatchAdd, t, XTestSysDataPermissionBatchAddM)
}

// TestSysDataPermissionDel defined TODO
func TestSysDataPermissionDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/data/permission/del", XTestSysDataPermissionDel, t, XTestSysDataPermissionDelM)
}

// TestSysDataPermissionBatchDel defined TODO
func TestSysDataPermissionBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/data/permission/batch_del", XTestSysDataPermissionBatchDel, t, XTestSysDataPermissionBatchDelM)
}

// TestSysDataPermissionUpdate defined TODO
func TestSysDataPermissionUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/data/permission/update", XTestSysDataPermissionUpdate, t, XTestSysDataPermissionUpdateM)
}

// TestSysDataPermissionBatchUpdate defined TODO
func TestSysDataPermissionBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/data/permission/batch_update", XTestSysDataPermissionBatchUpdate, t, XTestSysDataPermissionBatchUpdateM)
}

// TestSysDataPermissionPage defined TODO
func TestSysDataPermissionPage(t *testing.T) {
	x.Handle("GET", "/api/sys/data/permission/page", XTestSysDataPermissionPage, t, XTestSysDataPermissionPageM)
}

// TestSysDataPermissionGet defined TODO
func TestSysDataPermissionGet(t *testing.T) {
	x.Handle("GET", "/api/sys/data/permission/get", XTestSysDataPermissionGet, t, XTestSysDataPermissionGetM)
}

// TestDebugPprof defined TODO
func TestDebugPprof(t *testing.T) {
	x.Handle("GET", "/debug/pprof/", XTestDebugPprof, t, XTestDebugPprofM)
}

// TestDebugHeap defined TODO
func TestDebugHeap(t *testing.T) {
	x.Handle("GET", "/debug/pprof/heap", XTestDebugHeap, t, XTestDebugHeapM)
}

// TestDebugGoroutine defined TODO
func TestDebugGoroutine(t *testing.T) {
	x.Handle("GET", "/debug/pprof/goroutine", XTestDebugGoroutine, t, XTestDebugGoroutineM)
}

// TestDebugAllocs defined TODO
func TestDebugAllocs(t *testing.T) {
	x.Handle("GET", "/debug/pprof/allocs", XTestDebugAllocs, t, XTestDebugAllocsM)
}

// TestDebugBlock defined TODO
func TestDebugBlock(t *testing.T) {
	x.Handle("GET", "/debug/pprof/block", XTestDebugBlock, t, XTestDebugBlockM)
}

// TestDebugThreadcreate defined TODO
func TestDebugThreadcreate(t *testing.T) {
	x.Handle("GET", "/debug/pprof/threadcreate", XTestDebugThreadcreate, t, XTestDebugThreadcreateM)
}

// TestDebugCmdline defined TODO
func TestDebugCmdline(t *testing.T) {
	x.Handle("GET", "/debug/pprof/cmdline", XTestDebugCmdline, t, XTestDebugCmdlineM)
}

// TestDebugProfile defined TODO
func TestDebugProfile(t *testing.T) {
	x.Handle("GET", "/debug/pprof/profile", XTestDebugProfile, t, XTestDebugProfileM)
}

// TestDebugSymbol defined TODO
func TestDebugSymbol(t *testing.T) {
	x.Handle("GET,POST", "/debug/pprof/symbol", XTestDebugSymbol, t, XTestDebugSymbolM)
}

// TestDebugTrace defined TODO
func TestDebugTrace(t *testing.T) {
	x.Handle("GET", "/debug/pprof/trace", XTestDebugTrace, t, XTestDebugTraceM)
}

// TestDebugMutex defined TODO
func TestDebugMutex(t *testing.T) {
	x.Handle("GET", "/debug/pprof/mutex", XTestDebugMutex, t, XTestDebugMutexM)
}

// TestSysDingtalkOauth2 defined TODO
func TestSysDingtalkOauth2(t *testing.T) {
	x.Handle("GET", "/api/sys/dingtalk/oauth2", XTestSysDingtalkOauth2, t, XTestSysDingtalkOauth2M)
}

// TestSysDomainAdd defined TODO
func TestSysDomainAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/domain/add", XTestSysDomainAdd, t, XTestSysDomainAddM)
}

// TestSysDomainBatchAdd defined TODO
func TestSysDomainBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/domain/batch_add", XTestSysDomainBatchAdd, t, XTestSysDomainBatchAddM)
}

// TestSysDomainDel defined TODO
func TestSysDomainDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/domain/del", XTestSysDomainDel, t, XTestSysDomainDelM)
}

// TestSysDomainBatchDel defined TODO
func TestSysDomainBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/domain/batch_del", XTestSysDomainBatchDel, t, XTestSysDomainBatchDelM)
}

// TestSysDomainUpdate defined TODO
func TestSysDomainUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/domain/update", XTestSysDomainUpdate, t, XTestSysDomainUpdateM)
}

// TestSysDomainBatchUpdate defined TODO
func TestSysDomainBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/domain/batch_update", XTestSysDomainBatchUpdate, t, XTestSysDomainBatchUpdateM)
}

// TestSysDomainPage defined TODO
func TestSysDomainPage(t *testing.T) {
	x.Handle("GET", "/api/sys/domain/page", XTestSysDomainPage, t, XTestSysDomainPageM)
}

// TestSysDomainGet defined TODO
func TestSysDomainGet(t *testing.T) {
	x.Handle("GET", "/api/sys/domain/get", XTestSysDomainGet, t, XTestSysDomainGetM)
}

// TestSysMenuAdd defined TODO
func TestSysMenuAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/menu/add", XTestSysMenuAdd, t, XTestSysMenuAddM)
}

// TestSysMenuBatchAdd defined TODO
func TestSysMenuBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/menu/batch_add", XTestSysMenuBatchAdd, t, XTestSysMenuBatchAddM)
}

// TestSysMenuDel defined TODO
func TestSysMenuDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/menu/del", XTestSysMenuDel, t, XTestSysMenuDelM)
}

// TestSysMenuBatchDel defined TODO
func TestSysMenuBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/menu/batch_del", XTestSysMenuBatchDel, t, XTestSysMenuBatchDelM)
}

// TestSysMenuUpdate defined TODO
func TestSysMenuUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/menu/update", XTestSysMenuUpdate, t, XTestSysMenuUpdateM)
}

// TestSysMenuBatchUpdate defined TODO
func TestSysMenuBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/menu/batch_update", XTestSysMenuBatchUpdate, t, XTestSysMenuBatchUpdateM)
}

// TestSysMenuSidebar defined TODO
func TestSysMenuSidebar(t *testing.T) {
	x.Handle("GET", "/api/sys/menu/sidebar", XTestSysMenuSidebar, t, XTestSysMenuSidebarM)
}

// TestSysMenuPage defined TODO
func TestSysMenuPage(t *testing.T) {
	x.Handle("GET", "/api/sys/menu/page", XTestSysMenuPage, t, XTestSysMenuPageM)
}

// TestSysMenuTree defined TODO
func TestSysMenuTree(t *testing.T) {
	x.Handle("GET", "/api/sys/menu/tree", XTestSysMenuTree, t, XTestSysMenuTreeM)
}

// TestSysMenuGet defined TODO
func TestSysMenuGet(t *testing.T) {
	x.Handle("GET", "/api/sys/menu/get", XTestSysMenuGet, t, XTestSysMenuGetM)
}

// TestSysNotificationAdd defined TODO
func TestSysNotificationAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/notification/add", XTestSysNotificationAdd, t, XTestSysNotificationAddM)
}

// TestSysNotificationBatchAdd defined TODO
func TestSysNotificationBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/notification/batch_add", XTestSysNotificationBatchAdd, t, XTestSysNotificationBatchAddM)
}

// TestSysNotificationDel defined TODO
func TestSysNotificationDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/notification/del", XTestSysNotificationDel, t, XTestSysNotificationDelM)
}

// TestSysNotificationBatchDel defined TODO
func TestSysNotificationBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/notification/batch_del", XTestSysNotificationBatchDel, t, XTestSysNotificationBatchDelM)
}

// TestSysNotificationUpdate defined TODO
func TestSysNotificationUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/notification/update", XTestSysNotificationUpdate, t, XTestSysNotificationUpdateM)
}

// TestSysNotificationBatchUpdate defined TODO
func TestSysNotificationBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/notification/batch_update", XTestSysNotificationBatchUpdate, t, XTestSysNotificationBatchUpdateM)
}

// TestSysNotificationPage defined TODO
func TestSysNotificationPage(t *testing.T) {
	x.Handle("GET", "/api/sys/notification/page", XTestSysNotificationPage, t, XTestSysNotificationPageM)
}

// TestSysNotificationGet defined TODO
func TestSysNotificationGet(t *testing.T) {
	x.Handle("GET", "/api/sys/notification/get", XTestSysNotificationGet, t, XTestSysNotificationGetM)
}

// TestSysOptionsetAdd defined TODO
func TestSysOptionsetAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/optionset/add", XTestSysOptionsetAdd, t, XTestSysOptionsetAddM)
}

// TestSysOptionsetBatchAdd defined TODO
func TestSysOptionsetBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/optionset/batch_add", XTestSysOptionsetBatchAdd, t, XTestSysOptionsetBatchAddM)
}

// TestSysOptionsetDel defined TODO
func TestSysOptionsetDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/optionset/del", XTestSysOptionsetDel, t, XTestSysOptionsetDelM)
}

// TestSysOptionsetBatchDel defined TODO
func TestSysOptionsetBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/optionset/batch_del", XTestSysOptionsetBatchDel, t, XTestSysOptionsetBatchDelM)
}

// TestSysOptionsetUpdate defined TODO
func TestSysOptionsetUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/optionset/update", XTestSysOptionsetUpdate, t, XTestSysOptionsetUpdateM)
}

// TestSysOptionsetBatchUpdate defined TODO
func TestSysOptionsetBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/optionset/batch_update", XTestSysOptionsetBatchUpdate, t, XTestSysOptionsetBatchUpdateM)
}

// TestSysOptionsetPage defined TODO
func TestSysOptionsetPage(t *testing.T) {
	x.Handle("GET", "/api/sys/optionset/page", XTestSysOptionsetPage, t, XTestSysOptionsetPageM)
}

// TestSysOptionsetGet defined TODO
func TestSysOptionsetGet(t *testing.T) {
	x.Handle("GET", "/api/sys/optionset/get", XTestSysOptionsetGet, t, XTestSysOptionsetGetM)
}

// TestSysOrgAdd defined TODO
func TestSysOrgAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/org/add", XTestSysOrgAdd, t, XTestSysOrgAddM)
}

// TestSysOrgBatchAdd defined TODO
func TestSysOrgBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/org/batch_add", XTestSysOrgBatchAdd, t, XTestSysOrgBatchAddM)
}

// TestSysOrgDel defined TODO
func TestSysOrgDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/org/del", XTestSysOrgDel, t, XTestSysOrgDelM)
}

// TestSysOrgBatchDel defined TODO
func TestSysOrgBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/org/batch_del", XTestSysOrgBatchDel, t, XTestSysOrgBatchDelM)
}

// TestSysOrgUpdate defined TODO
func TestSysOrgUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/org/update", XTestSysOrgUpdate, t, XTestSysOrgUpdateM)
}

// TestSysOrgBatchUpdate defined TODO
func TestSysOrgBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/org/batch_update", XTestSysOrgBatchUpdate, t, XTestSysOrgBatchUpdateM)
}

// TestSysOrgPage defined TODO
func TestSysOrgPage(t *testing.T) {
	x.Handle("GET", "/api/sys/org/page", XTestSysOrgPage, t, XTestSysOrgPageM)
}

// TestSysOrgTree defined TODO
func TestSysOrgTree(t *testing.T) {
	x.Handle("GET", "/api/sys/org/tree", XTestSysOrgTree, t, XTestSysOrgTreeM)
}

// TestSysOrgGet defined TODO
func TestSysOrgGet(t *testing.T) {
	x.Handle("GET", "/api/sys/org/get", XTestSysOrgGet, t, XTestSysOrgGetM)
}

// TestSysPermissionAdd defined TODO
func TestSysPermissionAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/permission/add", XTestSysPermissionAdd, t, XTestSysPermissionAddM)
}

// TestSysPermissionBatchAdd defined TODO
func TestSysPermissionBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/permission/batch_add", XTestSysPermissionBatchAdd, t, XTestSysPermissionBatchAddM)
}

// TestSysPermissionDel defined TODO
func TestSysPermissionDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/permission/del", XTestSysPermissionDel, t, XTestSysPermissionDelM)
}

// TestSysPermissionBatchDel defined TODO
func TestSysPermissionBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/permission/batch_del", XTestSysPermissionBatchDel, t, XTestSysPermissionBatchDelM)
}

// TestSysPermissionUpdate defined TODO
func TestSysPermissionUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/permission/update", XTestSysPermissionUpdate, t, XTestSysPermissionUpdateM)
}

// TestSysPermissionBatchUpdate defined TODO
func TestSysPermissionBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/permission/batch_update", XTestSysPermissionBatchUpdate, t, XTestSysPermissionBatchUpdateM)
}

// TestSysPermissionPage defined TODO
func TestSysPermissionPage(t *testing.T) {
	x.Handle("GET", "/api/sys/permission/page", XTestSysPermissionPage, t, XTestSysPermissionPageM)
}

// TestSysPermissionGet defined TODO
func TestSysPermissionGet(t *testing.T) {
	x.Handle("GET", "/api/sys/permission/get", XTestSysPermissionGet, t, XTestSysPermissionGetM)
}

// TestSysRoleAdd defined TODO
func TestSysRoleAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/role/add", XTestSysRoleAdd, t, XTestSysRoleAddM)
}

// TestSysRoleBatchAdd defined TODO
func TestSysRoleBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/role/batch_add", XTestSysRoleBatchAdd, t, XTestSysRoleBatchAddM)
}

// TestSysRoleDel defined TODO
func TestSysRoleDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/role/del", XTestSysRoleDel, t, XTestSysRoleDelM)
}

// TestSysRoleBatchDel defined TODO
func TestSysRoleBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/role/batch_del", XTestSysRoleBatchDel, t, XTestSysRoleBatchDelM)
}

// TestSysRoleUpdate defined TODO
func TestSysRoleUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/role/update", XTestSysRoleUpdate, t, XTestSysRoleUpdateM)
}

// TestSysRoleBatchUpdate defined TODO
func TestSysRoleBatchUpdate(t *testing.T) {
	x.Handle("POST", "/api/sys/role/batch_update", XTestSysRoleBatchUpdate, t, XTestSysRoleBatchUpdateM)
}

// TestSysRolePage defined TODO
func TestSysRolePage(t *testing.T) {
	x.Handle("GET", "/api/sys/role/page", XTestSysRolePage, t, XTestSysRolePageM)
}

// TestSysRoleRoleMenuTree defined TODO
func TestSysRoleRoleMenuTree(t *testing.T) {
	x.Handle("GET", "/api/sys/role/role_menu_tree", XTestSysRoleRoleMenuTree, t, XTestSysRoleRoleMenuTreeM)
}

// TestSysRoleRoleAppFunTree defined TODO
func TestSysRoleRoleAppFunTree(t *testing.T) {
	x.Handle("GET", "/api/sys/role/role_app_fun_tree", XTestSysRoleRoleAppFunTree, t, XTestSysRoleRoleAppFunTreeM)
}

// TestSysRoleGet defined TODO
func TestSysRoleGet(t *testing.T) {
	x.Handle("GET", "/api/sys/role/get", XTestSysRoleGet, t, XTestSysRoleGetM)
}

// TestSysRoleMenuAdd defined TODO
func TestSysRoleMenuAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/role/menu/add", XTestSysRoleMenuAdd, t, XTestSysRoleMenuAddM)
}

// TestSysRoleMenuBatchAdd defined TODO
func TestSysRoleMenuBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/role/menu/batch_add", XTestSysRoleMenuBatchAdd, t, XTestSysRoleMenuBatchAddM)
}

// TestSysRoleMenuDel defined TODO
func TestSysRoleMenuDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/role/menu/del", XTestSysRoleMenuDel, t, XTestSysRoleMenuDelM)
}

// TestSysRoleMenuBatchDel defined TODO
func TestSysRoleMenuBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/role/menu/batch_del", XTestSysRoleMenuBatchDel, t, XTestSysRoleMenuBatchDelM)
}

// TestSysRoleMenuUpdate defined TODO
func TestSysRoleMenuUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/role/menu/update", XTestSysRoleMenuUpdate, t, XTestSysRoleMenuUpdateM)
}

// TestSysRoleMenuBatchUpdate defined TODO
func TestSysRoleMenuBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/role/menu/batch_update", XTestSysRoleMenuBatchUpdate, t, XTestSysRoleMenuBatchUpdateM)
}

// TestSysRoleMenuPage defined TODO
func TestSysRoleMenuPage(t *testing.T) {
	x.Handle("GET", "/api/sys/role/menu/page", XTestSysRoleMenuPage, t, XTestSysRoleMenuPageM)
}

// TestSysRoleMenuGet defined TODO
func TestSysRoleMenuGet(t *testing.T) {
	x.Handle("GET", "/api/sys/role/menu/get", XTestSysRoleMenuGet, t, XTestSysRoleMenuGetM)
}

// TestSysScheduleAdd defined TODO
func TestSysScheduleAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/schedule/add", XTestSysScheduleAdd, t, XTestSysScheduleAddM)
}

// TestSysScheduleBatchAdd defined TODO
func TestSysScheduleBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/schedule/batch_add", XTestSysScheduleBatchAdd, t, XTestSysScheduleBatchAddM)
}

// TestSysScheduleDel defined TODO
func TestSysScheduleDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/schedule/del", XTestSysScheduleDel, t, XTestSysScheduleDelM)
}

// TestSysScheduleBatchDel defined TODO
func TestSysScheduleBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/schedule/batch_del", XTestSysScheduleBatchDel, t, XTestSysScheduleBatchDelM)
}

// TestSysScheduleUpdate defined TODO
func TestSysScheduleUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/schedule/update", XTestSysScheduleUpdate, t, XTestSysScheduleUpdateM)
}

// TestSysScheduleBatchUpdate defined TODO
func TestSysScheduleBatchUpdate(t *testing.T) {
	x.Handle("POST", "/api/sys/schedule/batch_update", XTestSysScheduleBatchUpdate, t, XTestSysScheduleBatchUpdateM)
}

// TestSysSchedulePage defined TODO
func TestSysSchedulePage(t *testing.T) {
	x.Handle("GET", "/api/sys/schedule/page", XTestSysSchedulePage, t, XTestSysSchedulePageM)
}

// TestSysScheduleGet defined TODO
func TestSysScheduleGet(t *testing.T) {
	x.Handle("GET", "/api/sys/schedule/get", XTestSysScheduleGet, t, XTestSysScheduleGetM)
}

// TestSysScheduleHistoryPage defined TODO
func TestSysScheduleHistoryPage(t *testing.T) {
	x.Handle("GET", "/api/sys/schedule/history/page", XTestSysScheduleHistoryPage, t, XTestSysScheduleHistoryPageM)
}

// TestSysSchedulingAdd defined TODO
func TestSysSchedulingAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/scheduling/add", XTestSysSchedulingAdd, t, XTestSysSchedulingAddM)
}

// TestSysSchedulingDel defined TODO
func TestSysSchedulingDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/scheduling/del", XTestSysSchedulingDel, t, XTestSysSchedulingDelM)
}

// TestSysSchedulingUpdate defined TODO
func TestSysSchedulingUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/scheduling/update", XTestSysSchedulingUpdate, t, XTestSysSchedulingUpdateM)
}

// TestSysSchedulingPage defined TODO
func TestSysSchedulingPage(t *testing.T) {
	x.Handle("GET", "/api/sys/scheduling/page", XTestSysSchedulingPage, t, XTestSysSchedulingPageM)
}

// TestSysSchedulingGet defined TODO
func TestSysSchedulingGet(t *testing.T) {
	x.Handle("GET", "/api/sys/scheduling/get", XTestSysSchedulingGet, t, XTestSysSchedulingGetM)
}

// TestSysSettingAdd defined TODO
func TestSysSettingAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/setting/add", XTestSysSettingAdd, t, XTestSysSettingAddM)
}

// TestSysSettingBatchAdd defined TODO
func TestSysSettingBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/setting/batch_add", XTestSysSettingBatchAdd, t, XTestSysSettingBatchAddM)
}

// TestSysSettingDel defined TODO
func TestSysSettingDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/setting/del", XTestSysSettingDel, t, XTestSysSettingDelM)
}

// TestSysSettingBatchDel defined TODO
func TestSysSettingBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/setting/batch_del", XTestSysSettingBatchDel, t, XTestSysSettingBatchDelM)
}

// TestSysSettingUpdate defined TODO
func TestSysSettingUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/setting/update", XTestSysSettingUpdate, t, XTestSysSettingUpdateM)
}

// TestSysSettingBatchUpdate defined TODO
func TestSysSettingBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/setting/batch_update", XTestSysSettingBatchUpdate, t, XTestSysSettingBatchUpdateM)
}

// TestSysSettingPage defined TODO
func TestSysSettingPage(t *testing.T) {
	x.Handle("GET", "/api/sys/setting/page", XTestSysSettingPage, t, XTestSysSettingPageM)
}

// TestSysSettingGet defined TODO
func TestSysSettingGet(t *testing.T) {
	x.Handle("GET", "/api/sys/setting/get", XTestSysSettingGet, t, XTestSysSettingGetM)
}

// TestSysTableAdd defined TODO
func TestSysTableAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/table/add", XTestSysTableAdd, t, XTestSysTableAddM)
}

// TestSysTableBatchAdd defined TODO
func TestSysTableBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/table/batch_add", XTestSysTableBatchAdd, t, XTestSysTableBatchAddM)
}

// TestSysTableDel defined TODO
func TestSysTableDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/table/del", XTestSysTableDel, t, XTestSysTableDelM)
}

// TestSysTableBatchDel defined TODO
func TestSysTableBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/table/batch_del", XTestSysTableBatchDel, t, XTestSysTableBatchDelM)
}

// TestSysTableUpdate defined TODO
func TestSysTableUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/table/update", XTestSysTableUpdate, t, XTestSysTableUpdateM)
}

// TestSysTableBatchUpdate defined TODO
func TestSysTableBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/table/batch_update", XTestSysTableBatchUpdate, t, XTestSysTableBatchUpdateM)
}

// TestSysTablePage defined TODO
func TestSysTablePage(t *testing.T) {
	x.Handle("GET", "/api/sys/table/page", XTestSysTablePage, t, XTestSysTablePageM)
}

// TestSysTableGet defined TODO
func TestSysTableGet(t *testing.T) {
	x.Handle("GET", "/api/sys/table/get", XTestSysTableGet, t, XTestSysTableGetM)
}

// TestSysTableColumnAdd defined TODO
func TestSysTableColumnAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/table/column/add", XTestSysTableColumnAdd, t, XTestSysTableColumnAddM)
}

// TestSysTableColumnBatchAdd defined TODO
func TestSysTableColumnBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/table/column/batch_add", XTestSysTableColumnBatchAdd, t, XTestSysTableColumnBatchAddM)
}

// TestSysTableColumnDel defined TODO
func TestSysTableColumnDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/table/column/del", XTestSysTableColumnDel, t, XTestSysTableColumnDelM)
}

// TestSysTableColumnBatchDel defined TODO
func TestSysTableColumnBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/table/column/batch_del", XTestSysTableColumnBatchDel, t, XTestSysTableColumnBatchDelM)
}

// TestSysTableColumnUpdate defined TODO
func TestSysTableColumnUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/table/column/update", XTestSysTableColumnUpdate, t, XTestSysTableColumnUpdateM)
}

// TestSysTableColumnBatchUpdate defined TODO
func TestSysTableColumnBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/table/column/batch_update", XTestSysTableColumnBatchUpdate, t, XTestSysTableColumnBatchUpdateM)
}

// TestSysTableColumnPage defined TODO
func TestSysTableColumnPage(t *testing.T) {
	x.Handle("GET", "/api/sys/table/column/page", XTestSysTableColumnPage, t, XTestSysTableColumnPageM)
}

// TestSysTableColumnGet defined TODO
func TestSysTableColumnGet(t *testing.T) {
	x.Handle("GET", "/api/sys/table/column/get", XTestSysTableColumnGet, t, XTestSysTableColumnGetM)
}

// TestSysTagAdd defined TODO
func TestSysTagAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/tag/add", XTestSysTagAdd, t, XTestSysTagAddM)
}

// TestSysTagBatchAdd defined TODO
func TestSysTagBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/tag/batch_add", XTestSysTagBatchAdd, t, XTestSysTagBatchAddM)
}

// TestSysTagDel defined TODO
func TestSysTagDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/tag/del", XTestSysTagDel, t, XTestSysTagDelM)
}

// TestSysTagBatchDel defined TODO
func TestSysTagBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/tag/batch_del", XTestSysTagBatchDel, t, XTestSysTagBatchDelM)
}

// TestSysTagUpdate defined TODO
func TestSysTagUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/tag/update", XTestSysTagUpdate, t, XTestSysTagUpdateM)
}

// TestSysTagBatchUpdate defined TODO
func TestSysTagBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/tag/batch_update", XTestSysTagBatchUpdate, t, XTestSysTagBatchUpdateM)
}

// TestSysTagPage defined TODO
func TestSysTagPage(t *testing.T) {
	x.Handle("GET", "/api/sys/tag/page", XTestSysTagPage, t, XTestSysTagPageM)
}

// TestSysTagGet defined TODO
func TestSysTagGet(t *testing.T) {
	x.Handle("GET", "/api/sys/tag/get", XTestSysTagGet, t, XTestSysTagGetM)
}

// TestSysTagGroupAdd defined TODO
func TestSysTagGroupAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/tag/group/add", XTestSysTagGroupAdd, t, XTestSysTagGroupAddM)
}

// TestSysTagGroupBatchAdd defined TODO
func TestSysTagGroupBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/tag/group/batch_add", XTestSysTagGroupBatchAdd, t, XTestSysTagGroupBatchAddM)
}

// TestSysTagGroupDel defined TODO
func TestSysTagGroupDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/tag/group/del", XTestSysTagGroupDel, t, XTestSysTagGroupDelM)
}

// TestSysTagGroupBatchDel defined TODO
func TestSysTagGroupBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/tag/group/batch_del", XTestSysTagGroupBatchDel, t, XTestSysTagGroupBatchDelM)
}

// TestSysTagGroupUpdate defined TODO
func TestSysTagGroupUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/tag/group/update", XTestSysTagGroupUpdate, t, XTestSysTagGroupUpdateM)
}

// TestSysTagGroupBatchUpdate defined TODO
func TestSysTagGroupBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/tag/group/batch_update", XTestSysTagGroupBatchUpdate, t, XTestSysTagGroupBatchUpdateM)
}

// TestSysTagGroupPage defined TODO
func TestSysTagGroupPage(t *testing.T) {
	x.Handle("GET", "/api/sys/tag/group/page", XTestSysTagGroupPage, t, XTestSysTagGroupPageM)
}

// TestSysTagGroupGet defined TODO
func TestSysTagGroupGet(t *testing.T) {
	x.Handle("GET", "/api/sys/tag/group/get", XTestSysTagGroupGet, t, XTestSysTagGroupGetM)
}

// TestSysTrackerPage defined TODO
func TestSysTrackerPage(t *testing.T) {
	x.Handle("GET", "/api/sys/tracker/page", XTestSysTrackerPage, t, XTestSysTrackerPageM)
}

// TestSysTrackerGet defined TODO
func TestSysTrackerGet(t *testing.T) {
	x.Handle("GET", "/api/sys/tracker/get", XTestSysTrackerGet, t, XTestSysTrackerGetM)
}

// TestSysUserAdd defined TODO
func TestSysUserAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/user/add", XTestSysUserAdd, t, XTestSysUserAddM)
}

// TestSysUserBatchAdd defined TODO
func TestSysUserBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/user/batch_add", XTestSysUserBatchAdd, t, XTestSysUserBatchAddM)
}

// TestSysUserDel defined TODO
func TestSysUserDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/user/del", XTestSysUserDel, t, XTestSysUserDelM)
}

// TestSysUserBatchDel defined TODO
func TestSysUserBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/user/batch_del", XTestSysUserBatchDel, t, XTestSysUserBatchDelM)
}

// TestSysUserUpdate defined TODO
func TestSysUserUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/user/update", XTestSysUserUpdate, t, XTestSysUserUpdateM)
}

// TestSysUserBatchUpdate defined TODO
func TestSysUserBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/user/batch_update", XTestSysUserBatchUpdate, t, XTestSysUserBatchUpdateM)
}

// TestSysUserPage defined TODO
func TestSysUserPage(t *testing.T) {
	x.Handle("GET", "/api/sys/user/page", XTestSysUserPage, t, XTestSysUserPageM)
}

// TestSysUserGet defined TODO
func TestSysUserGet(t *testing.T) {
	x.Handle("GET", "/api/sys/user/get", XTestSysUserGet, t, XTestSysUserGetM)
}

// TestSysUserLogout defined TODO
func TestSysUserLogout(t *testing.T) {
	x.Handle("GET", "/api/sys/user/logout", XTestSysUserLogout, t, XTestSysUserLogoutM)
}

// TestSysUserTemplateAdd defined TODO
func TestSysUserTemplateAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/user/template/add", XTestSysUserTemplateAdd, t, XTestSysUserTemplateAddM)
}

// TestSysUserTemplateBatchAdd defined TODO
func TestSysUserTemplateBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/user/template/batch_add", XTestSysUserTemplateBatchAdd, t, XTestSysUserTemplateBatchAddM)
}

// TestSysUserTemplateDel defined TODO
func TestSysUserTemplateDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/user/template/del", XTestSysUserTemplateDel, t, XTestSysUserTemplateDelM)
}

// TestSysUserTemplateBatchDel defined TODO
func TestSysUserTemplateBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/user/template/batch_del", XTestSysUserTemplateBatchDel, t, XTestSysUserTemplateBatchDelM)
}

// TestSysUserTemplateUpdate defined TODO
func TestSysUserTemplateUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/user/template/update", XTestSysUserTemplateUpdate, t, XTestSysUserTemplateUpdateM)
}

// TestSysUserTemplateBatchUpdate defined TODO
func TestSysUserTemplateBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/user/template/batch_update", XTestSysUserTemplateBatchUpdate, t, XTestSysUserTemplateBatchUpdateM)
}

// TestSysUserTemplatePage defined TODO
func TestSysUserTemplatePage(t *testing.T) {
	x.Handle("GET", "/api/sys/user/template/page", XTestSysUserTemplatePage, t, XTestSysUserTemplatePageM)
}

// TestSysUserTemplateGet defined TODO
func TestSysUserTemplateGet(t *testing.T) {
	x.Handle("GET", "/api/sys/user/template/get", XTestSysUserTemplateGet, t, XTestSysUserTemplateGetM)
}

// TestSysUserTemplateDetailAdd defined TODO
func TestSysUserTemplateDetailAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/user/template/detail/add", XTestSysUserTemplateDetailAdd, t, XTestSysUserTemplateDetailAddM)
}

// TestSysUserTemplateDetailBatchAdd defined TODO
func TestSysUserTemplateDetailBatchAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/user/template/detail/batch_add", XTestSysUserTemplateDetailBatchAdd, t, XTestSysUserTemplateDetailBatchAddM)
}

// TestSysUserTemplateDetailDel defined TODO
func TestSysUserTemplateDetailDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/user/template/detail/del", XTestSysUserTemplateDetailDel, t, XTestSysUserTemplateDetailDelM)
}

// TestSysUserTemplateDetailBatchDel defined TODO
func TestSysUserTemplateDetailBatchDel(t *testing.T) {
	x.Handle("DELETE", "/api/sys/user/template/detail/batch_del", XTestSysUserTemplateDetailBatchDel, t, XTestSysUserTemplateDetailBatchDelM)
}

// TestSysUserTemplateDetailUpdate defined TODO
func TestSysUserTemplateDetailUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/user/template/detail/update", XTestSysUserTemplateDetailUpdate, t, XTestSysUserTemplateDetailUpdateM)
}

// TestSysUserTemplateDetailBatchUpdate defined TODO
func TestSysUserTemplateDetailBatchUpdate(t *testing.T) {
	x.Handle("PUT", "/api/sys/user/template/detail/batch_update", XTestSysUserTemplateDetailBatchUpdate, t, XTestSysUserTemplateDetailBatchUpdateM)
}

// TestSysUserTemplateDetailPage defined TODO
func TestSysUserTemplateDetailPage(t *testing.T) {
	x.Handle("GET", "/api/sys/user/template/detail/page", XTestSysUserTemplateDetailPage, t, XTestSysUserTemplateDetailPageM)
}

// TestSysUserTemplateDetailGet defined TODO
func TestSysUserTemplateDetailGet(t *testing.T) {
	x.Handle("GET", "/api/sys/user/template/detail/get", XTestSysUserTemplateDetailGet, t, XTestSysUserTemplateDetailGetM)
}

// TestSysWechatOauth2 defined TODO
func TestSysWechatOauth2(t *testing.T) {
	x.Handle("GET", "/api/sys/wechat/oauth2", XTestSysWechatOauth2, t, XTestSysWechatOauth2M)
}

// TestSysWorkerAdd defined TODO
func TestSysWorkerAdd(t *testing.T) {
	x.Handle("POST", "/api/sys/worker/add", XTestSysWorkerAdd, t, XTestSysWorkerAddM)
}

// TestSysWorkerGet defined TODO
func TestSysWorkerGet(t *testing.T) {
	x.Handle("GET", "/api/sys/worker/get", XTestSysWorkerGet, t, XTestSysWorkerGetM)
}
