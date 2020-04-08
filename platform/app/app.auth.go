package app

import (
	"errors"
	"fmt"
	"net/http"
	gt "time"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/time"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/models"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/viper"
	xoauth2 "github.com/2637309949/dolphin/packages/xoauth2"
)

// expiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const expiryDelta = 10 * gt.Second

// var defined
var (
	authServerURL   = viper.GetString("oauth.server")
	authAuthURL     = "/api/cas/authorize"
	authTokenURL    = "/api/cas/token"
	authRedirectURL = "/api/cas/oauth2"
	oa2cfg          = xoauth2.Config{
		ClientID:     viper.GetString("oauth.id"),
		ClientSecret: viper.GetString("oauth.secret"),
		Scopes:       []string{"admin"},
		RedirectURL:  fmt.Sprintf("%v%v", authRedirectURL, viper.GetString("oauth.cli")),
		Endpoint: xoauth2.Endpoint{
			AuthURL:  authServerURL + authAuthURL,
			TokenURL: authServerURL + authTokenURL,
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

// Auth defined
func (auth *AuthOAuth2) Auth(req *http.Request) bool {
	if bearer, ok := auth.server.BearerAuth(req); ok {
		accessToken, err := auth.server.Manager.LoadAccessToken(bearer)
		if err != nil {
			logrus.WithError(err).Warning("load accessToken failed.")
		} else {
			return auth.
				parseToken(accessToken).
				GetAccessCreateAt().
				Add(auth.GetToken().GetAccessExpiresIn()).Round(0).Add(-expiryDelta).
				After(*time.Now().Value())
		}
	}
	return false
}

// GetToken defined
func (auth *AuthOAuth2) GetToken() TokenInfo {
	return auth.token
}

// NewClientStore create client store
func NewClientStore() *ClientStore {
	return &ClientStore{}
}

// ClientStore client information store
type ClientStore struct {
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	cli := model.SysClient{}
	ext, err := App.PlatformDB.Where("client=?", id).Get(&cli)
	if err != nil {
		return nil, err
	}
	if !ext {
		return nil, errors.New("the record does not exist")
	}
	return &models.Client{
		ID:     cli.Client.String,
		Secret: cli.Secret.String,
		Domain: cli.Domain.String,
	}, nil
}
