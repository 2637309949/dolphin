package main

import (
	"encoding/json"
	"scene/types"

	// "github.com/mattn/go-sqlite3"
	"github.com/2637309949/dolphin/packages/null"
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

// XTestArticleAdd, XTestArticleAddM defined TODO
var XTestArticleAdd, XTestArticleAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleAdd = %v want %v", ret.Code, 200)
	}
}, types.Article{Name: null.StringFrom("各种场景应用")}

// XTestArticleBatchAdd, XTestArticleBatchAddM defined TODO
var XTestArticleBatchAdd, XTestArticleBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleDel, XTestArticleDelM defined TODO
var XTestArticleDel, XTestArticleDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	ctx.testingT.Log(ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleDel = %v want %v", ret.Code, 200)
	}
}, types.Article{ID: null.IntFrom(1)}

// XTestArticleBatchDel, XTestArticleBatchDelM defined TODO
var XTestArticleBatchDel, XTestArticleBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleUpdate, XTestArticleUpdateM defined TODO
var XTestArticleUpdate, XTestArticleUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleBatchUpdate, XTestArticleBatchUpdateM defined TODO
var XTestArticleBatchUpdate, XTestArticleBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticlePage, XTestArticlePageM defined TODO
var XTestArticlePage, XTestArticlePageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleGet, XTestArticleGetM defined TODO
var XTestArticleGet, XTestArticleGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticlePayment, XTestArticlePaymentM defined TODO
var XTestArticlePayment, XTestArticlePaymentRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePayment = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePayment = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestICacheInfo, XTestICacheInfoM defined TODO
var XTestICacheInfo, XTestICacheInfoRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestICacheInfo = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestICacheInfo = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestEncryptAdd, XTestEncryptAddM defined TODO
var XTestEncryptAdd, XTestEncryptAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestEncryptInfo, XTestEncryptInfoM defined TODO
var XTestEncryptInfo, XTestEncryptInfoRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptInfo = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptInfo = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestKafkaAdd, XTestKafkaAddM defined TODO
var XTestKafkaAdd, XTestKafkaAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestKafkaGet, XTestKafkaGetM defined TODO
var XTestKafkaGet, XTestKafkaGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestNsqAdd, XTestNsqAddM defined TODO
var XTestNsqAdd, XTestNsqAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestNsqAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestNsqAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestNsqGet, XTestNsqGetM defined TODO
var XTestNsqGet, XTestNsqGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestNsqGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestNsqGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisLockLock, XTestRedisLockLockM defined TODO
var XTestRedisLockLock, XTestRedisLockLockRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockLock = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockLock = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisLockUnlock, XTestRedisLockUnlockM defined TODO
var XTestRedisLockUnlock, XTestRedisLockUnlockRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockUnlock = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockUnlock = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisMqAdd, XTestRedisMqAddM defined TODO
var XTestRedisMqAdd, XTestRedisMqAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisMqGet, XTestRedisMqGetM defined TODO
var XTestRedisMqGet, XTestRedisMqGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRPCMessage, XTestRPCMessageM defined TODO
var XTestRPCMessage, XTestRPCMessageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRPCMessage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRPCMessage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSqlmapSelectone, XTestSqlmapSelectoneM defined TODO
var XTestSqlmapSelectone, XTestSqlmapSelectoneRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSqlmapSelectone = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSqlmapSelectone = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestUserInfo, XTestUserInfoM defined TODO
var XTestUserInfo, XTestUserInfoRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestUserInfo = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestUserInfo = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestViewFile, XTestViewFileM defined TODO
var XTestViewFile, XTestViewFileRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewFile = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewFile = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestViewHTML, XTestViewHTMLM defined TODO
var XTestViewHTML, XTestViewHTMLRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewHTML = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewHTML = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestViewXML, XTestViewXMLM defined TODO
var XTestViewXML, XTestViewXMLRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewXML = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewXML = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestVoteLike, XTestVoteLikeM defined TODO
var XTestVoteLike, XTestVoteLikeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestVoteLike = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestVoteLike = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestJwtCheck, XTestJwtCheckM defined TODO
var XTestJwtCheck, XTestJwtCheckRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestJwtCheck = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestJwtCheck = %v want %v", ret.Code, 200)
	}
}, Payload{}
