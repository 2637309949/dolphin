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

// TeachingTime defined
type TeachingTime struct {
	//
	TTId null.Int `xorm:"int(11) pk notnull autoincr 't_t_id'" json:"t_t_id" form:"t_t_id" xml:"t_t_id"`
	//
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	//
	TeaId null.Int `xorm:"int(11) 'tea_id'" json:"tea_id" form:"tea_id" xml:"tea_id"`
	//
	BeginDate null.Time `xorm:"datetime 'begin_date'" json:"begin_date" form:"begin_date" xml:"begin_date"`
	//
	BeginTime null.Time `xorm:"datetime 'begin_time'" json:"begin_time" form:"begin_time" xml:"begin_time"`
	//
	Endtime null.Time `xorm:"datetime 'endtime'" json:"endtime" form:"endtime" xml:"endtime"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	Minutes null.Float `xorm:"float(50,2) 'minutes'" json:"minutes" form:"minutes" xml:"minutes"`
	//
	TeaType null.Int `xorm:"int(11) 'tea_type'" json:"tea_type" form:"tea_type" xml:"tea_type"`
	//
	TeachingTimeDesc null.String `xorm:"varchar(500) 'teaching_time_desc'" json:"teaching_time_desc" form:"teaching_time_desc" xml:"teaching_time_desc"`
	//
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
	//
	ClassHour null.Float `xorm:"float(50,2) 'class_hour'" json:"class_hour" form:"class_hour" xml:"class_hour"`
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
