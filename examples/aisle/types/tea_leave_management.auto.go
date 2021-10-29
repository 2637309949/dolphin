// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// TeaLeaveManagement defined
type TeaLeaveManagement struct {
	// TLMId defined
	TLMId null.Int `xorm:"int(11) pk notnull autoincr 't_l_m_id'" json:"t_l_m_id" form:"t_l_m_id" xml:"t_l_m_id"`
	// TeaId defined
	TeaId null.Int `xorm:"int(11) 'tea_id'" json:"tea_id" form:"tea_id" xml:"tea_id"`
	// LeaveType defined
	LeaveType null.Int `xorm:"int(11) 'leave_type'" json:"leave_type" form:"leave_type" xml:"leave_type"`
	// LeaveStartDate defined
	LeaveStartDate null.Time `xorm:"datetime 'leave_start_date'" json:"leave_start_date" form:"leave_start_date" xml:"leave_start_date"`
	// LeaveStartTime defined
	LeaveStartTime null.Time `xorm:"datetime 'leave_start_time'" json:"leave_start_time" form:"leave_start_time" xml:"leave_start_time"`
	// LeaveEndDate defined
	LeaveEndDate null.Time `xorm:"datetime 'leave_end_date'" json:"leave_end_date" form:"leave_end_date" xml:"leave_end_date"`
	// LeaveEndTime defined
	LeaveEndTime null.Time `xorm:"datetime 'leave_end_time'" json:"leave_end_time" form:"leave_end_time" xml:"leave_end_time"`
	// LeaveWordDay defined
	LeaveWordDay null.Float `xorm:"float(11,2) 'leave_word_day'" json:"leave_word_day" form:"leave_word_day" xml:"leave_word_day"`
	// TeaStuNum defined
	TeaStuNum null.Float `xorm:"float(11,2) 'tea_stu_num'" json:"tea_stu_num" form:"tea_stu_num" xml:"tea_stu_num"`
	// Remark defined
	Remark null.String `xorm:"varchar(1000) 'remark'" json:"remark" form:"remark" xml:"remark"`
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

// TableName table name of defined TeaLeaveManagement
func (m *TeaLeaveManagement) TableName() string {
	return "tea_leave_management"
}

func (r *TeaLeaveManagement) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTeaLeaveManagement(data []byte) (TeaLeaveManagement, error) {
	var r TeaLeaveManagement
	err := json.Unmarshal(data, &r)
	return r, err
}
