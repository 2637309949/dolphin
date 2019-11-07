// Code generated by dol build. Only Generate by tools if not existed.
// source: pusher_activity.go

package app

import (
	"example/model"
	"example/util"
	"net/http"

	platformApp "github.com/2637309949/dolphin/cli/platform/app"
	"github.com/gin-gonic/gin/binding"
)

// PusherActivity struct
type PusherActivity struct {
	*Engine
}

// BuildPusherActivity return PusherActivity
func BuildPusherActivity(build func(*PusherActivity)) func(engine *platformApp.Engine) {
	return BuildEngine(func(engine *Engine) {
		build(&PusherActivity{Engine: engine})
	})
}

// Increase 增加次数
// @Title Increase
// @Description 增加次数
// @Param	id        记录id
// @Param	action        点赞（like），分享（share）和收藏(collect) 收藏和点赞第一次调用次数增一，调第二次减一，分享每一次都加一
// @Success 200 {object} Account
// @Failure 403 :id is empty
// @Router /api/pusher_activity/increase [get]
func (ctr *PusherActivity) Increase(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.M{"code": 500, "message": err.Error()})
		return
	}
	ret, err := util.AppAction(form)

	if err != nil {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

// Add 添加活动
// @Title Add
// @Description 添加活动
// @Param	pusher_activity        活动对象
// @Success 200 {object} Account
// @Failure 403 :id is empty
// @Router /api/pusher_activity/add [post]
func (ctr *PusherActivity) Add(ctx *Context) {
	var bean model.PusherActivity
	if err := ctx.ShouldBindBodyWith(&bean, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.M{"code": 500, "message": err.Error()})
		return
	}
	ret, err := ctx.DB.Insert(&bean)
	if err != nil {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

// Update 更新活动
// @Title Update
// @Description 更新活动
// @Param	pusher_activity        活动对象
// @Success 200 {object} Account
// @Failure 403 :id is empty
// @Router /api/pusher_activity/update [post]
func (ctr *PusherActivity) Update(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.M{"code": 500, "message": err.Error()})
		return
	}
	ret, err := util.AppAction(form)

	if err != nil {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

// Delete 删除活动
// @Title Delete
// @Description 删除活动
// @Param	ids        活动id对象数组
// @Success 200 {object} Account
// @Failure 403 :id is empty
// @Router /api/pusher_activity/delete [post]
func (ctr *PusherActivity) Delete(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.M{"code": 500, "message": err.Error()})
		return
	}
	ret, err := util.AppAction(form)

	if err != nil {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

// Page 活动分页查询
// @Title Page
// @Description 活动分页查询
// @Param	page        页码
// @Param	rows        单页数
// @Param	title        标题筛选
// @Param	hidden        是否隐藏筛选
// @Success 200 {object} Account
// @Failure 403 :id is empty
// @Router /api/pusher_activity/page [get]
func (ctr *PusherActivity) Page(ctx *Context) {
	q := ctr.Query(ctx)
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetInt("del_flag")
	q.SetString("title")
	q.SetString("campus")
	q.SetString("city")
	q.SetString("hidden")

	ret, err := ctr.PageSearch(ctx.DB, "pusher_activity", "page", "pusher_activity", q.Value())
	if err != nil {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

// PageByArea 活动分页按区域查询
// @Title PageByArea
// @Description 活动分页按区域查询
// @Param	page        页码
// @Param	rows        单页数
// @Param	title        标题筛选
// @Param	campus        校区id筛选
// @Param	city        城市筛选
// @Param	hidden        是否隐藏筛选
// @Success 200 {object} Account
// @Failure 403 :id is empty
// @Router /api/pusher_activity/page_by_area [get]
func (ctr *PusherActivity) PageByArea(ctx *Context) {
	q := ctr.Query(ctx)
	q.SetInt("page", 1)
	q.SetInt("size", 20)
	q.SetInt("del_flag")
	q.SetString("title")
	q.SetString("campus")
	q.SetString("city")
	q.SetString("hidden")

	ret, err := ctr.PageSearch(ctx.DB, "pusher_activity", "page_by_area", "pusher_activity", q.Value())
	if err != nil {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

// Get 获取活动
// @Title Get
// @Description 获取活动
// @Param	id        活动id
// @Success 200 {object} Account
// @Failure 403 :id is empty
// @Router /api/pusher_activity/get [get]
func (ctr *PusherActivity) Get(ctx *Context) {
	var form = &struct{}{}
	if err := ctx.ShouldBindBodyWith(form, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.M{"code": 500, "message": err.Error()})
		return
	}
	ret, err := util.AppAction(form)

	if err != nil {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ret)
}
