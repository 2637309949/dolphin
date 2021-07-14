package svc

import (
	"github.com/2637309949/dolphin/platform/svc"
)

type SvcHepler struct {
	svc.Svc
}

func NewSvcHepler(svc svc.Svc) Svc {
	return &SvcHepler{Svc: svc}
}
