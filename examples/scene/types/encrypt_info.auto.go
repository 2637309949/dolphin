// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import "github.com/2637309949/dolphin/packages/null"

// EncryptInfo defined 文章信息
type EncryptInfo struct {
	// 应用ID
	AppId null.String `json:"app_id" xml:"app_id"`
	// 签名数据
	Sign null.String `json:"sign" xml:"sign"`
	// 时间戳
	Timestamp null.Int `json:"timestamp" xml:"timestamp"`
	// 请求参数
	BizContent null.String `json:"biz_content" xml:"biz_content"`
}
