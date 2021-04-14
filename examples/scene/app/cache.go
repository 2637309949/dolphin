// Code generated by dol build. Only Generate by tools if not existed.
// source: cache.go

package app

import (
	"context"
	"scene/srv"

	"github.com/sirupsen/logrus"
)

// ICacheInfo api implementation
// @Summary Cache info
// @Tags cache controller
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/i/cache/info [get]
func ICacheInfo(ctx *Context) {
	q := ctx.TypeQuery()
	q.Value()
	ret, err := srv.ICacheTODO(ctx.Raw(), ctx.DB, context.Background(), struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
