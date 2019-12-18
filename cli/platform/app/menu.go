// Code generated by dol build. Only Generate by tools if not existed.
// source: menu.go

package app

import (
	"github.com/2637309949/dolphin/cli/platform/model"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"

	"github.com/2637309949/dolphin/cli/platform/util"
	"github.com/2637309949/dolphin/cli/platform/util/uuid"
)

// Menu struct
type Menu struct {
	*Engine
}

// BuildMenu return Menu
func BuildMenu(build func(*Menu)) func(engine *Engine) {
	return BuildEngine(func(engine *Engine) {
		build(&Menu{Engine: engine})
	})
}

// Add api implementation
// @Summary 添加菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param menu body model.Menu false "菜单信息"
// @Failure 403 {object} model.Response
// @Router /api/menu/add [post]
func (ctr *Menu) Add(ctx *Context) {
	var form model.Menu
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFrom(uuid.Must(uuid.NewRandom()).String())
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Update api implementation
// @Summary 更新菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param menu body model.Menu false "菜单信息"
// @Failure 403 {object} model.Response
// @Router /api/menu/update [post]
func (ctr *Menu) Update(ctx *Context) {
	var form model.Menu
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.ID(form.ID).Update(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// List api implementation
// @Summary 菜单分页查询
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Router /api/menu/list [get]
func (ctr *Menu) List(ctx *Context) {
	q := ctr.Query(ctx)
	q.SetInt("page")
	q.SetInt("size")
	ret, err := ctr.PageSearch(ctx.DB, "menu", "list", "menu", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Tree api implementation
// @Summary 菜单树形结构
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Response
// @Router /api/menu/tree [get]
func (ctr *Menu) Tree(ctx *Context) {
	q := ctr.Query(ctx)
	ret, err := util.AppAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// Get api implementation
// @Summary 获取菜单信息
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Param id query string false "菜单id"
// @Failure 403 {object} model.Response
// @Router /api/menu/get [get]
func (ctr *Menu) Get(ctx *Context) {
	var entity model.Menu
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
