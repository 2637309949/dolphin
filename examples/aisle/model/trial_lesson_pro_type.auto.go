// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// TrialLessonProType defined
type TrialLessonProType struct {
	// TLPTId defined
	TLPTId null.Int `xorm:"int(11) pk notnull autoincr 't_l_p_t_id'" json:"t_l_p_t_id" form:"t_l_p_t_id" xml:"t_l_p_t_id"`
	// TlId defined
	TlId null.Int `xorm:"int(11) 'tl_id'" json:"tl_id" form:"tl_id" xml:"tl_id"`
	// PtId defined
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" form:"pt_id" xml:"pt_id"`
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

// TableName table name of defined TrialLessonProType
func (m *TrialLessonProType) TableName() string {
	return "trial_lesson_pro_type"
}
