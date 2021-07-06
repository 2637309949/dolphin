// Code generated by dol build. Only Generate by tools if not existed.
// source: ami.go

package api

import (
	"scene/types"

	"github.com/sirupsen/logrus"
)

// RedisMqAdd api implementation
// @Summary Add ami
// @Tags Ami controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param ami_info body types.AmiInfo false "Ami info"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/ami/add [post]
func (ctr *RedisMq) RedisMqAdd(ctx *Context) {
	var payload types.AmiInfo
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctr.Srv.Producer(ctx, ctx.DB, payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// RedisMqGet api implementation
// @Summary Get ami info
// @Tags Ami controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "ami id"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/ami/get [get]
func (ctr *RedisMq) RedisMqGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	ret, err := ctr.Srv.Consumer(ctx, ctx.DB, q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}