package svc

import (
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/types"
)

// Db defined TODO
type Db interface {
	PageSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (*types.PageList, error)
	TreeSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...svc.Formatter) (interface{}, error)
	GetOptions(*xorm.Engine, ...string) (map[string]map[string]interface{}, error)
}
