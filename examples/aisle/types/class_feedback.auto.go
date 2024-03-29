// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ClassFeedback defined
type ClassFeedback struct {
	// CFId defined
	CFId null.Int `xorm:"int(11) pk notnull autoincr 'c_f_id'" json:"c_f_id" form:"c_f_id" xml:"c_f_id"`
	// CfContent defined
	CfContent null.String `xorm:"'cf_content'" json:"cf_content" form:"cf_content" xml:"cf_content"`
	// CsId defined
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// CfTime defined
	CfTime null.Time `xorm:"datetime 'cf_time'" json:"cf_time" form:"cf_time" xml:"cf_time"`
	// Hear defined
	Hear null.Int `xorm:"int(11) 'hear'" json:"hear" form:"hear" xml:"hear"`
	// Say defined
	Say null.Int `xorm:"int(11) 'say'" json:"say" form:"say" xml:"say"`
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
	// CfRead defined
	CfRead null.Int `xorm:"int(11) 'cf_read'" json:"cf_read" form:"cf_read" xml:"cf_read"`
	// CfWrite defined
	CfWrite null.Int `xorm:"int(11) 'cf_write'" json:"cf_write" form:"cf_write" xml:"cf_write"`
	// CfZhnl defined
	CfZhnl null.Int `xorm:"int(11) 'cf_zhnl'" json:"cf_zhnl" form:"cf_zhnl" xml:"cf_zhnl"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) default(54) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// CheckTime defined
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" form:"check_time" xml:"check_time"`
	// CheckNoReason defined
	CheckNoReason null.String `xorm:"varchar(5000) 'check_no_reason'" json:"check_no_reason" form:"check_no_reason" xml:"check_no_reason"`
	// CheckUser defined
	CheckUser null.Int `xorm:"int(11) 'check_user'" json:"check_user" form:"check_user" xml:"check_user"`
	// AddUserType defined
	AddUserType null.String `xorm:"varchar(1000) 'add_user_type'" json:"add_user_type" form:"add_user_type" xml:"add_user_type"`
	// See defined
	See null.Int `xorm:"int(11) 'see'" json:"see" form:"see" xml:"see"`
	// IfSendStu defined
	IfSendStu null.Int `xorm:"int(11) default(3) 'if_send_stu'" json:"if_send_stu" form:"if_send_stu" xml:"if_send_stu"`
}

// TableName table name of defined ClassFeedback
func (m *ClassFeedback) TableName() string {
	return "class_feedback"
}

func (r *ClassFeedback) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClassFeedback(data []byte) (ClassFeedback, error) {
	var r ClassFeedback
	err := json.Unmarshal(data, &r)
	return r, err
}
