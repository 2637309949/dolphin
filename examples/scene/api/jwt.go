// Code generated by dol build. Only Generate by tools if not existed.
// source: jwt.go

package api

import (
	"github.com/sirupsen/logrus"
)

// JwtLogin api implementation
// @Summary 登陆
// @Tags jwt认证
// @Accept application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/jwt/login [post]
func (ctr *Jwt) JwtLogin(ctx *Context) {
	var payload struct{}
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctr.Srv.CreateJWT(ctx, ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// JwtCheck api implementation
// @Summary 验证
// @Tags jwt认证
// @Param Authorization header string false "认证令牌"
// @Param id  query  int false "id"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/jwt/check [get]
func (ctr *Jwt) JwtCheck(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("id")
	q.Value()
	ret, err := ctr.Srv.TODO(ctx, ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}