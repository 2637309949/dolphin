// Code generated by dol build. Only Generate by tools if not existed.
// source: flarum_users.go

package app

import (
	"flarum/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// FlarumUsersAdd api implementation
// @Summary Add flarum_users
// @Tags FlarumUser controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.FlarumUsers false "FlarumUser info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/flarum/users/add [post]
func FlarumUsersAdd(ctx *Context) {
	var payload model.FlarumUsers
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
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

// FlarumUsersDel api implementation
// @Summary Delete flarum_users
// @Tags FlarumUser controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param flarum_users body model.FlarumUsers false "flarum_users"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/flarum/users/del [delete]
func FlarumUsersDel(ctx *Context) {
	var payload model.FlarumUsers
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&model.FlarumUsers{
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

// FlarumUsersUpdate api implementation
// @Summary Update flarum_users
// @Tags FlarumUser controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.FlarumUsers false "FlarumUser info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/flarum/users/update [put]
func FlarumUsersUpdate(ctx *Context) {
	var payload model.FlarumUsers
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// FlarumUsersPage api implementation
// @Summary FlarumUser page query
// @Tags FlarumUser controller
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "Page number"
// @Param size  query  int false "Page size"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/flarum/users/page [get]
func FlarumUsersPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetRule("flarum_users_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "flarum_users", "page", "flarum_users", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// FlarumUsersGet api implementation
// @Summary Get flarum_users info
// @Tags FlarumUser controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "FlarumUser id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/flarum/users/get [get]
func FlarumUsersGet(ctx *Context) {
	var entity model.FlarumUsers
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
