// Code generated by dol build. Only Generate by tools if not existed.
// source: caching.go

package api

import (
	"github.com/2637309949/dolphin/packages/logrus"
)

// CachingInfo api implementation
// @Summary Cache info
// @Tags cache controller
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/caching/info [get]
func (ctr *Caching) CachingInfo(ctx *Context) {
	q := ctx.TypeQuery()
	q.Value()

	db := ctx.MustDB()
	ret, err := ctr.Srv.TODO(ctx, db, struct{}{})
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
