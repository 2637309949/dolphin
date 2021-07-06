// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MarketExpend defined
type MarketExpend struct {
	// MEId defined
	MEId null.Int `xorm:"int(11) pk notnull autoincr 'm_e_id'" json:"m_e_id" form:"m_e_id" xml:"m_e_id"`
	// MarketId defined
	MarketId null.Int `xorm:"int(11) 'market_id'" json:"market_id" form:"market_id" xml:"market_id"`
	// ExpendMoney defined
	ExpendMoney null.Float `xorm:"float(11,2) 'expend_money'" json:"expend_money" form:"expend_money" xml:"expend_money"`
	// Remake defined
	Remake null.String `xorm:"varchar(2000) 'remake'" json:"remake" form:"remake" xml:"remake"`
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
	// ExpendType defined
	ExpendType null.Int `xorm:"int(11) 'expend_type'" json:"expend_type" form:"expend_type" xml:"expend_type"`
}

// TableName table name of defined MarketExpend
func (m *MarketExpend) TableName() string {
	return "market_expend"
}