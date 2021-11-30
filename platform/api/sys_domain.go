// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_domain.go

package api

import (
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// SysDomainAdd api implementation
// @Summary 添加域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysDomain false "域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/domain/add [post]
func (ctr *SysDomain) SysDomainAdd(ctx *Context) {
	var payload types.SysDomain
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFromNow()
	payload.Creater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	payload.AppName = null.StringFrom(viper.GetString("app.name"))
	ret, err := App.PlatformDB.Insert(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainBatchAdd api implementation
// @Summary 批量添加域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_domain body []types.SysDomain false "域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/domain/batch_add [post]
func (ctr *SysDomain) SysDomainBatchAdd(ctx *Context) {
	var payload []types.SysDomain
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
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

	db := ctx.MustDB()
	ret, err := db.Insert(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainDel api implementation
// @Summary 删除域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_domain body types.SysDomain false "域"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/domain/del [delete]
func (ctr *SysDomain) SysDomainDel(ctx *Context) {
	var payload types.SysDomain
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ret, err := App.PlatformDB.In("id", payload.ID.Int64).Update(&types.SysDomain{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainBatchDel api implementation
// @Summary 批量删除域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_domain body []types.SysDomain false "域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/domain/batch_del [delete]
func (ctr *SysDomain) SysDomainBatchDel(ctx *Context) {
	var payload []types.SysDomain
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	var ids = funk.Map(payload, func(form types.SysDomain) int64 { return form.ID.Int64 }).([]int64)
	ret, err := db.In("id", ids).Update(&types.SysDomain{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainUpdate api implementation
// @Summary 更新域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysRole false "域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/domain/update [put]
func (ctr *SysDomain) SysDomainUpdate(ctx *Context) {
	var payload types.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	ret, err := App.PlatformDB.ID(payload.ID).Update(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainBatchUpdate api implementation
// @Summary 批量更新域
// @Tags 域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_domain body []types.SysDomain false "域信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/domain/batch_update [put]
func (ctr *SysDomain) SysDomainBatchUpdate(ctx *Context) {
	var payload []types.SysDomain
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	s := db.NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.MustToken().GetUserID())
		r, err = s.ID(payload[i].ID.Int64).Update(&payload[i])
		if err != nil {
			s.Rollback()
			logrus.Error(ctx, err)
			ctx.Fail(err)
			return
		}
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	err = s.Commit()
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainPage api implementation
// @Summary 域分页查询
// @Tags 域
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/domain/page [get]
func (ctr *SysDomain) SysDomainPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("app_name", viper.GetString("app.name"))()
	q.SetRule("sys_domain_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.DB.PageSearch(App.PlatformDB, "sys_domain", "page", "sys_domain", q.Value())
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysDomainGet api implementation
// @Summary 获取域信息
// @Tags 域
// @Param Authorization header string false "认证令牌"
// @Param id query string false "域id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/domain/get [get]
func (ctr *SysDomain) SysDomainGet(ctx *Context) {
	var entity types.SysDomain
	err := ctx.ShouldBindWith(&entity)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	ext, err := db.Get(&entity)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(errors.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
