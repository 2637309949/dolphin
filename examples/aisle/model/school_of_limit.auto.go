// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SchoolOfLimit defined
type SchoolOfLimit struct {
	// SOLId defined
	SOLId null.Int `xorm:"int(11) pk notnull autoincr 's_o_l_id'" json:"s_o_l_id" form:"s_o_l_id" xml:"s_o_l_id"`
	// PkSch defined
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
	// StartDate defined
	StartDate null.Time `xorm:"datetime 'start_date'" json:"start_date" form:"start_date" xml:"start_date"`
	// EndDate defined
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" form:"end_date" xml:"end_date"`
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
	// OfYjType defined
	OfYjType null.Int `xorm:"int(11) 'of_yj_type'" json:"of_yj_type" form:"of_yj_type" xml:"of_yj_type"`
	// OfBase defined
	OfBase decimal.Decimal `xorm:"decimal(50,3) 'of_base'" json:"of_base" form:"of_base" xml:"of_base"`
	// SolName defined
	SolName null.String `xorm:"varchar(20) 'sol_name'" json:"sol_name" form:"sol_name" xml:"sol_name"`
	// PkYyb defined
	PkYyb null.Int `xorm:"int(11) 'pk_yyb'" json:"pk_yyb" form:"pk_yyb" xml:"pk_yyb"`
}

// With defined
func (m *SchoolOfLimit) With(s interface{}) (interface{}, error) {
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
func (m *SchoolOfLimit) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SchoolOfLimit) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SchoolOfLimit) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SchoolOfLimit) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SchoolOfLimit) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *SchoolOfLimit) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SchoolOfLimit
func (m *SchoolOfLimit) TableName() string {
	return "school_of_limit"
}
