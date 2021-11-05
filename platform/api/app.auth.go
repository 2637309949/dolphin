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

// Identity deifned TODO
type Identity struct {
	jwt    *JWT
	oauth2 *server.Server
	ticket *Token
}

// parseOAuth2Token defined TODO
func (i *Identity) parseOAuth2Token(t oauth2.TokenInfo) TokenInfo {
	i.ticket = &Token{
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
	return i.ticket
}

// parseJWTToken defined TODO
func (i *Identity) parseJWTToken(t *jwt.MapClaims) TokenInfo {
	i.ticket = &Token{
		UserID: (*t)["userId"].(string),
		Domain: (*t)["domain"].(string),
	}
	return i.ticket
}

// VerifyToken defined TODO
func (i *Identity) VerifyToken(req *http.Request) bool {
	if bearer, ok := i.oauth2.BearerAuth(req); ok {
		tk, err := i.oauth2.Manager.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return false
		}
		return i.parseOAuth2Token(tk).
			GetAccessCreateAt().
			Add(i.GetToken().GetAccessExpiresIn()).Round(0).Add(-TokenExpiryDelta).
			After(time.Now())
	}
	return false
}

// VerifyJWT defined TODO
func (i *Identity) VerifyJWT(req *http.Request) bool {
	if bearer, ok := i.jwt.BearerAuth(req); ok {
		tk, err := i.jwt.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return false
		}
		i.parseJWTToken(tk)
		return true
	}
	return false
}

// VerifyEncrypt defined TODO
func (i *Identity) VerifyEncrypt(ctx *Context) bool {
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

// GetToken defined TODO
func (i *Identity) GetToken() TokenInfo {
	return i.ticket
}

// AuthToken defined TODO
func AuthToken(ctx *Context) {
	identity := Identity{oauth2: App.OAuth2}
	if !identity.VerifyToken(ctx.Request()) {
		ctx.Fail(errors.ErrInvalidAccessToken)
		ctx.Abort()
		return
	}

	db := App.Manager.GetBusinessDB(identity.GetToken().GetDomain())
	if db == nil {
		ctx.Fail(errors.ErrInvalidDomain)
		return
	}
	ctx.Set("DB", db)
	ctx.Set("AuthInfo", identity.GetToken())
	ctx.Next()
}

// AuthJWT defined TODO
func AuthJWT(ctx *Context) {
	identity := Identity{jwt: App.JWT}
	if !identity.VerifyJWT(ctx.Request()) {
		ctx.Fail(errors.ErrInvalidAccessToken)
		ctx.Abort()
		return
	}
	db := App.Manager.GetBusinessDB(identity.GetToken().GetDomain())
	if db == nil {
		ctx.Fail(errors.ErrInvalidDomain)
		return
	}
	ctx.Set("DB", ctx.DB)
	ctx.Set("AuthInfo", identity.GetToken())
	ctx.Next()
}

// AuthEncrypt defined TODO
func AuthEncrypt(ctx *Context) {
	id := Identity{}
	if !id.VerifyEncrypt(ctx) {
		ctx.Fail(errors.ErrInvalidEncryData)
		return
	}
	ctx.Next()
}

// Roles middles TODO
func Roles(roles ...string) func(ctx *Context) {
	return func(ctx *Context) {
		svc, userId := svc.NewXDB(), ctx.MustToken().GetUserID()
		if !svc.InRole(ctx.MustDB(), userId, roles...) {
			ctx.Fail(errors.ErrAccessDenied)
			return
		}
		ctx.Next()
	}
}

// Auth middles TODO
func Auth(auth ...string) func(ctx *Context) {
	hlfs := []func(ctx *Context){}
	switch true {
	case slice.StrSliceContains(auth, TokenType):
		hlfs = append(hlfs, AuthToken)
		fallthrough
	case slice.StrSliceContains(auth, EncryptType):
		hlfs = append(hlfs, AuthEncrypt)
		fallthrough
	case slice.StrSliceContains(auth, JWTType):
		hlfs = append(hlfs, AuthJWT)
	}
	return func(ctx *Context) {
		for i := range hlfs {
			hlfs[i](ctx)
		}
	}
}
