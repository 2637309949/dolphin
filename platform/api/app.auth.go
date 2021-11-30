// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2/server"
)

// Provider deifned TODO
type (
	Provider interface {
		GetName() string
		Config(*Identity)
		Verify(*Context) (TokenInfo, bool)
		Ticket(userId, extra string, ctx *Context) (TokenInfo, error)
	}
	Identity struct {
		JWT       *JWT
		OAuth2    *server.Server
		providers []Provider
	}
)

// RegisterProvider register auth provider
func (i *Identity) RegisterProvider(provider Provider) {
	name := provider.GetName()
	for _, p := range i.providers {
		if p.GetName() == name {
			logrus.Infof(context.TODO(), "warning: auth provider %v already registered", name)
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
