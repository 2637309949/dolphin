// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_optionset.go

package app

import (
	"fmt"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysOptionsetAdd api implementation
// @Summary 添加字典
// @Tags 字典
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysOptionset false "字典信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/optionset/add [post]
func SysOptionsetAdd(ctx *Context) {
	var payload model.SysOptionset
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

// SysOptionsetDel api implementation
// @Summary 删除字典
// @Tags 字典
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_optionset body model.SysOptionset false "字典"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/optionset/del [delete]
func SysOptionsetDel(ctx *Context) {
	var payload model.SysOptionset
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysOptionset{
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

// SysOptionsetUpdate api implementation
// @Summary 更新字典
// @Tags 字典
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "字典信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/optionset/update [put]
func SysOptionsetUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysOptionsetPage api implementation
// @Summary 字典分页查询
// @Tags 字典
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/optionset/page [get]
func SysOptionsetPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_optionset", "page", "sys_optionset", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysOptionsetGet api implementation
// @Summary 获取字典信息
// @Tags 字典
// @Param Authorization header string false "认证令牌"
// @Param id query string false "字典id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/optionset/get [get]
func SysOptionsetGet(ctx *Context) {
	var entity model.SysOptionset
	id := ctx.Query("id")
	code := ctx.Query("code")
	sql := "select id, name, code, value from sys_optionset where del_flag=0"
	if id != "" {
		sql += fmt.Sprintf(" and id='%v'", id)
	}
	if code != "" {
		sql += fmt.Sprintf(" and code='%v'", code)
	}
	_, err := ctx.DB.Sql(sql).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ctx.OmitByZero(entity))
}
