package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
)

// TableInfo defined inital system data
func (m *SysTable) TableInfo(info *schemas.Table) SysTable {
	return SysTable{
		ID:            null.StringFromUUID(),
		Name:          null.StringFrom(info.Name),
		Desc:          null.StringFrom(info.Comment),
		Charset:       null.StringFrom(info.Charset),
		AutoIncrement: null.StringFrom(info.AutoIncrement),
		StoreEngine:   null.StringFrom(info.StoreEngine),
		CreateTime:    null.TimeFrom(time.Now()),
		CreateBy:      DefaultAdmin.ID,
		UpdateTime:    null.TimeFrom(time.Now()),
		UpdateBy:      DefaultAdmin.ID,
		DelFlag:       null.IntFrom(0),
	}
}
