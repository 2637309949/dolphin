// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_table_column.go

package api

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysTableColumnAdd api implementation
// @Summary 添加表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysTableColumn false "表字段信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/table/column/add [post]
func (ctr *SysTableColumn) SysTableColumnAdd(ctx *Context) {
	var payload types.SysTableColumn
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

// SysTableColumnBatchAdd api implementation
// @Summary 批量添加表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_table_column body []types.SysTableColumn false "表字段"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/table/column/batch_add [post]
func (ctr *SysTableColumn) SysTableColumnBatchAdd(ctx *Context) {
	var payload []types.SysTableColumn
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

// SysTableColumnDel api implementation
// @Summary 删除表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_table_column body types.SysTableColumn false "表字段"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/table/column/del [delete]
func (ctr *SysTableColumn) SysTableColumnDel(ctx *Context) {
	var payload types.SysTableColumn
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.MustDB().In("id", payload.ID.Int64).Update(&types.SysTableColumn{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTableColumnBatchDel api implementation
// @Summary 批量删除表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_table_column body []types.SysTableColumn false "表字段"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/table/column/batch_del [delete]
func (ctr *SysTableColumn) SysTableColumnBatchDel(ctx *Context) {
	var payload []*types.SysTableColumn
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysTableColumn) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.MustDB().In("id", ids).Update(&types.SysTableColumn{
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

// SysTableColumnUpdate api implementation
// @Summary 更新表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body types.SysRole false "表字段信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/table/column/update [put]
func (ctr *SysTableColumn) SysTableColumnUpdate(ctx *Context) {
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

// SysTableColumnBatchUpdate api implementation
// @Summary 批量更新表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_table_column body []types.SysTableColumn false "表字段"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/table/column/batch_update [put]
func (ctr *SysTableColumn) SysTableColumnBatchUpdate(ctx *Context) {
	var payload []types.SysTableColumn
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

// SysTableColumnPage api implementation
// @Summary 表字段分页查询
// @Tags 表字段
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/table/column/page [get]
func (ctr *SysTableColumn) SysTableColumnPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_table_column_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.DB.PageSearch(ctx.MustDB(), "sys_table_column", "page", "sys_table_column", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTableColumnGet api implementation
// @Summary 获取表字段信息
// @Tags 表字段
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "表字段id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/table/column/get [get]
func (ctr *SysTableColumn) SysTableColumnGet(ctx *Context) {
	var entity types.SysTableColumn
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
