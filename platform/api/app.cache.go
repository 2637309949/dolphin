// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/2637309949/dolphin/packages/cache"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var PageCachePrefix = "dolphin.page.cache"

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

type cachedWriter struct {
	gin.ResponseWriter
	status  int
	written bool
	store   cache.CacheStore
	expire  time.Duration
	key     string
}

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

func newCachedWriter(store cache.CacheStore, expire time.Duration, writer gin.ResponseWriter, key string) *cachedWriter {
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
func CachePage(store cache.CacheStore, expire time.Duration) func(*Context) {
	return func(c *Context) {
		var rsp responseCache
		url := c.Request().URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &rsp); err != nil {
			if err != cache.ErrCacheMiss {
				log.Println(err.Error())
			}
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			c.Next()

			if c.IsAborted() {
				store.Delete(key)
			}
		} else {
			c.Writer.WriteHeader(rsp.Status)
			for k, vals := range rsp.Header {
				for _, v := range vals {
					c.Writer.Header().Set(k, v)
				}
			}
			c.Writer.Write(rsp.Data)
			c.Abort()
		}
	}
}

// Cache middles
func Cache(time time.Duration) func(ctx *Context) {
	return func(ctx *Context) {
		CachePage(CacheStore, time)(ctx)
	}
}
