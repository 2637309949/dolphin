// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tracker.go

package app

import (
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util/slice"
)

// SysTrackerPage api implementation
// @Summary 日志分页查询
// @Tags 日志
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/tracker/page [get]
func SysTrackerPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetString("sort", "sys_tracker.update_time desc")
	q.SetRule("sys_tracker_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_tracker", "page", "sys_tracker", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}

	uids := slice.GetFieldSliceByName(ret.Data, "user_id").([]string)
	users := []model.SysUser{}
	err = ctx.PlatformDB.Cols("id", "name").In("id", uids).Find(&users)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}

	err = slice.PatchSliceByField(ret.Data, users, "user_id", "id", "user_name#name")(&ret.Data)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}

	if ctx.QueryBool("__export__") {
		cfg := NewBuildExcelConfig(ret.Data)
		cfg.Format = OptionsetsFormat(ctx.DB)
		ctx.SuccessWithExcel(cfg)
		return
	}
	ctx.Success(ret)
}

// SysTrackerGet api implementation
// @Summary 获取日志信息
// @Tags 日志
// @Param Authorization header string false "认证令牌"
// @Param id query string false "日志id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/tracker/get [get]
func SysTrackerGet(ctx *Context) {
	var entity model.SysTracker
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Cols("id", "header", "req_body", "res_body").Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ctx.OmitByZero(entity))
}
