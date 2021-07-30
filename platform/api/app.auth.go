// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"time"

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/golang-jwt/jwt"
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
	VerifyToken(*Context) bool
	VerifyEncrypt(*Context) bool
	VerifyJWT(*Context) bool
}

// AuthOAuth2 deifned
type AuthOAuth2 struct {
	oauth2 *server.Server
	jwt    *JWT
	ticket *Token
}

// Auth defined
func (auth *AuthOAuth2) parseOAuth2Token(t oauth2.TokenInfo) TokenInfo {
	auth.ticket = &Token{
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
	return auth.ticket
}

// Auth defined
func (auth *AuthOAuth2) parseJWTToken(t jwt.MapClaims) TokenInfo {
	auth.ticket = &Token{
		UserID: t["userId"].(string),
		Domain: t["domain"].(string),
	}
	return auth.ticket
}

// VerifyToken defined
func (auth *AuthOAuth2) VerifyToken(ctx *Context) bool {
	if bearer, ok := auth.oauth2.BearerAuth(ctx.Request); ok {
		accessToken, err := auth.oauth2.Manager.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return false
		}
		return auth.
			parseOAuth2Token(accessToken).
			GetAccessCreateAt().
			Add(auth.GetToken().GetAccessExpiresIn()).Round(0).Add(-TokenExpiryDelta).
			After(time.Now())
	}
	return false
}

// VerifyJWT defined
func (auth *AuthOAuth2) VerifyJWT(ctx *Context) bool {
	if bearer, ok := auth.jwt.BearerAuth(ctx.Request); ok {
		accessToken, err := auth.jwt.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return false
		}
		auth.parseJWTToken(*accessToken)
		return true
	}
	return false
}

// VerifyEncrypt defined
func (auth *AuthOAuth2) VerifyEncrypt(ctx *Context) bool {
	secret, err := NewSecret(ctx)
	if err != nil {
		logrus.Error(err)
		return false
	}
	cli, err := NewClientStore().GetByID(secret.AppID)
	if err != nil {
		logrus.Error(err)
		return false
	}
	valid, err := secret.Verify(cli)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return valid
}

// GetToken defined
func (auth *AuthOAuth2) GetToken() TokenInfo {
	return auth.ticket
}

// AuthToken defined
func AuthToken(ctx *Context) {
	if !ctx.VerifyToken(ctx) {
		ctx.Fail(types.ErrInvalidAccessToken, 401)
		ctx.Abort()
		return
	}
	if ctx.DB = App.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
		ctx.Fail(types.ErrInvalidDomain)
		ctx.Abort()
		return
	}
	ctx.Set("DB", ctx.DB)
	ctx.Set("AuthInfo", ctx.AuthInfo)
	ctx.Next()
}

// AuthJWT defined
func AuthJWT(ctx *Context) {
	if !ctx.VerifyJWT(ctx) {
		ctx.Fail(types.ErrInvalidAccessToken, 401)
		ctx.Abort()
		return
	}
	if ctx.DB = App.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
		ctx.Fail(types.ErrInvalidDomain)
		ctx.Abort()
		return
	}
	ctx.Set("DB", ctx.DB)
	ctx.Set("AuthInfo", ctx.AuthInfo)
	ctx.Next()
}

// AuthEncrypt defined
func AuthEncrypt(ctx *Context) {
	if !ctx.VerifyEncrypt(ctx) {
		ctx.Fail(types.ErrInvalidEncryData, 401)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// Roles middles
func Roles(roles ...string) HandlerFunc {
	return func(ctx *Context) {
		svcHelper, userId := new(svc.SvcHepler), ctx.GetToken().GetUserID()
		if !svcHelper.InRole(ctx.DB, userId, roles...) {
			ctx.Fail(types.ErrAccessDenied, 403)
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
	if slice.StrSliceContains(auth, JWTType) {
		middles = append(middles, AuthJWT)
	}
	return func(ctx *Context) {
		for i := range middles {
			middles[i](ctx)
		}
	}
}
