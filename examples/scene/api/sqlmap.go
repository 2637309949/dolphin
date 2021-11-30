// Code generated by dol build. Only Generate by tools if not existed.
// source: sqlmap.go

package api

import (
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/api"
	"github.com/2637309949/dolphin/platform/types"
)

// SqlmapSelectone api implementation
// @Summary 测试selectone
// @Tags SQL集合
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sqlmap/selectone [get]
func (ctr *Sqlmap) SqlmapSelectone(ctx *Context) {
	var user types.SysUser
	_, err := api.App.PlatformDB.SqlMapClient("selectone_sys_user", &map[string]interface{}{"id": types.DefaultAdmin.ID.ValueOrZero()}).Get(&user)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"nickname":   user.Nickname,
		"mobile":     user.Mobile,
		"email":      user.Email,
		"org_id":     user.OrgId,
		"temp_id":    user.TempId,
		"temp_value": user.TempValue,
	})
}
