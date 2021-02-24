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

// Gift defined
type Gift struct {
	// GiftId defined
	GiftId null.Int `xorm:"int(11) pk notnull autoincr 'gift_id'" json:"gift_id" form:"gift_id" xml:"gift_id"`
	// GiftName defined
	GiftName null.String `xorm:"varchar(100) 'gift_name'" json:"gift_name" form:"gift_name" xml:"gift_name"`
	// GiftNum defined
	GiftNum null.Int `xorm:"int(11) 'gift_num'" json:"gift_num" form:"gift_num" xml:"gift_num"`
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
	// Price defined
	Price null.Float `xorm:"float(50,2) 'price'" json:"price" form:"price" xml:"price"`
	// NeedIntegral defined
	NeedIntegral null.Float `xorm:"float(50,2) 'need_integral'" json:"need_integral" form:"need_integral" xml:"need_integral"`
	// GiftType defined
	GiftType null.Int `xorm:"int(11) 'gift_type'" json:"gift_type" form:"gift_type" xml:"gift_type"`
	// GiftMonth defined
	GiftMonth null.Time `xorm:"datetime 'gift_month'" json:"gift_month" form:"gift_month" xml:"gift_month"`
	// GiftEndMonth defined
	GiftEndMonth null.Time `xorm:"datetime 'gift_end_month'" json:"gift_end_month" form:"gift_end_month" xml:"gift_end_month"`
	// OrderType defined
	OrderType null.Int `xorm:"int(11) 'order_type'" json:"order_type" form:"order_type" xml:"order_type"`
	// QyState defined
	QyState null.Int `xorm:"int(11) 'qy_state'" json:"qy_state" form:"qy_state" xml:"qy_state"`
	// GiftFiled defined
	GiftFiled null.Int `xorm:"int(11) 'gift_filed'" json:"gift_filed" form:"gift_filed" xml:"gift_filed"`
	// PubCompany defined
	PubCompany null.String `xorm:"varchar(200) 'pub_company'" json:"pub_company" form:"pub_company" xml:"pub_company"`
	// Introduction defined
	Introduction null.String `xorm:"varchar(2000) 'introduction'" json:"introduction" form:"introduction" xml:"introduction"`
	// Conditions defined
	Conditions null.String `xorm:"varchar(2000) 'conditions'" json:"conditions" form:"conditions" xml:"conditions"`
	// Introduction2 defined
	Introduction2 null.String `xorm:"varchar(100) 'introduction2'" json:"introduction2" form:"introduction2" xml:"introduction2"`
}

// Parser defined
func (m *Gift) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *Gift) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined Gift
func (m *Gift) TableName() string {
	return "gift"
}
