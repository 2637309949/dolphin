// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/2637309949/dolphin/packages/cache"
	"github.com/2637309949/dolphin/packages/cache/persistence"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/store"
	"github.com/2637309949/dolphin/packages/redis"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/http"
	"github.com/2637309949/dolphin/platform/util/worker"
)

// Worker defined
type Worker interface {
	AddJob(j worker.Job)
	AddJobHandler(string, func(model.Worker) (interface{}, error))
	GetJobHandler(string) func(model.Worker) (interface{}, error)
	Save(model.Worker) error
	Update(model.Worker) error
	Find(string) (model.Worker, error)
}

// Manager Engine management interface
type Manager interface {
	MSet() MSeti
	Worker() Worker
	GetBusinessDBSet() map[string]*xorm.Engine
	GetBusinessDB(string) *xorm.Engine
	AddBusinessDB(string, *xorm.Engine)
	GetTokenStore() oauth2.TokenStore
}

// DefaultWorker defined
type DefaultWorker struct {
	store       map[string]model.Worker
	JobHandlers map[string]func(model.Worker) (interface{}, error)
	Dispatcher  *worker.Dispatcher
}

// AddJob defined
func (d *DefaultWorker) AddJob(j worker.Job) {
	d.Dispatcher.AddJob(j)
}

// Find defined
func (d *DefaultWorker) Find(code string) (model.Worker, error) {
	return d.store[code], nil
}

// Save defined
func (d *DefaultWorker) Save(w model.Worker) error {
	d.store[w.Code] = w
	return nil
}

// Update defined
func (d *DefaultWorker) Update(w model.Worker) error {
	d.store[w.Code] = w
	return nil
}

// AddJobHandler defined
func (d *DefaultWorker) AddJobHandler(code string, funk func(model.Worker) (interface{}, error)) {
	if d.JobHandlers[code] != nil {
		logrus.Fatalf("JobHandlers:%v already existed", code)
	}
	d.JobHandlers[code] = funk
}

// GetJobHandler defined
func (d *DefaultWorker) GetJobHandler(code string) func(model.Worker) (interface{}, error) {
	return func(worker model.Worker) (interface{}, error) {
		if d.JobHandlers[code] != nil {
			return d.JobHandlers[code](worker)
		}
		return nil, fmt.Errorf("not found! JobHandler:%v", code)
	}
}

// DefaultManager defined
type DefaultManager struct {
	BusinessDBSet map[string]*xorm.Engine
	MSeti         MSeti
	worker        Worker
}

// MSet defined
func (d *DefaultManager) MSet() MSeti {
	return d.MSeti
}

// GetBusinessDB defined
func (d *DefaultManager) GetBusinessDB(domain string) *xorm.Engine {
	return d.BusinessDBSet[domain]
}

// GetBusinessDBSet defined
func (d *DefaultManager) GetBusinessDBSet() map[string]*xorm.Engine {
	return d.BusinessDBSet
}

// AddBusinessDB defined
func (d *DefaultManager) AddBusinessDB(domain string, db *xorm.Engine) {
	if d.BusinessDBSet[domain] != nil {
		panic(fmt.Errorf("domain %v has exited", domain))
	}
	d.BusinessDBSet[domain] = db
}

// GetTokenStore defined
func (d *DefaultManager) GetTokenStore() oauth2.TokenStore {
	if RedisClient != nil {
		return store.NewRedisStoreWithCli(RedisClient, TokenkeyNamespace)
	}
	memo, _ := store.NewMemoryTokenStore()
	return memo
}

// Worker defined
func (d *DefaultManager) Worker() Worker {
	return d.worker
}

// NewDefaultManager defined
func NewDefaultManager() Manager {
	mg := &DefaultManager{}
	mg.BusinessDBSet = map[string]*xorm.Engine{}
	mg.MSeti = &MSet{m: map[string][]interface{}{}}

	dispatcher := worker.NewDispatcher(MaxWorkers)
	dispatcher.Run()

	woker := &DefaultWorker{store: map[string]model.Worker{}, JobHandlers: map[string]func(model.Worker) (interface{}, error){}}
	woker.Dispatcher = dispatcher
	mg.worker = woker
	return mg
}

// Cache middles
func Cache(time time.Duration) func(ctx *Context) {
	return func(ctx *Context) {
		cache.CachePage(CacheStore, time)(ctx.Context)
	}
}

// MaxWorkers defined
var MaxWorkers = 15

// RedisClient defined
var RedisClient *redis.Client

// CacheStore defined
var CacheStore persistence.CacheStore

func init() {
	if uri := util.EnsureLeft(http.Parse(viper.GetString("rd.dataSource"))).(*http.URI); uri.Laddr != "" {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     uri.Laddr,
			Password: uri.Passwd,
			DB:       util.EnsureLeft(strconv.Atoi(uri.DbName)).(int),
		})
		if _, err := RedisClient.Ping().Result(); err != nil {
			logrus.Warnf("Redis:%v connect failed", viper.GetString("rd.dataSource"))
			RedisClient = nil
		} else {
			logrus.Infof("Redis:%v connect successfully", viper.GetString("rd.dataSource"))
		}
	}
	CacheStore = persistence.NewInMemoryStore(60 * time.Second)
}
