// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// StudentGift defined
type StudentGift struct {
	// SGId defined
	SGId null.Int `xorm:"int(11) pk notnull autoincr 's_g_id'" json:"s_g_id" form:"s_g_id" xml:"s_g_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// GiftId defined
	GiftId null.Int `xorm:"int(11) 'gift_id'" json:"gift_id" form:"gift_id" xml:"gift_id"`
	// BuyWay defined
	BuyWay null.Int `xorm:"int(11) 'buy_way'" json:"buy_way" form:"buy_way" xml:"buy_way"`
	// ResourcePrice defined
	ResourcePrice null.Float `xorm:"float(11,2) 'resource_price'" json:"resource_price" form:"resource_price" xml:"resource_price"`
	// BuyPrice defined
	BuyPrice null.Float `xorm:"float(11,2) 'buy_price'" json:"buy_price" form:"buy_price" xml:"buy_price"`
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
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// RefundMoney defined
	RefundMoney null.Float `xorm:"float(10,2) 'refund_money'" json:"refund_money" form:"refund_money" xml:"refund_money"`
	// RefundPrice defined
	RefundPrice null.Float `xorm:"float(10,2) 'refund_price'" json:"refund_price" form:"refund_price" xml:"refund_price"`
	// SgNum defined
	SgNum null.Float `xorm:"float(10,2) 'sg_num'" json:"sg_num" form:"sg_num" xml:"sg_num"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// GiftExchangeType defined
	GiftExchangeType null.Int `xorm:"int(11) 'gift_exchange_type'" json:"gift_exchange_type" form:"gift_exchange_type" xml:"gift_exchange_type"`
	// OggId defined
	OggId null.Int `xorm:"int(11) 'ogg_id'" json:"ogg_id" form:"ogg_id" xml:"ogg_id"`
	// UsePoint defined
	UsePoint null.Float `xorm:"float(10,2) 'use_point'" json:"use_point" form:"use_point" xml:"use_point"`
	// StuExGift defined
	StuExGift null.Int `xorm:"int(11) 'stu_ex_gift'" json:"stu_ex_gift" form:"stu_ex_gift" xml:"stu_ex_gift"`
	// ZsGiveMonth defined
	ZsGiveMonth null.Time `xorm:"datetime 'zs_give_month'" json:"zs_give_month" form:"zs_give_month" xml:"zs_give_month"`
}

// Marshal defined
func (m *StudentGift) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StudentGift) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *StudentGift) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StudentGift) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentGift
func (m *StudentGift) TableName() string {
	return "student_gift"
}
