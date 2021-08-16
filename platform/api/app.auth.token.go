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

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/models"
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

// UserAuthorizationHandler defined
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

		// load domain
		// var domain string
		// var redirect_uri string
		// if v, ok := store.Get("ReturnUri"); ok {
		// 	uri, _ := url.ParseQuery(v.(url.Values).Get("state"))
		// 	domain = uri.Get("domain")
		// 	redirect_uri = uri.Get("redirect_uri")
		// }
		w.Header().Set("Location", viper.GetString("oauth.login"))
		// w.Header().Add("Set-Cookie", (&http.Cookie{Name: "domain", Value: url.QueryEscape(domain), MaxAge: 60 * 60 * 30, Path: "/", Domain: "", Secure: false, HttpOnly: false}).String())
		// w.Header().Add("Set-Cookie", (&http.Cookie{Name: "redirect_uri", Value: url.QueryEscape(redirect_uri), MaxAge: 60 * 60 * 30, Path: "/", Domain: "", Secure: false, HttpOnly: false}).String())
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
