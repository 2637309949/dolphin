// Code generated by dol build. Only Generate by tools if not existed.
// source: app.go

package app

import (
	"context"
	"path"

	"github.com/json-iterator/go/extra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	App *Dolphin
	Run func()
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

// NewLifeHook create lifecycle hook
func NewLifeHook(e *Dolphin) lifeHook {
	return &Hook{dol: e}
}

// init after NewEngine
func init() {
	extra.RegisterFuzzyDecoders()

	osHost, httpPrefix := viper.GetString("oauth.server"), viper.GetString("http.prefix")
	authUrl := osHost + path.Join(httpPrefix, SysCasInstance.Authorize.RelativePath)
	tokenUrl := osHost + path.Join(httpPrefix, SysCasInstance.Token.RelativePath)
	endpoint := oauth2.Endpoint{AuthURL: authUrl, TokenURL: tokenUrl}
	redirectURL := viper.GetString("oauth.cli") + path.Join(viper.GetString("http.prefix"), SysCasInstance.Oauth2.RelativePath)
	OA2Cfg = oauth2.Config{ClientID: viper.GetString("oauth.id"), ClientSecret: viper.GetString("oauth.secret"), Scopes: []string{"admin"}, RedirectURL: redirectURL, Endpoint: endpoint}

	app := NewDolphin()
	app.SyncModel()
	app.SyncController()
	app.SyncService()
	App, Run = app, app.Run
}
