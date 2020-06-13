// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user_template.go

package app

import (
	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/platform/model"
)

// SysUserTemplateAdd api implementation
// @Summary 添加用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/add [post]
func SysUserTemplateAdd(ctx *Context) {
	var payload model.SysUserTemplate
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

// SysUserTemplateDel api implementation
// @Summary 删除用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/del [delete]
func SysUserTemplateDel(ctx *Context) {
	var payload model.SysUserTemplate
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysUserTemplate{
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

// SysUserTemplateUpdate api implementation
// @Summary 更新用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/update [put]
func SysUserTemplateUpdate(ctx *Context) {
	var payload model.SysUserTemplate
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

// SysUserTemplatePage api implementation
// @Summary 用户模板分页查询
// @Tags 用户模板
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/page [get]
func SysUserTemplatePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetRule("sys_user_template_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_user_template", "page", "sys_user_template", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateGet api implementation
// @Summary 获取用户模板信息
// @Tags 用户模板
// @Param Authorization header string false "认证令牌"
// @Param id query string false "用户模板id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/get [get]
func SysUserTemplateGet(ctx *Context) {
	var entity model.SysUserTemplate
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
