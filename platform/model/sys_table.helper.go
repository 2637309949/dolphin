package model

import (
	"fmt"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
	"github.com/2637309949/dolphin/platform/util"
)

// TruncateTable defined
func (m *SysTable) TruncateTable(session *xorm.Session, driverName string) {
	if driverName != "sqlite3" {
		util.EnsureLeft(session.Exec(fmt.Sprintf("truncate table %v", new(SysTable).TableName())))
	} else {
		util.EnsureLeft(session.Exec(fmt.Sprintf("delete from %v", new(SysTable).TableName())))
	}
}

// TableInfo defined inital system data
func (m *SysTable) TableInfo(info *schemas.Table) SysTable {
	return SysTable{
		Name:          null.StringFrom(info.Name),
		Desc:          null.StringFrom(info.Comment),
		Charset:       null.StringFrom(info.Charset),
		AutoIncrement: null.StringFrom(info.AutoIncrement),
		StoreEngine:   null.StringFrom(info.StoreEngine),
		CreateTime:    null.TimeFrom(time.Now()),
		Creater:       null.IntFrom(0),
		UpdateTime:    null.TimeFrom(time.Now()),
		Updater:       null.IntFrom(0),
		IsDelete:      null.IntFrom(0),
	}
}
