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

// ClassTimeRange defined
type ClassTimeRange struct {
	// CTRId defined
	CTRId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_r_id'" json:"c_t_r_id" form:"c_t_r_id" xml:"c_t_r_id"`
	// TrId defined
	TrId null.Int `xorm:"int(11) 'tr_id'" json:"tr_id" form:"tr_id" xml:"tr_id"`
	// StartDate defined
	StartDate null.Time `xorm:"datetime 'start_date'" json:"start_date" form:"start_date" xml:"start_date"`
	// EndDate defined
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" form:"end_date" xml:"end_date"`
	// StartWeek defined
	StartWeek null.Int `xorm:"int(11) 'start_week'" json:"start_week" form:"start_week" xml:"start_week"`
	// EndWeek defined
	EndWeek null.Int `xorm:"int(11) 'end_week'" json:"end_week" form:"end_week" xml:"end_week"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// Hour defined
	Hour null.Float `xorm:"float(11,2) 'hour'" json:"hour" form:"hour" xml:"hour"`
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
	// TimeRemark defined
	TimeRemark null.String `xorm:"varchar(200) 'time_remark'" json:"time_remark" form:"time_remark" xml:"time_remark"`
	// IsThereATimeCourse defined
	IsThereATimeCourse null.Int `xorm:"int(11) 'is_there_a_time_course'" json:"is_there_a_time_course" form:"is_there_a_time_course" xml:"is_there_a_time_course"`
	// PkPkTime defined
	PkPkTime null.Int `xorm:"int(11) 'pk_pk_time'" json:"pk_pk_time" form:"pk_pk_time" xml:"pk_pk_time"`
}

// Marshal defined
func (m *ClassTimeRange) With(s interface{}) (interface{}, error) {
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
func (m *ClassTimeRange) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ClassTimeRange) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ClassTimeRange) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ClassTimeRange) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ClassTimeRange) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ClassTimeRange) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ClassTimeRange
func (m *ClassTimeRange) TableName() string {
	return "class_time_range"
}
