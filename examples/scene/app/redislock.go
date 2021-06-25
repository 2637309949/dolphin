// Code generated by dol build. Only Generate by tools if not existed.
// source: redis.go

package app

import (
	"github.com/sirupsen/logrus"
)

// RedisLockLock api implementation
// @Summary Add lock
// @Tags redis controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "redis info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/redis/lock [get]
func (ctr *RedisLock) RedisLockLock(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	q.Value()
	ret, err := ctr.Srv.TODO(ctx, ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// RedisLockUnlock api implementation
// @Summary del lock
// @Tags redis controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "redis info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/redis/unlock [get]
func (ctr *RedisLock) RedisLockUnlock(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	q.Value()
	ret, err := ctr.Srv.TODO(ctx, ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
