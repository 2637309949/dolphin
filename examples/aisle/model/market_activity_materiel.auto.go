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

// MarketActivityMateriel defined
type MarketActivityMateriel struct {
	// MaId defined
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
	// MaterielId defined
	MaterielId null.Int `xorm:"int(11) 'materiel_id'" json:"materiel_id" form:"materiel_id" xml:"materiel_id"`
	// PlanNum defined
	PlanNum null.Int `xorm:"int(11) 'plan_num'" json:"plan_num" form:"plan_num" xml:"plan_num"`
	// ReceiveNum defined
	ReceiveNum null.Int `xorm:"int(11) 'receive_num'" json:"receive_num" form:"receive_num" xml:"receive_num"`
	// UseNum defined
	UseNum null.Int `xorm:"int(11) 'use_num'" json:"use_num" form:"use_num" xml:"use_num"`
	// ReturnNum defined
	ReturnNum null.Int `xorm:"int(11) 'return_num'" json:"return_num" form:"return_num" xml:"return_num"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// MAMId defined
	MAMId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_m_id'" json:"m_a_m_id" form:"m_a_m_id" xml:"m_a_m_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// With defined
func (m *MarketActivityMateriel) With(s interface{}) (interface{}, error) {
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
func (m *MarketActivityMateriel) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketActivityMateriel) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *MarketActivityMateriel) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *MarketActivityMateriel) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *MarketActivityMateriel) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *MarketActivityMateriel) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketActivityMateriel
func (m *MarketActivityMateriel) TableName() string {
	return "market_activity_materiel"
}
