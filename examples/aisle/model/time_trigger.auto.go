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

// TimeTrigger defined
type TimeTrigger struct {
	//
	TgTjNamestr null.String `xorm:"varchar(500) notnull 'tg_tj_namestr'" json:"tg_tj_namestr" form:"tg_tj_namestr" xml:"tg_tj_namestr"`
	//
	T3240 null.Int `xorm:"int(11) pk notnull autoincr 't_324_0'" json:"t_324_0" form:"t_324_0" xml:"t_324_0"`
	//
	TgName null.String `xorm:"varchar(500) notnull 'tg_name'" json:"tg_name" form:"tg_name" xml:"tg_name"`
	//
	Tablecnname null.String `xorm:"varchar(500) notnull 'tablecnname'" json:"tablecnname" form:"tablecnname" xml:"tablecnname"`
	//
	ExeBetweentime null.String `xorm:"varchar(500) notnull 'exe_betweentime'" json:"exe_betweentime" form:"exe_betweentime" xml:"exe_betweentime"`
	//
	ExeTime null.String `xorm:"varchar(500) notnull 'exe_time'" json:"exe_time" form:"exe_time" xml:"exe_time"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
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

// TableName table name of defined TimeTrigger
func (m *TimeTrigger) TableName() string {
	return "time_trigger"
}
