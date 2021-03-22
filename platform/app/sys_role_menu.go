// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_role_menu.go

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

// SysRoleMenuAdd api implementation
// @Summary 添加角色菜单
// @Tags 角色菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role_menu body model.SysRoleMenu false "角色菜单信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/add [post]
func SysRoleMenuAdd(ctx *Context) {
	var payload model.SysRoleMenu
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
	payload.IsDelete = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleMenuBatchAdd api implementation
// @Summary 添加角色菜单
// @Tags 角色菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param role_menu body []model.SysRoleMenu false "角色菜单信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/batch_add [post]
func SysRoleMenuBatchAdd(ctx *Context) {
	var payload []*model.SysRoleMenu
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form *model.SysRoleMenu) {
		form.ID = null.StringFromUUID()
		form.CreateTime = null.TimeFrom(time.Now().Value())
		form.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
		form.UpdateTime = null.TimeFrom(time.Now().Value())
		form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		form.IsDelete = null.IntFrom(0)
	})
	payload = funk.Filter(payload, func(form *model.SysRoleMenu) bool {
		ext, _ := ctx.DB.Where("role_id=? and menu_id=?", form.RoleId.String, form.MenuId.String).Exist(new(model.SysRoleMenu))
		return !ext
	}).([]*model.SysRoleMenu)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleMenuDel api implementation
// @Summary 删除角色菜单
// @Tags 角色菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role_menu body model.SysRoleMenu false "角色菜单"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/del [delete]
func SysRoleMenuDel(ctx *Context) {
	var payload model.SysRoleMenu
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysRoleMenu{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleMenuBatchDel api implementation
// @Summary 删除角色菜单
// @Tags 角色菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []model.SysRoleMenu false "角色菜单信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/batch_del [delete]
func SysRoleMenuBatchDel(ctx *Context) {
	var payload []*model.SysRoleMenu
	var ids []string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form model.SysRoleMenu) {
		ids = append(ids, form.ID.String)
	})
	ret, err := ctx.DB.In("id", ids).Update(&model.SysRoleMenu{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleMenuUpdate api implementation
// @Summary 更新角色菜单
// @Tags 角色菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role_menu body model.SysRoleMenu false "角色菜单信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/update [put]
func SysRoleMenuUpdate(ctx *Context) {
	var payload model.SysRoleMenu
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

// SysRoleMenuBatchUpdate api implementation
// @Summary 更新角色菜单
// @Tags 角色菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_role_menu body []model.SysRoleMenu false "角色菜单信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/role/menu/batch_update [put]
func SysRoleMenuBatchUpdate(ctx *Context) {
	var payload []model.SysRoleMenu
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.String).Update(&payload[i])
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

// SysRoleMenuPage api implementation
// @Summary 角色菜单分页查询
// @Tags 角色菜单
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/page [get]
func SysRoleMenuPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_role_menu_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_role_menu", "page", "sys_role_menu", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysRoleMenuGet api implementation
// @Summary 获取角色菜单信息
// @Tags 角色菜单
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "角色菜单id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/get [get]
func SysRoleMenuGet(ctx *Context) {
	var entity model.SysRoleMenu
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
