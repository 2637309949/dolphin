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

// MarketModelMateriel defined
type MarketModelMateriel struct {
	// MMMId defined
	MMMId null.Int `xorm:"int(11) pk notnull autoincr 'm_m_m_id'" json:"m_m_m_id" form:"m_m_m_id" xml:"m_m_m_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// MmId defined
	MmId null.Int `xorm:"int(11) 'mm_id'" json:"mm_id" form:"mm_id" xml:"mm_id"`
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
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// MaterielId defined
	MaterielId null.Int `xorm:"int(11) 'materiel_id'" json:"materiel_id" form:"materiel_id" xml:"materiel_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *MarketModelMateriel) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MarketModelMateriel) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *MarketModelMateriel) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MarketModelMateriel) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketModelMateriel
func (m *MarketModelMateriel) TableName() string {
	return "market_model_materiel"
}
