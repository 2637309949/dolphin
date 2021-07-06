// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentMarketActivity defined
type StudentMarketActivity struct {
	// SMAId defined
	SMAId null.Int `xorm:"int(11) pk notnull autoincr 's_m_a_id'" json:"s_m_a_id" form:"s_m_a_id" xml:"s_m_a_id"`
	// SutId defined
	SutId null.Int `xorm:"int(11) 'sut_id'" json:"sut_id" form:"sut_id" xml:"sut_id"`
	// MaId defined
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
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

// TableName table name of defined StudentMarketActivity
func (m *StudentMarketActivity) TableName() string {
	return "student_market_activity"
}