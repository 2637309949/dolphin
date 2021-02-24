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

// GiftExchange defined
type GiftExchange struct {
	//
	GEId null.Int `xorm:"int(11) pk notnull autoincr 'g_e_id'" json:"g_e_id" form:"g_e_id" xml:"g_e_id"`
	//
	ExpressNumber null.String `xorm:"varchar(50) 'express_number'" json:"express_number" form:"express_number" xml:"express_number"`
	//
	StuAddress null.String `xorm:"varchar(100) 'stu_address'" json:"stu_address" form:"stu_address" xml:"stu_address"`
	//
	ReturnReason null.String `xorm:"varchar(200) 'return_reason'" json:"return_reason" form:"return_reason" xml:"return_reason"`
	//
	GiftType null.Int `xorm:"int(11) 'gift_type'" json:"gift_type" form:"gift_type" xml:"gift_type"`
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
	StuGiftId null.Int `xorm:"int(11) 'stu_gift_id'" json:"stu_gift_id" form:"stu_gift_id" xml:"stu_gift_id"`
}

// Parser defined
func (m *GiftExchange) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *GiftExchange) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined GiftExchange
func (m *GiftExchange) TableName() string {
	return "gift_exchange"
}
