// Code generated by dol build. Only Generate by tools if not existed.
// source: article.go

package app

import (
	"blogger/model"
	"blogger/srv"
	"errors"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// ArticleAdd api implementation
// @Summary 添加文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.Article false "文章信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/add [post]
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

// ArticleBatchAdd api implementation
// @Summary 添加文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []model.Article false "文章信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/batch_add [post]
func ArticleBatchAdd(ctx *Context) {
	var payload []model.Article
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].ID = null.StringFromUUID()
		payload[i].CreateTime = null.TimeFrom(time.Now().Value())
		payload[i].CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].DelFlag = null.IntFrom(0)
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
// @Summary 删除文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.Article false "文章"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/del [delete]
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

// ArticleBatchDel api implementation
// @Summary 删除文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []model.Article false "文章信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/batch_del [put]
func ArticleBatchDel(ctx *Context) {
	var payload []model.Article
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.Article) string { return form.ID.String }).([]string)
	ret, err := ctx.DB.In("id", ids).Update(&model.Article{
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
// @Summary 更新文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.Article false "文章信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/update [put]
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

// ArticleBatchUpdate api implementation
// @Summary 更新文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []model.Article false "文章信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/batch_update [put]
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

// ArticlePage api implementation
// @Summary 文章分页查询
// @Tags 文章
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/page [get]
func ArticlePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
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
// @Summary 获取文章信息
// @Tags 文章
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "文章id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/get [get]
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

// ArticlePayment api implementation
// @Summary 文章付费
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.ArticleInfo false "文章"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/article/payment [post]
func ArticlePayment(ctx *Context) {
	var payload model.ArticleInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := srv.ArticleAction(ctx.Raw(), ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
