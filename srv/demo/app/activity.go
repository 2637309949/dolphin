// Code generated by dol build. Only Generate by tools if not existed.
// source: activity.go

package app

import (
	"example/model"
	
	pApp "github.com/2637309949/dolphin/cli/platform/app"
	pUtil "github.com/2637309949/dolphin/cli/platform/util"
	"github.com/gin-gonic/gin/binding"
)

// Activity struct
type Activity struct {
	*Engine
}

// BuildActivity return Activity
func BuildActivity(build func(*Activity)) func(engine *pApp.Engine) {
	return BuildEngine(func(engine *Engine) {
		build(&Activity{Engine: engine})
	})
}

// AddMany api implementation
// @Summary 添加活动 
// @Tags 活动
// @Accept application/json
// @Param token header query string true "认证令牌"
// @Param activity body []model.Activity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/add_many [post]
func (ctr *Activity) AddMany(ctx *Context) {
	var form []model.Activity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// AddOne api implementation
// @Summary 添加活动 
// @Tags 活动
// @Accept application/json
// @Param token header query string true "认证令牌"
// @Param activity body model.Activity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/add_one [post]
func (ctr *Activity) AddOne(ctx *Context) {
	var form model.Activity
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// DeleteMany api implementation
// @Summary 删除活动 
// @Tags 活动
// @Accept application/json
// @Param token header query string true "认证令牌"
// @Param activity body []model.Activity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/delete_many [post]
func (ctr *Activity) DeleteMany(ctx *Context) {
	var form []model.Activity
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

// DeleteOne api implementation
// @Summary 删除活动 
// @Tags 活动
// @Accept application/json
// @Param token header query string true "认证令牌"
// @Param activity body model.Activity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/delete_one [post]
func (ctr *Activity) DeleteOne(ctx *Context) {
	var form model.Activity
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

// UpdateMany api implementation
// @Summary 更新活动 
// @Tags 活动
// @Accept application/json
// @Param token header query string true "认证令牌"
// @Param activity body []model.Activity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/update_many [post]
func (ctr *Activity) UpdateMany(ctx *Context) {
	var form []model.Activity
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

// UpdateOne api implementation
// @Summary 更新活动 
// @Tags 活动
// @Accept application/json
// @Param token header query string true "认证令牌"
// @Param activity body model.Activity false "活动对象"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/update_one [post]
func (ctr *Activity) UpdateOne(ctx *Context) {
	var form model.Activity
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

// List api implementation
// @Summary 活动分页查询 
// @Tags 活动
// @Param token header query string true "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Param title query string false "标题筛选"
// @Param hidden query int false "是否隐藏筛选"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/list [get]
func (ctr *Activity) List(ctx *Context) {
	q := ctr.Query(ctx)
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetString("title", "nn")
	q.SetInt("hidden")
	ret, err := ctr.PageSearch(ctx.DB, "activity", "list", "activity", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// One api implementation
// @Summary 获取活动 
// @Tags 活动
// @Param token header query string true "认证令牌"
// @Param id query string false "活动id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/activity/one [get]
func (ctr *Activity) One(ctx *Context) {
	var entity model.Activity
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// InCrease api implementation
// @Summary 增加次数 
// @Tags 活动
// @version 1.0
// @Param id query string false "记录id"
// @Param action query string false "点赞（like）,分享（share）和收藏(collect) 收藏和点赞第一次调用次数增一，调第二次减一，分享每一次都加一"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/v1/activity/in_crease [get]
func (ctr *Activity) InCrease(ctx *Context) {
	q := ctr.Query(ctx)
	q.SetString("id")
	q.SetString("action")
	ret, err := pUtil.AppAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// InCreaseV2 api implementation
// @Summary 增加次数 
// @Tags 活动
// @version 2.0
// @Accept application/json
// @Param id body string false "记录id"
// @Param action body string false "点赞（like）,分享（share）和收藏(collect) 收藏和点赞第一次调用次数增一，调第二次减一，分享每一次都加一"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/v2/activity/in_crease_v2 [post]
func (ctr *Activity) InCreaseV2(ctx *Context) {
	var form string
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := pUtil.AppAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

