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

// OrganAddAgreement defined
type OrganAddAgreement struct {
	//
	OAAId null.Int `xorm:"int(11) pk notnull autoincr 'o_a_a_id'" json:"o_a_a_id" form:"o_a_a_id" xml:"o_a_a_id"`
	//
	PkOrgan null.Int `xorm:"int(11) 'pk_organ'" json:"pk_organ" form:"pk_organ" xml:"pk_organ"`
	//
	PkAam null.Int `xorm:"int(11) 'pk_aam'" json:"pk_aam" form:"pk_aam" xml:"pk_aam"`
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
func (m *OrganAddAgreement) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *OrganAddAgreement) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OrganAddAgreement
func (m *OrganAddAgreement) TableName() string {
	return "organ_add_agreement"
}
