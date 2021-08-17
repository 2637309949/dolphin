// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// RedisClient defined TODO
var RedisClient redis.Cmdable

// CacheStore defined TODO
var CacheStore persistence.CacheStore

// NewRedisClient defined TODO
func NewRedisClient() redis.Cmdable {
	network := viper.GetString("redis.network")
	username := viper.GetString("redis.username")
	password := viper.GetString("redis.password")
	addr := viper.GetStringSlice("redis.addr")
	db := viper.GetInt("redis.db")
	mode := viper.GetString("redis.mode")
	max_redirects := viper.GetInt("redis.max_redirects")
	read_only := viper.GetBool("redis.read_only")
	max_retries := viper.GetInt("redis.max_redirects")
	dial_timeout := viper.GetInt("redis.dial_timeout")
	read_timeout := viper.GetInt("redis.read_timeout")
	write_timeout := viper.GetInt("redis.write_timeout")
	pool_size := viper.GetInt("redis.pool_size")
	min_idle_conns := viper.GetInt("redis.min_idle_conns")
	idle_conns := viper.GetInt("redis.idle_conns")

	if len(addr) == 0 || mode == "" {
		return nil
	}

	var rdsCli redis.Cmdable
	var err error
	switch mode {
	case "stub":
		rdsCli = redis.NewClient(&redis.Options{
			Network:      network,
			Addr:         addr[0],
			DB:           db,
			Username:     username,
			Password:     password,
			MaxRetries:   max_retries,
			DialTimeout:  time.Duration(dial_timeout) * time.Second,
			ReadTimeout:  time.Duration(read_timeout) * time.Second,
			WriteTimeout: time.Duration(write_timeout) * time.Second,
			PoolSize:     pool_size,
			MinIdleConns: min_idle_conns,
			IdleTimeout:  time.Duration(idle_conns) * time.Second,
		})
		_, err = rdsCli.Ping(context.Background()).Result()
	case "cluster":
		rdsCli = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        addr,
			MaxRedirects: max_redirects,
			ReadOnly:     read_only,
			Password:     password,
			MaxRetries:   max_retries,
			DialTimeout:  time.Duration(dial_timeout) * time.Second,
			ReadTimeout:  time.Duration(read_timeout) * time.Second,
			WriteTimeout: time.Duration(write_timeout) * time.Second,
			PoolSize:     pool_size,
			MinIdleConns: min_idle_conns,
			IdleTimeout:  time.Duration(idle_conns) * time.Second,
		})
		_, err = rdsCli.Ping(context.Background()).Result()
	default:
		logrus.Fatal("redis mode must be one of (stub, cluster)")
	}

	if err != nil {
		logrus.Warnf("Redis:%v connect failed, %v", addr[0], err)
		return nil
	}
	logrus.Infof("Redis:%v connect successfully", addr[0])
	return rdsCli
}

// InitCacheStore defined TODO
func InitCacheStore() {
	CacheStore = persistence.NewInMemoryStore(60 * time.Second)
	RedisClient = NewRedisClient()
	if RedisClient != nil {
		CacheStore = NewRedisCache(RedisClient, 60*time.Second)
	}
}
