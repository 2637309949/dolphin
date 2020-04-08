// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_cas.go

package app

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/go-session/cookie"
	"github.com/2637309949/dolphin/packages/go-session/session"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xoauth2"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/srv"
	"github.com/2637309949/dolphin/platform/util"
)

func init() {
	var hashKey = []byte(viper.GetString("http.hash"))
	session.InitManager(
		session.SetCookieName("session_id"),
		session.SetStore(
			cookie.NewCookieStore(
				cookie.SetCookieName("store_id"),
				cookie.SetHashKey(hashKey),
			),
		),
	)
	// fix: store.Set("ReturnUri", r.Form) can not be parsed.
	gob.Register(url.Values{})
}

// SysCasLogin api implementation
// @Summary 登录信息
// @Tags OAuth授权
// @Accept application/json
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/login [post]
func SysCasLogin(ctx *Context) {
	store, err := session.Start(nil, ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Error("SysCasLogin/Start:", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+err.Error())
		return
	}
	ctx.Request.ParseForm()
	f := ctx.Request.Form
	domain := f.Get("domain")
	username := f.Get("username")
	password := f.Get("password")
	account := model.SysUser{
		Name:   null.StringFrom(username),
		Domain: null.StringFrom(domain),
	}
	if account.Domain.String == "" {
		reg := regexp.MustCompile("^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$")
		base := reg.FindAllStringSubmatch(ctx.Request.Host, -1)
		account.Domain = null.StringFrom(base[0][2])
	}
	ext, err := ctx.engine.PlatformDB.Where("del_flag = 0 and status = 1").Get(&account)

	if err != nil {
		logrus.Error("SysCasLogin/Where:", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+err.Error())
		return
	}
	if !ext || !account.ValidPassword(password) {
		logrus.Error("SysCasLogin/ValidPassword:", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+util.ErrInvalidGrant.Error())
		return
	}
	store.Set("LoggedInUserID", account.ID.String)
	store.Set("LoggedInDomain", account.Domain.String)
	store.Save()
	ctx.Redirect(http.StatusFound, viper.GetString("oauth.affirm"))
}

