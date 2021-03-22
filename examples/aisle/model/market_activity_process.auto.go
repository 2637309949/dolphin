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

// MarketActivityProcess defined
type MarketActivityProcess struct {
	// MaId defined
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
	// MapName defined
	MapName null.String `xorm:"varchar(50) 'map_name'" json:"map_name" form:"map_name" xml:"map_name"`
	// MapContent defined
	MapContent null.String `xorm:"varchar(500) 'map_content'" json:"map_content" form:"map_content" xml:"map_content"`
	// MapResult defined
	MapResult null.String `xorm:"varchar(50) 'map_result'" json:"map_result" form:"map_result" xml:"map_result"`
	// RunId defined
	RunId null.Int `xorm:"int(11) 'run_id'" json:"run_id" form:"run_id" xml:"run_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// MAPId defined
	MAPId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_p_id'" json:"m_a_p_id" form:"m_a_p_id" xml:"m_a_p_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// With defined
func (m *MarketActivityProcess) With(s interface{}) (interface{}, error) {
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
func (m *MarketActivityProcess) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketActivityProcess) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *MarketActivityProcess) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *MarketActivityProcess) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *MarketActivityProcess) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *MarketActivityProcess) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketActivityProcess
func (m *MarketActivityProcess) TableName() string {
	return "market_activity_process"
}
