// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_client.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysClientAdd api implementation
// @Summary 添加客户端
// @Tags 客户端
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysClient false "客户端信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/client/add [post]
func SysClientAdd(ctx *Context) {
	var payload model.SysClient
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFromPtr(time.Now().Value())
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.DelFlag = null.IntFrom(0)
	ret, err := ctx.PlatformDB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysClientDel api implementation
// @Summary 删除客户端
// @Tags 客户端
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_client body model.SysClient false "客户端"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/client/del [delete]
func SysClientDel(ctx *Context) {
	var payload model.SysClient
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.PlatformDB.In("id", payload.ID.String).Update(&model.SysClient{
		UpdateTime: null.TimeFromPtr(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysClientUpdate api implementation
// @Summary 更新客户端
// @Tags 客户端
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "客户端信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/client/update [put]
func SysClientUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())
	ret, err := ctx.PlatformDB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysClientPage api implementation
// @Summary 客户端分页查询
// @Tags 客户端
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/client/page [get]
func SysClientPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.PlatformDB, "sys_client", "page", "sys_client", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysClientGet api implementation
// @Summary 获取客户端信息
// @Tags 客户端
// @Param Authorization header string false "认证令牌"
// @Param id query string false "客户端id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/client/get [get]
func SysClientGet(ctx *Context) {
	var entity model.SysClient
	id := ctx.Query("id")
	_, err := ctx.PlatformDB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
