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

// HourAllotRatio defined
type HourAllotRatio struct {
	//
	HARId null.Int `xorm:"int(11) pk notnull autoincr 'h_a_r_id'" json:"h_a_r_id" form:"h_a_r_id" xml:"h_a_r_id"`
	//
	Period null.Int `xorm:"int(11) 'period'" json:"period" form:"period" xml:"period"`
	//
	Ratio null.String `xorm:"varchar(10) 'ratio'" json:"ratio" form:"ratio" xml:"ratio"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	PkHamId null.Int `xorm:"int(11) 'pk_ham_id'" json:"pk_ham_id" form:"pk_ham_id" xml:"pk_ham_id"`
	//
	EndMolecule null.Int `xorm:"int(11) 'end_molecule'" json:"end_molecule" form:"end_molecule" xml:"end_molecule"`
	//
	StartMolecule null.Int `xorm:"int(11) 'start_molecule'" json:"start_molecule" form:"start_molecule" xml:"start_molecule"`
}

// Parser defined
func (m *HourAllotRatio) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *HourAllotRatio) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined HourAllotRatio
func (m *HourAllotRatio) TableName() string {
	return "hour_allot_ratio"
}
