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

// AutoCsTimeSet defined
type AutoCsTimeSet struct {
	// ACTSId defined
	ACTSId null.Int `xorm:"int(11) pk notnull autoincr 'a_c_t_s_id'" json:"a_c_t_s_id" form:"a_c_t_s_id" xml:"a_c_t_s_id"`
	// SetWeek defined
	SetWeek null.Int `xorm:"int(11) 'set_week'" json:"set_week" form:"set_week" xml:"set_week"`
	// SetHour defined
	SetHour null.Float `xorm:"float(11,2) 'set_hour'" json:"set_hour" form:"set_hour" xml:"set_hour"`
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
	// SetMinute defined
	SetMinute null.Int `xorm:"int(11) 'set_minute'" json:"set_minute" form:"set_minute" xml:"set_minute"`
	// IfOpen defined
	IfOpen null.Int `xorm:"int(11) 'if_open'" json:"if_open" form:"if_open" xml:"if_open"`
	// TimeName defined
	TimeName null.String `xorm:"varchar(20) 'time_name'" json:"time_name" form:"time_name" xml:"time_name"`
	// StartTimeStr defined
	StartTimeStr null.String `xorm:"varchar(5) 'start_time_str'" json:"start_time_str" form:"start_time_str" xml:"start_time_str"`
	// EndTimeStr defined
	EndTimeStr null.String `xorm:"varchar(5) 'end_time_str'" json:"end_time_str" form:"end_time_str" xml:"end_time_str"`
}

// Parser defined
func (m *AutoCsTimeSet) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *AutoCsTimeSet) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined AutoCsTimeSet
func (m *AutoCsTimeSet) TableName() string {
	return "auto_cs_time_set"
}
