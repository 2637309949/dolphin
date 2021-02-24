// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// MarketActivityMateriel defined
type MarketActivityMateriel struct {
	//
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
	//
	MaterielId null.Int `xorm:"int(11) 'materiel_id'" json:"materiel_id" form:"materiel_id" xml:"materiel_id"`
	//
	PlanNum null.Int `xorm:"int(11) 'plan_num'" json:"plan_num" form:"plan_num" xml:"plan_num"`
	//
	ReceiveNum null.Int `xorm:"int(11) 'receive_num'" json:"receive_num" form:"receive_num" xml:"receive_num"`
	//
	UseNum null.Int `xorm:"int(11) 'use_num'" json:"use_num" form:"use_num" xml:"use_num"`
	//
	ReturnNum null.Int `xorm:"int(11) 'return_num'" json:"return_num" form:"return_num" xml:"return_num"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	MAMId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_m_id'" json:"m_a_m_id" form:"m_a_m_id" xml:"m_a_m_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Parser defined
func (m *MarketActivityMateriel) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
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
