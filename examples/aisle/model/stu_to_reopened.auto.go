// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuToReopened defined
type StuToReopened struct {
	// STRId defined
	STRId null.Int `xorm:"int(11) pk notnull autoincr 's_t_r_id'" json:"s_t_r_id" form:"s_t_r_id" xml:"s_t_r_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// StrRemark defined
	StrRemark null.String `xorm:"varchar(2000) 'str_remark'" json:"str_remark" form:"str_remark" xml:"str_remark"`
	// StrDate defined
	StrDate null.Time `xorm:"datetime 'str_date'" json:"str_date" form:"str_date" xml:"str_date"`
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
	// StrBussType defined
	StrBussType null.Int `xorm:"int(11) 'str_buss_type'" json:"str_buss_type" form:"str_buss_type" xml:"str_buss_type"`
}

// TableName table name of defined StuToReopened
func (m *StuToReopened) TableName() string {
	return "stu_to_reopened"
}
