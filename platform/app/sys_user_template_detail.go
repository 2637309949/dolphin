// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user_template_detail.go

package app

import (
	"errors"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysUserTemplateDetailAdd api implementation
// @Summary 添加用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body model.SysUserTemplateDetail false "用户模板详情信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/detail/add [post]
func SysUserTemplateDetailAdd(ctx *Context) {
	var payload model.SysUserTemplateDetail
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
// @Param sys_user_template_detail body []model.SysUserTemplateDetail false "用户模板详情信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/template/detail/batch_add [post]
func SysUserTemplateDetailBatchAdd(ctx *Context) {
	var payload []model.SysUserTemplateDetail
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
// @Param sys_user_template_detail body model.SysUserTemplateDetail false "用户模板详情信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/detail/del [delete]
func SysUserTemplateDetailDel(ctx *Context) {
	var payload model.SysUserTemplateDetail
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysUserTemplateDetail{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
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
// @Param sys_user_template_detail body []model.SysUserTemplateDetail false "用户模板详情信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/template/detail/batch_del [delete]
func SysUserTemplateDetailBatchDel(ctx *Context) {
	var payload []model.SysUserTemplateDetail
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysUserTemplateDetail) string { return form.ID.String }).([]string)
	ret, err := ctx.DB.In("id", ids).Update(&model.SysUserTemplateDetail{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
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
// @Param sys_user_template_detail body model.SysUserTemplateDetail false "用户模板详情信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/detail/update [put]
func SysUserTemplateDetailUpdate(ctx *Context) {
	var payload model.SysUserTemplateDetail
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

// SysUserTemplateDetailBatchUpdate api implementation
// @Summary 更新用户模板详情
// @Tags 用户模板详情
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template_detail body []model.SysUserTemplateDetail false "用户模板详情信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/template/detail/batch_update [put]
func SysUserTemplateDetailBatchUpdate(ctx *Context) {
	var payload []model.SysUserTemplateDetail
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

// SysUserTemplateDetailPage api implementation
// @Summary 用户模板详情分页查询
// @Tags 用户模板详情
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/detail/page [get]
func SysUserTemplateDetailPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("temp_id")
	q.SetRule("sys_user_template_detail_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_user_template_detail", "page", "sys_user_template_detail", q.Value())
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/detail/get [get]
func SysUserTemplateDetailGet(ctx *Context) {
	var entity model.SysUserTemplateDetail
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
