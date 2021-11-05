// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_app_fun.go

package api

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysAppFunAdd api implementation
// @Summary 添加APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_app_fun body types.SysAppFun false "APP功能信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/app/fun/add [post]
func (ctr *SysAppFun) SysAppFunAdd(ctx *Context) {
	var payload types.SysAppFun
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFromNow()
	payload.Creater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	ret, err := ctx.MustDB().Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunBatchAdd api implementation
// @Summary 批量添加APP功能
// @Tags APP功能
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_app_fun body []types.SysAppFun false "APP功能信息"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/app/fun/batch_add [post]
func (ctr *SysAppFun) SysAppFunBatchAdd(ctx *Context) {
	var payload []types.SysAppFun
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].CreateTime = null.TimeFromNow()
		payload[i].Creater = null.IntFromStr(ctx.MustToken().GetUserID())
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.MustToken().GetUserID())
		payload[i].IsDelete = null.IntFrom(0)
	}
	ret, err := ctx.MustDB().Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunDel api implementation
// @Summary 删除APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_app_fun body types.SysAppFun false "APP功能"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/app/fun/del [delete]
func (ctr *SysAppFun) SysAppFunDel(ctx *Context) {
	var payload types.SysAppFun
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.MustDB().In("id", payload.ID.Int64).Update(&types.SysAppFun{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunBatchDel api implementation
// @Summary 批量删除APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_app_fun body []types.SysAppFun false "APP功能"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/app/fun/batch_del [delete]
func (ctr *SysAppFun) SysAppFunBatchDel(ctx *Context) {
	var payload []types.SysAppFun
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysAppFun) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.MustDB().In("id", ids).Update(&types.SysAppFun{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunUpdate api implementation
// @Summary 更新APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysRole false "APP功能信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/app/fun/update [put]
func (ctr *SysAppFun) SysAppFunUpdate(ctx *Context) {
	var payload types.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	ret, err := ctx.MustDB().ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunBatchUpdate api implementation
// @Summary 批量更新APP功能
// @Tags APP功能
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_app_fun body []types.SysAppFun false "APP功能信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/app/fun/batch_update [put]
func (ctr *SysAppFun) SysAppFunBatchUpdate(ctx *Context) {
	var payload []types.SysAppFun
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.MustDB().NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.MustToken().GetUserID())
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

// SysAppFunPage api implementation
// @Summary APP功能分页查询
// @Tags APP功能
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/app/fun/page [get]
func (ctr *SysAppFun) SysAppFunPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_app_fun_page")
	q.SetTags()
	ret, err := ctr.Srv.DB.PageSearch(ctx.MustDB(), "sys_app_fun", "page", "sys_app_fun", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunTree api implementation
// @Summary 菜单树形结构
// @Tags APP功能
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Router /api/sys/app/fun/tree [get]
func (ctr *SysAppFun) SysAppFunTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetRule("sys_app_fun_tree")
	q.SetTags()
	ret, err := ctr.Srv.DB.TreeSearch(ctx.MustDB(), "sys_app_fun", "tree", "sys_app_fun", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAppFunGet api implementation
// @Summary 获取APP功能信息
// @Tags APP功能
// @Param Authorization header string false "认证令牌"
// @Param id query string false "APP功能id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/app/fun/get [get]
func (ctr *SysAppFun) SysAppFunGet(ctx *Context) {
	var entity types.SysAppFun
	err := ctx.ShouldBindWith(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ext, err := ctx.MustDB().Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(errors.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
