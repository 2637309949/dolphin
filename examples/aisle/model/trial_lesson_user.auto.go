// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// TrialLessonUser defined
type TrialLessonUser struct {
	//
	TLUId null.Int `xorm:"int(11) pk notnull autoincr 't_l_u_id'" json:"t_l_u_id" xml:"t_l_u_id"`
	//
	TlId null.Int `xorm:"int(11) 'tl_id'" json:"tl_id" xml:"tl_id"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
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

// TableName table name of defined TrialLessonUser
func (m *TrialLessonUser) TableName() string {
	return "trial_lesson_user"
}