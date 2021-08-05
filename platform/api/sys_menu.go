// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_menu.go

package api

import (
	"errors"
	"fmt"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysMenuAdd api implementation
// @Summary 添加菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body types.SysMenu false "菜单信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/menu/add [post]
func (ctr *SysMenu) SysMenuAdd(ctx *Context) {
	var payload types.SysMenu
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFromNow()
	payload.Creater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	if !payload.Parent.IsZero() {
		parent := types.SysMenu{}
		ext, err := ctx.DB.SqlMapClient("selectone_sys_menu", &map[string]int64{"id": payload.Parent.Int64}).Get(&parent)
		if err != nil || !ext {
			ctx.Fail(err)
			return
		} else if !ext {
			ctx.Fail(errors.New("the record does not exist"))
			return
		}
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%v%v", payload.ID.Int64, parent.Inheritance.String))
	} else {
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%v|", payload.ID.Int64))
	}

	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuBatchAdd api implementation
// @Summary 添加菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body []types.SysMenu false "菜单信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/menu/batch_add [post]
func (ctr *SysMenu) SysMenuBatchAdd(ctx *Context) {
	var payload []types.SysMenu
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].CreateTime = null.TimeFromNow()
		payload[i].Creater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].IsDelete = null.IntFrom(0)
	}
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuDel api implementation
// @Summary 删除菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body types.SysMenu false "菜单"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/menu/del [delete]
func (ctr *SysMenu) SysMenuDel(ctx *Context) {
	var payload types.SysMenu
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.SysMenu{
		UpdateTime: null.TimeFromNow(),
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

// SysMenuBatchDel api implementation
// @Summary 删除菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body []types.SysMenu false "菜单"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/menu/batch_del [delete]
func (ctr *SysMenu) SysMenuBatchDel(ctx *Context) {
	var payload []types.SysMenu
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysDomain) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.SysMenu{
		UpdateTime: null.TimeFromNow(),
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

// SysMenuUpdate api implementation
// @Summary 更新菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body types.SysMenu false "菜单信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/menu/update [put]
func (ctr *SysMenu) SysMenuUpdate(ctx *Context) {
	var payload types.SysMenu
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()

	if !payload.Parent.IsZero() {
		parent := types.SysMenu{}
		ext, err := ctx.DB.SqlMapClient("selectone_sys_menu", &map[string]int64{"id": payload.Parent.Int64}).Get(&parent)
		if err != nil || !ext {
			ctx.Fail(err)
			return
		} else if !ext {
			ctx.Fail(errors.New("the record does not exist"))
			return
		}
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%v%v", payload.ID.Int64, parent.Inheritance.String))
	} else {
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%v|", payload.ID.Int64))
	}

	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuBatchUpdate api implementation
// @Summary 更新菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body []types.SysMenu false "菜单信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/menu/batch_update [put]
func (ctr *SysMenu) SysMenuBatchUpdate(ctx *Context) {
	var payload []types.SysMenu
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
		payload[i].UpdateTime = null.TimeFromNow()
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

// SysMenuSidebar api implementation
// @Summary 系统菜单
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Router /api/sys/menu/sidebar [get]
func (ctr *SysMenu) SysMenuSidebar(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetBool("isAdmin", ctr.Srv.InAdmin(ctx.DB, ctx.GetToken().GetUserID()))
	q.SetUser()
	q.SetRule("sys_menu_sidebar")
	q.SetTags()
	ret, err := ctr.Srv.TreeSearch(ctx.DB, "sys_menu", "sidebar", "sys_menu", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuPage api implementation
// @Summary 菜单分页查询
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/menu/page [get]
func (ctr *SysMenu) SysMenuPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("name")
	q.SetString("code")
	q.SetString("sort", "update_time desc")
	q.SetRule("sys_menu_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	if ctr.Srv.Check(ctx.Context) {
		ctr.Srv.SetOptionsetsFormat(OptionsetsFormat(ctx.DB))
		ret, err := ctr.Srv.PageExport(ctx.DB, "sys_menu", "page", "sys_menu", q.Value())
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		ctx.Success(ret)
		return
	}
	ret, err := ctr.Srv.PageSearch(ctx.DB, "sys_menu", "page", "sys_menu", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuTree api implementation
// @Summary 菜单树形结构
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Router /api/sys/menu/tree [get]
func (ctr *SysMenu) SysMenuTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetRule("sys_menu_tree")
	q.SetTags()
	ret, err := ctr.Srv.TreeSearch(ctx.DB, "sys_menu", "tree", "sys_menu", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuGet api implementation
// @Summary 获取菜单信息
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Param id query string false "菜单id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/menu/get [get]
func (ctr *SysMenu) SysMenuGet(ctx *Context) {
	var entity types.SysMenu
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
		ctx.Fail(types.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
