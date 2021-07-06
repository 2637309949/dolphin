// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"time"

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
)

// TokenExpiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const (
	TokenExpiryDelta = 10 * time.Second
	TokenType        = "token"
	EncryptType      = "encrypt"
	JWTType          = "jwt"
)

// AuthInfo defined
type AuthInfo interface {
	GetToken() TokenInfo
	AuthToken(*Context) bool
	AuthEncrypt(*Context) (bool, error)
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
	nef := NewEncryptForm()
	parseForm, err := nef.ParseForm(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	cli, err := NewClientStore().GetByID(parseForm.AppID)
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

// Roles middles
func Roles(roles ...string) HandlerFunc {
	return func(ctx *Context) {
		if !ctx.InRole(roles...) {
			ctx.Fail(util.ErrAccessDenied, 403)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
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
	return func(ctx *Context) {
		for i := range middles {
			middles[i](ctx)
		}
	}
}
