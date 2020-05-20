// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_domain.go

package app

import (
	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/model"
)

// SysDomainAdd api implementation
// @Summary 添加域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysDomain false "域信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/domain/add [post]
func SysDomainAdd(ctx *Context) {
	var payload model.SysDomain
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
	payload.AppName = null.StringFrom(viper.GetString("app.name"))
	ret, err := ctx.PlatformDB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainDel api implementation
// @Summary 删除域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_domain body model.SysDomain false "域"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/domain/del [delete]
func SysDomainDel(ctx *Context) {
	var payload model.SysDomain
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.PlatformDB.In("id", payload.ID.String).Update(&model.SysDomain{
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

// SysDomainUpdate api implementation
// @Summary 更新域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "域信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/domain/update [put]
func SysDomainUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.PlatformDB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainPage api implementation
// @Summary 域分页查询
// @Tags 域
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/domain/page [get]
func SysDomainPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetString("app_name", viper.GetString("app.name"))
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.PlatformDB, "sys_domain", "page", "sys_domain", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainGet api implementation
// @Summary 获取域信息
// @Tags 域
// @Param Authorization header string false "认证令牌"
// @Param id query string false "域id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/domain/get [get]
func SysDomainGet(ctx *Context) {
	var entity model.SysDomain
	id := ctx.Query("id")
	_, err := ctx.PlatformDB.ID(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
