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
func (m *SysTableColumn) TruncateTable(session *xorm.Session, driverName string) {
	if driverName != "sqlite3" {
		util.EnsureLeft(session.Exec(fmt.Sprintf("truncate table %v", new(SysTableColumn).TableName())))
	} else {
		util.EnsureLeft(session.Exec(fmt.Sprintf("delete from %v", new(SysTableColumn).TableName())))
	}
}

// ColumnInfo defined inital system data
func (m *SysTableColumn) ColumnInfo(info *schemas.Column) SysTableColumn {
	return SysTableColumn{
		ID:           null.StringFromUUID(),
		Name:         null.StringFrom(info.Name),
		Type:         null.StringFrom(info.SQLType.Name),
		Desc:         null.StringFrom(info.Comment),
		IsPrimaryKey: null.BoolFrom(info.IsPrimaryKey),
		Nullable:     null.BoolFrom(info.Nullable),
		Default:      null.StringFrom(info.Default),
		CreateTime:   null.TimeFrom(time.Now()),
		Creater:      DefaultAdmin.ID,
		UpdateTime:   null.TimeFrom(time.Now()),
		Updater:      DefaultAdmin.ID,
		IsDelete:     null.IntFrom(0),
	}
}
