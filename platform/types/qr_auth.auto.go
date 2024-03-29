// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// QrAuth defined 扫码认证
type QrAuth struct {
	// 回调地址
	RedirectUri null.String `json:"redirect_uri" xml:"redirect_uri"`
	// 域
	Domain null.String `json:"domain" xml:"domain"`
	// 用户ID
	UserId null.Int `json:"user_id" xml:"user_id"`
}

func (r *QrAuth) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalQrAuth(data []byte) (QrAuth, error) {
	var r QrAuth
	err := json.Unmarshal(data, &r)
	return r, err
}
