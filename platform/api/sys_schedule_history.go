// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_schedule_history.go

package api

import (
	"github.com/sirupsen/logrus"
)

// SysScheduleHistoryPage api implementation
// @Summary 调度分页查询
// @Tags 调度历史
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/schedule/history/page [get]
func (ctr *SysScheduleHistory) SysScheduleHistoryPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_schedule_history_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.DB.PageSearch(ctx.MustDB(), "sys_schedule_history", "page", "sys_schedule", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
