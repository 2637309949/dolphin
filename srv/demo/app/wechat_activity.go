// Code generated by dol build. Only Generate by tools if not existed.
// source: wechat_activity.go

package app

import (
	"example/model"
	"example/srv"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
)

// WechatActivityBatchAdd api implementation
// @Summary 添加活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/batch_add [post]
func WechatActivityBatchAdd(ctx *Context) {
	var form []model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	for _, f := range form {
		f.ID = null.StringFromUUID()
	}
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityAdd api implementation
// @Summary 添加活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/add [post]
func WechatActivityAdd(ctx *Context) {
	var form model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFromUUID()
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityBatchDel api implementation
// @Summary 删除活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/batch_del [post]
func WechatActivityBatchDel(ctx *Context) {
	var form []model.WechatActivity
	var err error
	var ret []int64
	var r int64
	if err = ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for _, f := range form {
		r, err = s.ID(f.ID).Delete(&f)
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityDel api implementation
// @Summary 删除活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/del [post]
func WechatActivityDel(ctx *Context) {
	var form model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.ID(form.ID).Delete(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityBatchUpdate api implementation
// @Summary 更新活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/batch_update [post]
func WechatActivityBatchUpdate(ctx *Context) {
	var form []model.WechatActivity
	var err error
	var ret []int64
	var r int64
	if err = ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for _, f := range form {
		r, err = s.ID(f.ID).Update(&f)
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityUpdate api implementation
// @Summary 更新活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/update [post]
func WechatActivityUpdate(ctx *Context) {
	var form model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.ID(form.ID).Update(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityList api implementation
// @Summary 活动分页查询
// @Tags 活动
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Param title query string false "标题筛选"
// @Param hidden query int false "是否隐藏筛选"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/list [get]
func WechatActivityList(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetString("title", "nn")
	q.SetInt("hidden")
	ret, err := ctx.PageSearch(ctx.DB, "wechat_activity", "list", "wechat_activity", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityOne api implementation
// @Summary 获取活动
// @Tags 活动
// @Param Authorization header string false "认证令牌"
// @Param id query string false "活动id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/wechat/activity/one [get]
func WechatActivityOne(ctx *Context) {
	var entity model.WechatActivity
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityIncrease api implementation
// @Summary 增加次数
// @Tags 活动
// @version 1.0
// @Accept application/json
// @Param id body string false "记录id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/v1/wechat/activity/increase [post]
func WechatActivityIncrease(ctx *Context) {
	var form string
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.WechatActivityAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityIncreaseV2 api implementation
// @Summary 增加次数
// @Tags 活动
// @version 2.0
// @Accept application/json
// @Param id body string false "记录id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/v2/wechat/activity/increase_v2 [post]
func WechatActivityIncreaseV2(ctx *Context) {
	var form string
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.WechatActivityAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
