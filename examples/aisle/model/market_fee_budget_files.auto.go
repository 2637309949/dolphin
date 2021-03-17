// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// MarketFeeBudgetFiles defined
type MarketFeeBudgetFiles struct {
	// MFBFId defined
	MFBFId null.Int `xorm:"int(11) pk notnull autoincr 'm_f_b_f_id'" json:"m_f_b_f_id" form:"m_f_b_f_id" xml:"m_f_b_f_id"`
	// MfbId defined
	MfbId null.Int `xorm:"int(11) 'mfb_id'" json:"mfb_id" form:"mfb_id" xml:"mfb_id"`
	// MfbFile defined
	MfbFile null.Int `xorm:"int(11) 'mfb_file'" json:"mfb_file" form:"mfb_file" xml:"mfb_file"`
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
}

// Marshal defined
func (m *MarketFeeBudgetFiles) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *MarketFeeBudgetFiles) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketFeeBudgetFiles) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *MarketFeeBudgetFiles) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *MarketFeeBudgetFiles) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *MarketFeeBudgetFiles) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MarketFeeBudgetFiles) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketFeeBudgetFiles
func (m *MarketFeeBudgetFiles) TableName() string {
	return "market_fee_budget_files"
}
