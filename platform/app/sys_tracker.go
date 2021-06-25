// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tracker.go

package app

import (
	"errors"

	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
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
	ret, err := ctx.PageSearch(ctx.DB, "sys_tracker", "page", "sys_tracker", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}

	if uids, ok := slice.GetFieldSliceByName(ret.Data, "user_id").([]string); ok {
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
func (ctr *SysTracker) SysTrackerGet(ctx *Context) {
	var entity model.SysTracker
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
		ctx.Fail(errors.New("not found"))
		return
	}
	ctx.Success(entity)
}
