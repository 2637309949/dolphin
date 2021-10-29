package persistence

import (
	"context"
	"strconv"
	"time"

	"github.com/2637309949/dolphin/packages/persistence/utils"
	"github.com/gin-contrib/cache/persistence"
	"github.com/go-redis/redis/v8"
)

// RedisStore represents the cache with redis persistence
type RedisStore struct {
	redisClient       redis.Cmdable
	defaultExpiration time.Duration
}

// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCache(redis redis.Cmdable, defaultExpiration time.Duration) *RedisStore {
	return &RedisStore{redis, defaultExpiration}
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
		return ErrNotStored
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

func exists(client redis.Cmdable, key string) bool {
	iresult := client.Exists(context.Background(), key)
	if err := iresult.Err(); err != nil && err != redis.Nil {
		return false
	}
	return iresult.Val() == 0
}

// Delete (see CacheStore interface)
func (c *RedisStore) Delete(key string) error {
	if !exists(c.redisClient, key) {
		return ErrCacheMiss
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
		return 0, ErrCacheMiss
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
		return 0, ErrCacheMiss
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
