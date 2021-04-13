// Code generated by dol build. Only Generate by tools if not existed.
// source: ami.go

package app

import (
	"scene/model"
	"scene/srv"

	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

// AmiAdd api implementation
// @Summary Add ami
// @Tags Ami controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param ami_info body model.AmiInfo false "Ami info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/ami/add [post]
func AmiAdd(ctx *Context) {
	var payload model.AmiInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := srv.AmiProducer(ctx.Raw(), ctx.DB, payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AmiGet api implementation
// @Summary Get ami info
// @Tags Ami controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "ami id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/ami/get [get]
func AmiGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	ret, err := srv.AmiConsumer(ctx.Raw(), ctx.DB, q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
