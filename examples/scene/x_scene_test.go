package main

import (
	"encoding/json"

	"scene/types"

	"github.com/2637309949/dolphin/packages/null"

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

// XTestArticleAdd, XTestArticleAddRequest defined TODO
var XTestArticleAdd, XTestArticleAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleAdd = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleAdd = %v want %v", ret.Code, 200)
	}
}, types.Article{Name: null.StringFrom("各种场景应用")}

// XTestArticleBatchAdd, XTestArticleBatchAddRequest defined TODO
var XTestArticleBatchAdd, XTestArticleBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchAdd = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleDel, XTestArticleDelRequest defined TODO
var XTestArticleDel, XTestArticleDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleDel = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleBatchDel, XTestArticleBatchDelRequest defined TODO
var XTestArticleBatchDel, XTestArticleBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchDel = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleUpdate, XTestArticleUpdateRequest defined TODO
var XTestArticleUpdate, XTestArticleUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleUpdate = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleBatchUpdate, XTestArticleBatchUpdateRequest defined TODO
var XTestArticleBatchUpdate, XTestArticleBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchUpdate = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticlePage, XTestArticlePageRequest defined TODO
var XTestArticlePage, XTestArticlePageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePage = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticleGet, XTestArticleGetRequest defined TODO
var XTestArticleGet, XTestArticleGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleGet = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestArticlePayment, XTestArticlePaymentRequest defined TODO
var XTestArticlePayment, XTestArticlePaymentRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePayment = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticlePayment = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestCachingInfo, XTestCachingInfoRequest defined TODO
var XTestCachingInfo, XTestCachingInfoRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestCachingInfo = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestCachingInfo = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDtmTcc, XTestDtmTccRequest defined TODO
var XTestDtmTcc, XTestDtmTccRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTcc = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTcc = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDtmTransOut, XTestDtmTransOutRequest defined TODO
var XTestDtmTransOut, XTestDtmTransOutRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransOut = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransOut = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDtmTransOutConfirm, XTestDtmTransOutConfirmRequest defined TODO
var XTestDtmTransOutConfirm, XTestDtmTransOutConfirmRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransOutConfirm = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransOutConfirm = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDtmTransOutRevert, XTestDtmTransOutRevertRequest defined TODO
var XTestDtmTransOutRevert, XTestDtmTransOutRevertRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransOutRevert = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransOutRevert = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDtmTransIn, XTestDtmTransInRequest defined TODO
var XTestDtmTransIn, XTestDtmTransInRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransIn = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransIn = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDtmTransInConfirm, XTestDtmTransInConfirmRequest defined TODO
var XTestDtmTransInConfirm, XTestDtmTransInConfirmRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransInConfirm = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransInConfirm = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestDtmTransInRevert, XTestDtmTransInRevertRequest defined TODO
var XTestDtmTransInRevert, XTestDtmTransInRevertRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransInRevert = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestDtmTransInRevert = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestEncryptAdd, XTestEncryptAddRequest defined TODO
var XTestEncryptAdd, XTestEncryptAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptAdd = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestEncryptInfo, XTestEncryptInfoRequest defined TODO
var XTestEncryptInfo, XTestEncryptInfoRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptInfo = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestEncryptInfo = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestJwtCheck, XTestJwtCheckRequest defined TODO
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

// XTestKafkaAdd, XTestKafkaAddRequest defined TODO
var XTestKafkaAdd, XTestKafkaAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaAdd = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestKafkaGet, XTestKafkaGetRequest defined TODO
var XTestKafkaGet, XTestKafkaGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaGet = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestKafkaGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestMiddlesGlobalTest, XTestMiddlesGlobalTestRequest defined TODO
var XTestMiddlesGlobalTest, XTestMiddlesGlobalTestRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestMiddlesGlobalTest = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestMiddlesGlobalTest = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestMiddlesLocalTest, XTestMiddlesLocalTestRequest defined TODO
var XTestMiddlesLocalTest, XTestMiddlesLocalTestRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestMiddlesLocalTest = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestMiddlesLocalTest = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestNsqAdd, XTestNsqAddRequest defined TODO
var XTestNsqAdd, XTestNsqAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestNsqAdd = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestNsqAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestNsqGet, XTestNsqGetRequest defined TODO
var XTestNsqGet, XTestNsqGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestNsqGet = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestNsqGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisLockLock, XTestRedisLockLockRequest defined TODO
var XTestRedisLockLock, XTestRedisLockLockRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockLock = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockLock = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisLockUnlock, XTestRedisLockUnlockRequest defined TODO
var XTestRedisLockUnlock, XTestRedisLockUnlockRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockUnlock = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisLockUnlock = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisMqAdd, XTestRedisMqAddRequest defined TODO
var XTestRedisMqAdd, XTestRedisMqAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqAdd = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRedisMqGet, XTestRedisMqGetRequest defined TODO
var XTestRedisMqGet, XTestRedisMqGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqGet = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRedisMqGet = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestRPCMessage, XTestRPCMessageRequest defined TODO
var XTestRPCMessage, XTestRPCMessageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestRPCMessage = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestRPCMessage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestSqlmapSelectone, XTestSqlmapSelectoneRequest defined TODO
var XTestSqlmapSelectone, XTestSqlmapSelectoneRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestSqlmapSelectone = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestSqlmapSelectone = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestUserInfo, XTestUserInfoRequest defined TODO
var XTestUserInfo, XTestUserInfoRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestUserInfo = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestUserInfo = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestViewFile, XTestViewFileRequest defined TODO
var XTestViewFile, XTestViewFileRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewFile = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewFile = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestViewHTML, XTestViewHTMLRequest defined TODO
var XTestViewHTML, XTestViewHTMLRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewHTML = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewHTML = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestViewXML, XTestViewXMLRequest defined TODO
var XTestViewXML, XTestViewXMLRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestViewXML = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestViewXML = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestVoteLike, XTestVoteLikeRequest defined TODO
var XTestVoteLike, XTestVoteLikeRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestVoteLike = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestVoteLike = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestXlsxImport, XTestXlsxImportRequest defined TODO
var XTestXlsxImport, XTestXlsxImportRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestXlsxImport = %v want %v", ctx.Code, 200)
	}
	err := ctx.ParseBody(&ret)
	if err != nil {
		ctx.testingT.Error(err)
	}
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestXlsxImport = %v want %v", ret.Code, 200)
	}
}, Payload{}
