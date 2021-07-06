// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ClassFeedbackFile defined
type ClassFeedbackFile struct {
	// CFFId defined
	CFFId null.Int `xorm:"int(11) pk notnull autoincr 'c_f_f_id'" json:"c_f_f_id" form:"c_f_f_id" xml:"c_f_f_id"`
	// CfId defined
	CfId null.Int `xorm:"int(11) 'cf_id'" json:"cf_id" form:"cf_id" xml:"cf_id"`
	// Filed defined
	Filed null.Int `xorm:"int(11) 'filed'" json:"filed" form:"filed" xml:"filed"`
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

// TableName table name of defined ClassFeedbackFile
func (m *ClassFeedbackFile) TableName() string {
	return "class_feedback_file"
}