// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_comment.go

package api

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SysCommentAdd api implementation
// @Summary 添加评论
// @Tags 评论
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_comment body types.SysComment false "评论信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/comment/add [post]
func (ctr *SysComment) SysCommentAdd(ctx *Context) {
	var payload types.SysComment
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFromNow()
	payload.Creater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	ret, err := ctx.MustDB().Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCommentBatchAdd api implementation
// @Summary 批量添加评论
// @Tags 评论
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_comment body []types.SysComment false "评论信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/comment/batch_add [post]
func (ctr *SysComment) SysCommentBatchAdd(ctx *Context) {
	var payload []types.SysComment
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].CreateTime = null.TimeFromNow()
		payload[i].Creater = null.IntFromStr(ctx.MustToken().GetUserID())
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.MustToken().GetUserID())
		payload[i].IsDelete = null.IntFrom(0)
	}
	ret, err := ctx.MustDB().Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCommentDel api implementation
// @Summary 删除评论
// @Tags 评论
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_comment body types.SysUserTemplate false "评论"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/comment/del [delete]
func (ctr *SysComment) SysCommentDel(ctx *Context) {
	var payload types.SysUserTemplate
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.MustDB().In("id", payload.ID.Int64).Update(&types.SysUserTemplate{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCommentBatchDel api implementation
// @Summary 批量删除评论
// @Tags 评论
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_comment body []types.SysUserTemplate false "评论"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/comment/batch_del [delete]
func (ctr *SysComment) SysCommentBatchDel(ctx *Context) {
	var payload []types.SysUserTemplate
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form types.SysUserTemplate) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.MustDB().In("id", ids).Update(&types.SysUserTemplate{
		UpdateTime: null.TimeFromNow(),
		Updater:    null.IntFromStr(ctx.MustToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCommentUpdate api implementation
// @Summary 更新评论
// @Tags 评论
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_comment body types.SysComment false "评论信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/comment/update [put]
func (ctr *SysComment) SysCommentUpdate(ctx *Context) {
	var payload types.SysComment
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.MustToken().GetUserID())
	payload.UpdateTime = null.TimeFromNow()
	ret, err := ctx.MustDB().ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCommentBatchUpdate api implementation
// @Summary 批量更新评论
// @Tags 评论
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_comment body []types.SysComment false "评论信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router/api/sys/comment/batch_update [put]
func (ctr *SysComment) SysCommentBatchUpdate(ctx *Context) {
	var payload []types.SysComment
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.MustDB().NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFromNow()
		payload[i].Updater = null.IntFromStr(ctx.MustToken().GetUserID())
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

// SysCommentPage api implementation
// @Summary 评论分页查询
// @Tags 评论
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/comment/page [get]
func (ctr *SysComment) SysCommentPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("sys_comment_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	ret, err := ctr.Srv.DB.PageSearch(ctx.MustDB(), "sys_comment", "page", "sys_comment", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCommentGet api implementation
// @Summary 获取评论信息
// @Tags 评论
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "评论id"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/comment/get [get]
func (ctr *SysComment) SysCommentGet(ctx *Context) {
	var entity types.SysComment
	err := ctx.ShouldBindWith(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ext, err := ctx.MustDB().Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(errors.ErrNotFound)
		return
	}
	ctx.Success(entity)
}
