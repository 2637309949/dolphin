// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// StudentCustomer defined
type StudentCustomer struct {
	// SCId defined
	SCId null.Int `xorm:"int(11) pk notnull autoincr 's_c_id'" json:"s_c_id" form:"s_c_id" xml:"s_c_id"`
	// Customer defined
	Customer null.Int `xorm:"int(11) 'customer'" json:"customer" form:"customer" xml:"customer"`
	// ScType defined
	ScType null.Int `xorm:"int(11) 'sc_type'" json:"sc_type" form:"sc_type" xml:"sc_type"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// StuTypeId defined
	StuTypeId null.Int `xorm:"int(11) 'stu_type_id'" json:"stu_type_id" form:"stu_type_id" xml:"stu_type_id"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// ScDate defined
	ScDate null.Time `xorm:"datetime 'sc_date'" json:"sc_date" form:"sc_date" xml:"sc_date"`
	// HistoryKf defined
	HistoryKf null.Int `xorm:"int(11) 'history_kf'" json:"history_kf" form:"history_kf" xml:"history_kf"`
	// CancelDate defined
	CancelDate null.Time `xorm:"datetime 'cancel_date'" json:"cancel_date" form:"cancel_date" xml:"cancel_date"`
	// Remark defined
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// IfPl defined
	IfPl null.Int `xorm:"int(11) 'if_pl'" json:"if_pl" form:"if_pl" xml:"if_pl"`
}

// Parser defined
func (m *StudentCustomer) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StudentCustomer) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentCustomer
func (m *StudentCustomer) TableName() string {
	return "student_customer"
}
