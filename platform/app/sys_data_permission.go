// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_data_permission.go

package app

import (
	"errors"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/platform/model"
)

// SysDataPermissionAdd api implementation
// @Summary 添加数据权限
// @Tags 数据权限
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysDataPermission false "数据权限信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/add [post]
func SysDataPermissionAdd(ctx *Context) {
	var payload model.SysDataPermission
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
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
// @Param sys_data_permission body []model.SysDataPermission false "数据权限信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/data/permission/batch_add [post]
func SysDataPermissionBatchAdd(ctx *Context) {
	var payload []model.SysDataPermission
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].ID = null.StringFromUUID()
		payload[i].CreateTime = null.TimeFrom(time.Now().Value())
		payload[i].CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].DelFlag = null.IntFrom(0)
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
// @Param sys_data_permission body model.SysDataPermission false "数据权限"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/del [delete]
func SysDataPermissionDel(ctx *Context) {
	var payload model.SysDataPermission
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysDataPermission{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
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
// @Param sys_data_permission body []model.SysDataPermission false "数据权限"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/data/permission/batch_del [delete]
func SysDataPermissionBatchDel(ctx *Context) {
	var payload []model.SysDataPermission
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysDataPermission) string { return form.ID.String }).([]string)
	ret, err := ctx.DB.In("id", ids).Update(&model.SysDataPermission{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
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
// @Param user body model.SysRole false "数据权限信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/update [put]
func SysDataPermissionUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.String).Update(&payload)
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
// @Param sys_data_permission body []model.SysDataPermission false "数据权限信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/data/permission/batch_update [put]
func SysDataPermissionBatchUpdate(ctx *Context) {
	var payload []model.SysDataPermission
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.String).Update(&payload[i])
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/page [get]
func SysDataPermissionPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_data_permission_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_data_permission", "page", "sys_data_permission", q.Value())
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/data/permission/get [get]
func SysDataPermissionGet(ctx *Context) {
	var entity model.SysDataPermission
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
