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

// OtmClassroom defined
type OtmClassroom struct {
	//
	OCId null.Int `xorm:"int(11) pk notnull autoincr 'o_c_id'" json:"o_c_id" form:"o_c_id" xml:"o_c_id"`
	//
	ClassroomName null.String `xorm:"varchar(500) 'classroom_name'" json:"classroom_name" form:"classroom_name" xml:"classroom_name"`
	//
	SchoolId null.Int `xorm:"int(11) 'school_id'" json:"school_id" form:"school_id" xml:"school_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	ClassroomContain null.Float `xorm:"float(11,2) 'classroom_contain'" json:"classroom_contain" form:"classroom_contain" xml:"classroom_contain"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	Remarke null.String `xorm:"varchar(2000) 'remarke'" json:"remarke" form:"remarke" xml:"remarke"`
	//
	EnglishName null.String `xorm:"varchar(100) 'english_name'" json:"english_name" form:"english_name" xml:"english_name"`
}

// Parser defined
func (m *OtmClassroom) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *OtmClassroom) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OtmClassroom
func (m *OtmClassroom) TableName() string {
	return "otm_classroom"
}
