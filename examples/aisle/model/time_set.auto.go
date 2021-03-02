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

// TimeSet defined
type TimeSet struct {
	// T3250 defined
	T3250 null.Int `xorm:"int(11) pk notnull autoincr 't_325_0'" json:"t_325_0" form:"t_325_0" xml:"t_325_0"`
	// TimeName defined
	TimeName null.String `xorm:"varchar(500) notnull 'time_name'" json:"time_name" form:"time_name" xml:"time_name"`
	// TimeCode defined
	TimeCode null.String `xorm:"varchar(500) notnull 'time_code'" json:"time_code" form:"time_code" xml:"time_code"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *TimeSet) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TimeSet) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *TimeSet) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TimeSet) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TimeSet
func (m *TimeSet) TableName() string {
	return "time_set"
}
