// Code generated by dol build. Only Generate by tools if not existed.
// source: article.go

package app

import (
	"oauth/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// ArticleAdd api implementation
// @Summary Add article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.Article false "Article info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/add [post]
func ArticleAdd(ctx *Context) {
	var payload model.Article
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

// ArticleDel api implementation
// @Summary Delete article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.Article false "article"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/del [delete]
func ArticleDel(ctx *Context) {
	var payload model.Article
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.Article{
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

// ArticleUpdate api implementation
// @Summary Update article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.Article false "Article info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/update [put]
func ArticleUpdate(ctx *Context) {
	var payload model.Article
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

// ArticlePage api implementation
// @Summary Article page query
// @Tags Article controller
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "Page number"
// @Param size  query  int false "Page size"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/page [get]
func ArticlePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetRule("article_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "article", "page", "article", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// ArticleGet api implementation
// @Summary Get article info
// @Tags Article controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "Article id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/get [get]
func ArticleGet(ctx *Context) {
	var entity model.Article
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
