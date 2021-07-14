// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user_template_detail.go

package api

import (
	"errors"

	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysUserTemplateDetailAdd api implementation
// @Summary 添加用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body types.SysUserTemplateDetail false "用户模板详情信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/detail/add [post]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailAdd(ctx *Context) {
	var payload types.SysUserTemplateDetail
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

// SysUserTemplateDetailBatchAdd api implementation
// @Summary 添加用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body []types.SysUserTemplateDetail false "用户模板详情信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/user/template/detail/batch_add [post]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailBatchAdd(ctx *Context) {
	var payload []types.SysUserTemplateDetail
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

// SysUserTemplateDetailDel api implementation
// @Summary 删除用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body types.SysUserTemplateDetail false "用户模板详情信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/detail/del [delete]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailDel(ctx *Context) {
	var payload types.SysUserTemplateDetail
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.SysUserTemplateDetail{
		UpdateTime: null.TimeFrom(time.Now()),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateDetailBatchDel api implementation
// @Summary 删除用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body []types.SysUserTemplateDetail false "用户模板详情信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/user/template/detail/batch_del [delete]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailBatchDel(ctx *Context) {
	var payload []types.SysUserTemplateDetail
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysUserTemplateDetail) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.SysUserTemplateDetail{
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

// SysUserTemplateDetailUpdate api implementation
// @Summary 更新用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body types.SysUserTemplateDetail false "用户模板详情信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/detail/update [put]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailUpdate(ctx *Context) {
	var payload types.SysUserTemplateDetail
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

// SysUserTemplateDetailBatchUpdate api implementation
// @Summary 更新用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body []types.SysUserTemplateDetail false "用户模板详情信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/user/template/detail/batch_update [put]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailBatchUpdate(ctx *Context) {
	var payload []types.SysUserTemplateDetail
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

// SysUserTemplateDetailPage api implementation
// @Summary 用户模板详情分页查询
// @Tags 用户模板详情
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/detail/page [get]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("temp_id")
	q.SetRule("sys_user_template_detail_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_user_template_detail", "page", "sys_user_template_detail", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateDetailGet api implementation
// @Summary 获取用户模板详情信息
// @Tags 用户模板详情
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "用户模板详情id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/detail/get [get]
func (ctr *SysUserTemplateDetail) SysUserTemplateDetailGet(ctx *Context) {
	var entity types.SysUserTemplateDetail
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
