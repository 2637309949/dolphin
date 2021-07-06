// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MaterialApply defined
type MaterialApply struct {
	// MAId defined
	MAId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_id'" json:"m_a_id" form:"m_a_id" xml:"m_a_id"`
	// MtId defined
	MtId null.Int `xorm:"int(11) 'mt_id'" json:"mt_id" form:"mt_id" xml:"mt_id"`
	// MaNum defined
	MaNum null.Float `xorm:"float(10,2) 'ma_num'" json:"ma_num" form:"ma_num" xml:"ma_num"`
	// MaMoney defined
	MaMoney null.Float `xorm:"float(10,2) 'ma_money'" json:"ma_money" form:"ma_money" xml:"ma_money"`
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
	// MbId defined
	MbId null.Int `xorm:"int(11) 'mb_id'" json:"mb_id" form:"mb_id" xml:"mb_id"`
}

// TableName table name of defined MaterialApply
func (m *MaterialApply) TableName() string {
	return "material_apply"
}