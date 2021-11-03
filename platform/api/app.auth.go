// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"time"

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/util/errors"
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

// AuthProtocol defined
type AuthProtocol interface {
	GetToken() TokenInfo
	VerifyToken(*http.Request) bool
	VerifyEncrypt(*http.Request) bool
	VerifyJWT(*http.Request) bool
}

// AuthOAuth2 deifned
type AuthOAuth2 struct {
	oauth2 *server.Server
	jwt    *JWT
	ticket *Token
}

// parseOAuth2Token defined TODO
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

// parseJWTToken defined TODO
func (auth *AuthOAuth2) parseJWTToken(t *jwt.MapClaims) TokenInfo {
	auth.ticket = &Token{
		UserID: (*t)["userId"].(string),
		Domain: (*t)["domain"].(string),
	}
	return auth.ticket
}

// VerifyToken defined TODO
func (auth *AuthOAuth2) VerifyToken(req *http.Request) bool {
	if bearer, ok := auth.oauth2.BearerAuth(req); ok {
		tk, err := auth.oauth2.Manager.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return false
		}
		return auth.
			parseOAuth2Token(tk).
			GetAccessCreateAt().
			Add(auth.GetToken().GetAccessExpiresIn()).Round(0).Add(-TokenExpiryDelta).
			After(time.Now())
	}
	return false
}

// VerifyJWT defined TODO
func (auth *AuthOAuth2) VerifyJWT(req *http.Request) bool {
	if bearer, ok := auth.jwt.BearerAuth(req); ok {
		tk, err := auth.jwt.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return false
		}
		auth.parseJWTToken(tk)
		return true
	}
	return false
}

// VerifyEncrypt defined TODO
func (auth *AuthOAuth2) VerifyEncrypt(req *http.Request) bool {
	secret, err := NewSecret(req)
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

// GetToken defined TODO
func (auth *AuthOAuth2) GetToken() TokenInfo {
	return auth.ticket
}

// AuthToken defined TODO
func AuthToken(ctx *Context) {
	if !ctx.VerifyToken(ctx.Request()) {
		ctx.Fail(errors.ErrInvalidAccessToken)
		ctx.Abort()
		return
	}
	if ctx.DB = App.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
		ctx.Fail(errors.ErrInvalidDomain)
		ctx.Abort()
		return
	}
	ctx.Set("DB", ctx.DB)
	ctx.Set("AuthProtocol", ctx.AuthProtocol)
	ctx.Next()
}

// AuthJWT defined TODO
func AuthJWT(ctx *Context) {
	if !ctx.VerifyJWT(ctx.Request()) {
		ctx.Fail(errors.ErrInvalidAccessToken)
		ctx.Abort()
		return
	}
	if ctx.DB = App.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
		ctx.Fail(errors.ErrInvalidDomain)
		ctx.Abort()
		return
	}
	ctx.Set("DB", ctx.DB)
	ctx.Set("AuthProtocol", ctx.AuthProtocol)
	ctx.Next()
}

// AuthEncrypt defined TODO
func AuthEncrypt(ctx *Context) {
	if !ctx.VerifyEncrypt(ctx.Request()) {
		ctx.Fail(errors.ErrInvalidEncryData)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// Roles middles TODO
func Roles(roles ...string) HandlerFunc {
	return func(ctx *Context) {
		svc, userId := svc.NewXDB(), ctx.GetToken().GetUserID()
		if !svc.InRole(ctx.DB, userId, roles...) {
			ctx.Fail(errors.ErrAccessDenied)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// Auth middles TODO
func Auth(auth ...string) HandlerFunc {
	hlfs := []HandlerFunc{}
	if slice.StrSliceContains(auth, TokenType) {
		hlfs = append(hlfs, AuthToken)
	}
	if slice.StrSliceContains(auth, EncryptType) {
		hlfs = append(hlfs, AuthEncrypt)
	}
	if slice.StrSliceContains(auth, JWTType) {
		hlfs = append(hlfs, AuthJWT)
	}
	return func(ctx *Context) {
		for i := range hlfs {
			hlfs[i](ctx)
		}
	}
}
