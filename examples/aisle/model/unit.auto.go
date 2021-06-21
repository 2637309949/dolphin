// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Unit defined
type Unit struct {
	// UnitId defined
	UnitId null.Int `xorm:"int(11) pk notnull autoincr 'unit_id'" json:"unit_id" form:"unit_id" xml:"unit_id"`
	// UnitName defined
	UnitName null.String `xorm:"varchar(1000) 'unit_name'" json:"unit_name" form:"unit_name" xml:"unit_name"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"notnull 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined Unit
func (m *Unit) TableName() string {
	return "unit"
}
