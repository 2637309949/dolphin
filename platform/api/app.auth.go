// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/sirupsen/logrus"
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
