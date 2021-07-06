// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// GiftExchange defined
type GiftExchange struct {
	// GEId defined
	GEId null.Int `xorm:"int(11) pk notnull autoincr 'g_e_id'" json:"g_e_id" form:"g_e_id" xml:"g_e_id"`
	// ExpressNumber defined
	ExpressNumber null.String `xorm:"varchar(50) 'express_number'" json:"express_number" form:"express_number" xml:"express_number"`
	// StuAddress defined
	StuAddress null.String `xorm:"varchar(100) 'stu_address'" json:"stu_address" form:"stu_address" xml:"stu_address"`
	// ReturnReason defined
	ReturnReason null.String `xorm:"varchar(200) 'return_reason'" json:"return_reason" form:"return_reason" xml:"return_reason"`
	// GiftType defined
	GiftType null.Int `xorm:"int(11) 'gift_type'" json:"gift_type" form:"gift_type" xml:"gift_type"`
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
	// StuGiftId defined
	StuGiftId null.Int `xorm:"int(11) 'stu_gift_id'" json:"stu_gift_id" form:"stu_gift_id" xml:"stu_gift_id"`
}

// TableName table name of defined GiftExchange
func (m *GiftExchange) TableName() string {
	return "gift_exchange"
}