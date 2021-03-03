// Code generated by dol build. Only Generate by tools if not existed.
// source: encrypt.go

package app

import (
	"scene/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
)

// EncryptAdd api implementation
// @Summary Add Encrypt
// @Tags Encrypt controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param encrypt_info body model.EncryptInfo false "Encrypt info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/encrypt/add [post]
func EncryptAdd(ctx *Context) {
	var payload model.EncryptInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(payload)
}