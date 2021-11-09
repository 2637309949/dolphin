// Code generated by dol build. Only Generate by tools if not existed.
// source: user.go

package api

import (
	"github.com/2637309949/dolphin/platform/api"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
)

// UserInfo api implementation
// @Summary
// @Tags 用户信息
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/user/info [get]
func (ctr *User) UserInfo(ctx *Context) {
	q := ctx.TypeQuery()
	q.Value()

	db := ctx.MustDB()
	articles, err := db.SqlMapClient("selectall_article").Query().List()
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if uids, ok := slice.GetFieldSliceByName(articles, "creater", "%v").([]float64); ok {
		users := []types.SysUser{}
		err := api.App.PlatformDB.Table("sys_user").In("id", uids).Where("is_delete != !").Cols("id", "name").Find(&users)
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		articles, err = slice.PatchSliceByField(articles, users, "creater", "id", "creater_name#name")
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
	}
	ctx.Success(articles)
}
