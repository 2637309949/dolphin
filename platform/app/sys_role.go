// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_role.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
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
func SysRoleAdd(ctx *Context) {
	var payload model.SysRole
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
func SysRoleDel(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysRole{
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
func SysRoleBatchDel(ctx *Context) {
	var payload []model.SysRole
	var ids []string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form model.SysRole) {
		ids = append(ids, form.ID.String)
	})
	ret, err := ctx.DB.In("id", ids).Update(&model.SysRole{
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
func SysRoleUpdate(ctx *Context) {
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
func SysRolePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetRule("sys_role_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_role", "page", "sys_role", q.Value())
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
func SysRoleRoleMenuTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetString("role_id")
	q.SetBool("is_admin", q.GetString("role_id") == model.AdminRole.ID.String)
	q.SetRule("sys_role_menu_tree")
	q.SetTags()
	ret, err := ctx.TreeSearch(ctx.DB, "sys_role", "menu_tree", "sys_org", q.Value())
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
func SysRoleRoleAppFunTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetString("role_id")
	q.SetBool("is_admin", q.GetString("role_id") == model.AdminRole.ID.String)
	q.SetRule("sys_role_app_fun_tree")
	q.SetTags()
	ret, err := ctx.TreeSearch(ctx.DB, "sys_role", "app_fun_tree", "sys_org", q.Value())
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
func SysRoleGet(ctx *Context) {
	var entity model.SysRole
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
