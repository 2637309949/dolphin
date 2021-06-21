// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// OrganAddAgreement defined
type OrganAddAgreement struct {
	// OAAId defined
	OAAId null.Int `xorm:"int(11) pk notnull autoincr 'o_a_a_id'" json:"o_a_a_id" form:"o_a_a_id" xml:"o_a_a_id"`
	// PkOrgan defined
	PkOrgan null.Int `xorm:"int(11) 'pk_organ'" json:"pk_organ" form:"pk_organ" xml:"pk_organ"`
	// PkAam defined
	PkAam null.Int `xorm:"int(11) 'pk_aam'" json:"pk_aam" form:"pk_aam" xml:"pk_aam"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"notnull 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined OrganAddAgreement
func (m *OrganAddAgreement) TableName() string {
	return "organ_add_agreement"
}
