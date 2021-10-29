// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_area.go

package api

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysAreaAdd api implementation
// @Summary 添加区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysArea false "区域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/area/add [post]
func (ctr *SysArea) SysAreaAdd(ctx *Context) {
	var payload types.SysArea
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}

	payload.CreateTime = null.TimeFromNow()
	payload.Creater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
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

// SysAreaBatchAdd api implementation
// @Summary 批量添加区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body []types.SysArea false "区域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/area/batch_add [post]
func (ctr *SysArea) SysAreaBatchAdd(ctx *Context) {
	var payload []types.SysArea
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].CreateTime = null.TimeFromNow()
		payload[i].Creater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFromNow()
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

// SysAreaDel api implementation
// @Summary 删除区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body types.SysArea false "区域"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/area/del [delete]
func (ctr *SysArea) SysAreaDel(ctx *Context) {
	var payload types.SysArea
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.SysArea{
		UpdateTime: null.TimeFromNow(),
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

// SysAreaBatchDel api implementation
// @Summary 批量删除文章
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body []types.SysArea false "区域"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/area/batch_del [delete]
func (ctr *SysArea) SysAreaBatchDel(ctx *Context) {
	var payload []types.SysArea
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysArea) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.SysArea{
		UpdateTime: null.TimeFromNow(),
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

// SysAreaUpdate api implementation
// @Summary 更新区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysRole false "区域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/area/update [put]
func (ctr *SysArea) SysAreaUpdate(ctx *Context) {
	var payload types.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAreaBatchUpdate api implementation
// @Summary 批量更新文章
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body []types.SysArea false "区域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/area/batch_update [put]
func (ctr *SysArea) SysAreaBatchUpdate(ctx *Context) {
	var payload []types.SysArea
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
		payload[i].UpdateTime = null.TimeFromNow()
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

// SysAreaPage api implementation
// @Summary 区域分页查询
// @Tags 区域
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/area/page [get]
func (ctr *SysArea) SysAreaPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_area_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.DB.PageSearch(ctx.DB, "sys_area", "page", "sys_area", q.Value())
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
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/area/get [get]
func (ctr *SysArea) SysAreaGet(ctx *Context) {
	var entity types.SysArea
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
		ctx.Fail(errors.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
