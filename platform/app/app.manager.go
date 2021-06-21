// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/store"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/http"
	"github.com/2637309949/dolphin/platform/util/worker"
	"github.com/gin-contrib/cache/persistence"
	"github.com/go-errors/errors"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

// Cron defined
type Cron interface {
	AddFunc(string, func()) (int, error)
	RefreshFunc(int, string) (int, error)
	DelFunc(int) error
	TryFunc(int) error
}

// Manager Engine management interface
type Manager interface {
	MSet() MSeti
	Worker() Worker
	Cron() Cron
	GetBusinessDBSet() map[string]*xorm.Engine
	GetBusinessDB(string) *xorm.Engine
	AddBusinessDB(string, *xorm.Engine)
	GetTokenStore() oauth2.TokenStore
}

// DefaultCron defined
type DefaultCron struct {
	cron *cron.Cron
}

// AddFunc defined
func (d *DefaultCron) AddFunc(spec string, cmd func()) (int, error) {
	entry, err := d.cron.AddFunc(spec, cmd)
	return int(entry), err
}

// RefreshFunc defined
func (d *DefaultCron) RefreshFunc(id int, spec string) (int, error) {
	var entry cron.EntryID
	s, err := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor).Parse(spec)
	if err != nil {
		return 0, err
	}
	if !d.cron.Entry(cron.EntryID(id)).Valid() {
		return 0, errors.Errorf("invalid entry#%v", id)
	}
	job := d.cron.Entry(cron.EntryID(id)).Job
	entry = d.cron.Schedule(s, job)
	d.cron.Remove(cron.EntryID(id))
	return int(entry), nil
}

// DelFunc defined
func (d *DefaultCron) DelFunc(id int) error {
	d.cron.Remove(cron.EntryID(id))
	return nil
}

// TryFunc defined
func (d *DefaultCron) TryFunc(id int) error {
	entry := d.cron.Entry(cron.EntryID(id))
	if !entry.Valid() {
		return fmt.Errorf("invalid id#%v", id)
	}
	entry.Job.Run()
	return nil
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
	cron          Cron
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
	ots, _ := store.NewMemoryTokenStore()
	if RedisClient != nil {
		ots = store.NewRedisStoreWithCli(RedisClient, TokenkeyNamespace)
	}
	return ots
}

// Worker defined
func (d *DefaultManager) Worker() Worker {
	return d.worker
}

// Cron defined
func (d *DefaultManager) Cron() Cron {
	return d.cron
}

// NewDefaultManager defined
func NewDefaultManager() Manager {
	mg := &DefaultManager{}
	mg.BusinessDBSet = map[string]*xorm.Engine{}
	mg.MSeti = &MSet{m: map[string][]interface{}{}, lock: new(sync.RWMutex)}

	dispatcher := worker.NewDispatcher(MaxWorkers)
	dispatcher.Run()

	woker := &DefaultWorker{store: map[string]model.Worker{}, JobHandlers: map[string]func(model.Worker) (interface{}, error){}}
	woker.Dispatcher = dispatcher
	corn := cron.New(cron.WithSeconds())
	if viper.GetBool("app.cron") {
		corn.Start()
	}
	mg.worker = woker
	mg.cron = &DefaultCron{corn}
	return mg
}

// MaxWorkers defined
var MaxWorkers = 15

// RedisClient defined
var RedisClient *redis.Client

// CacheStore defined
var CacheStore persistence.CacheStore = persistence.NewInMemoryStore(60 * time.Second)

func init() {
	uri := util.EnsureLeft(http.Parse(viper.GetString("rd.dataSource"))).(*http.URI)
	if uri.Laddr != "" {
		db := util.EnsureLeft(strconv.Atoi(uri.DbName)).(int)
		opts := redis.Options{Addr: uri.Laddr, Password: uri.Passwd, DB: db}
		RedisClient = redis.NewClient(&opts)
		if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
			logrus.Warnf("Redis:%v connect failed", viper.GetString("rd.dataSource"))
			RedisClient = nil
		} else {
			CacheStore = NewRedisCache(RedisClient, 60*time.Second)
			logrus.Infof("Redis:%v connect successfully", viper.GetString("rd.dataSource"))
		}
	}
}
