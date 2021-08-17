// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tracker.go

package api

import (
	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
)

// SysTrackerPage api implementation
// @Summary 日志分页查询
// @Tags 日志
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tracker/page [get]
func (ctr *SysTracker) SysTrackerPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("sort", "sys_tracker.update_time desc")
	q.SetRule("sys_tracker_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	if ctr.Srv.Check(ctx.Request) {
		ctr.Srv.SetOptionsetsFormat(OptionsetsFormat(ctx.DB))
		ret, err := ctr.Srv.PageExport(ctx.DB, "sys_tracker", "page", "sys_tracker", q.Value(), ctr.Srv.PageFormatter(ctx.PlatformDB))
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		ctx.Success(ret)
		return
	}
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_tracker", "page", "sys_tracker", q.Value(), ctr.Srv.PageFormatter(ctx.PlatformDB))
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if err != nil {
		logrus.Error(err)
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
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tracker/get [get]
func (ctr *SysTracker) SysTrackerGet(ctx *Context) {
	var entity types.SysTracker
	err := ctx.ShouldBindWith(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ext, err := ctx.DB.Cols("id", "header", "req_body", "res_body").Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(types.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
