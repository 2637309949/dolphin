package app

import (
	"fmt"
	"net/http"

	"github.com/2637309949/dolphin/cli/packages/logrus"
	"github.com/2637309949/dolphin/cli/packages/oauth2"
	"github.com/2637309949/dolphin/cli/packages/oauth2/server"
	"github.com/2637309949/dolphin/cli/packages/viper"
	xoauth2 "golang.org/x/oauth2"
)

// var defined
var (
	authServerURL = viper.GetString("oauth.server")
	oa2cfg        = xoauth2.Config{
		ClientID:     viper.GetString("oauth.id"),
		ClientSecret: viper.GetString("oauth.login"),
		Scopes:       []string{"all"},
		RedirectURL:  fmt.Sprintf("%v/api/oauth2/oauth2", viper.GetString("oauth.cli")),
		Endpoint: xoauth2.Endpoint{
			AuthURL:  authServerURL + "/api/oauth2/authorize",
			TokenURL: authServerURL + "/api/oauth2/token",
		},
	}
)

// define authorization model
var (
	OAuth2            AuthType = "oauth2"
	TokenkeyNamespace          = "dolphin:token"
)

// AuthType authorization model
type AuthType string

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
func (auth *AuthOAuth2) parseToken(t oauth2.TokenInfo) {
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
}

// Auth defined
func (auth *AuthOAuth2) Auth(req *http.Request) bool {
	if accessToken, ok := auth.server.BearerAuth(req); ok {
		token, err := auth.server.Manager.LoadAccessToken(accessToken)
		if err != nil {
			logrus.WithError(err).Warning("load accessToken failed.")
		} else {
			ok = true
			auth.parseToken(token)
			return true
		}
	}
	return false
}

// GetToken defined
func (auth *AuthOAuth2) GetToken() TokenInfo {
	return auth.token
}
