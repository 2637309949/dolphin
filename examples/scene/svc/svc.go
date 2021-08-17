package svc

import (
	"io"
	"time"

	nh "net/http"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/http"
)

type (
	// Cache defined TODO
	Cache interface {
		SetCache(key string, v interface{}, expire time.Duration) error
		GetCache(key string, v interface{}) error
	}

	// Export defined TODO
	Export interface {
		Check(*nh.Request) bool
		PageExport(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (*types.ExportInfo, error)
		ParseExcel(r io.Reader, sheet interface{}, header ...[]map[string]interface{}) ([]map[string]string, error)
		SetOptionsetsFormat(func(interface{}) func(interface{}) interface{})
	}

	// Db defined TODO
	Db interface {
		PageSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (*types.PageList, error)
		TreeSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (interface{}, error)
		GetOptions(*xorm.Engine, ...string) (map[string]map[string]interface{}, error)
		InRole(*xorm.Engine, string, ...string) bool
		InAdmin(*xorm.Engine, string, ...string) bool
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
	Oss
	Weapp
	Kafka
}
