// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// RefundProcessSet defined
type RefundProcessSet struct {
	// RPSId defined
	RPSId null.Int `xorm:"int(11) pk notnull autoincr 'r_p_s_id'" json:"r_p_s_id" form:"r_p_s_id" xml:"r_p_s_id"`
	// RpsName defined
	RpsName null.String `xorm:"varchar(1000) 'rps_name'" json:"rps_name" form:"rps_name" xml:"rps_name"`
	// RpsDesc defined
	RpsDesc null.String `xorm:"varchar(1000) 'rps_desc'" json:"rps_desc" form:"rps_desc" xml:"rps_desc"`
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

// TableName table name of defined RefundProcessSet
func (m *RefundProcessSet) TableName() string {
	return "refund_process_set"
}
