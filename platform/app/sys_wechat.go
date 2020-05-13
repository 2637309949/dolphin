// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_wechat.go

package app

import (
	"github.com/2637309949/dolphin/platform/srv"
)

// SysWechatOauth2 api implementation
// @Summary 授权回调
// @Tags 微信
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/wechat/oauth2 [get]
func SysWechatOauth2(ctx *Context) {
	q := ctx.TypeQuery()
	ret, err := srv.SysWechatAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
