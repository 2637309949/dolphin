// Code generated by dol build. Only Generate by tools if not existed.
// source: article.go

package api

import (
	"scene/types"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/api"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// ArticleAdd api implementation
// @Summary 添加文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body types.Article false "文章信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/add [post]
func (ctr *Article) ArticleAdd(ctx *Context) {
	var payload types.Article
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
// @Param article body []types.Article false "文章信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/batch_add [post]
func (ctr *Article) ArticleBatchAdd(ctx *Context) {
	var payload []types.Article
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

// ArticleDel api implementation
// @Summary 删除文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body types.Article false "文章"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/del [delete]
func (ctr *Article) ArticleDel(ctx *Context) {
	var payload types.Article
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&types.Article{
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

// ArticleBatchDel api implementation
// @Summary 删除文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []types.Article false "文章信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/batch_del [put]
func (ctr *Article) ArticleBatchDel(ctx *Context) {
	var payload []types.Article
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.Article) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&types.Article{
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

// ArticleUpdate api implementation
// @Summary 更新文章
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body types.Article false "文章信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/update [put]
func (ctr *Article) ArticleUpdate(ctx *Context) {
	var payload types.Article
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
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
// @Param article body []types.Article false "文章信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/batch_update [put]
func (ctr *Article) ArticleBatchUpdate(ctx *Context) {
	var payload []types.Article
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
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFromNow()
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
// @Summary 文章分页查询
// @Tags 文章
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/page [get]
func (ctr *Article) ArticlePage(ctx *Context) {
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
	if ctr.Srv.Check(ctx.Context.Context) {
		ctr.Srv.SetOptionsetsFormat(api.OptionsetsFormat(ctx.DB))
		ret, err := ctr.Srv.PageExport(ctx.DB, "article", "page", "article", q.Value())
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		ctx.Success(ret)
		return
	}
	ret, err := ctr.Srv.PageSearch(ctx.DB, "article", "page", "article", q.Value())
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
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/get [get]
func (ctr *Article) ArticleGet(ctx *Context) {
	var entity types.Article
	err := ctx.ShouldBindWith(&entity)
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
		ctx.Fail(types.ErrNotFound)
		return
	}
	ctx.Success(entity)
}

// ArticlePayment api implementation
// @Summary 文章付费
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body types.ArticleInfo false "文章"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/article/payment [post]
func (ctr *Article) ArticlePayment(ctx *Context) {
	var payload types.ArticleInfo
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctr.Srv.TODO(ctx, ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
