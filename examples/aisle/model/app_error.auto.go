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

// AppError defined
type AppError struct {
	// AppErrorId defined
	AppErrorId null.Int `xorm:"int(11) pk notnull autoincr 'app_error_id'" json:"app_error_id" form:"app_error_id" xml:"app_error_id"`
	// ErrorContent defined
	ErrorContent null.String `xorm:"varchar(2000) 'error_content'" json:"error_content" form:"error_content" xml:"error_content"`
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
	// ErrorType defined
	ErrorType null.Int `xorm:"int(11) 'error_type'" json:"error_type" form:"error_type" xml:"error_type"`
	// StudentId defined
	StudentId null.Int `xorm:"int(11) 'student_id'" json:"student_id" form:"student_id" xml:"student_id"`
	// PersonId defined
	PersonId null.Int `xorm:"int(11) 'person_id'" json:"person_id" form:"person_id" xml:"person_id"`
}

// Parser defined
func (m *AppError) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *AppError) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined AppError
func (m *AppError) TableName() string {
	return "app_error"
}
