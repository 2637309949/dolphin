// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// EmailSendSet defined
type EmailSendSet struct {
	// ESSId defined
	ESSId null.Int `xorm:"int(11) pk notnull autoincr 'e_s_s_id'" json:"e_s_s_id" form:"e_s_s_id" xml:"e_s_s_id"`
	// SetName defined
	SetName null.String `xorm:"varchar(50) 'set_name'" json:"set_name" form:"set_name" xml:"set_name"`
	// IfOpen defined
	IfOpen null.Int `xorm:"int(11) 'if_open'" json:"if_open" form:"if_open" xml:"if_open"`
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

// TableName table name of defined EmailSendSet
func (m *EmailSendSet) TableName() string {
	return "email_send_set"
}