// SysCasAffirm api implementation
// @Summary 用户授权
// @Tags OAuth授权
// @Accept application/json
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/affirm [post]
func SysCasAffirm(ctx *Context) {
	store, err := session.Start(nil, ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Error("SysCasAffirm/Start:", err)
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
	err = ctx.engine.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Error("SysCasAffirm/HandleAuthorizeRequest:", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.affirm")+"?error="+err.Error())
		return
	}
}

// SysCasAuthorize api implementation
// @Summary 用户授权
// @Tags OAuth授权
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/authorize [get]
func SysCasAuthorize(ctx *Context) {
	var form url.Values
	store, err := session.Start(nil, ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Error("SysCasAuthorize/Start:", err)
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

	err = ctx.engine.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Error("SysCasAuthorize/HandleAuthorizeRequest:", err)
		ctx.Fail(err)
		return
	}
}

// SysCasToken api implementation
// @Summary 获取令牌
// @Tags OAuth授权
// @Accept application/json
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/token [post]
func SysCasToken(ctx *Context) {
	err := ctx.engine.OAuth2.HandleTokenRequest(ctx.Writer, ctx.Request)
	if err != nil {
		logrus.Error("SysCasToken/HandleTokenRequest:", err)
		ctx.Fail(err)
	}
}

// SysCasURL api implementation
// @Summary 授权地址
// @Tags 认证中心
// @Param redirect_uri query string false "定向URL"
// @Param state query string false "状态"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/url [get]
func SysCasURL(ctx *Context) {
	state := "redirect_uri=" + ctx.Query("redirect_uri") + "&state=" + ctx.Query("state")
	ctx.Redirect(302, OA2Cfg.AuthCodeURL(state))
}

// SysCasOauth2 api implementation
// @Summary 授权回调
// @Tags 认证中心
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/oauth2 [get]
func SysCasOauth2(ctx *Context) {
	ctx.Request.ParseForm()
	f := ctx.Request.Form
	state := f.Get("state")
	code := f.Get("code")

	if code == "" {
		logrus.Error("SysCasOauth2/code:", errors.New("Code not found"))
		ctx.Fail(errors.New("Code not found"))
		return
	}
	ret, err := OA2Cfg.Exchange(context.Background(), code)
	if err != nil {
		logrus.Error("SysCasOauth2/Exchange:", errors.New("Code not found"))
		ctx.Fail(err)
		return
	}

	reg := regexp.MustCompile("redirect_uri=([^&]*)?&state=([^&]*)?$")
	groups := reg.FindAllStringSubmatch(state, -1)
	rawRedirect := groups[0][1]
	rawState := groups[0][2]

	ctx.SetCookie("access_token", ret.AccessToken, 60*60*30, "/", "", false, true)
	ctx.SetCookie("refresh_token", ret.RefreshToken, 60*60*30, "/", "", false, true)
	if strings.TrimSpace(rawRedirect) != "" {
		urlRedirect, err := url.Parse(rawRedirect)
		if err != nil {
			logrus.Error("SysCasOauth2/Parse:", err)
			ctx.Fail(err)
			return
		}
		redirect := urlRedirect.Path
		ctx.Redirect(http.StatusFound, redirect+"?state="+rawState)
	} else {
		ctx.Redirect(http.StatusFound, "/?state="+rawState)
	}
	ctx.Success(ret)
}

// SysCasRefresh api implementation
// @Summary 刷新令牌
// @Tags OAuth授权
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/refresh [get]
func SysCasRefresh(ctx *Context) {
	refreshtoken, ok := ctx.engine.OAuth2.BearerAuth(ctx.Request)
	if !ok {
		logrus.Error("SysCasRefresh/BearerAuth:", ok)
		ctx.Fail(util.ErrInvalidAccessToken)
		return
	}
	token := xoauth2.Token{}
	token.Expiry = time.Now()
	token.RefreshToken = refreshtoken
	ret, err := OA2Cfg.TokenSource(context.Background(), &token).Token()
	if err != nil {
		logrus.Error("SysCasRefresh/TokenSource:", err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCasCheck api implementation
// @Summary 检验令牌
// @Tags 认证中心
// @Param openid query string false "openid"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/check [get]
func SysCasCheck(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("openid")
	ret, err := srv.SysCasAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCasProfile api implementation
// @Summary 用户信息
// @Tags 认证中心
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/profile [get]
func SysCasProfile(ctx *Context) {
	q := ctx.TypeQuery()
	ret, err := srv.SysCasAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysCasQrcode api implementation
// @Summary 扫码登录(绑定第三方)
// @Tags 认证中心
// @Param Authorization header string false "认证令牌"
// @Param type query int false "类型(0: 微信 1：叮叮)"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/cas/qrcode [get]
func SysCasQrcode(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("type", 0)
	if q.GetInt("type") == 0 {
		wcQrURL := "https://open.weixin.qq.com/connect/qrconnect?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_login&state=%v#wechat_redirect"
		wcQrURL = fmt.Sprintf(wcQrURL, viper.GetString("wc.appid"), "http://127.0.0.1:8082/api/sys/wechat/oauth2", "")
		ctx.Redirect(302, wcQrURL)
	} else if q.GetInt("type") == 1 {
		dtQrURL := "https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_login&state=%v"
		dtQrURL = fmt.Sprintf(dtQrURL, viper.GetString("wc.appid"), "http://127.0.0.1:8082/api/sys/wechat/oauth2", "")
		ctx.Redirect(302, dtQrURL)
	} else {
		ctx.Fail(fmt.Errorf("not found type %v", q.GetInt("type")))
		return
	}
}
