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

// XTestOrganAdd, XTestOrganAddM defined TODO
var XTestOrganAdd, XTestOrganAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestOrganBatchAdd, XTestOrganBatchAddM defined TODO
var XTestOrganBatchAdd, XTestOrganBatchAddRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchAdd = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestOrganDel, XTestOrganDelM defined TODO
var XTestOrganDel, XTestOrganDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestOrganBatchDel, XTestOrganBatchDelM defined TODO
var XTestOrganBatchDel, XTestOrganBatchDelRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchDel = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestOrganUpdate, XTestOrganUpdateM defined TODO
var XTestOrganUpdate, XTestOrganUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestOrganBatchUpdate, XTestOrganBatchUpdateM defined TODO
var XTestOrganBatchUpdate, XTestOrganBatchUpdateRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchUpdate = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestOrganPage, XTestOrganPageM defined TODO
var XTestOrganPage, XTestOrganPageRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganPage = %v want %v", ret.Code, 200)
	}
}, Payload{}

// XTestOrganGet, XTestOrganGetM defined TODO
var XTestOrganGet, XTestOrganGetRequest = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganGet = %v want %v", ret.Code, 200)
	}
}, Payload{}
