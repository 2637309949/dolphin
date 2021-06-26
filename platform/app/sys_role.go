// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_role.go

package app

import (
	"errors"

	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysRoleAdd api implementation
// @Summary 添加角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body model.SysRole false "角色信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/add [post]
func (ctr *SysRole) SysRoleAdd(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFrom(time.Now())
	payload.Creater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now())
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	cnt, err := ctx.DB.Where("code=? and is_delete !=1", payload.Code.String).Count(new(model.SysRole))
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if cnt > 0 {
		ctx.Fail(errors.New("the record already exists"))
		return
	}
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleBatchAdd api implementation
// @Summary 添加角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body []model.SysRole false "角色信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/role/batch_add [post]
func (ctr *SysRole) SysRoleBatchAdd(ctx *Context) {
	var payload []model.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {

		payload[i].CreateTime = null.TimeFrom(time.Now())
		payload[i].Creater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFrom(time.Now())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].IsDelete = null.IntFrom(0)
	}
	cnt, err := ctx.DB.Where("is_delete !=1").In("code", (funk.Map(payload, func(item model.SysRole) interface{} {
		return item.Code.String
	}).([]interface{}))...).Count(new(model.SysRole))
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if cnt > 0 {
		ctx.Fail(errors.New("the record already exists"))
		return
	}
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
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
// @Param sys_role body model.SysRole false "角色"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/del [delete]
func (ctr *SysRole) SysRoleDel(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&model.SysRole{
		UpdateTime: null.TimeFrom(time.Now()),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleBatchDel api implementation
// @Summary 删除角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []model.SysRole false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/batch_del [delete]
func (ctr *SysRole) SysRoleBatchDel(ctx *Context) {
	var payload []*model.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysRole) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&model.SysRole{
		UpdateTime: null.TimeFrom(time.Now()),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
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
// @Param sys_role body model.SysRole false "角色信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/update [put]
func (ctr *SysRole) SysRoleUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now())
	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleBatchUpdate api implementation
// @Summary 更新角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role body []model.SysRole false "角色信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/role/batch_update [post]
func (ctr *SysRole) SysRoleBatchUpdate(ctx *Context) {
	var payload []model.SysRole
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
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

// SysRolePage api implementation
// @Summary 角色分页查询
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
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
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_role", "page", "sys_role", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if ctx.QueryBool("__export__") {
		cfg := NewBuildExcelConfig(ret.Data)
		cfg.Format = OptionsetsFormat(ctx.DB)
		ctx.SuccessWithExcel(cfg)
		return
	}
	ctx.Success(ret)
}

// SysRoleRoleMenuTree api implementation
// @Summary 角色菜单树形结构
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Fail
// @Router /api/sys/role/role_menu_tree [get]
func (ctr *SysRole) SysRoleRoleMenuTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetString("role_id")
	q.SetBool("is_admin", ctx.InAdmin())
	q.SetRule("sys_role_menu_tree")
	q.SetTags()
	ret, err := ctr.Srv.TreeSearch(ctx.DB, "sys_role", "menu_tree", "sys_org", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleRoleAppFunTree api implementation
// @Summary 角色App功能树形结构
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Fail
// @Router /api/sys/role/role_app_fun_tree [get]
func (ctr *SysRole) SysRoleRoleAppFunTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetString("role_id")
	q.SetBool("is_admin", ctx.InAdmin())
	q.SetRule("sys_role_app_fun_tree")
	q.SetTags()
	ret, err := ctr.Srv.TreeSearch(ctx.DB, "sys_role", "app_fun_tree", "sys_org", q.Value())
	if err != nil {
		logrus.Error(err)
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/get [get]
func (ctr *SysRole) SysRoleGet(ctx *Context) {
	var entity model.SysRole
	err := ctx.ShouldBindWith(&entity)
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
