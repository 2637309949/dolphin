// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_cas.go

package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/go-session/session"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
	xoauth2 "golang.org/x/oauth2"
)

// SysCasLogin api implementation
// @Summary 用户认证
// @Tags 认证中心
// @Accept multipart/form-data
// @Param username formData string false "用户名称"
// @Param password formData string false "用户密码"
// @Param domain formData string false "用户域"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/login [post]
func (ctr *SysCas) SysCasLogin(ctx *Context) {
	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("SysCasLogin/Start:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+err.Error())
		return
	}
	ctx.Request.ParseForm()
	f := ctx.Request.Form
	domain, username, password := f.Get("domain"), f.Get("username"), f.Get("password")
	account := types.SysUser{
		Name:   null.StringFrom(username),
		Domain: null.StringFrom(domain),
	}
	if account.Domain.String == "" {
		reg := regexp.MustCompile(`^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$`)
		base := reg.FindAllStringSubmatch(ctx.Request.Host, -1)
		account.Domain = null.StringFrom(base[0][2])
	}
	ext, err := ctx.PlatformDB.Where("is_delete = 0 and status = 1").Get(&account)

	if err != nil {
		logrus.Errorf("SysCasLogin/Where:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+err.Error())
		return
	}
	if !ext || !account.ValidPassword(password) {
		err = errors.New("account doesn't exist or password error")
		logrus.Errorf("SysCasLogin/ValidPassword:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+types.ErrInvalidGrant.Error())
		return
	}
	store.Set("LoggedInUserID", fmt.Sprintf("%v", account.ID.Int64))
	store.Set("LoggedInDomain", account.Domain.String)
	store.Save()
	ctx.Redirect(http.StatusFound, viper.GetString("oauth.affirm"))
}

// SysCasLogout api implementation
// @Summary 注销信息
// @Tags 认证中心
// @Param Authorization header string false "认证令牌"
// @Param redirect_uri query string false "定向URL"
// @Param state query string false "状态"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/logout [get]
func (ctr *SysCas) SysCasLogout(ctx *Context) {
	state := "redirect_uri=" + ctx.Query("redirect_uri") + "&state=" + ctx.Query("state")
	session.Destroy(context.Background(), ctx.Writer, ctx.Request)
	ctx.Redirect(302, OA2Cfg.AuthCodeURL(state))
}

// SysCasAffirm api implementation
// @Summary 用户授权
// @Tags 认证中心
// @Accept application/json
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/affirm [post]
func (ctr *SysCas) SysCasAffirm(ctx *Context) {
	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("SysCasAffirm/Start:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.affirm")+"?error="+err.Error())
		return
	}
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	ctx.Request.Form = form
	store.Delete("ReturnUri")
	store.Save()
	err = App.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("SysCasAffirm/HandleAuthorizeRequest:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.affirm")+"?error="+err.Error())
		return
	}
}

// SysCasAuthorize api implementation
// @Summary 用户授权
// @Tags 认证中心
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/authorize [get]
func (ctr *SysCas) SysCasAuthorize(ctx *Context) {
	var form url.Values
	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("SysCasAuthorize/Start:%v", err)
		ctx.Fail(err)
		return
	}
	ctx.Request.ParseForm()
	// ensure state, like state change
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
		for k, v := range ctx.Request.Form {
			form.Set(k, v[0])
		}
		ctx.Request.Form = form
	}
	store.Delete("ReturnUri")
	store.Save()
	err = App.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("SysCasAuthorize/HandleAuthorizeRequest:%v", err)
		ctx.Fail(err)
		return
	}
}

// SysCasToken api implementation
// @Summary 获取令牌
// @Tags 认证中心
// @Accept application/json
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/token [post]
func (ctr *SysCas) SysCasToken(ctx *Context) {
	err := App.OAuth2.HandleTokenRequest(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Errorf("SysCasToken/HandleTokenRequest:%v", err)
		ctx.Fail(err)
	}
}

// SysCasURL api implementation
// @Summary 授权地址
// @Tags 认证中心
// @Param redirect_uri query string false "定向URL"
// @Param state query string false "状态"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/url [get]
func (ctr *SysCas) SysCasURL(ctx *Context) {
	state := "redirect_uri=" + ctx.Query("redirect_uri") + "&state=" + ctx.Query("state")
	ctx.Redirect(302, OA2Cfg.AuthCodeURL(state))
}

