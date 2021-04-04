// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"time"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-session/session"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/models"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xoauth2"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/encrypt"
	"github.com/2637309949/dolphin/platform/util/slice"
)

// TokenExpiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const (
	TokenExpiryDelta = 10 * time.Second
	TokenType        = "token"
	EncryptType      = "encrypt"
)

// var defined
var (
	AuthServerURL   = viper.GetString("oauth.server")
	AuthAuthURL     = "/api/cas/authorize"
	AuthTokenURL    = "/api/cas/token"
	AuthRedirectURL = "/api/cas/oauth2"
	OA2Cfg          = xoauth2.Config{
		ClientID:     viper.GetString("oauth.id"),
		ClientSecret: viper.GetString("oauth.secret"),
		Scopes:       []string{"admin"},
		RedirectURL:  fmt.Sprintf("%v%v", AuthRedirectURL, viper.GetString("oauth.cli")),
		Endpoint: xoauth2.Endpoint{
			AuthURL:  AuthServerURL + AuthAuthURL,
			TokenURL: AuthServerURL + AuthTokenURL,
		},
	}
)

// TokenkeyNamespace define
var TokenkeyNamespace = "dolphin:token:"

// AuthInfo defined
type AuthInfo interface {
	GetToken() TokenInfo
	AuthToken(*Context) bool
	AuthEncrypt(*Context) (bool, error)
}

// EncryptForm defines Common request parameter
type EncryptForm struct {
	AppID      string `form:"app_id" json:"app_id" xml:"app_id" binding:"required"`
	Sign       string `form:"sign" json:"sign" xml:"sign" binding:"required"`
	TimeStamp  int64  `form:"timestamp" json:"timestamp" xml:"timestamp" binding:"required"`
	BizContent string `form:"biz_content" json:"biz_content" xml:"biz_content" binding:"required"`
}

// ParseForm defined
func (ec *EncryptForm) ParseForm(ctx *Context) (*EncryptForm, error) {
	puData := EncryptForm{}
	if ctx.Request.Method != "POST" {
		if err := ctx.BindQuery(&puData); err != nil {
			return nil, err
		}
	} else {
		if err := ctx.ShouldBindBodyWith(&puData, binding.JSON); err != nil {
			return nil, err
		}
	}
	return &puData, nil
}

// form2Uri defined
func (ec *EncryptForm) form2Uri() string {
	var puJSON map[string]string
	var puKeys = make([]string, 0, len(puJSON))
	puByte, _ := json.Marshal(ec)
	json.Unmarshal(puByte, &puJSON)
	for k := range puJSON {
		if k != "sign" {
			puKeys = append(puKeys, k)
		}
	}
	sort.Strings(puKeys)
	var signString = ""
	for _, k := range puKeys {
		if signString != "" {
			signString = signString + "&" + k + "=" + puJSON[k]
		} else {
			signString = signString + k + "=" + puJSON[k]
		}
	}
	return signString
}

func (ec *EncryptForm) sign(cli oauth2.ClientInfo) ([]byte, error) {
	uri := ec.form2Uri()
	ecyt, err := encrypt.AesEncrypt([]byte(uri), []byte(cli.GetSecret()))
	if err != nil {
		logrus.Error(err)
		return []byte{}, err
	}
	return ecyt, nil
}

// Verify defined
func (ec *EncryptForm) Verify(cli oauth2.ClientInfo) (bool, error) {
	nowTs := time.Now().Unix()
	ts := ec.TimeStamp
	if ts > nowTs || nowTs-ts >= 60 {
		return false, errors.New("timestamp error")
	}
	sn, err := ec.sign(cli)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	if string(sn) != ec.Sign {
		return false, errors.New("sign error")
	}
	return true, nil
}

// AuthOAuth2 deifned
type AuthOAuth2 struct {
	server *server.Server
	token  *Token
}

// Auth defined
func (auth *AuthOAuth2) parseToken(t oauth2.TokenInfo) TokenInfo {
	auth.token = &Token{
		ClientID:        t.GetClientID(),
		UserID:          t.GetUserID(),
		Domain:          t.GetDomain(),
		RedirectURI:     t.GetRedirectURI(),
		Scope:           t.GetScope(),
		Code:            t.GetCode(),
		CodeCreateAt:    t.GetCodeCreateAt(),
		CodeExpiresIn:   t.GetCodeExpiresIn(),
		Access:          t.GetAccess(),
		AccessCreateAt:  t.GetAccessCreateAt(),
		AccessExpiresIn: t.GetAccessExpiresIn(),
		Refresh:         t.GetRefresh(),
		RefreshCreateAt: t.GetRefreshCreateAt(),
	}
	return auth.token
}

