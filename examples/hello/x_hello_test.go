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
}, M{}

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
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleDel = %v want %v", ret.Code, 200)
	}
}, M{}

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

// XTestArticleTodo, XTestArticleTodoM defined TODO
var XTestArticleTodo, XTestArticleTodoM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestArticleTodo = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestArticleTodo = %v want %v", ret.Code, 200)
	}
}, M{}
