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

// StuclasstypeExtrafee defined
type StuclasstypeExtrafee struct {
	// SEId defined
	SEId null.Int `xorm:"int(11) pk notnull autoincr 's_e_id'" json:"s_e_id" form:"s_e_id" xml:"s_e_id"`
	// StuClassTypeId defined
	StuClassTypeId null.Int `xorm:"int(11) 'stu_class_type_id'" json:"stu_class_type_id" form:"stu_class_type_id" xml:"stu_class_type_id"`
	// ExtraFee defined
	ExtraFee null.Float `xorm:"float(10,2) 'extra_fee'" json:"extra_fee" form:"extra_fee" xml:"extra_fee"`
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
	// OrderExtraId defined
	OrderExtraId null.Int `xorm:"int(11) 'order_extra_id'" json:"order_extra_id" form:"order_extra_id" xml:"order_extra_id"`
}

// Parser defined
func (m *StuclasstypeExtrafee) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuclasstypeExtrafee) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuclasstypeExtrafee
func (m *StuclasstypeExtrafee) TableName() string {
	return "stuclasstype_extrafee"
}
