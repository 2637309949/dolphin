// Code generated by dol build. Only Generate by tools if not existed.
// source: rpc.go

package api

import (
	"github.com/sirupsen/logrus"
)

// RPCMessage api implementation
// @Summary Get message
// @Tags rpc controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "rpc info"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/rpc/message [get]
func (ctr *RPC) RPCMessage(ctx *Context) {
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
