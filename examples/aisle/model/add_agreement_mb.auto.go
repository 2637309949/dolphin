// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// AddAgreementMb defined
type AddAgreementMb struct {
	// AAMId defined
	AAMId null.Int `xorm:"int(11) pk notnull autoincr 'a_a_m_id'" json:"a_a_m_id" form:"a_a_m_id" xml:"a_a_m_id"`
	// AamName defined
	AamName null.String `xorm:"varchar(20) 'aam_name'" json:"aam_name" form:"aam_name" xml:"aam_name"`
	// AamHead defined
	AamHead null.Int `xorm:"int(11) 'aam_head'" json:"aam_head" form:"aam_head" xml:"aam_head"`
	// AamMidd defined
	AamMidd null.Int `xorm:"int(11) 'aam_midd'" json:"aam_midd" form:"aam_midd" xml:"aam_midd"`
	// AamLast defined
	AamLast null.Int `xorm:"int(11) 'aam_last'" json:"aam_last" form:"aam_last" xml:"aam_last"`
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
	// AamTitle defined
	AamTitle null.String `xorm:"varchar(500) 'aam_title'" json:"aam_title" form:"aam_title" xml:"aam_title"`
	// OpenOrClose defined
	OpenOrClose null.Int `xorm:"int(11) 'open_or_close'" json:"open_or_close" form:"open_or_close" xml:"open_or_close"`
}

// TableName table name of defined AddAgreementMb
func (m *AddAgreementMb) TableName() string {
	return "add_agreement_mb"
}
