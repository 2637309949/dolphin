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

// ClassTypeStageCourse defined
type ClassTypeStageCourse struct {
	// CTSCId defined
	CTSCId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_s_c_id'" json:"c_t_s_c_id" form:"c_t_s_c_id" xml:"c_t_s_c_id"`
	// CtsId defined
	CtsId null.Int `xorm:"int(11) 'cts_id'" json:"cts_id" form:"cts_id" xml:"cts_id"`
	// CourseId defined
	CourseId null.Int `xorm:"int(11) 'course_id'" json:"course_id" form:"course_id" xml:"course_id"`
	// CtscHour defined
	CtscHour null.Int `xorm:"int(11) 'ctsc_hour'" json:"ctsc_hour" form:"ctsc_hour" xml:"ctsc_hour"`
	// CtscMinute defined
	CtscMinute null.Int `xorm:"int(11) 'ctsc_minute'" json:"ctsc_minute" form:"ctsc_minute" xml:"ctsc_minute"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// Price defined
	Price null.Float `xorm:"float(10,2) 'price'" json:"price" form:"price" xml:"price"`
	// AllPrice defined
	AllPrice null.Float `xorm:"float(10,2) 'all_price'" json:"all_price" form:"all_price" xml:"all_price"`
	// CtId defined
	CtId null.Int `xorm:"int(11) 'ct_id'" json:"ct_id" form:"ct_id" xml:"ct_id"`
}

// Marshal defined
func (m *ClassTypeStageCourse) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ClassTypeStageCourse) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ClassTypeStageCourse) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ClassTypeStageCourse) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ClassTypeStageCourse) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ClassTypeStageCourse) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ClassTypeStageCourse
func (m *ClassTypeStageCourse) TableName() string {
	return "class_type_stage_course"
}
