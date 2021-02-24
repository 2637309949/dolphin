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

// ClassTimeRangeTea defined
type ClassTimeRangeTea struct {
	//
	CTRTId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_r_t_id'" json:"c_t_r_t_id" form:"c_t_r_t_id" xml:"c_t_r_t_id"`
	//
	PkCtr null.Int `xorm:"int(11) 'pk_ctr'" json:"pk_ctr" form:"pk_ctr" xml:"pk_ctr"`
	//
	PkTea null.Int `xorm:"int(11) 'pk_tea'" json:"pk_tea" form:"pk_tea" xml:"pk_tea"`
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

// TableName table name of defined ClassTimeRangeTea
func (m *ClassTimeRangeTea) TableName() string {
	return "class_time_range_tea"
}
