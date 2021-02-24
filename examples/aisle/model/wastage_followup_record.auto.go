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

// WastageFollowupRecord defined
type WastageFollowupRecord struct {
	//
	WFRId null.Int `xorm:"int(11) pk notnull autoincr 'w_f_r_id'" json:"w_f_r_id" form:"w_f_r_id" xml:"w_f_r_id"`
	//
	EnglishLevel null.String `xorm:"varchar(20) 'english_level'" json:"english_level" form:"english_level" xml:"english_level"`
	//
	EnglishSchool null.String `xorm:"varchar(50) 'english_school'" json:"english_school" form:"english_school" xml:"english_school"`
	//
	IfStudy null.Int `xorm:"int(11) 'if_study'" json:"if_study" form:"if_study" xml:"if_study"`
	//
	TrackingProgram null.String `xorm:"varchar(200) 'tracking_program'" json:"tracking_program" form:"tracking_program" xml:"tracking_program"`
	//
	WastageMonth null.Float `xorm:"float(10,2) 'wastage_month'" json:"wastage_month" form:"wastage_month" xml:"wastage_month"`
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
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
}

// Parser defined
func (m *WastageFollowupRecord) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *WastageFollowupRecord) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined WastageFollowupRecord
func (m *WastageFollowupRecord) TableName() string {
	return "wastage_followup_record"
}
