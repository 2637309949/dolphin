// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tag_group.go

package app

import (
	"github.com/2637309949/dolphin/cli/platform/model"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/time"
)

// SysTagGroupAdd api implementation
// @Summary 添加标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param role body model.SysTagGroup false "标签组信息"
// @Failure 403 {object} model.Response
// @Router /api/sys/tag/group/add [post]
func SysTagGroupAdd(ctx *Context) {
	var form model.SysTagGroup
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFromUUID()
	form.CreateTime = null.TimeFromPtr(time.Now().Value())
	form.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGroupUpdate api implementation
// @Summary 更新标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param role body model.SysTagGroup false "标签组信息"
// @Failure 403 {object} model.Response
// @Router /api/sys/tag/group/update [post]
func SysTagGroupUpdate(ctx *Context) {
	var form model.SysTagGroup
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

// SysTagGroupPage api implementation
// @Summary 标签组分页查询
// @Tags 标签组
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Router /api/sys/tag/group/page [get]
func SysTagGroupPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page")
	q.SetInt("size")
	ret, err := ctx.PageSearch(ctx.DB, "sys_tag_group", "page", "sys_tag_group", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGroupGet api implementation
// @Summary 获取标签组信息
// @Tags 标签组
// @Param Authorization header string false "认证令牌"
// @Param id query string false "标签组id"
// @Failure 403 {object} model.Response
// @Router /api/sys/tag/group/get [get]
func SysTagGroupGet(ctx *Context) {
	var entity model.SysTagGroup
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
