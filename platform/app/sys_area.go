// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_area.go

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

// SysAreaAdd api implementation
// @Summary 添加区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysArea false "区域信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/area/add [post]
func SysAreaAdd(ctx *Context) {
	var payload model.SysArea
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

// SysAreaBatchAdd api implementation
// @Summary 添加区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body []model.SysArea false "区域信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/area/batch_add [post]
func SysAreaBatchAdd(ctx *Context) {
	var payload []model.SysArea
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

// SysAreaDel api implementation
// @Summary 删除区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body model.SysArea false "区域"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/area/del [delete]
func SysAreaDel(ctx *Context) {
	var payload model.SysArea
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysArea{
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

// SysAreaBatchDel api implementation
// @Summary 删除文章
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body []model.SysArea false "区域"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/area/batch_del [delete]
func SysAreaBatchDel(ctx *Context) {
	var payload []model.SysArea
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysArea) string { return form.ID.String }).([]string)
	ret, err := ctx.DB.In("id", ids).Update(&model.SysArea{
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

// SysAreaUpdate api implementation
// @Summary 更新区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "区域信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/area/update [put]
func SysAreaUpdate(ctx *Context) {
	var payload model.SysRole
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

// SysAreaBatchUpdate api implementation
// @Summary 更新文章
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body []model.SysArea false "区域信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/area/batch_update [put]
func SysAreaBatchUpdate(ctx *Context) {
	var payload []model.SysArea
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
		payload[i].Updater = null.StringFrom(ctx.GetToken().GetUserID())
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

// SysAreaPage api implementation
// @Summary 区域分页查询
// @Tags 区域
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/area/page [get]
func SysAreaPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_area_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_area", "page", "sys_area", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAreaGet api implementation
// @Summary 获取区域信息
// @Tags 区域
// @Param Authorization header string false "认证令牌"
// @Param id query string false "区域id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/area/get [get]
func SysAreaGet(ctx *Context) {
	var entity model.SysArea
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
