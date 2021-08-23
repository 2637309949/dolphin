package svc

import (
	"io"
	"net/http"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/http/client"
	"github.com/gin-contrib/cache/persistence"
)

type (
	// Store defined TODO
	Store interface {
		Get(key string, value interface{}) error
		Set(key string, value interface{}, expire time.Duration) error
		Add(key string, value interface{}, expire time.Duration) error
		Replace(key string, data interface{}, expire time.Duration) error
		Delete(key string) error
		Increment(key string, data uint64) (uint64, error)
		Decrement(key string, data uint64) (uint64, error)
		Flush() error
	}

	// Report defined TODO
	Report interface {
		Check(*http.Request) bool
		PageExport(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (*types.ExportInfo, error)
		ParseExcel(r io.Reader, sheet interface{}, header ...[]map[string]interface{}) ([]map[string]string, error)
		SetOptionsetsFormat(func(interface{}) func(interface{}) interface{})
	}

	// XReport defined TODO
	XReport struct {
		xlsx *svc.Xlsx
	}

	// DB defined TODO
	DB interface {
		PageSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (*types.PageList, error)
		TreeSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (interface{}, error)
		GetOptions(*xorm.Engine, ...string) (map[string]map[string]interface{}, error)
		InRole(*xorm.Engine, string, ...string) bool
		InAdmin(*xorm.Engine, string, ...string) bool
		Persist(db *xorm.Session, ids ...string) (int64, error)
		PersistFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error)
		Remove(db *xorm.Session, ids ...string) (int64, error)
		RemoveFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error)
	}

	// XDB defined TODO
	XDB struct{}

	// XClient defined TODO
	Client interface {
		// Get returns *HttpRequest with GET method.
		Get(url string) *client.HttpRequest
		// Post returns *HttpRequest with POST method.
		Post(url string) *client.HttpRequest
		// Put returns *HttpRequest with PUT method.
		Put(url string) *client.HttpRequest
		// Delete returns *HttpRequest DELETE method.
		Delete(url string) *client.HttpRequest
		// Head returns *HttpRequest with HEAD method.
		Head(url string) *client.HttpRequest
	}
	// ServiceContext defined TODO
	ServiceContext struct {
		Report Report
		Store  Store
		Client Client
		DB     DB
		Oss    Oss
		Weapp  Weapp
		Kafka  Kafka
	}
)

// NewServiceContext defined TODO
func NewServiceContext(cache persistence.CacheStore) *ServiceContext {
	svc := svc.NewServiceContext(cache)
	return &ServiceContext{
		Store:  svc.Store,
		Report: svc.Report,
		Client: svc.Client,
		DB:     svc.DB,
		Oss:    NewXOss(),
		Weapp:  NewXWeapp(),
		Kafka:  NewXKafka(),
	}
}
