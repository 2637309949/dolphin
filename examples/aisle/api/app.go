// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: api.go

package api

import (
	"aisle/svc"

	"github.com/2637309949/dolphin/platform/api"
)

var App = api.App

type Context struct {
	*api.Context
}

// Run defined TODO
func Run() {
	SyncModel()
	SyncSrv(svc.NewServiceContext(api.CacheStore))
	SyncController()
	err := api.Run()
	if err != nil {
		panic(err)
	}
}
