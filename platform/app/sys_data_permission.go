// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_data_permission.go

package app

import (
	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/platform/model"
)

// SysDataPermissionAdd api implementation
// @Summary 添加数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysDataPermission false "数据权限信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/add [post]
func SysDataPermissionAdd(ctx *Context) {
	var payload model.SysDataPermission
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFrom(time.Now().Value())
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.DelFlag = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDataPermissionDel api implementation
// @Summary 删除数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_data_permission body model.SysDataPermission false "数据权限"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/del [delete]
func SysDataPermissionDel(ctx *Context) {
	var payload model.SysDataPermission
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysDataPermission{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDataPermissionUpdate api implementation
// @Summary 更新数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "数据权限信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/update [put]
func SysDataPermissionUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.String).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDataPermissionPage api implementation
// @Summary 数据权限分页查询
// @Tags 数据权限
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/page [get]
func SysDataPermissionPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_data_permission", "page", "sys_data_permission", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDataPermissionGet api implementation
// @Summary 获取数据权限信息
// @Tags 数据权限
// @Param Authorization header string false "认证令牌"
// @Param id query string false "数据权限id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/get [get]
func SysDataPermissionGet(ctx *Context) {
	var entity model.SysDataPermission
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
