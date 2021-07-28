package svc

import (
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/go-redis/redis/v8"
)

type SvcHepler struct {
	svc.Svc
}

func NewSvcHepler(rds redis.Cmdable) Svc {
	return &SvcHepler{Svc: svc.NewSvcHepler(rds)}
}
