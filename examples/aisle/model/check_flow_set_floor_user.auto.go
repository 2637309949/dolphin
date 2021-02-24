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

// CheckFlowSetFloorUser defined
type CheckFlowSetFloorUser struct {
	// CFSFUId defined
	CFSFUId null.Int `xorm:"int(11) pk notnull autoincr 'c_f_s_f_u_id'" json:"c_f_s_f_u_id" form:"c_f_s_f_u_id" xml:"c_f_s_f_u_id"`
	// PkCkfl defined
	PkCkfl null.Int `xorm:"int(11) 'pk_ckfl'" json:"pk_ckfl" form:"pk_ckfl" xml:"pk_ckfl"`
	// PkUser defined
	PkUser null.Int `xorm:"int(11) 'pk_user'" json:"pk_user" form:"pk_user" xml:"pk_user"`
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
}

// Parser defined
func (m *CheckFlowSetFloorUser) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *CheckFlowSetFloorUser) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CheckFlowSetFloorUser
func (m *CheckFlowSetFloorUser) TableName() string {
	return "check_flow_set_floor_user"
}
