package svc

import (
	"github.com/2637309949/dolphin/packages/cache"
	"github.com/2637309949/dolphin/platform/svc"
)

// ServiceContext defined TODO
type ServiceContext struct {
	Report svc.Report
	Store  svc.Store
	Client svc.Client
	DB     svc.DB
}

// NewServiceContext defined TODO
func NewServiceContext(cache cache.CacheStore) *ServiceContext {
	svc := svc.NewServiceContext(cache)
	return &ServiceContext{
		Store:  svc.Store,
		Report: svc.Report,
		Client: svc.Client,
		DB:     svc.DB,
	}
}
