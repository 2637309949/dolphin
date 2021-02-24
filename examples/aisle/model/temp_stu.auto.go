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

// TempStu defined
type TempStu struct {
	//
	Fid null.Int `xorm:"int(11) pk notnull 'fid'" json:"fid" form:"fid" xml:"fid"`
	//
	StuName null.String `xorm:"varchar(255) 'stu_name'" json:"stu_name" form:"stu_name" xml:"stu_name"`
	//
	StuPhone null.String `xorm:"varchar(255) 'stu_phone'" json:"stu_phone" form:"stu_phone" xml:"stu_phone"`
}

// Parser defined
func (m *TempStu) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TempStu) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TempStu
func (m *TempStu) TableName() string {
	return "temp_stu"
}
