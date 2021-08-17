package svc

import (
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/gin-contrib/cache/persistence"
)

type SvcHepler struct {
	svc.Svc
}

func NewSvcHepler(cacheStore persistence.CacheStore) Svc {
	return &SvcHepler{Svc: svc.NewSvcHepler(cacheStore)}
}
