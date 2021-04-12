// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_schedule.go

package app

import (
	"errors"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/thoas/go-funk"
)

// SysScheduleAdd api implementation
// @Summary 添加调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_schedule body model.SysSchedule false "调度信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/schedule/add [post]
func SysScheduleAdd(ctx *Context) {
	var payload model.SysSchedule
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFrom(time.Now().Value())
	payload.Creater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.Updater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysScheduleBatchAdd api implementation
// @Summary 添加调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_schedule body []model.SysSchedule false "调度信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/schedule/batch_add [post]
func SysScheduleBatchAdd(ctx *Context) {
	var payload []model.SysSchedule
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].ID = null.StringFromUUID()
		payload[i].CreateTime = null.TimeFrom(time.Now().Value())
		payload[i].Creater = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].Updater = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].IsDelete = null.IntFrom(0)
	}
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysScheduleDel api implementation
// @Summary 删除调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_schedule body model.SysSchedule false "调度"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/schedule/del [delete]
func SysScheduleDel(ctx *Context) {
	var payload model.SysSchedule
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysSchedule{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		Updater:    null.StringFrom(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysScheduleBatchDel api implementation
// @Summary 删除调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []model.SysSchedule false "调度信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/schedule/batch_del [delete]
func SysScheduleBatchDel(ctx *Context) {
	var payload []*model.SysSchedule
	var ids []string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form model.SysSchedule) {
		ids = append(ids, form.ID.String)
	})
	ret, err := ctx.DB.In("id", ids).Update(&model.SysSchedule{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		Updater:    null.StringFrom(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysScheduleUpdate api implementation
// @Summary 更新调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_schedule body model.SysSchedule false "调度信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/schedule/update [put]
func SysScheduleUpdate(ctx *Context) {
	var payload model.SysSchedule
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.String).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysScheduleBatchUpdate api implementation
// @Summary 更新调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_schedule body []model.SysSchedule false "调度信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/schedule/batch_update [post]
func SysScheduleBatchUpdate(ctx *Context) {
	var payload []model.SysSchedule
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].Updater = null.StringFrom(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.String).Update(&payload[i])
		if err != nil {
			s.Rollback()
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	err = s.Commit()
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysSchedulePage api implementation
// @Summary 调度分页查询
// @Tags 调度
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/schedule/page [get]
func SysSchedulePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_schedule_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_schedule", "page", "sys_schedule", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysScheduleGet api implementation
// @Summary 获取调度信息
// @Tags 调度
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "调度id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/schedule/get [get]
func SysScheduleGet(ctx *Context) {
	var entity model.SysSchedule
	err := ctx.ShouldBindQuery(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ext, err := ctx.DB.Get(&entity)
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
