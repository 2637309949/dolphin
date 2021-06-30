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

// XTestOrganAdd, XTestOrganAddM defined TODO
var XTestOrganAdd, XTestOrganAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestOrganBatchAdd, XTestOrganBatchAddM defined TODO
var XTestOrganBatchAdd, XTestOrganBatchAddM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchAdd = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchAdd = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestOrganDel, XTestOrganDelM defined TODO
var XTestOrganDel, XTestOrganDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestOrganBatchDel, XTestOrganBatchDelM defined TODO
var XTestOrganBatchDel, XTestOrganBatchDelM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchDel = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchDel = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestOrganUpdate, XTestOrganUpdateM defined TODO
var XTestOrganUpdate, XTestOrganUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestOrganBatchUpdate, XTestOrganBatchUpdateM defined TODO
var XTestOrganBatchUpdate, XTestOrganBatchUpdateM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchUpdate = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganBatchUpdate = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestOrganPage, XTestOrganPageM defined TODO
var XTestOrganPage, XTestOrganPageM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganPage = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganPage = %v want %v", ret.Code, 200)
	}
}, M{}

// XTestOrganGet, XTestOrganGetM defined TODO
var XTestOrganGet, XTestOrganGetM = func(ctx *Context) {
	ret := Response{}
	if ctx.Code != 200 {
		ctx.testingT.Errorf("XTestOrganGet = %v want %v", ctx.Code, 200)
	}
	ctx.ParseBody(&ret)
	if ret.Code != 200 {
		ctx.testingT.Errorf("XTestOrganGet = %v want %v", ret.Code, 200)
	}
}, M{}
