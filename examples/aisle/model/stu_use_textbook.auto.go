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

// StuUseTextbook defined
type StuUseTextbook struct {
	//
	SUTId null.Int `xorm:"int(11) pk notnull autoincr 's_u_t_id'" json:"s_u_t_id" form:"s_u_t_id" xml:"s_u_t_id"`
	//
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" form:"pk_stu" xml:"pk_stu"`
	//
	UseStartDate null.Time `xorm:"datetime 'use_start_date'" json:"use_start_date" form:"use_start_date" xml:"use_start_date"`
	//
	UseEndDate null.Time `xorm:"datetime 'use_end_date'" json:"use_end_date" form:"use_end_date" xml:"use_end_date"`
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
	PkTm null.Int `xorm:"int(11) 'pk_tm'" json:"pk_tm" form:"pk_tm" xml:"pk_tm"`
	//
	PkLevel null.Int `xorm:"int(11) 'pk_level'" json:"pk_level" form:"pk_level" xml:"pk_level"`
}

// Parser defined
func (m *StuUseTextbook) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuUseTextbook) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuUseTextbook
func (m *StuUseTextbook) TableName() string {
	return "stu_use_textbook"
}
