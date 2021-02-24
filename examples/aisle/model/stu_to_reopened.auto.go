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

// StuToReopened defined
type StuToReopened struct {
	//
	STRId null.Int `xorm:"int(11) pk notnull autoincr 's_t_r_id'" json:"s_t_r_id" form:"s_t_r_id" xml:"s_t_r_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	//
	StrRemark null.String `xorm:"varchar(2000) 'str_remark'" json:"str_remark" form:"str_remark" xml:"str_remark"`
	//
	StrDate null.Time `xorm:"datetime 'str_date'" json:"str_date" form:"str_date" xml:"str_date"`
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
	StrBussType null.Int `xorm:"int(11) 'str_buss_type'" json:"str_buss_type" form:"str_buss_type" xml:"str_buss_type"`
}

// Parser defined
func (m *StuToReopened) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuToReopened) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuToReopened
func (m *StuToReopened) TableName() string {
	return "stu_to_reopened"
}
