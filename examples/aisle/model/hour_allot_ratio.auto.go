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

// HourAllotRatio defined
type HourAllotRatio struct {
	// HARId defined
	HARId null.Int `xorm:"int(11) pk notnull autoincr 'h_a_r_id'" json:"h_a_r_id" form:"h_a_r_id" xml:"h_a_r_id"`
	// Period defined
	Period null.Int `xorm:"int(11) 'period'" json:"period" form:"period" xml:"period"`
	// Ratio defined
	Ratio null.String `xorm:"varchar(10) 'ratio'" json:"ratio" form:"ratio" xml:"ratio"`
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
	// PkHamId defined
	PkHamId null.Int `xorm:"int(11) 'pk_ham_id'" json:"pk_ham_id" form:"pk_ham_id" xml:"pk_ham_id"`
	// EndMolecule defined
	EndMolecule null.Int `xorm:"int(11) 'end_molecule'" json:"end_molecule" form:"end_molecule" xml:"end_molecule"`
	// StartMolecule defined
	StartMolecule null.Int `xorm:"int(11) 'start_molecule'" json:"start_molecule" form:"start_molecule" xml:"start_molecule"`
}

// Marshal defined
func (m *HourAllotRatio) With(s interface{}) (interface{}, error) {
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
func (m *HourAllotRatio) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *HourAllotRatio) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *HourAllotRatio) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *HourAllotRatio) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
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
