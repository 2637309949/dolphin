// Code generated by dol build. Only Generate by tools if not existed.
// source: article.go

package app

import (
	"errors"
	"mdclub/model"

	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// ArticleAdd api implementation
// @Summary Add article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.Article false "Article info"
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
	payload.CreateTime = null.TimeFrom(time.Now())
	payload.Creater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now())
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// ArticleBatchAdd api implementation
// @Summary Add article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []model.Article false "Article info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/batch_add [post]
func ArticleBatchAdd(ctx *Context) {
	var payload []model.Article
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
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
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&model.Article{
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

// ArticleBatchDel api implementation
// @Summary Delete article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []model.Article false "article"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/batch_del [delete]
func ArticleBatchDel(ctx *Context) {
	var payload []model.Article
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.Article) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&model.Article{
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

// ArticleUpdate api implementation
// @Summary Update article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.Article false "Article info"
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

// ArticleBatchUpdate api implementation
// @Summary Update article
// @Tags Article controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []model.Article false "Article info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/batch_update [put]
func ArticleBatchUpdate(ctx *Context) {
	var payload []model.Article
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	s.Begin()
	defer s.Close()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.Int64).Update(&payload[i])
		if err != nil {
			logrus.Error(err)
			s.Rollback()
			ctx.Fail(err)
			return
		}
		ret = append(ret, r)
	}
	if err != nil {
		logrus.Error(err)
		s.Rollback()
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
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetString("sort", "update_time desc")
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
	err := ctx.ShouldBindQuery(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if ext, err := ctx.DB.Get(&entity); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	} else if !ext {
		ctx.Fail(errors.New("not found"))
		return
	}
	ctx.Success(entity)
}