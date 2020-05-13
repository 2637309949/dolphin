// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_scheduling.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/srv"

	"github.com/2637309949/dolphin/packages/gin/binding"
)

// SysSchedulingAdd api implementation
// @Summary 添加调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param scheduling body model.Scheduling false "调度信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/scheduling/add [post]
func SysSchedulingAdd(ctx *Context) {
	var payload model.Scheduling
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.SysSchedulingAction(payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysSchedulingDel api implementation
// @Summary 删除调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param scheduling body model.Scheduling false "调度"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/scheduling/del [delete]
func SysSchedulingDel(ctx *Context) {
	var payload model.Scheduling
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.SysSchedulingAction(payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysSchedulingUpdate api implementation
// @Summary 更新调度
// @Tags 调度
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param scheduling body model.Scheduling false "调度信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/scheduling/update [put]
func SysSchedulingUpdate(ctx *Context) {
	var payload model.Scheduling
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.SysSchedulingAction(payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysSchedulingPage api implementation
// @Summary 调度分页查询
// @Tags 调度
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/scheduling/page [get]
func SysSchedulingPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	ret, err := srv.SysSchedulingAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysSchedulingGet api implementation
// @Summary
// @Tags 调度
// @Param Authorization header string false "认证令牌"
// @Param id query string false "调度id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/scheduling/get [get]
func SysSchedulingGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	ret, err := srv.SysSchedulingAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
