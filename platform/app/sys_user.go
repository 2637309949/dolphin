// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user.go

package app

import (
	"errors"
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"

	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/srv"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// SysUserAdd api implementation
// @Summary 添加用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail "Invalid token"
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/add [post]
func SysUserAdd(ctx *Context) {
	var payload model.SysUser
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Domain = null.StringFrom(ctx.GetToken().GetDomain())
	payload.CreateTime = null.TimeFrom(time.Now())
	payload.Creater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now())
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	if payload.Avatar.IsZero() {
		payload.Avatar = null.StringFrom("http://pic.616pic.com/ys_bnew_img/00/06/27/TWk2P5YJ5k.jpg?imageView2/1/w/80/h/80")
	}
	payload.SetPassword("123456")
	ret, err := ctx.PlatformDB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserBatchAdd api implementation
// @Summary 添加用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/batch_add [post]
func SysUserBatchAdd(ctx *Context) {
	var payload []model.SysUser
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

// SysUserDel api implementation
// @Summary 删除用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/del [delete]
func SysUserDel(ctx *Context) {
	var payload model.SysUser
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.PlatformDB.In("id", payload.ID.Int64).Update(&model.SysUser{
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

// SysUserBatchDel api implementation
// @Summary 删除用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/batch_del [delete]
func SysUserBatchDel(ctx *Context) {
	var payload []*model.SysUser
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysUser) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.PlatformDB.In("id", ids).Update(&model.SysUser{
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

// SysUserUpdate api implementation
// @Summary 更新用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/update [put]
func SysUserUpdate(ctx *Context) {
	var payload struct {
		model.SysUser `xorm:"extends"`
		UserRole      null.String `json:"user_role" xml:"user_role"`
	}
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now())
	payload.Password.Valid = false
	payload.Salt.Valid = false

	ps := ctx.PlatformDB.NewSession()
	ds := ctx.DB.NewSession()
	defer ps.Close()
	defer ds.Close()
	ret, err := ps.Table(new(model.SysUser)).ID(payload.ID.Int64).Omit("user_role").Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		ps.Rollback()
		return
	}

	_, err = ds.SQL("delete from sys_role_user where user_id=?", payload.ID.Int64).Execute()
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		ps.Rollback()
		ds.Rollback()
		return
	}

	roleUsers := funk.Map(strings.Split(payload.UserRole.String, ","), func(role string) model.SysRoleUser {
		r, _ := strconv.ParseInt(role, 10, 64)
		return model.SysRoleUser{
			UserId:     payload.ID,
			RoleId:     null.IntFrom(r),
			CreateTime: null.TimeFrom(time.Now()),
			Creater:    null.IntFromStr(ctx.GetToken().GetUserID()),
			UpdateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
			IsDelete:   null.IntFrom(0),
		}
	}).([]model.SysRoleUser)
	_, err = ds.Insert(&roleUsers)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		ps.Rollback()
		ds.Rollback()
		return
	}
	err = ps.Commit()
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		ps.Rollback()
		ds.Rollback()
		return
	}
	err = ds.Commit()
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		ps.Rollback()
		ds.Rollback()
		return
	}
	ctx.Success(ret)
}

// SysUserBatchUpdate api implementation
// @Summary 更新用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/user/batch_update [put]
func SysUserBatchUpdate(ctx *Context) {
	var payload []model.SysUser
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
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
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

// SysUserPage api implementation
// @Summary 用户分页查询
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/page [get]
func SysUserPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("org_id")
	q.SetString("mobile")
	q.SetString("name")
	q.SetString("cn_org_id")
	q.SetRule("sys_user_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetTags()
	if q.GetString("cn_org_id") != "" {
		ids, err := srv.SysUserGetOrgsFromInheritance(ctx.DB, q.GetString("cn_org_id"))
		if err != nil {
			ctx.Fail(err)
			return
		}
		q.SetString("cn_org_id", template.HTML(strings.Join(ids, ",")))()
	}
	ret, err := ctx.PageSearch(ctx.PlatformDB, "sys_user", "page", "sys_user", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}

	if uids, ok := slice.GetFieldSliceByName(ret.Data, "id", "'%v'").([]string); ok {
		roles, err := srv.SysUserGetUserRolesByUID(ctx.DB, strings.Join(uids, ","))
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		err = slice.PatchSliceByField(ret.Data, roles, "id", "user_id", "role_name", "user_role")(&ret.Data)
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
	}

	if uorgs, ok := slice.GetFieldSliceByName(ret.Data, "org_id", "'%v'").([]string); ok {
		orgs, err := srv.SysUserGetUserOrgsByUID(ctx.DB, strings.Join(uorgs, ","))
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		err = slice.PatchSliceByField(ret.Data, orgs, "org_id", "id", "org_id#id", "org_name#name")(&ret.Data)
		if err != nil {
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
	}

	ret = ret.With(new([]struct {
		ID        null.Int    `json:"id" xml:"id"`
		Name      null.String `json:"name" xml:"name"`
		NickName  null.String `json:"nickname" xml:"nickname"`
		Mobile    null.String `json:"mobile" xml:"mobile"`
		Email     null.String `json:"email" xml:"email"`
		RoleName  null.String `json:"role_name" xml:"role_name"`
		UserRole  null.String `json:"user_role" xml:"user_role"`
		OrgName   null.String `json:"org_name" xml:"org_name"`
		OrgID     null.Int    `json:"org_id" xml:"org_id"`
		TempID    null.Int    `json:"temp_id" xml:"temp_id"`
		TempValue null.String `json:"temp_value" xml:"temp_value"`
	}))
	if ctx.QueryBool("__export__") {
		cfg := NewBuildExcelConfig(ret.Data)
		cfg.Format = OptionsetsFormat(ctx.DB)
		ctx.SuccessWithExcel(cfg)
		return
	}
	ctx.Success(ret)
}

// SysUserGet api implementation
// @Summary 获取用户信息
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Param id query string false "用户id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/get [get]
func SysUserGet(ctx *Context) {
	var entity model.SysUser
	id := ctx.Query("id")
	_, err := ctx.PlatformDB.ID(id).Get(&entity)
	entity.Password.Valid = false
	entity.Salt.Valid = false
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}

// SysUserLogin api implementation
// @Summary 用户认证
// @Tags 用户
// @Accept application/json
// @Param payload body model.Login false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/login [post]
func SysUserLogin(ctx *Context) {
	var payload, account = model.Login{}, model.SysUser{}
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if payload.Domain.String == "" {
		reg := regexp.MustCompile(`^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$`)
		base := reg.FindAllStringSubmatch(ctx.Request.Host, -1)
		payload.Domain = null.StringFrom(base[0][2])
	}
	account.Domain = payload.Domain
	account.Name = payload.Name

	ext, err := ctx.PlatformDB.Where("is_delete != 1 and status = 1").Get(&account)
	if err != nil || !ext || !account.ValidPassword(payload.Password.String) {
		if err == nil {
			err = errors.New("account doesn't exist or password error")
		}
		logrus.Errorf("SysUserLogin/ValidPassword:%v", err)
		ctx.Fail(err)
		return
	}
	token, err := App.OAuth2.Manager.GenerateAccessToken(oauth2.PasswordCredentials, &oauth2.TokenGenerateRequest{
		UserID:       fmt.Sprintf("%v", account.ID.Int64),
		Domain:       account.Domain.String,
		ClientID:     viper.GetString("oauth.id"),
		ClientSecret: viper.GetString("oauth.secret"),
		Request:      ctx.Request,
	})
	if err != nil {
		logrus.Errorf("SysUserLogin/GenerateAccessToken:%v", err)
		ctx.Fail(err)
		return
	}
	ctx.Success(map[string]interface{}{
		"access_token":  token.GetAccess(),
		"refresh_token": token.GetRefresh(),
	})
}

// SysUserLogout api implementation
// @Summary 用户退出登录
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/logout [get]
func SysUserLogout(ctx *Context) {
	err := App.OAuth2.Manager.RemoveAccessToken(ctx.GetToken().GetAccess())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success("successful")
}
