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

// MaterielMinus defined
type MaterielMinus struct {
	//
	MMId null.Int `xorm:"int(11) pk notnull autoincr 'm_m_id'" json:"m_m_id" form:"m_m_id" xml:"m_m_id"`
	//
	MaterielId null.Int `xorm:"int(11) 'materiel_id'" json:"materiel_id" form:"materiel_id" xml:"materiel_id"`
	//
	MinusNum null.Int `xorm:"int(11) 'minus_num'" json:"minus_num" form:"minus_num" xml:"minus_num"`
	//
	MinusMoney null.Float `xorm:"float(11,2) 'minus_money'" json:"minus_money" form:"minus_money" xml:"minus_money"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	MinusReason null.Int `xorm:"int(11) 'minus_reason'" json:"minus_reason" form:"minus_reason" xml:"minus_reason"`
	//
	MinusObject null.Int `xorm:"int(11) 'minus_object'" json:"minus_object" form:"minus_object" xml:"minus_object"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Parser defined
func (m *MaterielMinus) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MaterielMinus) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MaterielMinus
func (m *MaterielMinus) TableName() string {
	return "materiel_minus"
}
