// Code generated by dol build. Only Generate by tools if not existed.
// source: user.go

package app

import (
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
)

// UserInfo api implementation
// @Summary
// @Tags 用户信息
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/user/info [get]
func UserInfo(ctx *Context) {
	q := ctx.TypeQuery()
	q.Value()
	articles, err := ctx.DB.SqlMapClient("selectall_article").Query().List()
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if uids, ok := slice.GetFieldSliceByName(articles, "creater", "%v").([]string); ok {
		users, err := ctx.PlatformDB.Table("sys_user").In("id", uids).Where("is_delete != !").Cols("id", "name").QueryString()
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		err = slice.PatchSliceByField(articles, users, "creater", "id", "creater_name#name")(&articles)
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
	}
	ctx.Success(articles)
}
