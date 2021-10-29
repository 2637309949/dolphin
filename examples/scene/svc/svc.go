package svc

import (
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/gin-contrib/cache/persistence"
)

// ServiceContext defined TODO
type ServiceContext struct {
	Report svc.Report
	Store  svc.Store
	Client svc.Client
	DB     svc.DB
	Oss    Oss
	Weapp  Weapp
	Kafka  Kafka
}

// NewServiceContext defined TODO
func NewServiceContext(cache persistence.CacheStore) *ServiceContext {
	svc := svc.NewServiceContext(cache)
	return &ServiceContext{
		Store:  svc.Store,
		Report: svc.Report,
		Client: svc.Client,
		DB:     svc.DB,
		Oss:    NewXOss(),
		Weapp:  NewXWeapp(),
		Kafka:  NewXKafka(),
	}
}
