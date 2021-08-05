// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tag_group.go

package api

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysTagGroupAdd api implementation
// @Summary 添加标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body types.SysTagGroup false "标签组信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/group/add [post]
func (ctr *SysTagGroup) SysTagGroupAdd(ctx *Context) {
	var payload types.SysTagGroup
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

// SysTagGroupBatchAdd api implementation
// @Summary 添加标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body []types.SysTagGroup false "标签组信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/tag/group/batch_add [post]
func (ctr *SysTagGroup) SysTagGroupBatchAdd(ctx *Context) {
	var payload []types.SysTagGroup
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

// SysTagGroupDel api implementation
// @Summary 删除标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body types.SysTagGroup false "标签"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/group/del [delete]
func (ctr *SysTagGroup) SysTagGroupDel(ctx *Context) {
	var payload types.SysTagGroup
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.SysTagGroup{
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

// SysTagGroupBatchDel api implementation
// @Summary 删除标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body []types.SysTagGroup false "标签组信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/tag/group/batch_del [delete]
func (ctr *SysTagGroup) SysTagGroupBatchDel(ctx *Context) {
	var payload []types.SysTagGroup
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysTagGroup) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.SysTagGroup{
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

// SysTagGroupUpdate api implementation
// @Summary 更新标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body types.SysTagGroup false "标签组信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/group/update [put]
func (ctr *SysTagGroup) SysTagGroupUpdate(ctx *Context) {
	var payload types.SysTagGroup
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

// SysTagGroupBatchUpdate api implementation
// @Summary 更新标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body []types.SysTagGroup false "标签组信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/tag/group/batch_update [put]
func (ctr *SysTagGroup) SysTagGroupBatchUpdate(ctx *Context) {
	var payload []types.SysTagGroup
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

// SysTagGroupPage api implementation
// @Summary 标签组分页查询
// @Tags 标签组
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/group/page [get]
func (ctr *SysTagGroup) SysTagGroupPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_tag_group_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_tag_group", "page", "sys_tag_group", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGroupGet api implementation
// @Summary 获取标签组信息
// @Tags 标签组
// @Param Authorization header string false "认证令牌"
// @Param id query string false "标签组id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/tag/group/get [get]
func (ctr *SysTagGroup) SysTagGroupGet(ctx *Context) {
	var entity types.SysTagGroup
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
		ctx.Fail(types.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
