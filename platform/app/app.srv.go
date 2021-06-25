package app

import "time"

type Srv interface {
	SetCache(key string, v interface{}) error
	GetCache(key string, v interface{}, expire time.Duration) error
}

type srvHepler struct{}

func NewSrvHepler() Srv {
	return &srvHepler{}
}

func (srv *srvHepler) SetCache(key string, v interface{}) error {
	return nil
}

func (srv *srvHepler) GetCache(key string, v interface{}, expire time.Duration) error {
	return nil
}
