// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuclasstypeExtrafee defined
type StuclasstypeExtrafee struct {
	//
	SEId null.Int `xorm:"int(11) pk notnull autoincr 's_e_id'" json:"s_e_id" xml:"s_e_id"`
	//
	StuClassTypeId null.Int `xorm:"int(11) 'stu_class_type_id'" json:"stu_class_type_id" xml:"stu_class_type_id"`
	//
	ExtraFee null.Float `xorm:"float(10,2) 'extra_fee'" json:"extra_fee" xml:"extra_fee"`
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
	//
	OrderExtraId null.Int `xorm:"int(11) 'order_extra_id'" json:"order_extra_id" xml:"order_extra_id"`
}

// TableName table name of defined StuclasstypeExtrafee
func (m *StuclasstypeExtrafee) TableName() string {
	return "stuclasstype_extrafee"
}