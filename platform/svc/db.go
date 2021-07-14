package svc

import (
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
)

// Db defined TODO
type Db interface {
	PageSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (*types.PageList, error)
	TreeSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (interface{}, error)
	GetOptions(*xorm.Engine, ...string) (map[string]map[string]interface{}, error)
	InRole(*xorm.Engine, string, ...string) bool
	InAdmin(*xorm.Engine, string, ...string) bool
}

// Formatter defined TODO
type Formatter func(*xorm.Engine, []map[string]interface{}) ([]map[string]interface{}, error)