// AuthToken defined
func (auth *AuthOAuth2) AuthToken(ctx *Context) bool {
	if bearer, ok := auth.server.BearerAuth(ctx.Request); ok {
		accessToken, err := auth.server.Manager.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return false
		}
		return auth.
			parseToken(accessToken).
			GetAccessCreateAt().
			Add(auth.GetToken().GetAccessExpiresIn()).Round(0).Add(-TokenExpiryDelta).
			After(time.Now())
	}
	return false
}

// AuthEncrypt defined
func (auth *AuthOAuth2) AuthEncrypt(ctx *Context) (bool, error) {
	parseForm, err := new(EncryptForm).ParseForm(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	cli, err := new(ClientStore).GetByID(parseForm.AppID)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	valid, err := parseForm.Verify(cli)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return valid, nil
}

// GetToken defined
func (auth *AuthOAuth2) GetToken() TokenInfo {
	return auth.token
}

// NewClientStore create client store
func NewClientStore() *ClientStore {
	return &ClientStore{}
}

// ClientStore client information store
type ClientStore struct {
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	cli := model.SysClient{}
	exist, err := App.PlatformDB.Where("client=?", id).Cols("client", "domain", "secret").Get(&cli)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("the record does not exist")
	}
	return &models.Client{
		ID:     cli.Client.String,
		Secret: cli.Secret.String,
		Domain: cli.Domain.String,
	}, nil
}

// UserAuthorizationHandler defined
var UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (uid string, dm string, err error) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		return
	}
	userID, uok := store.Get("LoggedInUserID")
	domain, dok := store.Get("LoggedInDomain")
	if !uok || !dok {
		if r.Form == nil {
			r.ParseForm()
		}
		store.Set("ReturnUri", r.Form)
		store.Save()
		w.Header().Set("Location", viper.GetString("oauth.login"))
		w.WriteHeader(http.StatusFound)
		return
	}
	uid = userID.(string)
	dm = domain.(string)
	store.Save()
	return
}

// ValidateURIHandler defined
var ValidateURIHandler = func(baseURI string, redirectURI string) error {
	reg := regexp.MustCompile(`^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$`)
	base := reg.FindAllStringSubmatch(baseURI, -1)
	redirect := reg.FindAllStringSubmatch(redirectURI, -1)
	if base[0][2] != redirect[0][2] {
		logrus.Errorf("baseURI=%v, redirectURI=%v", base[0][2], redirect[0][2])
		return errors.New("invalid redirect uri")
	}
	return nil
}

// AuthToken defined
func AuthToken(ctx *Context) {
	if !ctx.AuthToken(ctx) {
		ctx.Fail(util.ErrInvalidAccessToken, 401)
		ctx.Abort()
		return
	}
	if ctx.DB = App.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
		ctx.Fail(util.ErrInvalidDomain)
		ctx.Abort()
		return
	}
	ctx.Set("DB", ctx.DB)
	ctx.Set("AuthInfo", ctx.AuthInfo)
	ctx.Next()
}

// AuthEncrypt defined
func AuthEncrypt(ctx *Context) {
	valid, err := ctx.AuthEncrypt(ctx)
	if err != nil || !valid {
		ctx.Fail(err, 401)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// Auth middles
func Auth(auth ...string) HandlerFunc {
	middles := []func(ctx *Context){}
	if slice.StrSliceContains(auth, TokenType) {
		middles = append(middles, AuthToken)
	}
	if slice.StrSliceContains(auth, EncryptType) {
		middles = append(middles, AuthEncrypt)
	}
	return HF2Handler(func(ctx *Context) {
		for i := range middles {
			middles[i](ctx)
		}
	})
}

// Roles middles
func Roles(roles ...string) HandlerFunc {
	return HF2Handler(func(ctx *Context) {
		if !ctx.InRole(roles...) {
			ctx.Fail(util.ErrAccessDenied, 403)
			ctx.Abort()
			return
		}
		ctx.Next()
	})
}
