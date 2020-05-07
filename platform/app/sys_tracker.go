// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tracker.go

package app

import (
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/model"
)

// SysTrackerPage api implementation
// @Summary 日志分页查询
// @Tags 日志
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/tracker/page [get]
func SysTrackerPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetString("app_name", viper.GetString("app.name"))
	q.SetString("domain", ctx.GetToken().GetDomain())
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.PlatformDB, "sys_tracker", "page", "sys_tracker", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTrackerGet api implementation
// @Summary 获取日志信息
// @Tags 日志
// @Param Authorization header string false "认证令牌"
// @Param id query string false "日志id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/tracker/get [get]
func SysTrackerGet(ctx *Context) {
	var entity model.SysTracker
	id := ctx.Query("id")
	_, err := ctx.PlatformDB.Id(id).Cols("id", "header", "req_body", "res_body").Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ctx.OmitByZero(entity))
}
