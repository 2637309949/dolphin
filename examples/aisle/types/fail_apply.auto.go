// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// FailApply defined
type FailApply struct {
	// FailApplyId defined
	FailApplyId null.Int `xorm:"int(11) pk notnull autoincr 'fail_apply_id'" json:"fail_apply_id" form:"fail_apply_id" xml:"fail_apply_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// FaReason defined
	FaReason null.String `xorm:"varchar(1000) 'fa_reason'" json:"fa_reason" form:"fa_reason" xml:"fa_reason"`
	// FaRemark defined
	FaRemark null.String `xorm:"varchar(1000) 'fa_remark'" json:"fa_remark" form:"fa_remark" xml:"fa_remark"`
	// FaOpponent defined
	FaOpponent null.String `xorm:"varchar(100) 'fa_opponent'" json:"fa_opponent" form:"fa_opponent" xml:"fa_opponent"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// CheckTwoState defined
	CheckTwoState null.Int `xorm:"int(11) 'check_two_state'" json:"check_two_state" form:"check_two_state" xml:"check_two_state"`
	// FailTime defined
	FailTime null.Time `xorm:"datetime 'fail_time'" json:"fail_time" form:"fail_time" xml:"fail_time"`
	// CheckRemark defined
	CheckRemark null.String `xorm:"varchar(300) 'check_remark'" json:"check_remark" form:"check_remark" xml:"check_remark"`
	// CheckTwoRemark defined
	CheckTwoRemark null.String `xorm:"varchar(300) 'check_two_remark'" json:"check_two_remark" form:"check_two_remark" xml:"check_two_remark"`
	// EnteringCheckUser defined
	EnteringCheckUser null.Int `xorm:"int(11) 'entering_check_user'" json:"entering_check_user" form:"entering_check_user" xml:"entering_check_user"`
	// EnteringCheckTime defined
	EnteringCheckTime null.Time `xorm:"datetime 'entering_check_time'" json:"entering_check_time" form:"entering_check_time" xml:"entering_check_time"`
	// SalesSupervisorCheckUser defined
	SalesSupervisorCheckUser null.Int `xorm:"int(11) 'sales_supervisor_check_user'" json:"sales_supervisor_check_user" form:"sales_supervisor_check_user" xml:"sales_supervisor_check_user"`
	// SalesSupervisorCheckTime defined
	SalesSupervisorCheckTime null.Time `xorm:"datetime 'sales_supervisor_check_time'" json:"sales_supervisor_check_time" form:"sales_supervisor_check_time" xml:"sales_supervisor_check_time"`
	// PrincipalCheckUser defined
	PrincipalCheckUser null.Int `xorm:"int(11) 'principal_check_user'" json:"principal_check_user" form:"principal_check_user" xml:"principal_check_user"`
	// PrincipalCheckTime defined
	PrincipalCheckTime null.Time `xorm:"datetime 'principal_check_time'" json:"principal_check_time" form:"principal_check_time" xml:"principal_check_time"`
	// FaBussType defined
	FaBussType null.Int `xorm:"int(11) 'fa_buss_type'" json:"fa_buss_type" form:"fa_buss_type" xml:"fa_buss_type"`
	// PaCheckState defined
	PaCheckState null.Int `xorm:"int(11) 'pa_check_state'" json:"pa_check_state" form:"pa_check_state" xml:"pa_check_state"`
	// PaRemark defined
	PaRemark null.String `xorm:"varchar(2000) 'pa_remark'" json:"pa_remark" form:"pa_remark" xml:"pa_remark"`
	// IfFromTmk defined
	IfFromTmk null.Int `xorm:"int(11) 'if_from_tmk'" json:"if_from_tmk" form:"if_from_tmk" xml:"if_from_tmk"`
	// OldStuType defined
	OldStuType null.Int `xorm:"int(11) 'old_stu_type'" json:"old_stu_type" form:"old_stu_type" xml:"old_stu_type"`
	// OldStuState defined
	OldStuState null.Int `xorm:"int(11) 'old_stu_state'" json:"old_stu_state" form:"old_stu_state" xml:"old_stu_state"`
	// HjTmkjlCheckstate defined
	HjTmkjlCheckstate null.Int `xorm:"int(11) 'hj_tmkjl_checkstate'" json:"hj_tmkjl_checkstate" form:"hj_tmkjl_checkstate" xml:"hj_tmkjl_checkstate"`
	// HjTmkjlCheckDesc defined
	HjTmkjlCheckDesc null.String `xorm:"varchar(1000) 'hj_tmkjl_check_desc'" json:"hj_tmkjl_check_desc" form:"hj_tmkjl_check_desc" xml:"hj_tmkjl_check_desc"`
}

// TableName table name of defined FailApply
func (m *FailApply) TableName() string {
	return "fail_apply"
}

func (r *FailApply) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFailApply(data []byte) (FailApply, error) {
	var r FailApply
	err := json.Unmarshal(data, &r)
	return r, err
}
