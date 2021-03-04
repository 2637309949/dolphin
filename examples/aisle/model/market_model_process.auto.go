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

// MarketModelProcess defined
type MarketModelProcess struct {
	// MMPId defined
	MMPId null.Int `xorm:"int(11) pk notnull autoincr 'm_m_p_id'" json:"m_m_p_id" form:"m_m_p_id" xml:"m_m_p_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// MmId defined
	MmId null.Int `xorm:"int(11) 'mm_id'" json:"mm_id" form:"mm_id" xml:"mm_id"`
	// MmpName defined
	MmpName null.String `xorm:"varchar(500) 'mmp_name'" json:"mmp_name" form:"mmp_name" xml:"mmp_name"`
	// MmpContent defined
	MmpContent null.String `xorm:"varchar(1000) 'mmp_content'" json:"mmp_content" form:"mmp_content" xml:"mmp_content"`
	// MmpResult defined
	MmpResult null.String `xorm:"varchar(500) 'mmp_result'" json:"mmp_result" form:"mmp_result" xml:"mmp_result"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *MarketModelProcess) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketModelProcess) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *MarketModelProcess) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *MarketModelProcess) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *MarketModelProcess) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MarketModelProcess) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketModelProcess
func (m *MarketModelProcess) TableName() string {
	return "market_model_process"
}
