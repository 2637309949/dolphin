// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"time"

	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/util/errors"
	"github.com/sirupsen/logrus"
)

// TokenExpiryDelta determines how earlier a token should be considered
const (
	TokenExpiryDelta = 10 * time.Second
)

// Provider deifned TODO
type Provider interface {
	GetName() string
	Config(*Identity)
	Verify(*Context) (TokenInfo, bool)
	Ticket(userId, extra string, ctx *Context) (TokenInfo, error)
}

// Identity deifned TODO
type Identity struct {
	jwt       *JWT
	OAuth2    *server.Server
	providers []Provider
}

// RegisterProvider register auth provider
func (i *Identity) RegisterProvider(provider Provider) {
	name := provider.GetName()
	for _, p := range i.providers {
		if p.GetName() == name {
			logrus.Infof("warning: auth provider %v already registered", name)
			return
		}
	}
	provider.Config(i)
	i.providers = append(i.providers, provider)
}

// GetProvider get provider with name
func (i *Identity) GetProvider(name string) Provider {
	for _, provider := range i.providers {
		if provider.GetName() == name {
			return provider
		}
	}
	return nil
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
	return func(ctx *Context) {
		for _, a := range auth {
			p := App.Identity.GetProvider(a)
			if p == nil {
				ctx.Fail(errors.ErrNotFoundProvider)
				return
			}

			tk, ok := p.Verify(ctx)
			if !ok {
				ctx.Fail(errors.ErrAuthenticationFailed)
				return
			}

			if tk.GetDomain() != "" {
				db := App.Manager.GetBusinessDB(tk.GetDomain())
				if db == nil {
					ctx.Fail(errors.ErrInvalidDomain)
					return
				}
				ctx.Set("DB", db)
			}

			ctx.Set("AuthInfo", tk)
			ctx.Next()
		}
	}
}
