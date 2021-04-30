package main

import (
	"encoding/json"
)

// XTestSysUserLogin defined
func XTestSysUserLogin(ctx *Context) {
	ret := struct {
		Code int `json:"code"`
		Data struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}{}
	json.Unmarshal(ctx.Body.Bytes(), &ret)
	SetToken(ret.Data.AccessToken)
}

// XTestSysMenuPage defined
func XTestSysMenuPage(ctx *Context) {
	if ctx.Code != 200 {
		ctx.T.Errorf("XTestSysMenuPage = %v want %v", ctx.Code, 200)
	}
	ret := Response{}
	json.Unmarshal(ctx.Body.Bytes(), &ret)
	if ret.Code != 200 {
		ctx.T.Errorf("XTestSysMenuPage = %v want %v", ret.Code, 200)
	}
}