// SysCasOauth2 api implementation
// @Summary 授权回调
// @Tags 认证中心
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/oauth2 [get]
func (ctr *SysCas) SysCasOauth2(ctx *Context) {
	ctx.Request.ParseForm()
	f := ctx.Request.Form
	state := f.Get("state")
	code := f.Get("code")
	state, _ = url.QueryUnescape(state)

	if code == "" {
		logrus.Errorf("SysCasOauth2/code:%v", errors.New("code not found"))
		ctx.Fail(errors.New("code not found"))
		return
	}
	ret, err := OA2Cfg.Exchange(context.Background(), code)
	if err != nil {
		logrus.Errorf("SysCasOauth2/Exchange:%v", errors.New("code not found"))
		ctx.Fail(err)
		return
	}
	reg := regexp.MustCompile("redirect_uri=([^&]*)?&state=([^&]*)?$")
	groups := reg.FindAllStringSubmatch(state, -1)
	rawRedirect := groups[0][1]
	rawState := groups[0][2]

	ctx.SetCookie("access_token", ret.AccessToken, 60*60*30, "/", "", false, false)
	ctx.SetCookie("refresh_token", ret.RefreshToken, 60*60*30, "/", "", false, false)
	if strings.TrimSpace(rawRedirect) != "" {
		ctx.Redirect(http.StatusFound, rawRedirect+"?state="+rawState)
		return
	}
	ctx.Redirect(http.StatusFound, "/?state="+rawState)
}

// SysCasRefresh api implementation
// @Summary 刷新令牌
// @Tags 认证中心
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/refresh [get]
func (ctr *SysCas) SysCasRefresh(ctx *Context) {
	refreshtoken, ok := App.OAuth2.BearerAuth(ctx.Request)
	if !ok {
		logrus.Errorf("SysCasRefresh/BearerAuth:%v", ok)
		ctx.Fail(types.ErrInvalidAccessToken)
		return
	}
	token := xoauth2.Token{}
	token.Expiry = time.Now()
	token.RefreshToken = refreshtoken
	ret, err := OA2Cfg.TokenSource(context.Background(), &token).Token()
	if err != nil {
		logrus.Errorf("SysCasRefresh/TokenSource:%v", err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCasCheck api implementation
// @Summary 检验令牌
// @Tags 认证中心
// @Param openid query string false "openid"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/check [get]
func (ctr *SysCas) SysCasCheck(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("openid")
	ret, err := ctr.Srv.TODO(ctx, ctx.DB, struct{}{})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCasProfile api implementation
// @Summary 用户信息
// @Tags 认证中心
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/profile [get]
func (ctr *SysCas) SysCasProfile(ctx *Context) {
	user := types.SysUser{}
	_, err := ctx.LoginInInfo(&user)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	roles := []types.SysRole{}
	err = ctx.DB.SqlTemplateClient("sys_cas_role.tpl", &map[string]interface{}{"user_id": user.ID.Int64}).Find(&roles)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	rolesArr := funk.Map(roles, func(role types.SysRole) string {
		return role.Code.String
	}).([]string)
	ctx.Success(map[string]interface{}{
		"roles":  rolesArr,
		"avatar": user.Avatar.String,
		"name":   user.Name.String,
	})
}

// SysCasQrcode api implementation
// @Summary 扫码登录(绑定第三方)
// @Tags 认证中心
// @Param Authorization header string false "认证令牌"
// @Param type query int false "类型(0: 微信 1：叮叮)"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/qrcode [get]
func (ctr *SysCas) SysCasQrcode(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("type", 0)
	if q.GetInt("type") == 0 {
		wcQrURL := "https://open.weixin.qq.com/connect/qrconnect?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_login&state=%v#wechat_redirect"
		wcQrURL = fmt.Sprintf(wcQrURL, viper.GetString("wc.appid"), "http://localhost:8082/api/sys/wechat/oauth2", "")
		ctx.Redirect(302, wcQrURL)
	} else if q.GetInt("type") == 1 {
		dtQrURL := "https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_login&state=%v"
		dtQrURL = fmt.Sprintf(dtQrURL, viper.GetString("wc.appid"), "http://localhost:8082/api/sys/wechat/oauth2", "")
		ctx.Redirect(302, dtQrURL)
	} else {
		ctx.Fail(fmt.Errorf("not found type %v", q.GetInt("type")))
		return
	}
}
