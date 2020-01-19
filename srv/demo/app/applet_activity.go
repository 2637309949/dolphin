// Code generated by dol build. Only Generate by tools if not existed.
// source: applet_activity.go

package app

import (
	"example/model"
	"example/srv"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/time"
)

// AppletActivityBatchAdd api implementation
// @Summary 添加活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.AppletActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/applet/activity/batch_add [post]
func AppletActivityBatchAdd(ctx *Context) {
	var form []model.AppletActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	for _, f := range form {
		f.ID = null.StringFromUUID()
		f.CreateTime = null.TimeFromPtr(time.Now().Value())
		f.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	}
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityAdd api implementation
// @Summary 添加活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.AppletActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/applet/activity/add [post]
func AppletActivityAdd(ctx *Context) {
	var form model.AppletActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFromUUID()
	form.CreateTime = null.TimeFromPtr(time.Now().Value())
	form.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityBatchDel api implementation
// @Summary 删除活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.AppletActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/applet/activity/batch_del [post]
func AppletActivityBatchDel(ctx *Context) {
	var form []model.AppletActivity
	var ids []string
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	for _, f := range form {
		ids = append(ids, f.ID.String)
	}
	ret, err := ctx.DB.Table(new(model.AppletActivity)).In("id", ids).Update(map[string]interface{}{
		"delete_time": null.TimeFromPtr(time.Now().Value()),
		"delete_by":   null.StringFrom(ctx.GetToken().GetUserID()),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityDel api implementation
// @Summary 删除活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.AppletActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/applet/activity/del [post]
func AppletActivityDel(ctx *Context) {
	var form model.AppletActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Table(new(model.AppletActivity)).In("id", form.ID.String).Update(map[string]interface{}{
		"delete_time": null.TimeFromPtr(time.Now().Value()),
		"delete_by":   null.StringFrom(ctx.GetToken().GetUserID()),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityBatchUpdate api implementation
// @Summary 更新活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body []model.AppletActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/applet/activity/batch_update [post]
func AppletActivityBatchUpdate(ctx *Context) {
	var form []model.AppletActivity
	var err error
	var ret []int64
	var r int64
	if err = ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for _, f := range form {
		f.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		f.UpdateTime = null.TimeFromPtr(time.Now().Value())
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

// AppletActivityUpdate api implementation
// @Summary 更新活动
// @Tags 活动
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param activity body model.AppletActivity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/applet/activity/update [post]
func AppletActivityUpdate(ctx *Context) {
	var form model.AppletActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	form.UpdateTime = null.TimeFromPtr(time.Now().Value())
	ret, err := ctx.DB.ID(form.ID).Update(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityList api implementation
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
// @Router /api/applet/activity/list [get]
func AppletActivityList(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetString("title", "nn")
	q.SetInt("hidden")
	ret, err := ctx.PageSearch(ctx.DB, "applet_activity", "list", "applet_activity", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityOne api implementation
// @Summary 获取活动
// @Tags 活动
// @Param Authorization header string false "认证令牌"
// @Param id query string false "活动id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/applet/activity/one [get]
func AppletActivityOne(ctx *Context) {
	var entity model.AppletActivity
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityIncrease api implementation
// @Summary 增加次数
// @Tags 活动
// @version 1.0
// @Accept application/json
// @Param applet_activity body model.AppletActivity false "记录id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/v1/applet/activity/increase [post]
func AppletActivityIncrease(ctx *Context) {
	var form model.AppletActivity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.AppletActivityAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AppletActivityIncreaseV2 api implementation
// @Summary 增加次数
// @Tags 活动
// @version 2.0
// @Accept application/json
// @Param id body string false "记录id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/v2/applet/activity/increase_v2 [post]
func AppletActivityIncreaseV2(ctx *Context) {
	var form string
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.AppletActivityAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
