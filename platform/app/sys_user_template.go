// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user_template.go

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

// SysUserTemplateAdd api implementation
// @Summary 添加用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/add [post]
func SysUserTemplateAdd(ctx *Context) {
	var payload model.SysUserTemplate
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
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
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateBatchAdd api implementation
// @Summary 添加用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body []model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/template/batch_add [post]
func SysUserTemplateBatchAdd(ctx *Context) {
	var payload []model.SysUserTemplate
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

// SysUserTemplateDel api implementation
// @Summary 删除用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/del [delete]
func SysUserTemplateDel(ctx *Context) {
	var payload model.SysUserTemplate
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysUserTemplate{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		Updater:    null.StringFrom(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateBatchDel api implementation
// @Summary 删除用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body []model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/template/batch_del [delete]
func SysUserTemplateBatchDel(ctx *Context) {
	var payload []model.SysUserTemplate
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysUserTemplate) string { return form.ID.String }).([]string)
	ret, err := ctx.DB.In("id", ids).Update(&model.SysUserTemplate{
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

// SysUserTemplateUpdate api implementation
// @Summary 更新用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/update [put]
func SysUserTemplateUpdate(ctx *Context) {
	var payload model.SysUserTemplate
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.Updater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.String).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserTemplateBatchUpdate api implementation
// @Summary 更新用户模板
// @Tags 用户模板
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_user_template body []model.SysUserTemplate false "用户模板信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/template/batch_update [put]
func SysUserTemplateBatchUpdate(ctx *Context) {
	var payload []model.SysUserTemplate
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].Updater = null.StringFrom(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.String).Update(&payload[i])
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/page [get]
func SysUserTemplatePage(ctx *Context) {
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
	ret, err := ctx.PageSearch(ctx.DB, "sys_user_template", "page", "sys_user_template", q.Value())
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/template/get [get]
func SysUserTemplateGet(ctx *Context) {
	var entity model.SysUserTemplate
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
