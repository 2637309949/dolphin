// Code generated by dol build. Only Generate by tools if not existed.
// source: app.go

package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/2637309949/dolphin/packages/json-iterator/extra"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/util"

	// github.com/2637309949/dolphin/platform/conf
	_ "github.com/2637309949/dolphin/platform/conf"
)

// HTTPServer defined http.Server
var HTTPServer *http.Server

// RPCListener defined net
var RPCListener net.Listener

// OnStart defined OnStart
func OnStart(e *Engine) func(context.Context) error {
	return func(context.Context) error {
		go func() {
			logrus.Infof("http listen on port:%v", viper.GetString("http.port"))
			HTTPServer.Handler = e.Gin
			if err := HTTPServer.ListenAndServe(); err != nil {
				logrus.Fatal(err)
			}
		}()
		go func() {
			logrus.Infof("grpc listen on port:%v", viper.GetString("grpc.port"))
			if err := e.GRPC.Serve(RPCListener); err != nil {
				logrus.Fatal(err)
			}
		}()
		return nil
	}
}

// OnStop defined OnStop
func OnStop(e *Engine) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		if err := HTTPServer.Shutdown(ctx); err != nil {
			logrus.Fatal(err)
			return err
		}
		if err := RPCListener.Close(); err != nil {
			logrus.Fatal(err)
			return err
		}
		return nil
	}
}

// NewLifeHook create lifecycle hook
func NewLifeHook(e *Engine) Hook {
	return Hook{
		OnStart: OnStart(e),
		OnStop:  OnStop(e),
	}
}

func init() {
	extra.RegisterFuzzyDecoders()
	AuthServerURL = viper.GetString("oauth.server")
	OA2Cfg.ClientID = viper.GetString("oauth.id")
	OA2Cfg.ClientSecret = viper.GetString("oauth.secret")
	OA2Cfg.RedirectURL = fmt.Sprintf("%v/api/sys/cas/oauth2", viper.GetString("oauth.cli"))
	OA2Cfg.Endpoint.AuthURL = AuthServerURL + "/api/sys/cas/authorize"
	OA2Cfg.Endpoint.TokenURL = AuthServerURL + "/api/sys/cas/token"
	HTTPServer = &http.Server{Addr: fmt.Sprintf(":%v", viper.GetString("http.port"))}
	RPCListener = util.EnsureLeft(net.Listen("tcp", fmt.Sprintf(":%v", viper.GetString("grpc.port")))).(net.Listener)
	SyncModel()
	SyncController()
	SyncService()
}
