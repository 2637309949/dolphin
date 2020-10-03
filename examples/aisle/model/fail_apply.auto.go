// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// FailApply defined
type FailApply struct {
	//
	FailApplyId null.Int `xorm:"int(11) pk notnull autoincr 'fail_apply_id'" json:"fail_apply_id" xml:"fail_apply_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	FaReason null.String `xorm:"varchar(1000) 'fa_reason'" json:"fa_reason" xml:"fa_reason"`
	//
	FaRemark null.String `xorm:"varchar(1000) 'fa_remark'" json:"fa_remark" xml:"fa_remark"`
	//
	FaOpponent null.String `xorm:"varchar(100) 'fa_opponent'" json:"fa_opponent" xml:"fa_opponent"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" xml:"check_state"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	CheckTwoState null.Int `xorm:"int(11) 'check_two_state'" json:"check_two_state" xml:"check_two_state"`
	//
	FailTime null.Time `xorm:"datetime 'fail_time'" json:"fail_time" xml:"fail_time"`
	//
	CheckRemark null.String `xorm:"varchar(300) 'check_remark'" json:"check_remark" xml:"check_remark"`
	//
	CheckTwoRemark null.String `xorm:"varchar(300) 'check_two_remark'" json:"check_two_remark" xml:"check_two_remark"`
	//
	EnteringCheckUser null.Int `xorm:"int(11) 'entering_check_user'" json:"entering_check_user" xml:"entering_check_user"`
	//
	EnteringCheckTime null.Time `xorm:"datetime 'entering_check_time'" json:"entering_check_time" xml:"entering_check_time"`
	//
	SalesSupervisorCheckUser null.Int `xorm:"int(11) 'sales_supervisor_check_user'" json:"sales_supervisor_check_user" xml:"sales_supervisor_check_user"`
	//
	SalesSupervisorCheckTime null.Time `xorm:"datetime 'sales_supervisor_check_time'" json:"sales_supervisor_check_time" xml:"sales_supervisor_check_time"`
	//
	PrincipalCheckUser null.Int `xorm:"int(11) 'principal_check_user'" json:"principal_check_user" xml:"principal_check_user"`
	//
	PrincipalCheckTime null.Time `xorm:"datetime 'principal_check_time'" json:"principal_check_time" xml:"principal_check_time"`
	//
	FaBussType null.Int `xorm:"int(11) 'fa_buss_type'" json:"fa_buss_type" xml:"fa_buss_type"`
	//
	PaCheckState null.Int `xorm:"int(11) 'pa_check_state'" json:"pa_check_state" xml:"pa_check_state"`
	//
	PaRemark null.String `xorm:"varchar(2000) 'pa_remark'" json:"pa_remark" xml:"pa_remark"`
	//
	IfFromTmk null.Int `xorm:"int(11) 'if_from_tmk'" json:"if_from_tmk" xml:"if_from_tmk"`
	//
	OldStuType null.Int `xorm:"int(11) 'old_stu_type'" json:"old_stu_type" xml:"old_stu_type"`
	//
	OldStuState null.Int `xorm:"int(11) 'old_stu_state'" json:"old_stu_state" xml:"old_stu_state"`
	//
	HjTmkjlCheckstate null.Int `xorm:"int(11) 'hj_tmkjl_checkstate'" json:"hj_tmkjl_checkstate" xml:"hj_tmkjl_checkstate"`
	//
	HjTmkjlCheckDesc null.String `xorm:"varchar(1000) 'hj_tmkjl_check_desc'" json:"hj_tmkjl_check_desc" xml:"hj_tmkjl_check_desc"`
}

// TableName table name of defined FailApply
func (m *FailApply) TableName() string {
	return "fail_apply"
}