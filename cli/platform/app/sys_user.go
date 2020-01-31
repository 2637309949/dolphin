// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user.go

package app

import (
	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/2637309949/dolphin/cli/platform/srv"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/time"
)

// SysUserAdd api implementation
// @Summary 添加用户 
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "用户信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/add [post]
func SysUserAdd(ctx *Context) {
	var form model.SysRole
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFromUUID()
	form.CreateTime = null.TimeFromPtr(time.Now().Value())
	form.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	form.UpdateTime = null.TimeFromPtr(time.Now().Value())
	form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserDel api implementation
// @Summary 删除用户 
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "用户信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/del [delete]
func SysUserDel(ctx *Context) {
	var form model.SysRole
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Table(new(model.SysRole)).In("id", form.ID.String).Update(map[string]interface{}{
		"delete_time": null.TimeFromPtr(time.Now().Value()),
		"delete_by":   null.StringFrom(ctx.GetToken().GetUserID()),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserUpdate api implementation
// @Summary 更新用户 
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "用户信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/update [put]
func SysUserUpdate(ctx *Context) {
	var form model.SysRole
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	form.UpdateTime = null.TimeFromPtr(time.Now().Value())
	ret, err := ctx.DB.ID(form.ID).Update(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserPage api implementation
// @Summary 用户分页查询 
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/page [get]
func SysUserPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page")
	q.SetInt("size")
	ret, err := ctx.PageSearch(ctx.DB, "sys_user", "page", "sys_user", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserGet api implementation
// @Summary 获取用户信息 
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Param id query string false "用户id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/get [get]
func SysUserGet(ctx *Context) {
	var entity model.SysUser
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserLogout api implementation
// @Summary 用户退出登录 
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/logout [get]
func SysUserLogout(ctx *Context) {
	q := ctx.TypeQuery()
	ret, err := srv.SysUserAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

