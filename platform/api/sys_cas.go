// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_cas.go

package api

import (
	"context"
	"fmt"
	"image/png"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/utils/uuid"
	"github.com/2637309949/dolphin/packages/web/gin"
	"github.com/2637309949/dolphin/platform/srv"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
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
func (ctr *SysCas) SysCasLogin(ctx *gin.Context) {
	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request())
	if err != nil {
		logrus.Errorf("SysCasLogin/Start:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+err.Error())
		return
	}
	ctx.Request().ParseForm()
	f := ctx.Request().Form
	domain, username, password := f.Get("domain"), f.Get("username"), f.Get("password")
	account := types.SysUser{
		Name:   null.StringFrom(username),
		Domain: null.StringFrom(domain),
	}

	// load domain from state
	if account.Domain.String == "" {
		if v, ok := store.Get("ReturnUri"); ok {
			sri, err := url.ParseQuery(v.(url.Values).Get("state"))
			if err != nil {
				logrus.Errorf("SysCasLogin/url.Parse:%v", err)
				ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+err.Error())
				return
			}
			account.Domain.String = sri.Get("domain")
		}
	}

	// load domain from request
	if account.Domain.String == "" {
		reg := regexp.MustCompile(`^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$`)
		base := reg.FindAllStringSubmatch(ctx.Request().Host, -1)
		account.Domain = null.StringFrom(base[0][2])
	}

	ext, err := App.PlatformDB.Where("is_delete = 0 and status = 1").Get(&account)
	if err != nil {
		logrus.Errorf("SysCasLogin/Where:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+err.Error())
		return
	}
	if !ext || !account.ValidPassword(password) {
		err = errors.ErrAccountOrPasswordError
		logrus.Errorf("SysCasLogin/ValidPassword:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.login")+"?error="+errors.ErrInvalidGrant.Error())
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
func (ctr *SysCas) SysCasLogout(ctx *gin.Context) {
	state := "domain=" + ctx.Query("domain") + "&redirect_uri=" + ctx.Query("redirect_uri") + "&state=" + ctx.Query("state")
	if ctx.Query("domain") == "" {
		logrus.Errorf("SysCasLogout/domain:%v", errors.ErrDomainNotFound)
		ctx.JSON(200, map[string]interface{}{
			"code":   "500",
			"detail": "ErrDomainNotFound",
			"status": "",
		})
		return
	}
	if ctx.Query("redirect_uri") == "" {
		logrus.Errorf("SysCasLogout/redirect_uri:%v", errors.ErrRedirectUriNotFound)
		ctx.JSON(200, map[string]interface{}{
			"code":   "500",
			"detail": "ErrRedirectUriNotFound",
			"status": "",
		})
		return
	}
	session.Destroy(context.Background(), ctx.Writer, ctx.Request())
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
func (ctr *SysCas) SysCasAffirm(ctx *gin.Context) {
	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request())
	if err != nil {
		logrus.Errorf("SysCasAffirm/Start:%v", err)
		ctx.Redirect(http.StatusFound, viper.GetString("oauth.affirm")+"?error="+err.Error())
		return
	}
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	ctx.Request().Form = form
	store.Delete("ReturnUri")
	store.Save()
	err = App.Identity.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request())
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
func (ctr *SysCas) SysCasAuthorize(ctx *gin.Context) {
	var form url.Values
	store, err := session.Start(context.Background(), ctx.Writer, ctx.Request())
	if err != nil {
		logrus.Errorf("SysCasAuthorize/Start:%v", err)
		ctx.Fail(err)
		return
	}
	ctx.Request().ParseForm()
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
		for k, v := range ctx.Request().Form {
			form.Set(k, v[0])
		}
		ctx.Request().Form = form
	}
	store.Delete("ReturnUri")
	store.Save()
	err = App.Identity.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request())
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
func (ctr *SysCas) SysCasToken(ctx *gin.Context) {
	err := App.Identity.OAuth2.HandleTokenRequest(ctx.Writer, ctx.Request())
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
func (ctr *SysCas) SysCasURL(ctx *gin.Context) {
	if ctx.Query("redirect_uri") == "" {
		ctx.JSON(200, map[string]interface{}{
			"code":   "500",
			"detail": "ErrRedirectUriNotFound",
			"status": "",
		})
		return
	}
	if ctx.Query("domain") == "" {
		ctx.JSON(200, map[string]interface{}{
			"code":   "500",
			"detail": "ErrDomainNotFound",
			"status": "ErrDomainNotFound",
		})
		return
	}
	state := "domain=" + ctx.Query("domain") + "&redirect_uri=" + ctx.Query("redirect_uri") + "&state=" + ctx.Query("state")
	uri, err := url.Parse(ctx.Query("redirect_uri"))
	if err != nil {
		ctx.Fail(err)
		return
	}
	if uri.Scheme == "" {
		uri.Scheme = "http"
	}
	redirectURL := uri.Scheme + "://" + uri.Host + "/oauth.html"
	ctx.Writer.Header().Add("Set-Cookie", (&http.Cookie{Name: "domain", Value: url.QueryEscape(ctx.Query("domain")), MaxAge: 60 * 60 * 30, Path: "/", Domain: "", Secure: false, HttpOnly: false}).String())
	ctx.Writer.Header().Add("Set-Cookie", (&http.Cookie{Name: "state", Value: url.QueryEscape(state), MaxAge: 60 * 60 * 30, Path: "/", Domain: "", Secure: false, HttpOnly: false}).String())
	ctx.Writer.Header().Add("Set-Cookie", (&http.Cookie{Name: "redirect_host", Value: uri.Scheme + "://" + uri.Host, MaxAge: 60 * 60 * 30, Path: "/", Domain: "", Secure: false, HttpOnly: false}).String())

	OA2Cfg.RedirectURL = redirectURL
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
func (ctr *SysCas) SysCasOauth2(ctx *gin.Context) {
	ctx.Request().ParseForm()
	f := ctx.Request().Form
	state, _ := url.QueryUnescape(f.Get("state"))
	code, _ := url.QueryUnescape(f.Get("code"))

	if code == "" {
		logrus.Errorf("SysCasOauth2/code:%v", errors.ErrCodeNotFound)
		ctx.JSON(200, map[string]interface{}{
			"code":   "500",
			"detail": "ErrCodeNotFound",
			"status": "",
		})
		return
	}
	ret, err := OA2Cfg.Exchange(context.Background(), code)
	if err != nil {
		logrus.Errorf("SysCasOauth2/Exchange:%v", err)
		ctx.JSON(200, map[string]interface{}{
			"code":   "500",
			"detail": err.Error(),
			"status": "",
		})
		return
	}
	reg := regexp.MustCompile("redirect_uri=([^&]*)?&state=([^&]*)?$")
	groups := reg.FindAllStringSubmatch(state, -1)
	redirect := groups[0][1]
	state = groups[0][2]
	ctx.JSON(200, map[string]interface{}{
		"code": "200",
		"data": map[string]interface{}{
			"access_token":  ret.AccessToken,
			"refresh_token": ret.RefreshToken,
			"redirect_uri":  redirect + "?state=" + state,
		},
	})
}

// SysCasQrOauth2 api implementation
// @Summary 授权回调
// @Tags 认证中心
// @Produce application/json
// @Param state  query  string false "状态码"
// @Param code  query  string false "临时令牌"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/qr_oauth2 [get]
func (ctr *SysCas) SysCasQrOauth2(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("state")
	q.SetString("code")

	if q.GetString("code") == "" {
		ctx.Fail(errors.ErrCodeNotFound)
		return
	}

	qrToken := types.QrToken{}
	err := CacheStore.Get("wechat:qrtoken:"+q.GetString("code"), &qrToken)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	token, err := App.Identity.OAuth2.Manager.GenerateAccessToken(oauth2.PasswordCredentials, &oauth2.TokenGenerateRequest{
		UserID:       fmt.Sprintf("%v", qrToken.UserId.Int64),
		Domain:       qrToken.Domain.String,
		ClientID:     viper.GetString("oauth.id"),
		ClientSecret: viper.GetString("oauth.secret"),
		Request:      ctx.Request(),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(map[string]interface{}{
		"access_token":  token.GetAccess(),
		"refresh_token": token.GetRefresh(),
		"redirect_uri":  qrToken.RedirectUri.String,
	})
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
	refreshtoken, ok := App.Identity.OAuth2.BearerAuth(ctx.Request())
	if !ok {
		logrus.Errorf("SysCasRefresh/BearerAuth:%v", ok)
		ctx.Fail(errors.ErrInvalidAccessToken)
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
	ret, err := ctr.Srv.TODO(ctx, ctx.MustDB(), struct{}{})
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
	err = ctx.MustDB().SqlTemplateClient("sys_cas_role.tpl", &map[string]interface{}{"user_id": user.ID.Int64}).Find(&roles)
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
	q, qrconnect := ctx.TypeQuery(), ""
	q.SetInt("width", 320)
	q.SetInt("height", 320)
	q.SetInt("type", 0)
	q.SetString("state")
	q.SetString("uuid")

	if q.GetString("uuid") == "" {
		ctx.Fail(errors.ErrUUIDNotFound)
		return
	}
	if len(q.GetString("uuid")) < 10 {
		ctx.Fail(errors.ErrInvalidUUID)
		return
	}
	if q.GetString("state") == "" {
		ctx.Fail(errors.ErrNotFoundState)
		return
	}

	appHost, httpPrefix := viper.GetString("app.host"), viper.GetString("http.prefix")
	wechatAuthUrl := appHost + path.Join(httpPrefix, SysWechatInstance.Oauth2.RelativePath)
	dingtalkAuthUrl := appHost + path.Join(httpPrefix, SysDingtalkInstance.Oauth2.RelativePath)

	state := url.QueryEscape("uuid=" + q.GetString("uuid") + "&" + q.GetString("state"))
	switch t := srv.SysCasQrType(q.GetInt("type")); t {
	case srv.SysCasQrTypeDingTalk:
		qrconnect = "https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_base&state=%v"
		qrconnect = fmt.Sprintf(qrconnect, viper.GetString("wc.appid"), url.QueryEscape(dingtalkAuthUrl), state)
	case srv.SysCasQrTypeWeiXin:
		qrconnect = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_base&state=%v"
		qrconnect = fmt.Sprintf(qrconnect, viper.GetString("wc.appid"), url.QueryEscape(wechatAuthUrl), state)
	default:
		ctx.Fail(fmt.Errorf("not support type:%v", t))
		return
	}

	qrCode, _ := qr.Encode(qrconnect, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, q.GetInt("width"), q.GetInt("height"))
	png.Encode(ctx.ResponseWriter(), qrCode)
}

// SysCasQrconnect api implementation
// @Summary 扫码地址
// @Tags 认证中心
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/qrconnect [get]
func (ctr *SysCas) SysCasQrconnect(ctx *gin.Context) {
	rq := ctx.Request().URL.RawQuery
	if rq != "" {
		rq = "?" + rq
	}
	ctx.Redirect(http.StatusFound, viper.GetString("oauth.qrconnect")+rq)
}

// SysCasQrcodeLogin api implementation
// @Summary 扫码登陆
// @Tags 认证中心
// @Produce application/json
// @Param uuid  query  string false "唯一id"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/cas/qrcode_login [get]
func (ctr *SysCas) SysCasQrcodeLogin(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("uuid")

	if q.GetString("uuid") == "" {
		ctx.Fail(errors.ErrUUIDNotFound)
		return
	}

	cwt, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		select {
		case <-cwt.Done():
			ctx.ResponseWriter().WriteHeader(http.StatusNoContent)
			return
		default:
			qrAuth := types.QrAuth{}
			err := CacheStore.Get("wechat:qrcode:"+q.GetString("uuid"), &qrAuth)
			if err == nil {
				CacheStore.Delete(q.GetString("uuid"))
				code := uuid.Must(uuid.NewRandom()).String()
				qrToken := types.QrToken{UserId: qrAuth.UserId, Domain: qrAuth.Domain, RedirectUri: qrAuth.RedirectUri, Code: null.StringFrom(code), TokenUri: null.StringFrom("")}
				CacheStore.Set("wechat:qrtoken:"+code, &qrToken, 6*time.Second)
				ctx.Success(qrToken)
				return
			}
		}
	}
	ctx.ResponseWriter().WriteHeader(http.StatusNoContent)
}
