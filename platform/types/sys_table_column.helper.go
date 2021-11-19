package types

import (
	"fmt"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
)

// TruncateTable defined
func (m *SysTableColumn) TruncateTable(session *xorm.Session) error {
	switch session.Engine().DriverName() {
	case "sqlite3":
		_, err := session.Exec(fmt.Sprintf("delete from %v", new(SysTableColumn).TableName()))
		if err != nil {
			return err
		}
	default:
		_, err := session.Exec(fmt.Sprintf("truncate table %v", new(SysTableColumn).TableName()))
		if err != nil {
			return err
		}
	}
	return nil
}

// ColumnInfo defined inital system data
func (m *SysTableColumn) ColumnInfo(info *schemas.Column) SysTableColumn {
	return SysTableColumn{
		Name:         null.StringFrom(info.Name),
		Type:         null.StringFrom(info.SQLType.Name),
		Desc:         null.StringFrom(info.Comment),
		IsPrimaryKey: null.BoolFrom(info.IsPrimaryKey),
		Nullable:     null.BoolFrom(info.Nullable),
		Default:      null.StringFrom(info.Default),
		CreateTime:   null.TimeFromNow(),
		Creater:      null.IntFrom(0),
		UpdateTime:   null.TimeFromNow(),
		Updater:      null.IntFrom(0),
		IsDelete:     null.IntFrom(0),
	}
}
