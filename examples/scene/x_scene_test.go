package main

import (
	"encoding/json"

	"scene/model"
	// "github.com/mattn/go-sqlite3"
	"github.com/2637309949/dolphin/packages/null"
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

// XTestArticleAdd, XTestArticleAddM defined TODO
var XTestArticleAdd, XTestArticleAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleAdd = %v want %v", ret.Code, 200)
	}
}, model.Article{Name: null.StringFrom("各种场景应用")}

// XTestArticleBatchAdd, XTestArticleBatchAddM defined TODO
var XTestArticleBatchAdd, XTestArticleBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestArticleDel, XTestArticleDelM defined TODO
var XTestArticleDel, XTestArticleDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	ctx.testingT.Log(ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleDel = %v want %v", ret.Code, 200)
	}
}, model.Article{ID: null.IntFrom(1)}

// XTestArticleBatchDel, XTestArticleBatchDelM defined TODO
var XTestArticleBatchDel, XTestArticleBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestArticleUpdate, XTestArticleUpdateM defined TODO
var XTestArticleUpdate, XTestArticleUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestArticleBatchUpdate, XTestArticleBatchUpdateM defined TODO
var XTestArticleBatchUpdate, XTestArticleBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestArticlePage, XTestArticlePageM defined TODO
var XTestArticlePage, XTestArticlePageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestArticleGet, XTestArticleGetM defined TODO
var XTestArticleGet, XTestArticleGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestArticlePayment, XTestArticlePaymentM defined TODO
var XTestArticlePayment, XTestArticlePaymentM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePayment = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePayment = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestICacheInfo, XTestICacheInfoM defined TODO
var XTestICacheInfo, XTestICacheInfoM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestICacheInfo = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestICacheInfo = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestEncryptAdd, XTestEncryptAddM defined TODO
var XTestEncryptAdd, XTestEncryptAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestEncryptInfo, XTestEncryptInfoM defined TODO
var XTestEncryptInfo, XTestEncryptInfoM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptInfo = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptInfo = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestKafkaAdd, XTestKafkaAddM defined TODO
var XTestKafkaAdd, XTestKafkaAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestKafkaGet, XTestKafkaGetM defined TODO
var XTestKafkaGet, XTestKafkaGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestNsqAdd, XTestNsqAddM defined TODO
var XTestNsqAdd, XTestNsqAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestNsqAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestNsqAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestNsqGet, XTestNsqGetM defined TODO
var XTestNsqGet, XTestNsqGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestNsqGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestNsqGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestRedisLockLock, XTestRedisLockLockM defined TODO
var XTestRedisLockLock, XTestRedisLockLockM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockLock = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockLock = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestRedisLockUnlock, XTestRedisLockUnlockM defined TODO
var XTestRedisLockUnlock, XTestRedisLockUnlockM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockUnlock = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockUnlock = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestRedisMqAdd, XTestRedisMqAddM defined TODO
var XTestRedisMqAdd, XTestRedisMqAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestRedisMqGet, XTestRedisMqGetM defined TODO
var XTestRedisMqGet, XTestRedisMqGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqGet = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestRPCMessage, XTestRPCMessageM defined TODO
var XTestRPCMessage, XTestRPCMessageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRPCMessage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRPCMessage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestSqlmapSelectone, XTestSqlmapSelectoneM defined TODO
var XTestSqlmapSelectone, XTestSqlmapSelectoneM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSqlmapSelectone = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSqlmapSelectone = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestUserInfo, XTestUserInfoM defined TODO
var XTestUserInfo, XTestUserInfoM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestUserInfo = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestUserInfo = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestViewFile, XTestViewFileM defined TODO
var XTestViewFile, XTestViewFileM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewFile = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewFile = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestViewHTML, XTestViewHTMLM defined TODO
var XTestViewHTML, XTestViewHTMLM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewHTML = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewHTML = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestViewXML, XTestViewXMLM defined TODO
var XTestViewXML, XTestViewXMLM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewXML = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewXML = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestVoteLike, XTestVoteLikeM defined TODO
var XTestVoteLike, XTestVoteLikeM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestVoteLike = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestVoteLike = %v want %v", ret.Code, 200)
	}
}, M{}
