// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import "github.com/2637309949/dolphin/packages/null"

// Login defined 登录表单
type Login struct {
	// 名字
	Name null.String `json:"name" xml:"name"`
	// 密码
	Password null.String `json:"password" xml:"password"`
	// 域
	Domain null.String `json:"domain" xml:"domain"`
}
