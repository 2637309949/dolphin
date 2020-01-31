// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user_template.go

package app

import (
	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/2637309949/dolphin/cli/platform/srv"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/time"
)

// SysUserTemplateAdd api implementation
// @Summary 添加用户模板 
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/template/add [post]
func SysUserTemplateAdd(ctx *Context) {
	var form model.SysUserTemplate
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

// SysUserTemplateDel api implementation
// @Summary 删除用户模板 
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/template/del [delete]
func SysUserTemplateDel(ctx *Context) {
	var form model.SysUserTemplate
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Table(new(model.SysUserTemplate)).In("id", form.ID.String).Update(map[string]interface{}{
		"delete_time": null.TimeFromPtr(time.Now().Value()),
		"delete_by":   null.StringFrom(ctx.GetToken().GetUserID()),
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
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/template/update [put]
func SysUserTemplateUpdate(ctx *Context) {
	var form model.SysUserTemplate
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

// SysUserTemplatePage api implementation
// @Summary 用户模板分页查询 
// @Tags 用户模板
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/template/page [get]
func SysUserTemplatePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page")
	q.SetInt("size")
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
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/user/template/get [get]
func SysUserTemplateGet(ctx *Context) {
	var entity model.SysUserTemplate
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

