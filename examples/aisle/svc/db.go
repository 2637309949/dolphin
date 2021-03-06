package svc

import (
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
)

// Db defined TODO
type Db interface {
	PageSearch(db *xorm.Engine, ctr, api, table string, params map[string]interface{}) (*types.PageList, error)
	TreeSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error)
	GetOptions(db *xorm.Engine, keys ...string) (map[string]map[string]interface{}, error)
}
