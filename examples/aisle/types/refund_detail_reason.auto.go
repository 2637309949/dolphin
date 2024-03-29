// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// RefundDetailReason defined
type RefundDetailReason struct {
	// RDRId defined
	RDRId null.Int `xorm:"int(11) pk notnull autoincr 'r_d_r_id'" json:"r_d_r_id" form:"r_d_r_id" xml:"r_d_r_id"`
	// DetailReason defined
	DetailReason null.String `xorm:"varchar(1000) 'detail_reason'" json:"detail_reason" form:"detail_reason" xml:"detail_reason"`
	// OrderNum defined
	OrderNum null.Int `xorm:"int(11) 'order_num'" json:"order_num" form:"order_num" xml:"order_num"`
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
	// OnOff defined
	OnOff null.Int `xorm:"int(11) 'on_off'" json:"on_off" form:"on_off" xml:"on_off"`
}

// TableName table name of defined RefundDetailReason
func (m *RefundDetailReason) TableName() string {
	return "refund_detail_reason"
}

func (r *RefundDetailReason) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRefundDetailReason(data []byte) (RefundDetailReason, error) {
	var r RefundDetailReason
	err := json.Unmarshal(data, &r)
	return r, err
}
