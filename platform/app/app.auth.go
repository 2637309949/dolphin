// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/2637309949/dolphin/packages/go-session/session"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/models"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xoauth2"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
)

// TokenExpiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const TokenExpiryDelta = 10 * time.Second

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
	Auth(*http.Request) bool
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

// Auth defined
func (auth *AuthOAuth2) Auth(req *http.Request) bool {
	if bearer, ok := auth.server.BearerAuth(req); ok {
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
	ext, err := App.PlatformDB.Where("client=?", id).Get(&cli)
	if err != nil || !ext {
		if !ext {
			err = errors.New("the record does not exist")
		}
		return nil, err
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
	reg := regexp.MustCompile("^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$")
	base := reg.FindAllStringSubmatch(baseURI, -1)
	redirect := reg.FindAllStringSubmatch(redirectURI, -1)
	if base[0][2] != redirect[0][2] {
		logrus.Errorf("baseURI=%v, redirectURI=%v", base[0][2], redirect[0][2])
		return errors.New("invalid redirect uri")
	}
	return nil
}

// Auth middles
func Auth(ctx *Context) {
	if !ctx.Auth(ctx.Request) {
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

// Roles middles
func Roles(roles ...string) func(ctx *Context) {
	return func(ctx *Context) {
		if !ctx.InRole(roles...) {
			ctx.Fail(util.ErrAccessDenied, 403)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
