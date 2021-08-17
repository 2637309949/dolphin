// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"sync"

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/store"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/worker"
	"github.com/go-errors/errors"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// MaxWorkers defined TODO
var MaxWorkers = 15

// Worker defined TODO
type Worker interface {
	AddJob(j worker.Job)
	AddJobHandler(string, func(types.Worker) (interface{}, error))
	GetJobHandler(string) func(types.Worker) (interface{}, error)
	Save(types.Worker) error
	Update(types.Worker) error
	Find(string) (types.Worker, error)
}

// Cron defined TODO
type Cron interface {
	AddCron(string, func()) (int, error)
	RefreshCron(int, string) (int, error)
	DelCron(int) error
	TryCron(int) error
}

// Manager Engine management interface
type Manager interface {
	ModelSet() ModelSetter
	Worker() Worker
	Cron() Cron
	GetBusinessDBSet() map[string]*xorm.Engine
	GetBusinessDB(string) *xorm.Engine
	AddBusinessDB(string, *xorm.Engine)
	GetTokenStore() oauth2.TokenStore
}

// DefaultCron defined TODO
type DefaultCron struct {
	cron *cron.Cron
}

// AddCron defined TODO
func (d *DefaultCron) AddCron(spec string, cmd func()) (int, error) {
	entry, err := d.cron.AddFunc(spec, cmd)
	return int(entry), err
}

// RefreshCron defined TODO
func (d *DefaultCron) RefreshCron(id int, spec string) (int, error) {
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

// DelCron defined TODO
func (d *DefaultCron) DelCron(id int) error {
	d.cron.Remove(cron.EntryID(id))
	return nil
}

// TryCron defined TODO
func (d *DefaultCron) TryCron(id int) error {
	entry := d.cron.Entry(cron.EntryID(id))
	if !entry.Valid() {
		return fmt.Errorf("invalid id#%v", id)
	}
	entry.Job.Run()
	return nil
}

// DefaultWorker defined TODO
type DefaultWorker struct {
	store       map[string]types.Worker
	JobHandlers map[string]func(types.Worker) (interface{}, error)
	Dispatcher  *worker.Dispatcher
}

// AddJob defined TODO
func (d *DefaultWorker) AddJob(j worker.Job) {
	d.Dispatcher.AddJob(j)
}

// Find defined TODO
func (d *DefaultWorker) Find(code string) (types.Worker, error) {
	return d.store[code], nil
}

// Save defined TODO
func (d *DefaultWorker) Save(w types.Worker) error {
	d.store[w.Code] = w
	return nil
}

// Update defined TODO
func (d *DefaultWorker) Update(w types.Worker) error {
	d.store[w.Code] = w
	return nil
}

// AddJobHandler defined TODO
func (d *DefaultWorker) AddJobHandler(code string, funk func(types.Worker) (interface{}, error)) {
	if d.JobHandlers[code] != nil {
		logrus.Fatalf("JobHandlers:%v already existed", code)
	}
	d.JobHandlers[code] = funk
}

// GetJobHandler defined TODO
func (d *DefaultWorker) GetJobHandler(code string) func(types.Worker) (interface{}, error) {
	return func(worker types.Worker) (interface{}, error) {
		if d.JobHandlers[code] != nil {
			return d.JobHandlers[code](worker)
		}
		return nil, fmt.Errorf("not found! JobHandler:%v", code)
	}
}

// DefaultManager defined TODO
type DefaultManager struct {
	BusinessDBSet map[string]*xorm.Engine
	ModelSetter   ModelSetter
	worker        Worker
	cron          Cron
}

// ModelSet defined TODO
func (d *DefaultManager) ModelSet() ModelSetter {
	return d.ModelSetter
}

// GetBusinessDB defined TODO
func (d *DefaultManager) GetBusinessDB(domain string) *xorm.Engine {
	return d.BusinessDBSet[domain]
}

// GetBusinessDBSet defined TODO
func (d *DefaultManager) GetBusinessDBSet() map[string]*xorm.Engine {
	return d.BusinessDBSet
}

// AddBusinessDB defined TODO
func (d *DefaultManager) AddBusinessDB(domain string, db *xorm.Engine) {
	if d.BusinessDBSet[domain] != nil {
		panic(fmt.Errorf("domain %v has exited", domain))
	}
	d.BusinessDBSet[domain] = db
}

// GetTokenStore defined TODO
func (d *DefaultManager) GetTokenStore() oauth2.TokenStore {
	ots, _ := store.NewMemoryTokenStore()
	if RedisClient != nil {
		ots = store.NewRedisStoreWithCli(RedisClient, TokenkeyNamespace)
	}
	return ots
}

// Worker defined TODO
func (d *DefaultManager) Worker() Worker {
	return d.worker
}

// Cron defined TODO
func (d *DefaultManager) Cron() Cron {
	return d.cron
}

// NewDefaultManager defined TODO
func NewDefaultManager() Manager {
	mg := &DefaultManager{}
	mg.BusinessDBSet = map[string]*xorm.Engine{}
	mg.ModelSetter = &defaultModelSetter{m: map[string][]Table{}, lock: new(sync.RWMutex)}

	dispatcher := worker.NewDispatcher(MaxWorkers)
	dispatcher.Run()

	woker := &DefaultWorker{store: map[string]types.Worker{}, JobHandlers: map[string]func(types.Worker) (interface{}, error){}}
	woker.Dispatcher = dispatcher
	corn := cron.New(cron.WithSeconds())
	if viper.GetBool("app.cron") {
		corn.Start()
	}
	mg.worker = woker
	mg.cron = &DefaultCron{corn}
	return mg
}
