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

// MarketModel defined
type MarketModel struct {
	// MMId defined
	MMId null.Int `xorm:"int(11) pk notnull autoincr 'm_m_id'" json:"m_m_id" form:"m_m_id" xml:"m_m_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// MmName defined
	MmName null.String `xorm:"varchar(100) 'mm_name'" json:"mm_name" form:"mm_name" xml:"mm_name"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// MmType defined
	MmType null.Int `xorm:"int(11) 'mm_type'" json:"mm_type" form:"mm_type" xml:"mm_type"`
	// TsId defined
	TsId null.Int `xorm:"int(11) 'ts_id'" json:"ts_id" form:"ts_id" xml:"ts_id"`
	// SchoolId defined
	SchoolId null.Int `xorm:"int(11) 'school_id'" json:"school_id" form:"school_id" xml:"school_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *MarketModel) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketModel) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *MarketModel) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MarketModel) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketModel
func (m *MarketModel) TableName() string {
	return "market_model"
}
