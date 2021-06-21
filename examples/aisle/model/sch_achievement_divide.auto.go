// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SchAchievementDivide defined
type SchAchievementDivide struct {
	// SADId defined
	SADId null.Int `xorm:"int(11) pk notnull autoincr 's_a_d_id'" json:"s_a_d_id" form:"s_a_d_id" xml:"s_a_d_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// JxMoney defined
	JxMoney null.Float `xorm:"float(11,2) 'jx_money'" json:"jx_money" form:"jx_money" xml:"jx_money"`
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
	// Percent defined
	Percent null.Float `xorm:"float(50,2) 'percent'" json:"percent" form:"percent" xml:"percent"`
	// PkFee defined
	PkFee null.Int `xorm:"int(11) 'pk_fee'" json:"pk_fee" form:"pk_fee" xml:"pk_fee"`
}

// TableName table name of defined SchAchievementDivide
func (m *SchAchievementDivide) TableName() string {
	return "sch_achievement_divide"
}
