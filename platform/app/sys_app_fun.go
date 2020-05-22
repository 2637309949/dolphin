// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_app_fun.go

package app

import (
	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/platform/model"
)

// SysAppFunAdd api implementation
// @Summary 添加APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysAppFun false "APP功能信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/app/fun/add [post]
func SysAppFunAdd(ctx *Context) {
	var payload model.SysAppFun
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

// SysAppFunDel api implementation
// @Summary 删除APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_app_fun body model.SysAppFun false "APP功能"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/app/fun/del [delete]
func SysAppFunDel(ctx *Context) {
	var payload model.SysAppFun
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysAppFun{
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

// SysAppFunUpdate api implementation
// @Summary 更新APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "APP功能信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/app/fun/update [put]
func SysAppFunUpdate(ctx *Context) {
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

// SysAppFunPage api implementation
// @Summary APP功能分页查询
// @Tags APP功能
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/app/fun/page [get]
func SysAppFunPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_app_fun", "page", "sys_app_fun", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunTree api implementation
// @Summary 菜单树形结构
// @Tags APP功能
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Fail
// @Router /api/sys/app/fun/tree [get]
func SysAppFunTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetTags()
	ret, err := ctx.TreeSearch(ctx.DB, "sys_app_fun", "tree", "sys_app_fun", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunGet api implementation
// @Summary 获取APP功能信息
// @Tags APP功能
// @Param Authorization header string false "认证令牌"
// @Param id query string false "APP功能id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/app/fun/get [get]
func SysAppFunGet(ctx *Context) {
	var entity model.SysAppFun
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
