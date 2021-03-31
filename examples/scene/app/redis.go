// Code generated by dol build. Only Generate by tools if not existed.
// source: redis.go

package app

import (
	"scene/srv"

	"github.com/2637309949/dolphin/packages/logrus"
)

// RedisLock api implementation
// @Summary Add lock
// @Tags redis controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "redis info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/redis/lock [get]
func RedisLock(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	q.Value()
	ret, err := srv.RedisAction(ctx.Raw(), ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// RedisUnlock api implementation
// @Summary del lock
// @Tags redis controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "redis info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/redis/unlock [get]
func RedisUnlock(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	q.Value()
	ret, err := srv.RedisAction(ctx.Raw(), ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
