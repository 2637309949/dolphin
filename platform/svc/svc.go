package svc

import "github.com/go-redis/redis/v8"

// Svc defined TODO
type Svc interface {
	Cache
	Db
}

type SvcHepler struct {
	rds *redis.Client
}

func NewSvcHepler(rds *redis.Client) Svc {
	return &SvcHepler{rds: rds}
}
