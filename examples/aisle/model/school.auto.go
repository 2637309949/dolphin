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

// School defined
type School struct {
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// SchoolId defined
	SchoolId null.Int `xorm:"int(11) pk notnull autoincr 'school_id'" json:"school_id" form:"school_id" xml:"school_id"`
	// SchoolName defined
	SchoolName null.String `xorm:"varchar(500) 'school_name'" json:"school_name" form:"school_name" xml:"school_name"`
	// SchoolArea defined
	SchoolArea null.String `xorm:"varchar(2000) 'school_area'" json:"school_area" form:"school_area" xml:"school_area"`
	// SchoolRemark defined
	SchoolRemark null.String `xorm:"varchar(2000) 'school_remark'" json:"school_remark" form:"school_remark" xml:"school_remark"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// OrganId defined
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
}

// Marshal defined
func (m *School) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *School) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *School) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *School) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined School
func (m *School) TableName() string {
	return "school"
}
