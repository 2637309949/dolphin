package svc

import "time"

// Cache defined TODO
type Cache interface {
	SetCache(key string, v interface{}) error
	GetCache(key string, v interface{}, expire time.Duration) error
}

// SetCache defined TODO
func (svc *SvcHepler) SetCache(key string, v interface{}) error {
	return nil
}

// GetCache defined TODO
func (svc *SvcHepler) GetCache(key string, v interface{}, expire time.Duration) error {
	return nil
}
