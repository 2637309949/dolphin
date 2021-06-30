package main

import (
	"encoding/json"

	// "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

// XTestSysUserLogin defined TODO
var XTestSysUserLogin, XTestSysUserLoginM = func(ctx *Context) {
	ret := struct {
		Code int `json:"code"`
		Data struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}{}
	json.Unmarshal(ctx.Body.Bytes(), &ret)
	SetToken(ret.Data.AccessToken)
}, M{"domain": "localhost", "name": "admin", "password": "admin"}

// XTestSysAppFunAdd, XTestSysAppFunAddM defined TODO
var XTestSysAppFunAdd, XTestSysAppFunAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunBatchAdd, XTestSysAppFunBatchAddM defined TODO
var XTestSysAppFunBatchAdd, XTestSysAppFunBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunDel, XTestSysAppFunDelM defined TODO
var XTestSysAppFunDel, XTestSysAppFunDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunBatchDel, XTestSysAppFunBatchDelM defined TODO
var XTestSysAppFunBatchDel, XTestSysAppFunBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunUpdate, XTestSysAppFunUpdateM defined TODO
var XTestSysAppFunUpdate, XTestSysAppFunUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunBatchUpdate, XTestSysAppFunBatchUpdateM defined TODO
var XTestSysAppFunBatchUpdate, XTestSysAppFunBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunPage, XTestSysAppFunPageM defined TODO
var XTestSysAppFunPage, XTestSysAppFunPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunTree, XTestSysAppFunTreeM defined TODO
var XTestSysAppFunTree, XTestSysAppFunTreeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunTree = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAppFunGet, XTestSysAppFunGetM defined TODO
var XTestSysAppFunGet, XTestSysAppFunGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaAdd, XTestSysAreaAddM defined TODO
var XTestSysAreaAdd, XTestSysAreaAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaBatchAdd, XTestSysAreaBatchAddM defined TODO
var XTestSysAreaBatchAdd, XTestSysAreaBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaDel, XTestSysAreaDelM defined TODO
var XTestSysAreaDel, XTestSysAreaDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaBatchDel, XTestSysAreaBatchDelM defined TODO
var XTestSysAreaBatchDel, XTestSysAreaBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaUpdate, XTestSysAreaUpdateM defined TODO
var XTestSysAreaUpdate, XTestSysAreaUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaBatchUpdate, XTestSysAreaBatchUpdateM defined TODO
var XTestSysAreaBatchUpdate, XTestSysAreaBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaPage, XTestSysAreaPageM defined TODO
var XTestSysAreaPage, XTestSysAreaPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAreaGet, XTestSysAreaGetM defined TODO
var XTestSysAreaGet, XTestSysAreaGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentAdd, XTestSysAttachmentAddM defined TODO
var XTestSysAttachmentAdd, XTestSysAttachmentAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentBatchAdd, XTestSysAttachmentBatchAddM defined TODO
var XTestSysAttachmentBatchAdd, XTestSysAttachmentBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentUpload, XTestSysAttachmentUploadM defined TODO
var XTestSysAttachmentUpload, XTestSysAttachmentUploadM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpload = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpload = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentExport, XTestSysAttachmentExportM defined TODO
var XTestSysAttachmentExport, XTestSysAttachmentExportM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentExport = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentExport = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentDel, XTestSysAttachmentDelM defined TODO
var XTestSysAttachmentDel, XTestSysAttachmentDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentBatchDel, XTestSysAttachmentBatchDelM defined TODO
var XTestSysAttachmentBatchDel, XTestSysAttachmentBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentUpdate, XTestSysAttachmentUpdateM defined TODO
var XTestSysAttachmentUpdate, XTestSysAttachmentUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentBatchUpdate, XTestSysAttachmentBatchUpdateM defined TODO
var XTestSysAttachmentBatchUpdate, XTestSysAttachmentBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentPage, XTestSysAttachmentPageM defined TODO
var XTestSysAttachmentPage, XTestSysAttachmentPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysAttachmentGet, XTestSysAttachmentGetM defined TODO
var XTestSysAttachmentGet, XTestSysAttachmentGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasLogin, XTestSysCasLoginM defined TODO
var XTestSysCasLogin, XTestSysCasLoginM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogin = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogin = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasLogout, XTestSysCasLogoutM defined TODO
var XTestSysCasLogout, XTestSysCasLogoutM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogout = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogout = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasAffirm, XTestSysCasAffirmM defined TODO
var XTestSysCasAffirm, XTestSysCasAffirmM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAffirm = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAffirm = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasAuthorize, XTestSysCasAuthorizeM defined TODO
var XTestSysCasAuthorize, XTestSysCasAuthorizeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAuthorize = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAuthorize = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasToken, XTestSysCasTokenM defined TODO
var XTestSysCasToken, XTestSysCasTokenM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasToken = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasToken = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasURL, XTestSysCasURLM defined TODO
var XTestSysCasURL, XTestSysCasURLM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasURL = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasURL = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasOauth2, XTestSysCasOauth2M defined TODO
var XTestSysCasOauth2, XTestSysCasOauth2M = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasOauth2 = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasOauth2 = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasRefresh, XTestSysCasRefreshM defined TODO
var XTestSysCasRefresh, XTestSysCasRefreshM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasRefresh = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasRefresh = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasCheck, XTestSysCasCheckM defined TODO
var XTestSysCasCheck, XTestSysCasCheckM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasCheck = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasCheck = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasProfile, XTestSysCasProfileM defined TODO
var XTestSysCasProfile, XTestSysCasProfileM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasProfile = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasProfile = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCasQrcode, XTestSysCasQrcodeM defined TODO
var XTestSysCasQrcode, XTestSysCasQrcodeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasQrcode = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasQrcode = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientAdd, XTestSysClientAddM defined TODO
var XTestSysClientAdd, XTestSysClientAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientBatchAdd, XTestSysClientBatchAddM defined TODO
var XTestSysClientBatchAdd, XTestSysClientBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientDel, XTestSysClientDelM defined TODO
var XTestSysClientDel, XTestSysClientDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientBatchDel, XTestSysClientBatchDelM defined TODO
var XTestSysClientBatchDel, XTestSysClientBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientUpdate, XTestSysClientUpdateM defined TODO
var XTestSysClientUpdate, XTestSysClientUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientBatchUpdate, XTestSysClientBatchUpdateM defined TODO
var XTestSysClientBatchUpdate, XTestSysClientBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientPage, XTestSysClientPageM defined TODO
var XTestSysClientPage, XTestSysClientPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysClientGet, XTestSysClientGetM defined TODO
var XTestSysClientGet, XTestSysClientGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentAdd, XTestSysCommentAddM defined TODO
var XTestSysCommentAdd, XTestSysCommentAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentBatchAdd, XTestSysCommentBatchAddM defined TODO
var XTestSysCommentBatchAdd, XTestSysCommentBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentDel, XTestSysCommentDelM defined TODO
var XTestSysCommentDel, XTestSysCommentDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentBatchDel, XTestSysCommentBatchDelM defined TODO
var XTestSysCommentBatchDel, XTestSysCommentBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentUpdate, XTestSysCommentUpdateM defined TODO
var XTestSysCommentUpdate, XTestSysCommentUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentBatchUpdate, XTestSysCommentBatchUpdateM defined TODO
var XTestSysCommentBatchUpdate, XTestSysCommentBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentPage, XTestSysCommentPageM defined TODO
var XTestSysCommentPage, XTestSysCommentPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysCommentGet, XTestSysCommentGetM defined TODO
var XTestSysCommentGet, XTestSysCommentGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionAdd, XTestSysDataPermissionAddM defined TODO
var XTestSysDataPermissionAdd, XTestSysDataPermissionAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionBatchAdd, XTestSysDataPermissionBatchAddM defined TODO
var XTestSysDataPermissionBatchAdd, XTestSysDataPermissionBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionDel, XTestSysDataPermissionDelM defined TODO
var XTestSysDataPermissionDel, XTestSysDataPermissionDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionBatchDel, XTestSysDataPermissionBatchDelM defined TODO
var XTestSysDataPermissionBatchDel, XTestSysDataPermissionBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionUpdate, XTestSysDataPermissionUpdateM defined TODO
var XTestSysDataPermissionUpdate, XTestSysDataPermissionUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionBatchUpdate, XTestSysDataPermissionBatchUpdateM defined TODO
var XTestSysDataPermissionBatchUpdate, XTestSysDataPermissionBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionPage, XTestSysDataPermissionPageM defined TODO
var XTestSysDataPermissionPage, XTestSysDataPermissionPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDataPermissionGet, XTestSysDataPermissionGetM defined TODO
var XTestSysDataPermissionGet, XTestSysDataPermissionGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugPprof, XTestDebugPprofM defined TODO
var XTestDebugPprof, XTestDebugPprofM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugPprof = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugPprof = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugHeap, XTestDebugHeapM defined TODO
var XTestDebugHeap, XTestDebugHeapM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugHeap = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugHeap = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugGoroutine, XTestDebugGoroutineM defined TODO
var XTestDebugGoroutine, XTestDebugGoroutineM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugGoroutine = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugGoroutine = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugAllocs, XTestDebugAllocsM defined TODO
var XTestDebugAllocs, XTestDebugAllocsM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugAllocs = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugAllocs = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugBlock, XTestDebugBlockM defined TODO
var XTestDebugBlock, XTestDebugBlockM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugBlock = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugBlock = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugThreadcreate, XTestDebugThreadcreateM defined TODO
var XTestDebugThreadcreate, XTestDebugThreadcreateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugThreadcreate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugThreadcreate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugCmdline, XTestDebugCmdlineM defined TODO
var XTestDebugCmdline, XTestDebugCmdlineM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugCmdline = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugCmdline = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugProfile, XTestDebugProfileM defined TODO
var XTestDebugProfile, XTestDebugProfileM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugProfile = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugProfile = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugSymbol, XTestDebugSymbolM defined TODO
var XTestDebugSymbol, XTestDebugSymbolM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugSymbol = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugSymbol = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugTrace, XTestDebugTraceM defined TODO
var XTestDebugTrace, XTestDebugTraceM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugTrace = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugTrace = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestDebugMutex, XTestDebugMutexM defined TODO
var XTestDebugMutex, XTestDebugMutexM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugMutex = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugMutex = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDingtalkOauth2, XTestSysDingtalkOauth2M defined TODO
var XTestSysDingtalkOauth2, XTestSysDingtalkOauth2M = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDingtalkOauth2 = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDingtalkOauth2 = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainAdd, XTestSysDomainAddM defined TODO
var XTestSysDomainAdd, XTestSysDomainAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainBatchAdd, XTestSysDomainBatchAddM defined TODO
var XTestSysDomainBatchAdd, XTestSysDomainBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainDel, XTestSysDomainDelM defined TODO
var XTestSysDomainDel, XTestSysDomainDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainBatchDel, XTestSysDomainBatchDelM defined TODO
var XTestSysDomainBatchDel, XTestSysDomainBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainUpdate, XTestSysDomainUpdateM defined TODO
var XTestSysDomainUpdate, XTestSysDomainUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainBatchUpdate, XTestSysDomainBatchUpdateM defined TODO
var XTestSysDomainBatchUpdate, XTestSysDomainBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainPage, XTestSysDomainPageM defined TODO
var XTestSysDomainPage, XTestSysDomainPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysDomainGet, XTestSysDomainGetM defined TODO
var XTestSysDomainGet, XTestSysDomainGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuAdd, XTestSysMenuAddM defined TODO
var XTestSysMenuAdd, XTestSysMenuAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuBatchAdd, XTestSysMenuBatchAddM defined TODO
var XTestSysMenuBatchAdd, XTestSysMenuBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuDel, XTestSysMenuDelM defined TODO
var XTestSysMenuDel, XTestSysMenuDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuBatchDel, XTestSysMenuBatchDelM defined TODO
var XTestSysMenuBatchDel, XTestSysMenuBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuUpdate, XTestSysMenuUpdateM defined TODO
var XTestSysMenuUpdate, XTestSysMenuUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuBatchUpdate, XTestSysMenuBatchUpdateM defined TODO
var XTestSysMenuBatchUpdate, XTestSysMenuBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuSidebar, XTestSysMenuSidebarM defined TODO
var XTestSysMenuSidebar, XTestSysMenuSidebarM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuSidebar = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuSidebar = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuPage, XTestSysMenuPageM defined TODO
var XTestSysMenuPage, XTestSysMenuPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuTree, XTestSysMenuTreeM defined TODO
var XTestSysMenuTree, XTestSysMenuTreeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuTree = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysMenuGet, XTestSysMenuGetM defined TODO
var XTestSysMenuGet, XTestSysMenuGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationAdd, XTestSysNotificationAddM defined TODO
var XTestSysNotificationAdd, XTestSysNotificationAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationBatchAdd, XTestSysNotificationBatchAddM defined TODO
var XTestSysNotificationBatchAdd, XTestSysNotificationBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationDel, XTestSysNotificationDelM defined TODO
var XTestSysNotificationDel, XTestSysNotificationDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationBatchDel, XTestSysNotificationBatchDelM defined TODO
var XTestSysNotificationBatchDel, XTestSysNotificationBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationUpdate, XTestSysNotificationUpdateM defined TODO
var XTestSysNotificationUpdate, XTestSysNotificationUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationBatchUpdate, XTestSysNotificationBatchUpdateM defined TODO
var XTestSysNotificationBatchUpdate, XTestSysNotificationBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationPage, XTestSysNotificationPageM defined TODO
var XTestSysNotificationPage, XTestSysNotificationPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysNotificationGet, XTestSysNotificationGetM defined TODO
var XTestSysNotificationGet, XTestSysNotificationGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetAdd, XTestSysOptionsetAddM defined TODO
var XTestSysOptionsetAdd, XTestSysOptionsetAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetBatchAdd, XTestSysOptionsetBatchAddM defined TODO
var XTestSysOptionsetBatchAdd, XTestSysOptionsetBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetDel, XTestSysOptionsetDelM defined TODO
var XTestSysOptionsetDel, XTestSysOptionsetDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetBatchDel, XTestSysOptionsetBatchDelM defined TODO
var XTestSysOptionsetBatchDel, XTestSysOptionsetBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetUpdate, XTestSysOptionsetUpdateM defined TODO
var XTestSysOptionsetUpdate, XTestSysOptionsetUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetBatchUpdate, XTestSysOptionsetBatchUpdateM defined TODO
var XTestSysOptionsetBatchUpdate, XTestSysOptionsetBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetPage, XTestSysOptionsetPageM defined TODO
var XTestSysOptionsetPage, XTestSysOptionsetPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOptionsetGet, XTestSysOptionsetGetM defined TODO
var XTestSysOptionsetGet, XTestSysOptionsetGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgAdd, XTestSysOrgAddM defined TODO
var XTestSysOrgAdd, XTestSysOrgAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgBatchAdd, XTestSysOrgBatchAddM defined TODO
var XTestSysOrgBatchAdd, XTestSysOrgBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgDel, XTestSysOrgDelM defined TODO
var XTestSysOrgDel, XTestSysOrgDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgBatchDel, XTestSysOrgBatchDelM defined TODO
var XTestSysOrgBatchDel, XTestSysOrgBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgUpdate, XTestSysOrgUpdateM defined TODO
var XTestSysOrgUpdate, XTestSysOrgUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgBatchUpdate, XTestSysOrgBatchUpdateM defined TODO
var XTestSysOrgBatchUpdate, XTestSysOrgBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgPage, XTestSysOrgPageM defined TODO
var XTestSysOrgPage, XTestSysOrgPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgTree, XTestSysOrgTreeM defined TODO
var XTestSysOrgTree, XTestSysOrgTreeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgTree = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysOrgGet, XTestSysOrgGetM defined TODO
var XTestSysOrgGet, XTestSysOrgGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionAdd, XTestSysPermissionAddM defined TODO
var XTestSysPermissionAdd, XTestSysPermissionAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionBatchAdd, XTestSysPermissionBatchAddM defined TODO
var XTestSysPermissionBatchAdd, XTestSysPermissionBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionDel, XTestSysPermissionDelM defined TODO
var XTestSysPermissionDel, XTestSysPermissionDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionBatchDel, XTestSysPermissionBatchDelM defined TODO
var XTestSysPermissionBatchDel, XTestSysPermissionBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionUpdate, XTestSysPermissionUpdateM defined TODO
var XTestSysPermissionUpdate, XTestSysPermissionUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionBatchUpdate, XTestSysPermissionBatchUpdateM defined TODO
var XTestSysPermissionBatchUpdate, XTestSysPermissionBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionPage, XTestSysPermissionPageM defined TODO
var XTestSysPermissionPage, XTestSysPermissionPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysPermissionGet, XTestSysPermissionGetM defined TODO
var XTestSysPermissionGet, XTestSysPermissionGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleAdd, XTestSysRoleAddM defined TODO
var XTestSysRoleAdd, XTestSysRoleAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleBatchAdd, XTestSysRoleBatchAddM defined TODO
var XTestSysRoleBatchAdd, XTestSysRoleBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleDel, XTestSysRoleDelM defined TODO
var XTestSysRoleDel, XTestSysRoleDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleBatchDel, XTestSysRoleBatchDelM defined TODO
var XTestSysRoleBatchDel, XTestSysRoleBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleUpdate, XTestSysRoleUpdateM defined TODO
var XTestSysRoleUpdate, XTestSysRoleUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleBatchUpdate, XTestSysRoleBatchUpdateM defined TODO
var XTestSysRoleBatchUpdate, XTestSysRoleBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRolePage, XTestSysRolePageM defined TODO
var XTestSysRolePage, XTestSysRolePageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRolePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRolePage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleRoleMenuTree, XTestSysRoleRoleMenuTreeM defined TODO
var XTestSysRoleRoleMenuTree, XTestSysRoleRoleMenuTreeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleMenuTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleMenuTree = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleRoleAppFunTree, XTestSysRoleRoleAppFunTreeM defined TODO
var XTestSysRoleRoleAppFunTree, XTestSysRoleRoleAppFunTreeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleAppFunTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleAppFunTree = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleGet, XTestSysRoleGetM defined TODO
var XTestSysRoleGet, XTestSysRoleGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuAdd, XTestSysRoleMenuAddM defined TODO
var XTestSysRoleMenuAdd, XTestSysRoleMenuAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuBatchAdd, XTestSysRoleMenuBatchAddM defined TODO
var XTestSysRoleMenuBatchAdd, XTestSysRoleMenuBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuDel, XTestSysRoleMenuDelM defined TODO
var XTestSysRoleMenuDel, XTestSysRoleMenuDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuBatchDel, XTestSysRoleMenuBatchDelM defined TODO
var XTestSysRoleMenuBatchDel, XTestSysRoleMenuBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuUpdate, XTestSysRoleMenuUpdateM defined TODO
var XTestSysRoleMenuUpdate, XTestSysRoleMenuUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuBatchUpdate, XTestSysRoleMenuBatchUpdateM defined TODO
var XTestSysRoleMenuBatchUpdate, XTestSysRoleMenuBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuPage, XTestSysRoleMenuPageM defined TODO
var XTestSysRoleMenuPage, XTestSysRoleMenuPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysRoleMenuGet, XTestSysRoleMenuGetM defined TODO
var XTestSysRoleMenuGet, XTestSysRoleMenuGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleAdd, XTestSysScheduleAddM defined TODO
var XTestSysScheduleAdd, XTestSysScheduleAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleBatchAdd, XTestSysScheduleBatchAddM defined TODO
var XTestSysScheduleBatchAdd, XTestSysScheduleBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleDel, XTestSysScheduleDelM defined TODO
var XTestSysScheduleDel, XTestSysScheduleDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleBatchDel, XTestSysScheduleBatchDelM defined TODO
var XTestSysScheduleBatchDel, XTestSysScheduleBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleUpdate, XTestSysScheduleUpdateM defined TODO
var XTestSysScheduleUpdate, XTestSysScheduleUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleBatchUpdate, XTestSysScheduleBatchUpdateM defined TODO
var XTestSysScheduleBatchUpdate, XTestSysScheduleBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSchedulePage, XTestSysSchedulePageM defined TODO
var XTestSysSchedulePage, XTestSysSchedulePageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulePage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleGet, XTestSysScheduleGetM defined TODO
var XTestSysScheduleGet, XTestSysScheduleGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysScheduleHistoryPage, XTestSysScheduleHistoryPageM defined TODO
var XTestSysScheduleHistoryPage, XTestSysScheduleHistoryPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleHistoryPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleHistoryPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSchedulingAdd, XTestSysSchedulingAddM defined TODO
var XTestSysSchedulingAdd, XTestSysSchedulingAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSchedulingDel, XTestSysSchedulingDelM defined TODO
var XTestSysSchedulingDel, XTestSysSchedulingDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSchedulingUpdate, XTestSysSchedulingUpdateM defined TODO
var XTestSysSchedulingUpdate, XTestSysSchedulingUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSchedulingPage, XTestSysSchedulingPageM defined TODO
var XTestSysSchedulingPage, XTestSysSchedulingPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSchedulingGet, XTestSysSchedulingGetM defined TODO
var XTestSysSchedulingGet, XTestSysSchedulingGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingAdd, XTestSysSettingAddM defined TODO
var XTestSysSettingAdd, XTestSysSettingAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingBatchAdd, XTestSysSettingBatchAddM defined TODO
var XTestSysSettingBatchAdd, XTestSysSettingBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingDel, XTestSysSettingDelM defined TODO
var XTestSysSettingDel, XTestSysSettingDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingBatchDel, XTestSysSettingBatchDelM defined TODO
var XTestSysSettingBatchDel, XTestSysSettingBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingUpdate, XTestSysSettingUpdateM defined TODO
var XTestSysSettingUpdate, XTestSysSettingUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingBatchUpdate, XTestSysSettingBatchUpdateM defined TODO
var XTestSysSettingBatchUpdate, XTestSysSettingBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingPage, XTestSysSettingPageM defined TODO
var XTestSysSettingPage, XTestSysSettingPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysSettingGet, XTestSysSettingGetM defined TODO
var XTestSysSettingGet, XTestSysSettingGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableAdd, XTestSysTableAddM defined TODO
var XTestSysTableAdd, XTestSysTableAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableBatchAdd, XTestSysTableBatchAddM defined TODO
var XTestSysTableBatchAdd, XTestSysTableBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableDel, XTestSysTableDelM defined TODO
var XTestSysTableDel, XTestSysTableDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableBatchDel, XTestSysTableBatchDelM defined TODO
var XTestSysTableBatchDel, XTestSysTableBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableUpdate, XTestSysTableUpdateM defined TODO
var XTestSysTableUpdate, XTestSysTableUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableBatchUpdate, XTestSysTableBatchUpdateM defined TODO
var XTestSysTableBatchUpdate, XTestSysTableBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTablePage, XTestSysTablePageM defined TODO
var XTestSysTablePage, XTestSysTablePageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTablePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTablePage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableGet, XTestSysTableGetM defined TODO
var XTestSysTableGet, XTestSysTableGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnAdd, XTestSysTableColumnAddM defined TODO
var XTestSysTableColumnAdd, XTestSysTableColumnAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnBatchAdd, XTestSysTableColumnBatchAddM defined TODO
var XTestSysTableColumnBatchAdd, XTestSysTableColumnBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnDel, XTestSysTableColumnDelM defined TODO
var XTestSysTableColumnDel, XTestSysTableColumnDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnBatchDel, XTestSysTableColumnBatchDelM defined TODO
var XTestSysTableColumnBatchDel, XTestSysTableColumnBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnUpdate, XTestSysTableColumnUpdateM defined TODO
var XTestSysTableColumnUpdate, XTestSysTableColumnUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnBatchUpdate, XTestSysTableColumnBatchUpdateM defined TODO
var XTestSysTableColumnBatchUpdate, XTestSysTableColumnBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnPage, XTestSysTableColumnPageM defined TODO
var XTestSysTableColumnPage, XTestSysTableColumnPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTableColumnGet, XTestSysTableColumnGetM defined TODO
var XTestSysTableColumnGet, XTestSysTableColumnGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagAdd, XTestSysTagAddM defined TODO
var XTestSysTagAdd, XTestSysTagAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagBatchAdd, XTestSysTagBatchAddM defined TODO
var XTestSysTagBatchAdd, XTestSysTagBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagDel, XTestSysTagDelM defined TODO
var XTestSysTagDel, XTestSysTagDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagBatchDel, XTestSysTagBatchDelM defined TODO
var XTestSysTagBatchDel, XTestSysTagBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagUpdate, XTestSysTagUpdateM defined TODO
var XTestSysTagUpdate, XTestSysTagUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagBatchUpdate, XTestSysTagBatchUpdateM defined TODO
var XTestSysTagBatchUpdate, XTestSysTagBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagPage, XTestSysTagPageM defined TODO
var XTestSysTagPage, XTestSysTagPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGet, XTestSysTagGetM defined TODO
var XTestSysTagGet, XTestSysTagGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupAdd, XTestSysTagGroupAddM defined TODO
var XTestSysTagGroupAdd, XTestSysTagGroupAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupBatchAdd, XTestSysTagGroupBatchAddM defined TODO
var XTestSysTagGroupBatchAdd, XTestSysTagGroupBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupDel, XTestSysTagGroupDelM defined TODO
var XTestSysTagGroupDel, XTestSysTagGroupDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupBatchDel, XTestSysTagGroupBatchDelM defined TODO
var XTestSysTagGroupBatchDel, XTestSysTagGroupBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupUpdate, XTestSysTagGroupUpdateM defined TODO
var XTestSysTagGroupUpdate, XTestSysTagGroupUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupBatchUpdate, XTestSysTagGroupBatchUpdateM defined TODO
var XTestSysTagGroupBatchUpdate, XTestSysTagGroupBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupPage, XTestSysTagGroupPageM defined TODO
var XTestSysTagGroupPage, XTestSysTagGroupPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTagGroupGet, XTestSysTagGroupGetM defined TODO
var XTestSysTagGroupGet, XTestSysTagGroupGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTrackerPage, XTestSysTrackerPageM defined TODO
var XTestSysTrackerPage, XTestSysTrackerPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysTrackerGet, XTestSysTrackerGetM defined TODO
var XTestSysTrackerGet, XTestSysTrackerGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserAdd, XTestSysUserAddM defined TODO
var XTestSysUserAdd, XTestSysUserAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserBatchAdd, XTestSysUserBatchAddM defined TODO
var XTestSysUserBatchAdd, XTestSysUserBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserDel, XTestSysUserDelM defined TODO
var XTestSysUserDel, XTestSysUserDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserBatchDel, XTestSysUserBatchDelM defined TODO
var XTestSysUserBatchDel, XTestSysUserBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserUpdate, XTestSysUserUpdateM defined TODO
var XTestSysUserUpdate, XTestSysUserUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserBatchUpdate, XTestSysUserBatchUpdateM defined TODO
var XTestSysUserBatchUpdate, XTestSysUserBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserPage, XTestSysUserPageM defined TODO
var XTestSysUserPage, XTestSysUserPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserGet, XTestSysUserGetM defined TODO
var XTestSysUserGet, XTestSysUserGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserLogout, XTestSysUserLogoutM defined TODO
var XTestSysUserLogout, XTestSysUserLogoutM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserLogout = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserLogout = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateAdd, XTestSysUserTemplateAddM defined TODO
var XTestSysUserTemplateAdd, XTestSysUserTemplateAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateBatchAdd, XTestSysUserTemplateBatchAddM defined TODO
var XTestSysUserTemplateBatchAdd, XTestSysUserTemplateBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDel, XTestSysUserTemplateDelM defined TODO
var XTestSysUserTemplateDel, XTestSysUserTemplateDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateBatchDel, XTestSysUserTemplateBatchDelM defined TODO
var XTestSysUserTemplateBatchDel, XTestSysUserTemplateBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateUpdate, XTestSysUserTemplateUpdateM defined TODO
var XTestSysUserTemplateUpdate, XTestSysUserTemplateUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateBatchUpdate, XTestSysUserTemplateBatchUpdateM defined TODO
var XTestSysUserTemplateBatchUpdate, XTestSysUserTemplateBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplatePage, XTestSysUserTemplatePageM defined TODO
var XTestSysUserTemplatePage, XTestSysUserTemplatePageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplatePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplatePage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateGet, XTestSysUserTemplateGetM defined TODO
var XTestSysUserTemplateGet, XTestSysUserTemplateGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailAdd, XTestSysUserTemplateDetailAddM defined TODO
var XTestSysUserTemplateDetailAdd, XTestSysUserTemplateDetailAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailBatchAdd, XTestSysUserTemplateDetailBatchAddM defined TODO
var XTestSysUserTemplateDetailBatchAdd, XTestSysUserTemplateDetailBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailDel, XTestSysUserTemplateDetailDelM defined TODO
var XTestSysUserTemplateDetailDel, XTestSysUserTemplateDetailDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailBatchDel, XTestSysUserTemplateDetailBatchDelM defined TODO
var XTestSysUserTemplateDetailBatchDel, XTestSysUserTemplateDetailBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailUpdate, XTestSysUserTemplateDetailUpdateM defined TODO
var XTestSysUserTemplateDetailUpdate, XTestSysUserTemplateDetailUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailBatchUpdate, XTestSysUserTemplateDetailBatchUpdateM defined TODO
var XTestSysUserTemplateDetailBatchUpdate, XTestSysUserTemplateDetailBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailPage, XTestSysUserTemplateDetailPageM defined TODO
var XTestSysUserTemplateDetailPage, XTestSysUserTemplateDetailPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysUserTemplateDetailGet, XTestSysUserTemplateDetailGetM defined TODO
var XTestSysUserTemplateDetailGet, XTestSysUserTemplateDetailGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysWechatOauth2, XTestSysWechatOauth2M defined TODO
var XTestSysWechatOauth2, XTestSysWechatOauth2M = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysWechatOauth2 = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysWechatOauth2 = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysWorkerAdd, XTestSysWorkerAddM defined TODO
var XTestSysWorkerAdd, XTestSysWorkerAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSysWorkerGet, XTestSysWorkerGetM defined TODO
var XTestSysWorkerGet, XTestSysWorkerGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerGet = %v want %v", ret.Code, 200)
	}
}, M{}
