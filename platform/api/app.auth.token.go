// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"errors"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/models"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/go-session/session"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	xoauth2 "golang.org/x/oauth2"
)

var (
	// TokenkeyNamespace define TODO
	TokenkeyNamespace = "dolphin:token:"
	// OA2Cfg defined TODO
	OA2Cfg xoauth2.Config
	// QrOA2Cfg defined TODO
	QrOA2Cfg xoauth2.Config
)

// ClientStore client information store
type ClientStore struct {
}

// NewClientStore create client store
func NewClientStore() *ClientStore {
	return &ClientStore{}
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	cli := types.SysClient{}
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

// UserAuthorizationHandler defined TODO
var UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (uid string, dm string, err error) {
	store, err := session.Start(context.Background(), w, r)
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

// ValidateURIHandler defined TODO
var ValidateURIHandler = func(baseURI string, redirectURI string) error {
	reg := regexp.MustCompile(`^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$`)
	base := reg.FindAllStringSubmatch(baseURI, -1)
	redirect := reg.FindAllStringSubmatch(redirectURI, -1)

	// Skip check for ip redirect
	if util.IsIp(redirect[0][2]) {
		ip := net.ParseIP(redirect[0][2])
		if util.IsPrivate(ip) {
			return nil
		}
	}
	if base[0][2] != redirect[0][2] {
		logrus.Errorf("baseURI=%v, redirectURI=%v", base[0][2], redirect[0][2])
		return errors.New("invalid redirect uri")
	}
	return nil
}

var _ Provider = &TokenProvider{}

type TokenProvider struct {
	oauth2 *server.Server
}

// GetName defined TODO
func (p *TokenProvider) GetName() string {
	return "token"
}

// Config defined TODO
func (p *TokenProvider) Config(i *Identity) {
	p.oauth2 = i.OAuth2
}

// parseToken defined TODO
func (p *TokenProvider) parseToken(t oauth2.TokenInfo) TokenInfo {
	return &Token{
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
}

// Verify defined TODO
func (p *TokenProvider) Verify(ctx *Context) (TokenInfo, bool) {
	if bearer, ok := p.oauth2.BearerAuth(ctx.Request()); ok {
		tk, err := p.oauth2.Manager.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return nil, false
		}
		info := p.parseToken(tk)
		return info, info.
			GetAccessCreateAt().
			Add(info.GetAccessExpiresIn()).Round(0).Add(-10 * time.Second).
			After(time.Now())
	}
	return nil, false
}

// Verify defined TODO
func (p *TokenProvider) Ticket(userId, extra string, ctx *Context) (TokenInfo, error) {
	tk, err := p.oauth2.Manager.GenerateAccessToken(oauth2.PasswordCredentials, &oauth2.TokenGenerateRequest{
		UserID:       userId,
		Domain:       extra,
		ClientID:     viper.GetString("oauth.id"),
		ClientSecret: viper.GetString("oauth.secret"),
		Request:      ctx.Request(),
	})
	if err != nil {
		return nil, err
	}
	return p.parseToken(tk), nil
}
