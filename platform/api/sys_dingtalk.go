// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_dingtalk.go

package api

import (
	"github.com/sirupsen/logrus"
)

// SysDingtalkOauth2 api implementation
// @Summary 授权回调
// @Tags 钉钉
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/dingtalk/oauth2 [get]
func (ctr *SysDingtalk) SysDingtalkOauth2(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetUser()
	ret, err := ctr.Srv.TODO(ctx, ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}