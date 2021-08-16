// Code generated by dol build. Only Generate by tools if not existed.
// source: app.go

package api

import (
	"context"
	"path"

	"github.com/2637309949/dolphin/packages/oauth2/errors"
	"github.com/2637309949/dolphin/packages/oauth2/generates"
	"github.com/2637309949/dolphin/packages/oauth2/manage"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	App *Dolphin
	Run func() error
)

// A Hook is a pair of start and stop callbacks, either of which can be nil,
// plus a string identifying the supplier of the hook.
type Hook struct {
	dol *Dolphin
}

// OnStart defined OnStart
func (h *Hook) OnStart(ctx context.Context) error {
	h.dol.Http.OnStart(ctx)
	h.dol.RPC.OnStart(ctx)
	return nil
}

// OnStop defined OnStop
func (h *Hook) OnStop(ctx context.Context) error {
	h.dol.Http.OnStop(ctx)
	h.dol.RPC.OnStop(ctx)
	return nil
}

// NewLifeHook defined TODO
func NewLifeHook(e *Dolphin) lifeHook {
	return &Hook{dol: e}
}

// WithHttpHandler defined TODO
func WithHttpHandler() Option {
	return func(dol *Dolphin) {
		dol.Http = NewGinHandler(dol)
	}
}

// WithRpcHandler defined TODO
func WithRpcHandler() Option {
	return func(dol *Dolphin) {
		dol.RPC = NewRpcHandler()
	}
}

// WithLifecycle defined TODO
func WithLifecycle() Option {
	return func(dol *Dolphin) {
		dol.Lifecycle = &lifecycleWrapper{}
		dol.Append(NewLifeHook(dol))
	}
}

// WithManager defined TODO
func WithManager() Option {
	return func(dol *Dolphin) {
		dol.Manager = NewDefaultManager()
	}
}

// WithJWT defined TODO
func WithJWT() Option {
	return func(dol *Dolphin) {
		dol.JWT = NewJWT(viper.GetString("jwt.secret"), viper.GetInt64("jwt.expire"))
	}
}

// WithOAuth2 defined TODO
func WithOAuth2() Option {
	osHost, httpPrefix := viper.GetString("oauth.server"), viper.GetString("http.prefix")
	authUrl := osHost + path.Join(httpPrefix, SysCasInstance.Authorize.RelativePath)
	tokenUrl := osHost + path.Join(httpPrefix, SysCasInstance.Token.RelativePath)
	endpoint := oauth2.Endpoint{AuthURL: authUrl, TokenURL: tokenUrl}
	redirectURL := viper.GetString("oauth.client") + path.Join(viper.GetString("http.prefix"), SysCasInstance.Oauth2.RelativePath)
	OA2Cfg = oauth2.Config{ClientID: viper.GetString("oauth.id"), ClientSecret: viper.GetString("oauth.secret"), Scopes: []string{"admin"}, RedirectURL: redirectURL, Endpoint: endpoint}
	qrRedirectURL := viper.GetString("oauth.client") + path.Join(viper.GetString("http.prefix"), SysCasInstance.QrOauth2.RelativePath)
	QrOA2Cfg = oauth2.Config{Scopes: []string{"admin"}, RedirectURL: qrRedirectURL}
	return func(dol *Dolphin) {
		manager := manage.NewDefaultManager()
		manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
		manager.MapTokenStorage(dol.Manager.GetTokenStore())
		manager.MapAccessGenerate(generates.NewAccessGenerate())
		manager.MapClientStorage(NewClientStore())
		manager.SetValidateURIHandler(ValidateURIHandler)
		dol.OAuth2 = server.NewServer(server.NewConfig(), manager)
		dol.OAuth2.SetUserAuthorizationHandler(UserAuthorizationHandler)
		dol.OAuth2.SetInternalErrorHandler(func(err error) (re *errors.Response) { logrus.Error(err); return })
		dol.OAuth2.SetResponseErrorHandler(func(re *errors.Response) { logrus.Error(re.Error) })
	}
}

// IsDebugging returns true if the framework is running in debug mode.
func IsDebugging() bool {
	return viper.GetString("app.mode") != "release"
}

// init after NewEngine
func init() {
	InitViper()
	InitLogger()
	InitRedisClient()
	InitSession()

	opts := []Option{WithHttpHandler(), WithRpcHandler(), WithLifecycle(), WithManager(), WithOAuth2(), WithJWT()}
	app := NewDefault(opts...)
	StaticRoutes(app)
	App, Run = app, app.Run
}
