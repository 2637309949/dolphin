// Code generated by dol build. Only Generate by tools if not existed.
// source: vote.go

package app

import (
	"scene/model"
	"scene/srv"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
)

// VoteLike api implementation
// @Summary like article
// @Tags vote controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param id body model.VoteInfo false "like id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/vote/like [post]
func VoteLike(ctx *Context) {
	var payload model.VoteInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.UserId = null.StringFrom(ctx.GetToken().GetUserID())
	ret, err := srv.VoteLike(ctx.Raw(), ctx.DB, payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
