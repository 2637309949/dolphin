// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"time"

	"github.com/2637309949/dolphin/packages/cache"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// RedisClient defined TODO
var RedisClient redis.Cmdable

// CacheStore defined TODO
var CacheStore cache.CacheStore

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

	if len(addr) == 0 {
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
		logrus.Fatal(context.TODO(), "redis mode must be one of (stub, cluster)")
	}

	if err != nil {
		logrus.Warnf(context.TODO(), "redis:%v", err)
		return nil
	}
	logrus.Infof(context.TODO(), "redis:%v connect successfully", addr[0])
	return rdsCli
}

// InitCacheStore defined TODO
func InitCacheStore() {
	if RedisClient = NewRedisClient(); RedisClient != nil {
		CacheStore = cache.NewRedisCache(RedisClient, 60*time.Second)
		logrus.Infof(context.TODO(), "use Redis as CacheStore")
	} else {
		CacheStore = cache.NewInMemoryStore(60 * time.Second)
		logrus.Infof(context.TODO(), "use Memory as CacheStore")
	}
}
