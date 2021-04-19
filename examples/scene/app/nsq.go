// Code generated by dol build. Only Generate by tools if not existed.
// source: nsq.go

package app

import (
	"scene/model"
	"scene/srv"

	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

// NsqAdd api implementation
// @Summary Add article
// @Tags Nsq controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param nsq body model.NsqInfo false "nsq info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/nsq/add [post]
func NsqAdd(ctx *Context) {
	var payload model.NsqInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := srv.NProducer(ctx.Raw(), ctx.DB, payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// NsqGet api implementation
// @Summary Get Nsq info
// @Tags Nsq controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "nsq id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/nsq/get [get]
func NsqGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	q.Value()
	// ret, err := srv.NConsumer(ctx.Raw(), ctx.DB, q.Value())
	// if err != nil {
	// 	logrus.Error(err)
	// 	ctx.Fail(err)
	// 	return
	// }
	ctx.Success("ok")
}
