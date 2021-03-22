// Code generated by dol build. Only Generate by tools if not existed.
// source: organ.go

package app

import (
	"aisle/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// OrganAdd api implementation
// @Summary Add organ
// @Tags Organ controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.Organ false "Article info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/organ/add [post]
func OrganAdd(ctx *Context) {
	var payload model.Organ
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.CreateDate = null.TimeFrom(time.Now().Value())
	payload.Creater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateDate = null.TimeFrom(time.Now().Value())
	payload.Updater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.Isdelete = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// OrganDel api implementation
// @Summary Delete organ
// @Tags Organ controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param organ body model.Organ false "organ"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/organ/del [delete]
func OrganDel(ctx *Context) {
	var payload model.Organ
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("organ_id", payload.OrganId.Int64).Update(&model.Organ{
		UpdateDate: null.TimeFrom(time.Now().Value()),
		Updater:    null.StringFrom(ctx.GetToken().GetUserID()),
		Isdelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// OrganUpdate api implementation
// @Summary Update organ
// @Tags Organ controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.Organ false "Article info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/organ/update [put]
func OrganUpdate(ctx *Context) {
	var payload model.Organ
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	// payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateDate = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.OrganId.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// OrganPage api implementation
// @Summary Article page query
// @Tags Organ controller
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "Page number"
// @Param size  query  int false "Page size"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/organ/page [get]
func OrganPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetInt("is_delete", 0)
	q.SetRule("organ_page")
	q.SetRange("create_time")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "organ", "page", "organ", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// OrganGet api implementation
// @Summary Get organ info
// @Tags Organ controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "Article id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/organ/get [get]
func OrganGet(ctx *Context) {
	var entity model.Organ
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
