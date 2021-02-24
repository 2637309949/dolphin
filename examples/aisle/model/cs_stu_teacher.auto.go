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

// CsStuTeacher defined
type CsStuTeacher struct {
	// CSTId defined
	CSTId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_t_id'" json:"c_s_t_id" form:"c_s_t_id" xml:"c_s_t_id"`
	// PkCsStu defined
	PkCsStu null.Int `xorm:"int(11) 'pk_cs_stu'" json:"pk_cs_stu" form:"pk_cs_stu" xml:"pk_cs_stu"`
	// PkTea defined
	PkTea null.Int `xorm:"int(11) 'pk_tea'" json:"pk_tea" form:"pk_tea" xml:"pk_tea"`
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
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
}

// Parser defined
func (m *CsStuTeacher) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *CsStuTeacher) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CsStuTeacher
func (m *CsStuTeacher) TableName() string {
	return "cs_stu_teacher"
}
