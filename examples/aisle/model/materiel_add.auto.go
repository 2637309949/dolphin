// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MaterielAdd defined
type MaterielAdd struct {
	// MaterielId defined
	MaterielId null.Int `xorm:"int(11) 'materiel_id'" json:"materiel_id" form:"materiel_id" xml:"materiel_id"`
	// AddNum defined
	AddNum null.Int `xorm:"int(11) 'add_num'" json:"add_num" form:"add_num" xml:"add_num"`
	// AddMoney defined
	AddMoney null.Float `xorm:"float(11,2) 'add_money'" json:"add_money" form:"add_money" xml:"add_money"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// MAId defined
	MAId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_id'" json:"m_a_id" form:"m_a_id" xml:"m_a_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// AddReason defined
	AddReason null.Int `xorm:"int(11) 'add_reason'" json:"add_reason" form:"add_reason" xml:"add_reason"`
	// AddObject defined
	AddObject null.Int `xorm:"int(11) 'add_object'" json:"add_object" form:"add_object" xml:"add_object"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined MaterielAdd
func (m *MaterielAdd) TableName() string {
	return "materiel_add"
}
