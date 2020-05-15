package app

import (
	"fmt"
	"strconv"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/store"
	"github.com/2637309949/dolphin/packages/redis"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/http"
)

// DefaultManager defined
type DefaultManager struct {
	BusinessDBSet map[string]*xorm.Engine
	MSeti         MSeti
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
	if uri := util.EnsureLeft(http.Parse(viper.GetString("rd.dataSource"))).(*http.URI); uri.Laddr != "" {
		redis := redis.NewClient(&redis.Options{
			Addr:     uri.Laddr,
			Password: uri.Passwd,
			DB:       util.EnsureLeft(strconv.Atoi(uri.DbName)).(int),
		})
		if _, err := redis.Ping().Result(); err == nil {
			return store.NewRedisStoreWithCli(redis, TokenkeyNamespace)
		}
		logrus.Warnf("GetTokenStore:%v connect failed", viper.GetString("rd.dataSource"))
	}
	memo, _ := store.NewMemoryTokenStore()
	return memo
}

// NewDefaultManager defined
func NewDefaultManager() Manager {
	mg := &DefaultManager{}
	mg.BusinessDBSet = map[string]*xorm.Engine{}
	mg.MSeti = &MSet{m: map[string][]interface{}{}}
	return mg
}
