// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user_template.go

package api

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysUserTemplateAdd api implementation
// @Summary 添加用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body types.SysUserTemplate false "用户模板信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/add [post]
func (ctr *SysUserTemplate) SysUserTemplateAdd(ctx *Context) {
	var payload types.SysUserTemplate
	if err := ctx.ShouldBindWith(&payload); err != nil {
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
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateBatchAdd api implementation
// @Summary 批量添加用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body []types.SysUserTemplate false "用户模板信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/user/template/batch_add [post]
func (ctr *SysUserTemplate) SysUserTemplateBatchAdd(ctx *Context) {
	var payload []types.SysUserTemplate
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

// SysUserTemplateDel api implementation
// @Summary 删除用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body types.SysUserTemplate false "用户模板信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/del [delete]
func (ctr *SysUserTemplate) SysUserTemplateDel(ctx *Context) {
	var payload types.SysUserTemplate
	if err := ctx.ShouldBindWith(&payload); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.SysUserTemplate{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateBatchDel api implementation
// @Summary 批量删除用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body []types.SysUserTemplate false "用户模板信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/user/template/batch_del [delete]
func (ctr *SysUserTemplate) SysUserTemplateBatchDel(ctx *Context) {
	var payload []types.SysUserTemplate
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysUserTemplate) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.SysUserTemplate{
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

// SysUserTemplateUpdate api implementation
// @Summary 更新用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body types.SysUserTemplate false "用户模板信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/update [put]
func (ctr *SysUserTemplate) SysUserTemplateUpdate(ctx *Context) {
	var payload types.SysUserTemplate
	if err := ctx.ShouldBindWith(&payload); err != nil {
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateBatchUpdate api implementation
// @Summary 批量更新用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body []types.SysUserTemplate false "用户模板信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/user/template/batch_update [put]
func (ctr *SysUserTemplate) SysUserTemplateBatchUpdate(ctx *Context) {
	var payload []types.SysUserTemplate
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

// SysUserTemplatePage api implementation
// @Summary 用户模板分页查询
// @Tags 用户模板
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/page [get]
func (ctr *SysUserTemplate) SysUserTemplatePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_user_template_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_user_template", "page", "sys_user_template", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateGet api implementation
// @Summary 获取用户模板信息
// @Tags 用户模板
// @Param Authorization header string false "认证令牌"
// @Param id query string false "用户模板id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/user/template/get [get]
func (ctr *SysUserTemplate) SysUserTemplateGet(ctx *Context) {
	var entity types.SysUserTemplate
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
