package svc

import (
	"encoding/json"
	"time"

	"golang.org/x/net/context"
)

// Cache defined TODO
type Cache interface {
	SetCache(key string, v interface{}, expire time.Duration) error
	GetCache(key string, v interface{}) error
}

// SetCache defined TODO
func (svc *SvcHepler) SetCache(key string, v interface{}, expire time.Duration) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = svc.rds.Set(context.Background(), key, string(bytes), expire).Err()
	return err
}

// GetCache defined TODO
func (svc *SvcHepler) GetCache(key string, v interface{}) error {
	str, err := svc.rds.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), v)
	return err
}
