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

// PaActivityClassStu defined
type PaActivityClassStu struct {
	// PACSId defined
	PACSId null.Int `xorm:"int(11) pk notnull autoincr 'p_a_c_s_id'" json:"p_a_c_s_id" form:"p_a_c_s_id" xml:"p_a_c_s_id"`
	// ParId defined
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" form:"par_id" xml:"par_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
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
}

// Marshal defined
func (m *PaActivityClassStu) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *PaActivityClassStu) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *PaActivityClassStu) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *PaActivityClassStu) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined PaActivityClassStu
func (m *PaActivityClassStu) TableName() string {
	return "pa_activity_class_stu"
}
