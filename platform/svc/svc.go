package svc

import (
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/http"
	"github.com/gin-gonic/gin"
)

type (
	// Cache defined TODO
	Cache interface {
		SetCache(key string, v interface{}, expire time.Duration) error
		GetCache(key string, v interface{}) error
	}

	// Export defined TODO
	Export interface {
		Check(*gin.Context) bool
		PageExport(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (*types.ExportInfo, error)
		SetOptionsetsFormat(func(interface{}) func(interface{}) interface{})
	}

	// Db defined TODO
	Db interface {
		PageSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (*types.PageList, error)
		TreeSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (interface{}, error)
		GetOptions(*xorm.Engine, ...string) (map[string]map[string]interface{}, error)
		InRole(*xorm.Engine, string, ...string) bool
		InAdmin(*xorm.Engine, string, ...string) bool
		Persist(db *xorm.Session, ids ...string) (int64, error)
		PersistFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error)
		Remove(db *xorm.Session, ids ...string) (int64, error)
		RemoveFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error)
	}
	// Formatter defined TODO
	Formatter func(*xorm.Engine, []map[string]interface{}) ([]map[string]interface{}, error)
	// Http defined TODO
	Http interface {
		// Get returns *HttpRequest with GET method.
		Get(url string) *http.HttpRequest
		// Post returns *HttpRequest with POST method.
		Post(url string) *http.HttpRequest
		// Put returns *HttpRequest with PUT method.
		Put(url string) *http.HttpRequest
		// Delete returns *HttpRequest DELETE method.
		Delete(url string) *http.HttpRequest
		// Head returns *HttpRequest with HEAD method.
		Head(url string) *http.HttpRequest
	}
)

// Svc defined TODO
type Svc interface {
	Export
	Cache
	Http
	Db
}
