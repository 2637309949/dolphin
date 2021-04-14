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

// ClassManageStageCourse defined
type ClassManageStageCourse struct {
	// CMSCId defined
	CMSCId null.Int `xorm:"int(11) pk notnull autoincr 'c_m_s_c_id'" json:"c_m_s_c_id" form:"c_m_s_c_id" xml:"c_m_s_c_id"`
	// CmsId defined
	CmsId null.Int `xorm:"int(11) 'cms_id'" json:"cms_id" form:"cms_id" xml:"cms_id"`
	// CtscId defined
	CtscId null.Int `xorm:"int(11) 'ctsc_id'" json:"ctsc_id" form:"ctsc_id" xml:"ctsc_id"`
	// CourseId defined
	CourseId null.Int `xorm:"int(11) 'course_id'" json:"course_id" form:"course_id" xml:"course_id"`
	// CmscHour defined
	CmscHour null.Int `xorm:"int(11) 'cmsc_hour'" json:"cmsc_hour" form:"cmsc_hour" xml:"cmsc_hour"`
	// CmscMinute defined
	CmscMinute null.Int `xorm:"int(11) 'cmsc_minute'" json:"cmsc_minute" form:"cmsc_minute" xml:"cmsc_minute"`
	// PlanHour defined
	PlanHour null.Int `xorm:"int(11) 'plan_hour'" json:"plan_hour" form:"plan_hour" xml:"plan_hour"`
	// QrHour defined
	QrHour null.Int `xorm:"int(11) 'qr_hour'" json:"qr_hour" form:"qr_hour" xml:"qr_hour"`
	// UseHour defined
	UseHour null.Int `xorm:"int(11) 'use_hour'" json:"use_hour" form:"use_hour" xml:"use_hour"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// ClassId defined
	ClassId null.Int `xorm:"int(11) 'class_id'" json:"class_id" form:"class_id" xml:"class_id"`
}

// With defined
func (m *ClassManageStageCourse) With(s interface{}) (interface{}, error) {
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
func (m *ClassManageStageCourse) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ClassManageStageCourse) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ClassManageStageCourse) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ClassManageStageCourse) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ClassManageStageCourse) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *ClassManageStageCourse) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ClassManageStageCourse
func (m *ClassManageStageCourse) TableName() string {
	return "class_manage_stage_course"
}
