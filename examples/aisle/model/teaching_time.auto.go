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

// TeachingTime defined
type TeachingTime struct {
	// TTId defined
	TTId null.Int `xorm:"int(11) pk notnull autoincr 't_t_id'" json:"t_t_id" form:"t_t_id" xml:"t_t_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// TeaId defined
	TeaId null.Int `xorm:"int(11) 'tea_id'" json:"tea_id" form:"tea_id" xml:"tea_id"`
	// BeginDate defined
	BeginDate null.Time `xorm:"datetime 'begin_date'" json:"begin_date" form:"begin_date" xml:"begin_date"`
	// BeginTime defined
	BeginTime null.Time `xorm:"datetime 'begin_time'" json:"begin_time" form:"begin_time" xml:"begin_time"`
	// Endtime defined
	Endtime null.Time `xorm:"datetime 'endtime'" json:"endtime" form:"endtime" xml:"endtime"`
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
	// Minutes defined
	Minutes null.Float `xorm:"float(50,2) 'minutes'" json:"minutes" form:"minutes" xml:"minutes"`
	// TeaType defined
	TeaType null.Int `xorm:"int(11) 'tea_type'" json:"tea_type" form:"tea_type" xml:"tea_type"`
	// TeachingTimeDesc defined
	TeachingTimeDesc null.String `xorm:"varchar(500) 'teaching_time_desc'" json:"teaching_time_desc" form:"teaching_time_desc" xml:"teaching_time_desc"`
	// IfKouHour defined
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
	// ClassHour defined
	ClassHour null.Float `xorm:"float(50,2) 'class_hour'" json:"class_hour" form:"class_hour" xml:"class_hour"`
}

// Marshal defined
func (m *TeachingTime) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TeachingTime) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *TeachingTime) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TeachingTime) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TeachingTime
func (m *TeachingTime) TableName() string {
	return "teaching_time"
}
