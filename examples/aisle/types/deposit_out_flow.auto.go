// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// DepositOutFlow defined
type DepositOutFlow struct {
	// CFlowId defined
	CFlowId null.Int `xorm:"int(11) pk notnull autoincr 'c_flow_id'" json:"c_flow_id" form:"c_flow_id" xml:"c_flow_id"`
	// PkDOId defined
	PkDOId null.Int `xorm:"int(11) 'pk_d_o_id'" json:"pk_d_o_id" form:"pk_d_o_id" xml:"pk_d_o_id"`
	// PkCfs defined
	PkCfs null.Int `xorm:"int(11) 'pk_cfs'" json:"pk_cfs" form:"pk_cfs" xml:"pk_cfs"`
	// PkUser defined
	PkUser null.Int `xorm:"int(11) 'pk_user'" json:"pk_user" form:"pk_user" xml:"pk_user"`
	// ZdCheckState defined
	ZdCheckState null.Int `xorm:"int(11) 'zd_check_state'" json:"zd_check_state" form:"zd_check_state" xml:"zd_check_state"`
	// CheckDate defined
	CheckDate null.Time `xorm:"datetime 'check_date'" json:"check_date" form:"check_date" xml:"check_date"`
	// CheckRemark defined
	CheckRemark null.String `xorm:"varchar(500) 'check_remark'" json:"check_remark" form:"check_remark" xml:"check_remark"`
	// NowFloor defined
	NowFloor null.Int `xorm:"int(11) 'now_floor'" json:"now_floor" form:"now_floor" xml:"now_floor"`
	// PkCfPool defined
	PkCfPool null.Int `xorm:"int(11) 'pk_cf_pool'" json:"pk_cf_pool" form:"pk_cf_pool" xml:"pk_cf_pool"`
}

// TableName table name of defined DepositOutFlow
func (m *DepositOutFlow) TableName() string {
	return "deposit_out_flow"
}

func (r *DepositOutFlow) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDepositOutFlow(data []byte) (DepositOutFlow, error) {
	var r DepositOutFlow
	err := json.Unmarshal(data, &r)
	return r, err
}
