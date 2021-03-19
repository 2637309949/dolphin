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
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// ExpendType defined
	ExpendType null.Int `xorm:"int(11) 'expend_type'" json:"expend_type" form:"expend_type" xml:"expend_type"`
}

// With defined
func (m *MarketExpend) With(s interface{}) (interface{}, error) {
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
func (m *MarketExpend) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketExpend) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *MarketExpend) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *MarketExpend) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *MarketExpend) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *MarketExpend) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketExpend
func (m *MarketExpend) TableName() string {
	return "market_expend"
}
