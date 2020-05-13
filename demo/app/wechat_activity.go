// Code generated by dol build. Only Generate by tools if not existed.
// source: wechat_activity.go

package app

import (
	"demo/model"
	"demo/srv"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// WechatActivityBatchAdd api implementation
// @Summary 添加活动
// @Tags 微信活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/batch_add [post]
func WechatActivityBatchAdd(ctx *Context) {
	var payload []model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form model.WechatActivity) {
		form.ID = null.StringFromUUID()
		form.CreateTime = null.TimeFrom(time.Now().Value())
		form.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
		form.UpdateTime = null.TimeFrom(time.Now().Value())
		form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		form.DelFlag = null.IntFrom(0)
	})
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityAdd api implementation
// @Summary 添加活动
// @Tags 微信活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/add [post]
func WechatActivityAdd(ctx *Context) {
	var payload model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
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
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityBatchDel api implementation
// @Summary 删除活动
// @Tags 微信活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/batch_del [delete]
func WechatActivityBatchDel(ctx *Context) {
	var payload []model.WechatActivity
	var ids []string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form model.WechatActivity) {
		ids = append(ids, form.ID.String)
	})
	ret, err := ctx.DB.In("id", ids).Update(&model.WechatActivity{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityDel api implementation
// @Summary 删除活动
// @Tags 微信活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/del [delete]
func WechatActivityDel(ctx *Context) {
	var payload model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.WechatActivity{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityBatchUpdate api implementation
// @Summary 更新活动
// @Tags 微信活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/batch_update [put]
func WechatActivityBatchUpdate(ctx *Context) {
	var payload []model.WechatActivity
	var err error
	var ret []int64
	var r int64
	if err = ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	funk.ForEach(payload, func(form model.WechatActivity) {
		form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		form.UpdateTime = null.TimeFrom(time.Now().Value())
		r, err = s.ID(form.ID).Update(&form)
		ret = append(ret, r)
	})
	if err != nil {
		s.Rollback()
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityUpdate api implementation
// @Summary 更新活动
// @Tags 微信活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.WechatActivity false "活动对象"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/update [put]
func WechatActivityUpdate(ctx *Context) {
	var payload model.WechatActivity
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityList api implementation
// @Summary 活动分页查询
// @Tags 微信活动
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Param title query string false "标题筛选"
// @Param hidden query int false "是否隐藏筛选"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/list [get]
func WechatActivityList(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetString("title", "nn")
	q.SetInt("hidden")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "wechat_activity", "list", "wechat_activity", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityOne api implementation
// @Summary 获取活动
// @Tags 微信活动
// @Param Authorization header string false "认证令牌"
// @Param id query string false "活动id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/wechat/activity/one [get]
func WechatActivityOne(ctx *Context) {
	var entity model.WechatActivity
	id := ctx.Query("id")
	_, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}

// WechatActivityIncrease api implementation
// @Summary 增加次数
// @Tags 微信活动
// @version 1.0
// @Accept application/json
// @Param id body string false "记录id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/v1/wechat/activity/increase [post]
func WechatActivityIncrease(ctx *Context) {
	var payload string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.WechatActivityAction(payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// WechatActivityIncreaseV2 api implementation
// @Summary 增加次数
// @Tags 微信活动
// @version 2.0
// @Accept application/json
// @Param id body string false "记录id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/v2/wechat/activity/increase_v2 [post]
func WechatActivityIncreaseV2(ctx *Context) {
	var payload string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.WechatActivityAction(payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
