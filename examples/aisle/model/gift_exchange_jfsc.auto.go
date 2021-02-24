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

// GiftExchangeJfsc defined
type GiftExchangeJfsc struct {
	//
	GEJId null.Int `xorm:"int(11) pk notnull autoincr 'g_e_j_id'" json:"g_e_j_id" form:"g_e_j_id" xml:"g_e_j_id"`
	//
	ExchangeNum null.Float `xorm:"float(11,2) 'exchange_num'" json:"exchange_num" form:"exchange_num" xml:"exchange_num"`
	//
	XhJf null.Float `xorm:"float(11,2) 'xh_jf'" json:"xh_jf" form:"xh_jf" xml:"xh_jf"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	GiftId null.Int `xorm:"int(11) 'gift_id'" json:"gift_id" form:"gift_id" xml:"gift_id"`
	//
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" form:"pk_stu" xml:"pk_stu"`
}

// Parser defined
func (m *GiftExchangeJfsc) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *GiftExchangeJfsc) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined GiftExchangeJfsc
func (m *GiftExchangeJfsc) TableName() string {
	return "gift_exchange_jfsc"
}
