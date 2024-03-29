// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// RefundFlow defined
type RefundFlow struct {
	// CFlowId defined
	CFlowId null.Int `xorm:"int(11) pk notnull autoincr 'c_flow_id'" json:"c_flow_id" form:"c_flow_id" xml:"c_flow_id"`
	// PkRefundId defined
	PkRefundId null.Int `xorm:"int(11) 'pk_refund_id'" json:"pk_refund_id" form:"pk_refund_id" xml:"pk_refund_id"`
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

// TableName table name of defined RefundFlow
func (m *RefundFlow) TableName() string {
	return "refund_flow"
}

func (r *RefundFlow) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRefundFlow(data []byte) (RefundFlow, error) {
	var r RefundFlow
	err := json.Unmarshal(data, &r)
	return r, err
}
