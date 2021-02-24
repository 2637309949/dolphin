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

// StudentTypeJournal defined
type StudentTypeJournal struct {
	//
	STJId null.Int `xorm:"int(11) pk notnull autoincr 's_t_j_id'" json:"s_t_j_id" form:"s_t_j_id" xml:"s_t_j_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	//
	StuType null.Int `xorm:"int(11) 'stu_type'" json:"stu_type" form:"stu_type" xml:"stu_type"`
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
	StuTypeName null.String `xorm:"varchar(100) 'stu_type_name'" json:"stu_type_name" form:"stu_type_name" xml:"stu_type_name"`
	//
	Stutypedesc null.String `xorm:"varchar(500) 'stutypedesc'" json:"stutypedesc" form:"stutypedesc" xml:"stutypedesc"`
	//
	OldStuType null.Int `xorm:"int(11) 'old_stu_type'" json:"old_stu_type" form:"old_stu_type" xml:"old_stu_type"`
	//
	OldStuTypeName null.String `xorm:"varchar(1000) 'old_stu_type_name'" json:"old_stu_type_name" form:"old_stu_type_name" xml:"old_stu_type_name"`
}

// Parser defined
func (m *SysCommentReply) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysCommentReply) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentTypeJournal
func (m *StudentTypeJournal) TableName() string {
	return "student_type_journal"
}
