// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tag.go

package api

import (
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/thoas/go-funk"
)

// SysTagAdd api implementation
// @Summary 添加标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body types.SysTag false "标签信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/add [post]
func (ctr *SysTag) SysTagAdd(ctx *Context) {
	var payload types.SysTag
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFromNow()
	payload.Creater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)

	db := ctx.MustDB()
	ret, err := db.Insert(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagBatchAdd api implementation
// @Summary 批量添加标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body []types.SysTag false "标签信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/tag/batch_add [post]
func (ctr *SysTag) SysTagBatchAdd(ctx *Context) {
	var payload []types.SysTag
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].CreateTime = null.TimeFromNow()
		payload[i].Creater = null.IntFromStr(ctx.MustToken().GetUserID())
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.MustToken().GetUserID())
		payload[i].IsDelete = null.IntFrom(0)
	}

	db := ctx.MustDB()
	ret, err := db.Insert(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagDel api implementation
// @Summary 删除标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body types.SysUserTemplate false "标签"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/del [delete]
func (ctr *SysTag) SysTagDel(ctx *Context) {
	var payload types.SysUserTemplate
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	ret, err := db.In("id", payload.ID.Int64).Update(&types.SysUserTemplate{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagBatchDel api implementation
// @Summary 批量删除标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body []types.SysTag false "删除标签"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/batch_del [delete]
func (ctr *SysTag) SysTagBatchDel(ctx *Context) {
	var payload []*types.SysTag
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	var ids = funk.Map(payload, func(form types.SysTag) int64 { return form.ID.Int64 }).([]int64)
	ret, err := db.In("id", ids).Update(&types.SysTag{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagUpdate api implementation
// @Summary 更新标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body types.SysTag false "标签信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/update [put]
func (ctr *SysTag) SysTagUpdate(ctx *Context) {
	var payload types.SysTag
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()

	db := ctx.MustDB()
	ret, err := db.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagBatchUpdate api implementation
// @Summary 批量更新标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body []types.SysTag false "标签信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/tag/batch_update [put]
func (ctr *SysTag) SysTagBatchUpdate(ctx *Context) {
	var payload []types.SysTag
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	s := db.NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.MustToken().GetUserID())
		r, err = s.ID(payload[i].ID.Int64).Update(&payload[i])
		if err != nil {
			s.Rollback()
			logrus.Error(ctx, err)
			ctx.Fail(err)
			return
		}
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	err = s.Commit()
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagPage api implementation
// @Summary 标签分页查询
// @Tags 标签
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/page [get]
func (ctr *SysTag) SysTagPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_tag_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()

	db := ctx.MustDB()
	ret, err := ctr.Srv.DB.PageSearch(db, "sys_tag", "page", "sys_tag", q.Value())
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGet api implementation
// @Summary 获取标签信息
// @Tags 标签
// @Param Authorization header string false "认证令牌"
// @Param id query string false "标签id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/get [get]
func (ctr *SysTag) SysTagGet(ctx *Context) {
	var entity types.SysTag
	err := ctx.ShouldBindWith(&entity)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	ext, err := db.Get(&entity)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(errors.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
