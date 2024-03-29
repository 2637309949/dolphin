// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// CsExecuteContent defined
type CsExecuteContent struct {
	// CECId defined
	CECId null.Int `xorm:"int(11) pk notnull autoincr 'c_e_c_id'" json:"c_e_c_id" form:"c_e_c_id" xml:"c_e_c_id"`
	// CspcId defined
	CspcId null.Int `xorm:"int(11) 'cspc_id'" json:"cspc_id" form:"cspc_id" xml:"cspc_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// PlanType defined
	PlanType null.Int `xorm:"int(11) 'plan_type'" json:"plan_type" form:"plan_type" xml:"plan_type"`
	// ExecuteDate defined
	ExecuteDate null.Time `xorm:"datetime 'execute_date'" json:"execute_date" form:"execute_date" xml:"execute_date"`
	// Feedback defined
	Feedback null.String `xorm:"varchar(500) 'feedback'" json:"feedback" form:"feedback" xml:"feedback"`
	// UnfinishedReason defined
	UnfinishedReason null.String `xorm:"varchar(500) 'unfinished_reason'" json:"unfinished_reason" form:"unfinished_reason" xml:"unfinished_reason"`
	// NextExecuteDate defined
	NextExecuteDate null.Time `xorm:"datetime 'next_execute_date'" json:"next_execute_date" form:"next_execute_date" xml:"next_execute_date"`
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

// TableName table name of defined CsExecuteContent
func (m *CsExecuteContent) TableName() string {
	return "cs_execute_content"
}

func (r *CsExecuteContent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCsExecuteContent(data []byte) (CsExecuteContent, error) {
	var r CsExecuteContent
	err := json.Unmarshal(data, &r)
	return r, err
}
