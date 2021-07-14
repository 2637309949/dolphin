// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_data_permission.go

package api

import (
	"errors"

	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysDataPermissionAdd api implementation
// @Summary 添加数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysDataPermission false "数据权限信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/data/permission/add [post]
func (ctr *SysDataPermission) SysDataPermissionAdd(ctx *Context) {
	var payload types.SysDataPermission
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

// SysDataPermissionBatchAdd api implementation
// @Summary 添加数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_data_permission body []types.SysDataPermission false "数据权限信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/data/permission/batch_add [post]
func (ctr *SysDataPermission) SysDataPermissionBatchAdd(ctx *Context) {
	var payload []types.SysDataPermission
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

// SysDataPermissionDel api implementation
// @Summary 删除数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_data_permission body types.SysDataPermission false "数据权限"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/data/permission/del [delete]
func (ctr *SysDataPermission) SysDataPermissionDel(ctx *Context) {
	var payload types.SysDataPermission
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.SysDataPermission{
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

// SysDataPermissionBatchDel api implementation
// @Summary 删除数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_data_permission body []types.SysDataPermission false "数据权限"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/data/permission/batch_del [delete]
func (ctr *SysDataPermission) SysDataPermissionBatchDel(ctx *Context) {
	var payload []types.SysDataPermission
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysDataPermission) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.SysDataPermission{
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

// SysDataPermissionUpdate api implementation
// @Summary 更新数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysRole false "数据权限信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/data/permission/update [put]
func (ctr *SysDataPermission) SysDataPermissionUpdate(ctx *Context) {
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

// SysDataPermissionBatchUpdate api implementation
// @Summary 更新数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_data_permission body []types.SysDataPermission false "数据权限信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/data/permission/batch_update [put]
func (ctr *SysDataPermission) SysDataPermissionBatchUpdate(ctx *Context) {
	var payload []types.SysDataPermission
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

// SysDataPermissionPage api implementation
// @Summary 数据权限分页查询
// @Tags 数据权限
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/data/permission/page [get]
func (ctr *SysDataPermission) SysDataPermissionPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_data_permission_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_data_permission", "page", "sys_data_permission", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDataPermissionGet api implementation
// @Summary 获取数据权限信息
// @Tags 数据权限
// @Param Authorization header string false "认证令牌"
// @Param id query string false "数据权限id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/data/permission/get [get]
func (ctr *SysDataPermission) SysDataPermissionGet(ctx *Context) {
	var entity types.SysDataPermission
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
