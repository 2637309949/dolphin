// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ScClassTea defined
type ScClassTea struct {
	// SCTId defined
	SCTId null.Int `xorm:"int(11) pk notnull autoincr 's_c_t_id'" json:"s_c_t_id" form:"s_c_t_id" xml:"s_c_t_id"`
	// TeaId defined
	TeaId null.Int `xorm:"int(11) 'tea_id'" json:"tea_id" form:"tea_id" xml:"tea_id"`
	// TeaStartDate defined
	TeaStartDate null.Time `xorm:"datetime 'tea_start_date'" json:"tea_start_date" form:"tea_start_date" xml:"tea_start_date"`
	// TeaEndDate defined
	TeaEndDate null.Time `xorm:"datetime 'tea_end_date'" json:"tea_end_date" form:"tea_end_date" xml:"tea_end_date"`
	// ScId defined
	ScId null.Int `xorm:"int(11) 'sc_id'" json:"sc_id" form:"sc_id" xml:"sc_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// SctState defined
	SctState null.Int `xorm:"int(11) 'sct_state'" json:"sct_state" form:"sct_state" xml:"sct_state"`
	// LeaveNotes defined
	LeaveNotes null.String `xorm:"varchar(500) 'leave_notes'" json:"leave_notes" form:"leave_notes" xml:"leave_notes"`
	// IfCountHour defined
	IfCountHour null.Int `xorm:"int(11) 'if_count_hour'" json:"if_count_hour" form:"if_count_hour" xml:"if_count_hour"`
	// NotCountReason defined
	NotCountReason null.String `xorm:"varchar(1000) 'not_count_reason'" json:"not_count_reason" form:"not_count_reason" xml:"not_count_reason"`
	// KqHour defined
	KqHour null.Float `xorm:"float(10,2) 'kq_hour'" json:"kq_hour" form:"kq_hour" xml:"kq_hour"`
	// QjType defined
	QjType null.Int `xorm:"int(11) 'qj_type'" json:"qj_type" form:"qj_type" xml:"qj_type"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// ShenhePerson defined
	ShenhePerson null.Int `xorm:"int(11) 'shenhe_person'" json:"shenhe_person" form:"shenhe_person" xml:"shenhe_person"`
	// KqKc defined
	KqKc null.Float `xorm:"float(50,2) 'kq_kc'" json:"kq_kc" form:"kq_kc" xml:"kq_kc"`
	// ParId defined
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" form:"par_id" xml:"par_id"`
	// PkClass defined
	PkClass null.Int `xorm:"int(11) 'pk_class'" json:"pk_class" form:"pk_class" xml:"pk_class"`
	// ChangeTeaId defined
	ChangeTeaId null.Int `xorm:"int(11) 'change_tea_id'" json:"change_tea_id" form:"change_tea_id" xml:"change_tea_id"`
	// RealKc defined
	RealKc null.Float `xorm:"float(50,2) 'real_kc'" json:"real_kc" form:"real_kc" xml:"real_kc"`
	// Remark defined
	Remark null.String `xorm:"varchar(100) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// ZdNotReason defined
	ZdNotReason null.Int `xorm:"int(11) 'zd_not_reason'" json:"zd_not_reason" form:"zd_not_reason" xml:"zd_not_reason"`
}

// TableName table name of defined ScClassTea
func (m *ScClassTea) TableName() string {
	return "sc_class_tea"
}
