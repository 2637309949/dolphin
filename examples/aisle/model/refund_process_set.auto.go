// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// RefundProcessSet defined
type RefundProcessSet struct {
	//
	RPSId null.Int `xorm:"int(11) pk notnull autoincr 'r_p_s_id'" json:"r_p_s_id" xml:"r_p_s_id"`
	//
	RpsName null.String `xorm:"varchar(1000) 'rps_name'" json:"rps_name" xml:"rps_name"`
	//
	RpsDesc null.String `xorm:"varchar(1000) 'rps_desc'" json:"rps_desc" xml:"rps_desc"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined RefundProcessSet
func (m *RefundProcessSet) TableName() string {
	return "refund_process_set"
}