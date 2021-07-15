package main

import (
	"encoding/json"

	// "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

// XTestSysUserLogin defined TODO
var XTestSysUserLogin, XTestSysUserLoginRequest = func(ctx *Context) {
	ret := struct {
		Code int `json:"code"`
		Data struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}{}
	json.Unmarshal(ctx.Body.Bytes(), &ret)
	SetToken(ret.Data.AccessToken)
}, Payload{"domain": "localhost", "name": "admin", "password": "admin"}

// XTestSysAppFunAdd, XTestSysAppFunAddM defined TODO
var XTestSysAppFunAdd, XTestSysAppFunAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunBatchAdd, XTestSysAppFunBatchAddM defined TODO
var XTestSysAppFunBatchAdd, XTestSysAppFunBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunDel, XTestSysAppFunDelM defined TODO
var XTestSysAppFunDel, XTestSysAppFunDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunBatchDel, XTestSysAppFunBatchDelM defined TODO
var XTestSysAppFunBatchDel, XTestSysAppFunBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunUpdate, XTestSysAppFunUpdateM defined TODO
var XTestSysAppFunUpdate, XTestSysAppFunUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunBatchUpdate, XTestSysAppFunBatchUpdateM defined TODO
var XTestSysAppFunBatchUpdate, XTestSysAppFunBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunPage, XTestSysAppFunPageM defined TODO
var XTestSysAppFunPage, XTestSysAppFunPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunPage = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunTree, XTestSysAppFunTreeM defined TODO
var XTestSysAppFunTree, XTestSysAppFunTreeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunTree = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAppFunGet, XTestSysAppFunGetM defined TODO
var XTestSysAppFunGet, XTestSysAppFunGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAppFunGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaAdd, XTestSysAreaAddM defined TODO
var XTestSysAreaAdd, XTestSysAreaAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaBatchAdd, XTestSysAreaBatchAddM defined TODO
var XTestSysAreaBatchAdd, XTestSysAreaBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaDel, XTestSysAreaDelM defined TODO
var XTestSysAreaDel, XTestSysAreaDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaBatchDel, XTestSysAreaBatchDelM defined TODO
var XTestSysAreaBatchDel, XTestSysAreaBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaUpdate, XTestSysAreaUpdateM defined TODO
var XTestSysAreaUpdate, XTestSysAreaUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaBatchUpdate, XTestSysAreaBatchUpdateM defined TODO
var XTestSysAreaBatchUpdate, XTestSysAreaBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaPage, XTestSysAreaPageM defined TODO
var XTestSysAreaPage, XTestSysAreaPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAreaGet, XTestSysAreaGetM defined TODO
var XTestSysAreaGet, XTestSysAreaGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAreaGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentAdd, XTestSysAttachmentAddM defined TODO
var XTestSysAttachmentAdd, XTestSysAttachmentAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentBatchAdd, XTestSysAttachmentBatchAddM defined TODO
var XTestSysAttachmentBatchAdd, XTestSysAttachmentBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentUpload, XTestSysAttachmentUploadM defined TODO
var XTestSysAttachmentUpload, XTestSysAttachmentUploadRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpload = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpload = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentExport, XTestSysAttachmentExportM defined TODO
var XTestSysAttachmentExport, XTestSysAttachmentExportRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentExport = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentExport = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentDel, XTestSysAttachmentDelM defined TODO
var XTestSysAttachmentDel, XTestSysAttachmentDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentBatchDel, XTestSysAttachmentBatchDelM defined TODO
var XTestSysAttachmentBatchDel, XTestSysAttachmentBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentUpdate, XTestSysAttachmentUpdateM defined TODO
var XTestSysAttachmentUpdate, XTestSysAttachmentUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentBatchUpdate, XTestSysAttachmentBatchUpdateM defined TODO
var XTestSysAttachmentBatchUpdate, XTestSysAttachmentBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentPage, XTestSysAttachmentPageM defined TODO
var XTestSysAttachmentPage, XTestSysAttachmentPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysAttachmentGet, XTestSysAttachmentGetM defined TODO
var XTestSysAttachmentGet, XTestSysAttachmentGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysAttachmentGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasLogin, XTestSysCasLoginM defined TODO
var XTestSysCasLogin, XTestSysCasLoginRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogin = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogin = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasLogout, XTestSysCasLogoutM defined TODO
var XTestSysCasLogout, XTestSysCasLogoutRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogout = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasLogout = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasAffirm, XTestSysCasAffirmM defined TODO
var XTestSysCasAffirm, XTestSysCasAffirmRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAffirRequest = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAffirRequest = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasAuthorize, XTestSysCasAuthorizeM defined TODO
var XTestSysCasAuthorize, XTestSysCasAuthorizeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAuthorize = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasAuthorize = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasToken, XTestSysCasTokenM defined TODO
var XTestSysCasToken, XTestSysCasTokenRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasToken = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasToken = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasURL, XTestSysCasURLM defined TODO
var XTestSysCasURL, XTestSysCasURLRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasURL = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasURL = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasOauth2, XTestSysCasOauth2M defined TODO
var XTestSysCasOauth2, XTestSysCasOauth2Request = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasOauth2 = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasOauth2 = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasRefresh, XTestSysCasRefreshM defined TODO
var XTestSysCasRefresh, XTestSysCasRefreshRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasRefresh = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasRefresh = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasCheck, XTestSysCasCheckM defined TODO
var XTestSysCasCheck, XTestSysCasCheckRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasCheck = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasCheck = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasProfile, XTestSysCasProfileM defined TODO
var XTestSysCasProfile, XTestSysCasProfileRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasProfile = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasProfile = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCasQrcode, XTestSysCasQrcodeM defined TODO
var XTestSysCasQrcode, XTestSysCasQrcodeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasQrcode = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCasQrcode = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientAdd, XTestSysClientAddM defined TODO
var XTestSysClientAdd, XTestSysClientAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientBatchAdd, XTestSysClientBatchAddM defined TODO
var XTestSysClientBatchAdd, XTestSysClientBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientDel, XTestSysClientDelM defined TODO
var XTestSysClientDel, XTestSysClientDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientBatchDel, XTestSysClientBatchDelM defined TODO
var XTestSysClientBatchDel, XTestSysClientBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientUpdate, XTestSysClientUpdateM defined TODO
var XTestSysClientUpdate, XTestSysClientUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientBatchUpdate, XTestSysClientBatchUpdateM defined TODO
var XTestSysClientBatchUpdate, XTestSysClientBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientPage, XTestSysClientPageM defined TODO
var XTestSysClientPage, XTestSysClientPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysClientGet, XTestSysClientGetM defined TODO
var XTestSysClientGet, XTestSysClientGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysClientGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentAdd, XTestSysCommentAddM defined TODO
var XTestSysCommentAdd, XTestSysCommentAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentBatchAdd, XTestSysCommentBatchAddM defined TODO
var XTestSysCommentBatchAdd, XTestSysCommentBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentDel, XTestSysCommentDelM defined TODO
var XTestSysCommentDel, XTestSysCommentDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentBatchDel, XTestSysCommentBatchDelM defined TODO
var XTestSysCommentBatchDel, XTestSysCommentBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentUpdate, XTestSysCommentUpdateM defined TODO
var XTestSysCommentUpdate, XTestSysCommentUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentBatchUpdate, XTestSysCommentBatchUpdateM defined TODO
var XTestSysCommentBatchUpdate, XTestSysCommentBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentPage, XTestSysCommentPageM defined TODO
var XTestSysCommentPage, XTestSysCommentPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysCommentGet, XTestSysCommentGetM defined TODO
var XTestSysCommentGet, XTestSysCommentGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysCommentGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionAdd, XTestSysDataPermissionAddM defined TODO
var XTestSysDataPermissionAdd, XTestSysDataPermissionAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionBatchAdd, XTestSysDataPermissionBatchAddM defined TODO
var XTestSysDataPermissionBatchAdd, XTestSysDataPermissionBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionDel, XTestSysDataPermissionDelM defined TODO
var XTestSysDataPermissionDel, XTestSysDataPermissionDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionBatchDel, XTestSysDataPermissionBatchDelM defined TODO
var XTestSysDataPermissionBatchDel, XTestSysDataPermissionBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionUpdate, XTestSysDataPermissionUpdateM defined TODO
var XTestSysDataPermissionUpdate, XTestSysDataPermissionUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionBatchUpdate, XTestSysDataPermissionBatchUpdateM defined TODO
var XTestSysDataPermissionBatchUpdate, XTestSysDataPermissionBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionPage, XTestSysDataPermissionPageM defined TODO
var XTestSysDataPermissionPage, XTestSysDataPermissionPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDataPermissionGet, XTestSysDataPermissionGetM defined TODO
var XTestSysDataPermissionGet, XTestSysDataPermissionGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDataPermissionGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugPprof, XTestDebugPprofM defined TODO
var XTestDebugPprof, XTestDebugPprofRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugPprof = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugPprof = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugHeap, XTestDebugHeapM defined TODO
var XTestDebugHeap, XTestDebugHeapRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugHeap = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugHeap = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugGoroutine, XTestDebugGoroutineM defined TODO
var XTestDebugGoroutine, XTestDebugGoroutineRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugGoroutine = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugGoroutine = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugAllocs, XTestDebugAllocsM defined TODO
var XTestDebugAllocs, XTestDebugAllocsRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugAllocs = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugAllocs = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugBlock, XTestDebugBlockM defined TODO
var XTestDebugBlock, XTestDebugBlockRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugBlock = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugBlock = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugThreadcreate, XTestDebugThreadcreateM defined TODO
var XTestDebugThreadcreate, XTestDebugThreadcreateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugThreadcreate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugThreadcreate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugCmdline, XTestDebugCmdlineM defined TODO
var XTestDebugCmdline, XTestDebugCmdlineRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugCmdline = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugCmdline = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugProfile, XTestDebugProfileM defined TODO
var XTestDebugProfile, XTestDebugProfileRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugProfile = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugProfile = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugSymbol, XTestDebugSymbolM defined TODO
var XTestDebugSymbol, XTestDebugSymbolRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugSymbol = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugSymbol = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugTrace, XTestDebugTraceM defined TODO
var XTestDebugTrace, XTestDebugTraceRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugTrace = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugTrace = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDebugMutex, XTestDebugMutexM defined TODO
var XTestDebugMutex, XTestDebugMutexRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDebugMutex = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDebugMutex = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDingtalkOauth2, XTestSysDingtalkOauth2M defined TODO
var XTestSysDingtalkOauth2, XTestSysDingtalkOauth2Request = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDingtalkOauth2 = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDingtalkOauth2 = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainAdd, XTestSysDomainAddM defined TODO
var XTestSysDomainAdd, XTestSysDomainAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainBatchAdd, XTestSysDomainBatchAddM defined TODO
var XTestSysDomainBatchAdd, XTestSysDomainBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainDel, XTestSysDomainDelM defined TODO
var XTestSysDomainDel, XTestSysDomainDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainBatchDel, XTestSysDomainBatchDelM defined TODO
var XTestSysDomainBatchDel, XTestSysDomainBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainUpdate, XTestSysDomainUpdateM defined TODO
var XTestSysDomainUpdate, XTestSysDomainUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainBatchUpdate, XTestSysDomainBatchUpdateM defined TODO
var XTestSysDomainBatchUpdate, XTestSysDomainBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainPage, XTestSysDomainPageM defined TODO
var XTestSysDomainPage, XTestSysDomainPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysDomainGet, XTestSysDomainGetM defined TODO
var XTestSysDomainGet, XTestSysDomainGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysDomainGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuAdd, XTestSysMenuAddM defined TODO
var XTestSysMenuAdd, XTestSysMenuAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuBatchAdd, XTestSysMenuBatchAddM defined TODO
var XTestSysMenuBatchAdd, XTestSysMenuBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuDel, XTestSysMenuDelM defined TODO
var XTestSysMenuDel, XTestSysMenuDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuBatchDel, XTestSysMenuBatchDelM defined TODO
var XTestSysMenuBatchDel, XTestSysMenuBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuUpdate, XTestSysMenuUpdateM defined TODO
var XTestSysMenuUpdate, XTestSysMenuUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuBatchUpdate, XTestSysMenuBatchUpdateM defined TODO
var XTestSysMenuBatchUpdate, XTestSysMenuBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuSidebar, XTestSysMenuSidebarM defined TODO
var XTestSysMenuSidebar, XTestSysMenuSidebarRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuSidebar = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuSidebar = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuPage, XTestSysMenuPageM defined TODO
var XTestSysMenuPage, XTestSysMenuPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuTree, XTestSysMenuTreeM defined TODO
var XTestSysMenuTree, XTestSysMenuTreeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuTree = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysMenuGet, XTestSysMenuGetM defined TODO
var XTestSysMenuGet, XTestSysMenuGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysMenuGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationAdd, XTestSysNotificationAddM defined TODO
var XTestSysNotificationAdd, XTestSysNotificationAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationBatchAdd, XTestSysNotificationBatchAddM defined TODO
var XTestSysNotificationBatchAdd, XTestSysNotificationBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationDel, XTestSysNotificationDelM defined TODO
var XTestSysNotificationDel, XTestSysNotificationDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationBatchDel, XTestSysNotificationBatchDelM defined TODO
var XTestSysNotificationBatchDel, XTestSysNotificationBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationUpdate, XTestSysNotificationUpdateM defined TODO
var XTestSysNotificationUpdate, XTestSysNotificationUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationBatchUpdate, XTestSysNotificationBatchUpdateM defined TODO
var XTestSysNotificationBatchUpdate, XTestSysNotificationBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationPage, XTestSysNotificationPageM defined TODO
var XTestSysNotificationPage, XTestSysNotificationPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysNotificationGet, XTestSysNotificationGetM defined TODO
var XTestSysNotificationGet, XTestSysNotificationGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysNotificationGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetAdd, XTestSysOptionsetAddM defined TODO
var XTestSysOptionsetAdd, XTestSysOptionsetAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetBatchAdd, XTestSysOptionsetBatchAddM defined TODO
var XTestSysOptionsetBatchAdd, XTestSysOptionsetBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetDel, XTestSysOptionsetDelM defined TODO
var XTestSysOptionsetDel, XTestSysOptionsetDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetBatchDel, XTestSysOptionsetBatchDelM defined TODO
var XTestSysOptionsetBatchDel, XTestSysOptionsetBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetUpdate, XTestSysOptionsetUpdateM defined TODO
var XTestSysOptionsetUpdate, XTestSysOptionsetUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetBatchUpdate, XTestSysOptionsetBatchUpdateM defined TODO
var XTestSysOptionsetBatchUpdate, XTestSysOptionsetBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetPage, XTestSysOptionsetPageM defined TODO
var XTestSysOptionsetPage, XTestSysOptionsetPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOptionsetGet, XTestSysOptionsetGetM defined TODO
var XTestSysOptionsetGet, XTestSysOptionsetGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOptionsetGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgAdd, XTestSysOrgAddM defined TODO
var XTestSysOrgAdd, XTestSysOrgAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgBatchAdd, XTestSysOrgBatchAddM defined TODO
var XTestSysOrgBatchAdd, XTestSysOrgBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgDel, XTestSysOrgDelM defined TODO
var XTestSysOrgDel, XTestSysOrgDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgBatchDel, XTestSysOrgBatchDelM defined TODO
var XTestSysOrgBatchDel, XTestSysOrgBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgUpdate, XTestSysOrgUpdateM defined TODO
var XTestSysOrgUpdate, XTestSysOrgUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgBatchUpdate, XTestSysOrgBatchUpdateM defined TODO
var XTestSysOrgBatchUpdate, XTestSysOrgBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgPage, XTestSysOrgPageM defined TODO
var XTestSysOrgPage, XTestSysOrgPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgTree, XTestSysOrgTreeM defined TODO
var XTestSysOrgTree, XTestSysOrgTreeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgTree = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysOrgGet, XTestSysOrgGetM defined TODO
var XTestSysOrgGet, XTestSysOrgGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysOrgGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionAdd, XTestSysPermissionAddM defined TODO
var XTestSysPermissionAdd, XTestSysPermissionAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionBatchAdd, XTestSysPermissionBatchAddM defined TODO
var XTestSysPermissionBatchAdd, XTestSysPermissionBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionDel, XTestSysPermissionDelM defined TODO
var XTestSysPermissionDel, XTestSysPermissionDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionBatchDel, XTestSysPermissionBatchDelM defined TODO
var XTestSysPermissionBatchDel, XTestSysPermissionBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionUpdate, XTestSysPermissionUpdateM defined TODO
var XTestSysPermissionUpdate, XTestSysPermissionUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionBatchUpdate, XTestSysPermissionBatchUpdateM defined TODO
var XTestSysPermissionBatchUpdate, XTestSysPermissionBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionPage, XTestSysPermissionPageM defined TODO
var XTestSysPermissionPage, XTestSysPermissionPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysPermissionGet, XTestSysPermissionGetM defined TODO
var XTestSysPermissionGet, XTestSysPermissionGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysPermissionGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleAdd, XTestSysRoleAddM defined TODO
var XTestSysRoleAdd, XTestSysRoleAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleBatchAdd, XTestSysRoleBatchAddM defined TODO
var XTestSysRoleBatchAdd, XTestSysRoleBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleDel, XTestSysRoleDelM defined TODO
var XTestSysRoleDel, XTestSysRoleDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleBatchDel, XTestSysRoleBatchDelM defined TODO
var XTestSysRoleBatchDel, XTestSysRoleBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleUpdate, XTestSysRoleUpdateM defined TODO
var XTestSysRoleUpdate, XTestSysRoleUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleBatchUpdate, XTestSysRoleBatchUpdateM defined TODO
var XTestSysRoleBatchUpdate, XTestSysRoleBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRolePage, XTestSysRolePageM defined TODO
var XTestSysRolePage, XTestSysRolePageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRolePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRolePage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleRoleMenuTree, XTestSysRoleRoleMenuTreeM defined TODO
var XTestSysRoleRoleMenuTree, XTestSysRoleRoleMenuTreeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleMenuTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleMenuTree = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleRoleAppFunTree, XTestSysRoleRoleAppFunTreeM defined TODO
var XTestSysRoleRoleAppFunTree, XTestSysRoleRoleAppFunTreeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleAppFunTree = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleRoleAppFunTree = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleGet, XTestSysRoleGetM defined TODO
var XTestSysRoleGet, XTestSysRoleGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuAdd, XTestSysRoleMenuAddM defined TODO
var XTestSysRoleMenuAdd, XTestSysRoleMenuAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuBatchAdd, XTestSysRoleMenuBatchAddM defined TODO
var XTestSysRoleMenuBatchAdd, XTestSysRoleMenuBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuDel, XTestSysRoleMenuDelM defined TODO
var XTestSysRoleMenuDel, XTestSysRoleMenuDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuBatchDel, XTestSysRoleMenuBatchDelM defined TODO
var XTestSysRoleMenuBatchDel, XTestSysRoleMenuBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuUpdate, XTestSysRoleMenuUpdateM defined TODO
var XTestSysRoleMenuUpdate, XTestSysRoleMenuUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuBatchUpdate, XTestSysRoleMenuBatchUpdateM defined TODO
var XTestSysRoleMenuBatchUpdate, XTestSysRoleMenuBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuPage, XTestSysRoleMenuPageM defined TODO
var XTestSysRoleMenuPage, XTestSysRoleMenuPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysRoleMenuGet, XTestSysRoleMenuGetM defined TODO
var XTestSysRoleMenuGet, XTestSysRoleMenuGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysRoleMenuGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleAdd, XTestSysScheduleAddM defined TODO
var XTestSysScheduleAdd, XTestSysScheduleAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleBatchAdd, XTestSysScheduleBatchAddM defined TODO
var XTestSysScheduleBatchAdd, XTestSysScheduleBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleDel, XTestSysScheduleDelM defined TODO
var XTestSysScheduleDel, XTestSysScheduleDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleBatchDel, XTestSysScheduleBatchDelM defined TODO
var XTestSysScheduleBatchDel, XTestSysScheduleBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleUpdate, XTestSysScheduleUpdateM defined TODO
var XTestSysScheduleUpdate, XTestSysScheduleUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleBatchUpdate, XTestSysScheduleBatchUpdateM defined TODO
var XTestSysScheduleBatchUpdate, XTestSysScheduleBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSchedulePage, XTestSysSchedulePageM defined TODO
var XTestSysSchedulePage, XTestSysSchedulePageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulePage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleGet, XTestSysScheduleGetM defined TODO
var XTestSysScheduleGet, XTestSysScheduleGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysScheduleHistoryPage, XTestSysScheduleHistoryPageM defined TODO
var XTestSysScheduleHistoryPage, XTestSysScheduleHistoryPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleHistoryPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysScheduleHistoryPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSchedulingAdd, XTestSysSchedulingAddM defined TODO
var XTestSysSchedulingAdd, XTestSysSchedulingAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSchedulingDel, XTestSysSchedulingDelM defined TODO
var XTestSysSchedulingDel, XTestSysSchedulingDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSchedulingUpdate, XTestSysSchedulingUpdateM defined TODO
var XTestSysSchedulingUpdate, XTestSysSchedulingUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSchedulingPage, XTestSysSchedulingPageM defined TODO
var XTestSysSchedulingPage, XTestSysSchedulingPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSchedulingGet, XTestSysSchedulingGetM defined TODO
var XTestSysSchedulingGet, XTestSysSchedulingGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSchedulingGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingAdd, XTestSysSettingAddM defined TODO
var XTestSysSettingAdd, XTestSysSettingAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingBatchAdd, XTestSysSettingBatchAddM defined TODO
var XTestSysSettingBatchAdd, XTestSysSettingBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingDel, XTestSysSettingDelM defined TODO
var XTestSysSettingDel, XTestSysSettingDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingBatchDel, XTestSysSettingBatchDelM defined TODO
var XTestSysSettingBatchDel, XTestSysSettingBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingUpdate, XTestSysSettingUpdateM defined TODO
var XTestSysSettingUpdate, XTestSysSettingUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingBatchUpdate, XTestSysSettingBatchUpdateM defined TODO
var XTestSysSettingBatchUpdate, XTestSysSettingBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingPage, XTestSysSettingPageM defined TODO
var XTestSysSettingPage, XTestSysSettingPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysSettingGet, XTestSysSettingGetM defined TODO
var XTestSysSettingGet, XTestSysSettingGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysSettingGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableAdd, XTestSysTableAddM defined TODO
var XTestSysTableAdd, XTestSysTableAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableBatchAdd, XTestSysTableBatchAddM defined TODO
var XTestSysTableBatchAdd, XTestSysTableBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableDel, XTestSysTableDelM defined TODO
var XTestSysTableDel, XTestSysTableDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableBatchDel, XTestSysTableBatchDelM defined TODO
var XTestSysTableBatchDel, XTestSysTableBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableUpdate, XTestSysTableUpdateM defined TODO
var XTestSysTableUpdate, XTestSysTableUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableBatchUpdate, XTestSysTableBatchUpdateM defined TODO
var XTestSysTableBatchUpdate, XTestSysTableBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTablePage, XTestSysTablePageM defined TODO
var XTestSysTablePage, XTestSysTablePageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTablePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTablePage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableGet, XTestSysTableGetM defined TODO
var XTestSysTableGet, XTestSysTableGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnAdd, XTestSysTableColumnAddM defined TODO
var XTestSysTableColumnAdd, XTestSysTableColumnAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnBatchAdd, XTestSysTableColumnBatchAddM defined TODO
var XTestSysTableColumnBatchAdd, XTestSysTableColumnBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnDel, XTestSysTableColumnDelM defined TODO
var XTestSysTableColumnDel, XTestSysTableColumnDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnBatchDel, XTestSysTableColumnBatchDelM defined TODO
var XTestSysTableColumnBatchDel, XTestSysTableColumnBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnUpdate, XTestSysTableColumnUpdateM defined TODO
var XTestSysTableColumnUpdate, XTestSysTableColumnUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnBatchUpdate, XTestSysTableColumnBatchUpdateM defined TODO
var XTestSysTableColumnBatchUpdate, XTestSysTableColumnBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnPage, XTestSysTableColumnPageM defined TODO
var XTestSysTableColumnPage, XTestSysTableColumnPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTableColumnGet, XTestSysTableColumnGetM defined TODO
var XTestSysTableColumnGet, XTestSysTableColumnGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTableColumnGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagAdd, XTestSysTagAddM defined TODO
var XTestSysTagAdd, XTestSysTagAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagBatchAdd, XTestSysTagBatchAddM defined TODO
var XTestSysTagBatchAdd, XTestSysTagBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagDel, XTestSysTagDelM defined TODO
var XTestSysTagDel, XTestSysTagDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagBatchDel, XTestSysTagBatchDelM defined TODO
var XTestSysTagBatchDel, XTestSysTagBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagUpdate, XTestSysTagUpdateM defined TODO
var XTestSysTagUpdate, XTestSysTagUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagBatchUpdate, XTestSysTagBatchUpdateM defined TODO
var XTestSysTagBatchUpdate, XTestSysTagBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagPage, XTestSysTagPageM defined TODO
var XTestSysTagPage, XTestSysTagPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGet, XTestSysTagGetM defined TODO
var XTestSysTagGet, XTestSysTagGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupAdd, XTestSysTagGroupAddM defined TODO
var XTestSysTagGroupAdd, XTestSysTagGroupAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupBatchAdd, XTestSysTagGroupBatchAddM defined TODO
var XTestSysTagGroupBatchAdd, XTestSysTagGroupBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupDel, XTestSysTagGroupDelM defined TODO
var XTestSysTagGroupDel, XTestSysTagGroupDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupBatchDel, XTestSysTagGroupBatchDelM defined TODO
var XTestSysTagGroupBatchDel, XTestSysTagGroupBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupUpdate, XTestSysTagGroupUpdateM defined TODO
var XTestSysTagGroupUpdate, XTestSysTagGroupUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupBatchUpdate, XTestSysTagGroupBatchUpdateM defined TODO
var XTestSysTagGroupBatchUpdate, XTestSysTagGroupBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupPage, XTestSysTagGroupPageM defined TODO
var XTestSysTagGroupPage, XTestSysTagGroupPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTagGroupGet, XTestSysTagGroupGetM defined TODO
var XTestSysTagGroupGet, XTestSysTagGroupGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTagGroupGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTrackerPage, XTestSysTrackerPageM defined TODO
var XTestSysTrackerPage, XTestSysTrackerPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysTrackerGet, XTestSysTrackerGetM defined TODO
var XTestSysTrackerGet, XTestSysTrackerGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysTrackerGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserAdd, XTestSysUserAddM defined TODO
var XTestSysUserAdd, XTestSysUserAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserBatchAdd, XTestSysUserBatchAddM defined TODO
var XTestSysUserBatchAdd, XTestSysUserBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserDel, XTestSysUserDelM defined TODO
var XTestSysUserDel, XTestSysUserDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserBatchDel, XTestSysUserBatchDelM defined TODO
var XTestSysUserBatchDel, XTestSysUserBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserUpdate, XTestSysUserUpdateM defined TODO
var XTestSysUserUpdate, XTestSysUserUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserBatchUpdate, XTestSysUserBatchUpdateM defined TODO
var XTestSysUserBatchUpdate, XTestSysUserBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserPage, XTestSysUserPageM defined TODO
var XTestSysUserPage, XTestSysUserPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserGet, XTestSysUserGetM defined TODO
var XTestSysUserGet, XTestSysUserGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserLogout, XTestSysUserLogoutM defined TODO
var XTestSysUserLogout, XTestSysUserLogoutRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserLogout = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserLogout = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateAdd, XTestSysUserTemplateAddM defined TODO
var XTestSysUserTemplateAdd, XTestSysUserTemplateAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateBatchAdd, XTestSysUserTemplateBatchAddM defined TODO
var XTestSysUserTemplateBatchAdd, XTestSysUserTemplateBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDel, XTestSysUserTemplateDelM defined TODO
var XTestSysUserTemplateDel, XTestSysUserTemplateDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateBatchDel, XTestSysUserTemplateBatchDelM defined TODO
var XTestSysUserTemplateBatchDel, XTestSysUserTemplateBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateUpdate, XTestSysUserTemplateUpdateM defined TODO
var XTestSysUserTemplateUpdate, XTestSysUserTemplateUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateBatchUpdate, XTestSysUserTemplateBatchUpdateM defined TODO
var XTestSysUserTemplateBatchUpdate, XTestSysUserTemplateBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplatePage, XTestSysUserTemplatePageM defined TODO
var XTestSysUserTemplatePage, XTestSysUserTemplatePageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplatePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplatePage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateGet, XTestSysUserTemplateGetM defined TODO
var XTestSysUserTemplateGet, XTestSysUserTemplateGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailAdd, XTestSysUserTemplateDetailAddM defined TODO
var XTestSysUserTemplateDetailAdd, XTestSysUserTemplateDetailAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailBatchAdd, XTestSysUserTemplateDetailBatchAddM defined TODO
var XTestSysUserTemplateDetailBatchAdd, XTestSysUserTemplateDetailBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailDel, XTestSysUserTemplateDetailDelM defined TODO
var XTestSysUserTemplateDetailDel, XTestSysUserTemplateDetailDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailBatchDel, XTestSysUserTemplateDetailBatchDelM defined TODO
var XTestSysUserTemplateDetailBatchDel, XTestSysUserTemplateDetailBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailUpdate, XTestSysUserTemplateDetailUpdateM defined TODO
var XTestSysUserTemplateDetailUpdate, XTestSysUserTemplateDetailUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailBatchUpdate, XTestSysUserTemplateDetailBatchUpdateM defined TODO
var XTestSysUserTemplateDetailBatchUpdate, XTestSysUserTemplateDetailBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailPage, XTestSysUserTemplateDetailPageM defined TODO
var XTestSysUserTemplateDetailPage, XTestSysUserTemplateDetailPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysUserTemplateDetailGet, XTestSysUserTemplateDetailGetM defined TODO
var XTestSysUserTemplateDetailGet, XTestSysUserTemplateDetailGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysUserTemplateDetailGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysWechatOauth2, XTestSysWechatOauth2M defined TODO
var XTestSysWechatOauth2, XTestSysWechatOauth2Request = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysWechatOauth2 = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysWechatOauth2 = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysWorkerAdd, XTestSysWorkerAddM defined TODO
var XTestSysWorkerAdd, XTestSysWorkerAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSysWorkerGet, XTestSysWorkerGetM defined TODO
var XTestSysWorkerGet, XTestSysWorkerGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSysWorkerGet = %v want %v", ret.Code, 200)
	}
}, Payload{}
