// Code generated by dol build. Only Generate by tools if not existed.
// source: activity.go

package app

import (
	"example/model"

	platformApp "github.com/2637309949/dolphin/cli/platform/app"
	platformUtil "github.com/2637309949/dolphin/cli/platform/util"
	"github.com/gin-gonic/gin/binding"
)

// Activity struct
type Activity struct {
	*Engine
}

// BuildActivity return Activity
func BuildActivity(build func(*Activity)) func(engine *platformApp.Engine) {
	return BuildEngine(func(engine *Engine) {
		build(&Activity{Engine: engine})
	})
}

// InCrease api implementation
// @Summary 增加次数
// @Tags 活动
// @version 1.0
// @Accept application/json
// @Param query id  记录id
// @Param query action  点赞（like），分享（share）和收藏(collect) 收藏和点赞第一次调用次数增一，调第二次减一，分享每一次都加一
// @Success 200 {object} Account
// @Failure 500 :id is empty
// @Router /api1.0/activity/in_crease [get]
func (ctr *Activity) InCrease(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := platformUtil.AppAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Add api implementation
// @Summary 添加活动
// @Tags 活动
// @Accept application/json
// @Param body activity  活动对象
// @Success 200 {object} Account
// @Failure 500 :id is empty
// @Router /api/activity/add [post]
func (ctr *Activity) Add(ctx *Context) {
	var bean model.Activity
	if err := ctx.ShouldBindBodyWith(&bean, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Insert(&bean)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Update api implementation
// @Summary 更新活动
// @Tags 活动
// @Accept application/json
// @Param body activity  活动对象
// @Success 200 {object} Account
// @Failure 500 :id is empty
// @Router /api/activity/update [post]
func (ctr *Activity) Update(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := platformUtil.AppAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Delete api implementation
// @Summary 删除活动
// @Tags 活动
// @Accept application/json
// @Param body ids  活动id对象数组
// @Success 200 {object} Account
// @Failure 500 :id is empty
// @Router /api/activity/delete [post]
func (ctr *Activity) Delete(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := platformUtil.AppAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Page api implementation
// @Summary 活动分页查询
// @Tags 活动
// @Accept application/json
// @Param query page  页码
// @Param query rows  单页数
// @Param query title  标题筛选
// @Param query hidden  是否隐藏筛选
// @Success 200 {object} Account
// @Failure 500 :id is empty
// @Router /api/activity/page [get]
func (ctr *Activity) Page(ctx *Context) {
	q := ctr.Query(ctx)
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetInt("del_flag")
	q.SetString("title")
	q.SetString("campus")
	q.SetString("city")
	q.SetString("hidden")
	ret, err := ctr.PageSearch(ctx.DB, "activity", "page", "activity", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// PageByArea api implementation
// @Summary 活动分页按区域查询
// @Tags 活动
// @Accept application/json
// @Param query page  页码
// @Param query rows  单页数
// @Param query title  标题筛选
// @Param query campus  校区id筛选
// @Param query city  城市筛选
// @Param query hidden  是否隐藏筛选
// @Success 200 {object} Account
// @Failure 500 :id is empty
// @Router /api/activity/page_by_area [get]
func (ctr *Activity) PageByArea(ctx *Context) {
	q := ctr.Query(ctx)
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetInt("del_flag")
	q.SetString("title")
	q.SetString("campus")
	q.SetString("city")
	q.SetString("hidden")
	ret, err := ctr.PageSearch(ctx.DB, "activity", "page_by_area", "activity", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Get api implementation
// @Summary 获取活动
// @Tags 活动
// @Accept application/json
// @Param query id  活动id
// @Success 200 {object} Account
// @Failure 500 :id is empty
// @Router /api/activity/get [get]
func (ctr *Activity) Get(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := platformUtil.AppAction(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
