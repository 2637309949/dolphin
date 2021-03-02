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

// MarketFeeCost defined
type MarketFeeCost struct {
	// MFCId defined
	MFCId null.Int `xorm:"int(11) pk notnull autoincr 'm_f_c_id'" json:"m_f_c_id" form:"m_f_c_id" xml:"m_f_c_id"`
	// MfbId defined
	MfbId null.Int `xorm:"int(11) 'mfb_id'" json:"mfb_id" form:"mfb_id" xml:"mfb_id"`
	// CostMoney defined
	CostMoney null.Float `xorm:"float(10,2) 'cost_money'" json:"cost_money" form:"cost_money" xml:"cost_money"`
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
	// Remark defined
	Remark null.String `xorm:"varchar(1000) 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// Marshal defined
func (m *MarketFeeCost) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketFeeCost) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *MarketFeeCost) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MarketFeeCost) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketFeeCost
func (m *MarketFeeCost) TableName() string {
	return "market_fee_cost"
}
