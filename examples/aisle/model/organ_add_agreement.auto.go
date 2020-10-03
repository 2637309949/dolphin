// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// OrganAddAgreement defined
type OrganAddAgreement struct {
	//
	OAAId null.Int `xorm:"int(11) pk notnull autoincr 'o_a_a_id'" json:"o_a_a_id" xml:"o_a_a_id"`
	//
	PkOrgan null.Int `xorm:"int(11) 'pk_organ'" json:"pk_organ" xml:"pk_organ"`
	//
	PkAam null.Int `xorm:"int(11) 'pk_aam'" json:"pk_aam" xml:"pk_aam"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined OrganAddAgreement
func (m *OrganAddAgreement) TableName() string {
	return "organ_add_agreement"
}