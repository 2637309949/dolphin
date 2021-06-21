package app

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/gob"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cache/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

type cachedWriter struct {
	gin.ResponseWriter
	status  int
	written bool
	store   persistence.CacheStore
	expire  time.Duration
	key     string
}

// RedisStore represents the cache with redis persistence
type RedisStore struct {
	redisClient       *redis.Client
	defaultExpiration time.Duration
}

var (
	PageCachePrefix = "dolphin.page.cache"
)

// RegisterResponseCacheGob registers the responseCache type with the encoding/gob package
func RegisterResponseCacheGob() {
	gob.Register(responseCache{})
}

// CreateKey creates a package specific key for a given string
func CreateKey(u string) string {
	return urlEscape(PageCachePrefix, u)
}

func urlEscape(prefix string, u string) string {
	key := url.QueryEscape(u)
	if len(key) > 200 {
		h := sha1.New()
		io.WriteString(h, u)
		key = string(h.Sum(nil))
	}
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(":")
	buffer.WriteString(key)
	return buffer.String()
}

func newCachedWriter(store persistence.CacheStore, expire time.Duration, writer gin.ResponseWriter, key string) *cachedWriter {
	return &cachedWriter{writer, 0, false, store, expire, key}
}

func (w *cachedWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *cachedWriter) Status() int {
	return w.ResponseWriter.Status()
}

func (w *cachedWriter) Written() bool {
	return w.ResponseWriter.Written()
}

func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil {
		store := w.store
		var cache responseCache
		if err := store.Get(w.key, &cache); err == nil {
			data = append(cache.Data, data...)
		}
		//cache responses with a status code < 300
		if w.Status() < 300 {
			val := responseCache{
				w.Status(),
				w.Header(),
				data,
			}
			err = store.Set(w.key, val, w.expire)
			if err != nil {
				logrus.Errorln(err)
			}
		}
	}
	return ret, err
}

func (w *cachedWriter) WriteString(data string) (n int, err error) {
	ret, err := w.ResponseWriter.WriteString(data)
	//cache responses with a status code < 300
	if err == nil && w.Status() < 300 {
		store := w.store
		val := responseCache{
			w.Status(),
			w.Header(),
			[]byte(data),
		}
		store.Set(w.key, val, w.expire)
	}
	return ret, err
}

// CachePage Decorator
func CachePage(store persistence.CacheStore, expire time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cache responseCache
		url := c.Request.URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &cache); err != nil {
			if err != persistence.ErrCacheMiss {
				log.Println(err.Error())
			}
			// replace writer
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			c.Next()
			// Drop caches of aborted contexts
			if c.IsAborted() {
				store.Delete(key)
			}
		} else {
			c.Writer.WriteHeader(cache.Status)
			for k, vals := range cache.Header {
				for _, v := range vals {
					c.Writer.Header().Set(k, v)
				}
			}
			c.Writer.Write(cache.Data)
			c.Abort()
		}
	}
}

// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCache(redis *redis.Client, defaultExpiration time.Duration) *RedisStore {
	if _, err := redis.Ping(context.Background()).Result(); err != nil {
		logrus.Error("Redis: connect failed")
	}
	return &RedisStore{redis, defaultExpiration}
}

// NewRedisCacheWithPool returns a RedisStore using the provided pool
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCacheWithPool(redisClient *redis.Client, defaultExpiration time.Duration) *RedisStore {
	return &RedisStore{redisClient, defaultExpiration}
}

// Set (see CacheStore interface)
func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {
	return c.invoke(key, value, expires)
}

// Add (see CacheStore interface)
func (c *RedisStore) Add(key string, value interface{}, expires time.Duration) error {
	return c.invoke(key, value, expires)
}

// Replace (see CacheStore interface)
func (c *RedisStore) Replace(key string, value interface{}, expires time.Duration) error {
	err := c.invoke(key, value, expires)
	if value == nil {
		return persistence.ErrNotStored
	}
	return err
}

// Get (see CacheStore interface)
func (c *RedisStore) Get(key string, ptrValue interface{}) error {
	buf, err := c.redisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return utils.Deserialize(buf, ptrValue)
}

func exists(client *redis.Client, key string) bool {
	iresult := client.Exists(context.Background(), key)
	if err := iresult.Err(); err != nil && err != redis.Nil {
		return false
	}
	return iresult.Val() == 0
}

// Delete (see CacheStore interface)
func (c *RedisStore) Delete(key string) error {
	if !exists(c.redisClient, key) {
		return persistence.ErrCacheMiss
	}
	result := c.redisClient.Del(context.Background(), key)
	if err := result.Err(); err != nil && err != redis.Nil {
		return err
	}
	return nil
}

// Increment (see CacheStore interface)
func (c *RedisStore) Increment(key string, delta uint64) (uint64, error) {
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
	val, err := c.redisClient.Get(context.Background(), key).Result()
	if val == "" {
		return 0, persistence.ErrCacheMiss
	}
	if err == nil {
		currentVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		sum := currentVal + int64(delta)
		_, err = c.redisClient.Set(context.Background(), key, sum, 0).Result()
		return uint64(sum), err
	}

	return 0, err
}

// Decrement (see CacheStore interface)
func (c *RedisStore) Decrement(key string, delta uint64) (newValue uint64, err error) {
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
	if !exists(c.redisClient, key) {
		return 0, persistence.ErrCacheMiss
	}
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
	val, _ := c.redisClient.Get(context.Background(), key).Result()
	if val == "" {
		return 0, persistence.ErrCacheMiss
	}
	currentVal, err := strconv.ParseInt(val, 10, 64)
	if err == nil && delta > uint64(currentVal) {
		tempint, err := c.redisClient.DecrBy(context.Background(), key, currentVal).Result()
		return uint64(tempint), err
	}
	tempint, err := c.redisClient.DecrBy(context.Background(), key, int64(delta)).Result()
	return uint64(tempint), err
}

// Flush (see CacheStore interface)
func (c *RedisStore) Flush() error {
	_, err := c.redisClient.FlushAll(context.Background()).Result()
	return err
}

func (c *RedisStore) invoke(key string, value interface{}, expires time.Duration) error {

	switch expires {
	case persistence.DEFAULT:
		expires = c.defaultExpiration
	case persistence.FOREVER:
		expires = time.Duration(0)
	}

	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}

	if expires > 0 {
		_, err := c.redisClient.SetNX(context.Background(), key, b, expires).Result()
		return err
	}

	_, err = c.redisClient.Set(context.Background(), key, b, expires).Result()
	return err
}

// Cache middles
func Cache(time time.Duration) HandlerFunc {
	return func(ctx *Context) {
		CachePage(CacheStore, time)(ctx.Context)
	}
}
