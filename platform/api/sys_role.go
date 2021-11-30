// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_role.go

package api

import (
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/thoas/go-funk"
)

// SysRoleAdd api implementation
// @Summary 添加角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body types.SysRole false "角色信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/role/add [post]
func (ctr *SysRole) SysRoleAdd(ctx *Context) {
	var payload types.SysRole
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

	db := ctx.MustDB()
	cnt, err := db.Where("code=? and is_delete !=1", payload.Code.String).Count(new(types.SysRole))
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	if cnt > 0 {
		ctx.Fail(errors.ErrRecordExisted)
		return
	}
	ret, err := db.Insert(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleBatchAdd api implementation
// @Summary 批量添加角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body []types.SysRole false "角色信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/role/batch_add [post]
func (ctr *SysRole) SysRoleBatchAdd(ctx *Context) {
	var payload []types.SysRole
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
	cnt, err := db.Where("is_delete !=1").In("code", (funk.Map(payload, func(item types.SysRole) interface{} {
		return item.Code.String
	}).([]interface{}))...).Count(new(types.SysRole))
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	if cnt > 0 {
		ctx.Fail(errors.ErrRecordExisted)
		return
	}
	ret, err := db.Insert(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleDel api implementation
// @Summary 删除角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body types.SysRole false "角色"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/role/del [delete]
func (ctr *SysRole) SysRoleDel(ctx *Context) {
	var payload types.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	ret, err := db.In("id", payload.ID.Int64).Update(&types.SysRole{
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

// SysRoleBatchDel api implementation
// @Summary 批量删除角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []types.SysRole false "用户信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/role/batch_del [delete]
func (ctr *SysRole) SysRoleBatchDel(ctx *Context) {
	var payload []*types.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	db := ctx.MustDB()
	var ids = funk.Map(payload, func(form types.SysRole) int64 { return form.ID.Int64 }).([]int64)
	ret, err := db.In("id", ids).Update(&types.SysRole{
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

// SysRoleUpdate api implementation
// @Summary 更新角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body types.SysRole false "角色信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/role/update [put]
func (ctr *SysRole) SysRoleUpdate(ctx *Context) {
	var payload types.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()

	db := ctx.MustDB()
	ret, err := db.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleBatchUpdate api implementation
// @Summary 批量更新角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body []types.SysRole false "角色信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/role/batch_update [post]
func (ctr *SysRole) SysRoleBatchUpdate(ctx *Context) {
	var payload []types.SysRole
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

// SysRolePage api implementation
// @Summary 角色分页查询
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/role/page [get]
func (ctr *SysRole) SysRolePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_role_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()

	db := ctx.MustDB()
	if ctr.Srv.Report.Check(ctx.Request()) {
		ctr.Srv.Report.SetOptionsetsFormat(OptionsetsFormat(db))
		ret, err := ctr.Srv.Report.PageExport(db, "sys_role", "page", "sys_role", q.Value())
		if err != nil {
			logrus.Error(ctx, err)
			ctx.Fail(err)
			return
		}
		ctx.Success(ret)
		return
	}
	ret, err := ctr.Srv.DB.PageSearch(db, "sys_role", "page", "sys_role", q.Value())
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleRoleMenuTree api implementation
// @Summary 角色菜单树形结构
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Router /api/sys/role/role_menu_tree [get]
func (ctr *SysRole) SysRoleRoleMenuTree(ctx *Context) {
	db := ctx.MustDB()
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetString("role_id")
	q.SetBool("is_admin", ctr.Srv.DB.InAdmin(db, ctx.MustToken().GetUserID()))
	q.SetRule("sys_role_menu_tree")
	q.SetTags()
	ret, err := ctr.Srv.DB.TreeSearch(db, "sys_role", "menu_tree", "sys_org", q.Value())
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleRoleAppFunTree api implementation
// @Summary 角色App功能树形结构
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Router /api/sys/role/role_app_fun_tree [get]
func (ctr *SysRole) SysRoleRoleAppFunTree(ctx *Context) {
	db := ctx.MustDB()
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetString("role_id")
	q.SetBool("is_admin", ctr.Srv.DB.InAdmin(db, ctx.MustToken().GetUserID()))
	q.SetRule("sys_role_app_fun_tree")
	q.SetTags()
	ret, err := ctr.Srv.DB.TreeSearch(db, "sys_role", "app_fun_tree", "sys_org", q.Value())
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleGet api implementation
// @Summary 获取角色信息
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Param id query string false "角色id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/role/get [get]
func (ctr *SysRole) SysRoleGet(ctx *Context) {
	var entity types.SysRole
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
