package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
)

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
