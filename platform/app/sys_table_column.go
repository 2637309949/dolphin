// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_table_column.go

package app

import (
	"errors"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysTableColumnAdd api implementation
// @Summary 添加表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysTableColumn false "表字段信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/table/column/add [post]
func SysTableColumnAdd(ctx *Context) {
	var payload model.SysTableColumn
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
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
// @Param sys_table_column body model.SysTableColumn false "表字段"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/table/column/del [delete]
func SysTableColumnDel(ctx *Context) {
	var payload model.SysTableColumn
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysTableColumn{
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

// SysTableColumnBatchDel api implementation
// @Summary 删除表字段
// @Tags 表字段
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_table_column body []model.SysTableColumn false "表字段"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/table/column/batch_del [delete]
func SysTableColumnBatchDel(ctx *Context) {
	var payload []*model.SysTableColumn
	var ids []string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form model.SysTableColumn) {
		ids = append(ids, form.ID.String)
	})
	ret, err := ctx.DB.In("id", ids).Update(&model.SysTableColumn{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
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
// @Param user body model.SysRole false "表字段信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/table/column/update [put]
func SysTableColumnUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.String).Update(&payload)
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/table/column/page [get]
func SysTableColumnPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetRule("sys_table_column_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_table_column", "page", "sys_table_column", q.Value())
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/table/column/get [get]
func SysTableColumnGet(ctx *Context) {
	var entity model.SysTableColumn
	err := ctx.ShouldBindQuery(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ext, err := ctx.DB.Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(errors.New("not found"))
		return
	}
	ctx.Success(entity)
}
