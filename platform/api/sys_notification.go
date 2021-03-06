// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_notification.go

package api

import (
	"errors"

	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysNotificationAdd api implementation
// @Summary 添加站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param notification body types.SysNotification false "站内消息信息"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/notification/add [post]
func (ctr *SysNotification) SysNotificationAdd(ctx *Context) {
	var payload types.SysNotification
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFrom(time.Now())
	payload.Creater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now())
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationBatchAdd api implementation
// @Summary 添加站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body []types.SysNotification false "站内消息信息"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/notification/batch_add [post]
func (ctr *SysNotification) SysNotificationBatchAdd(ctx *Context) {
	var payload []types.SysNotification
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].CreateTime = null.TimeFrom(time.Now())
		payload[i].Creater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFrom(time.Now())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
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

// SysNotificationDel api implementation
// @Summary 删除站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body types.SysNotification false "站内消息"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/notification/del [delete]
func (ctr *SysNotification) SysNotificationDel(ctx *Context) {
	var payload types.SysNotification
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.SysNotification{
		UpdateTime: null.TimeFrom(time.Now()),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationBatchDel api implementation
// @Summary 删除站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body []types.SysNotification false "站内消息"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/notification/batch_del [delete]
func (ctr *SysNotification) SysNotificationBatchDel(ctx *Context) {
	var payload []types.SysNotification
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysNotification) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.SysNotification{
		UpdateTime: null.TimeFrom(time.Now()),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationUpdate api implementation
// @Summary 更新站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param notification body types.SysRole false "站内消息信息"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/notification/update [put]
func (ctr *SysNotification) SysNotificationUpdate(ctx *Context) {
	var payload types.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now())
	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationBatchUpdate api implementation
// @Summary 更新站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body []types.SysNotification false "站内消息信息"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/notification/batch_update [put]
func (ctr *SysNotification) SysNotificationBatchUpdate(ctx *Context) {
	var payload []types.SysNotification
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.Int64).Update(&payload[i])
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

// SysNotificationPage api implementation
// @Summary 站内消息分页查询
// @Tags 站内消息
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/notification/page [get]
func (ctr *SysNotification) SysNotificationPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_notification_page")
	q.SetTags()
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_notification", "page", "sys_notification", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationGet api implementation
// @Summary 获取站内消息信息
// @Tags 站内消息
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "站内消息id"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/notification/get [get]
func (ctr *SysNotification) SysNotificationGet(ctx *Context) {
	var entity types.SysNotification
	err := ctx.ShouldBindWith(&entity)
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
