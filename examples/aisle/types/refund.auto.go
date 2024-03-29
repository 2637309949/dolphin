// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// Refund defined
type Refund struct {
	// RefundId defined
	RefundId null.Int `xorm:"int(11) pk notnull autoincr 'refund_id'" json:"refund_id" form:"refund_id" xml:"refund_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// RefMoney defined
	RefMoney null.Float `xorm:"float(10,2) 'ref_money'" json:"ref_money" form:"ref_money" xml:"ref_money"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// RefWay defined
	RefWay null.Int `xorm:"int(11) 'ref_way'" json:"ref_way" form:"ref_way" xml:"ref_way"`
	// RefReason defined
	RefReason null.Int `xorm:"int(11) 'ref_reason'" json:"ref_reason" form:"ref_reason" xml:"ref_reason"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// CheckRenmark defined
	CheckRenmark null.String `xorm:"varchar(500) 'check_renmark'" json:"check_renmark" form:"check_renmark" xml:"check_renmark"`
	// RefRemark defined
	RefRemark null.String `xorm:"varchar(500) 'ref_remark'" json:"ref_remark" form:"ref_remark" xml:"ref_remark"`
	// RefType defined
	RefType null.Int `xorm:"int(11) 'ref_type'" json:"ref_type" form:"ref_type" xml:"ref_type"`
	// ZcStuId defined
	ZcStuId null.Int `xorm:"int(11) 'zc_stu_id'" json:"zc_stu_id" form:"zc_stu_id" xml:"zc_stu_id"`
	// RefundType defined
	RefundType null.Int `xorm:"int(11) 'refund_type'" json:"refund_type" form:"refund_type" xml:"refund_type"`
	// YhCount defined
	YhCount null.String `xorm:"varchar(50) 'yh_count'" json:"yh_count" form:"yh_count" xml:"yh_count"`
	// YhName defined
	YhName null.String `xorm:"varchar(1000) 'yh_name'" json:"yh_name" form:"yh_name" xml:"yh_name"`
	// CheckUserId defined
	CheckUserId null.Int `xorm:"int(11) 'check_user_id'" json:"check_user_id" form:"check_user_id" xml:"check_user_id"`
	// CheckTime defined
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" form:"check_time" xml:"check_time"`
	// ActualRefundTime defined
	ActualRefundTime null.Time `xorm:"datetime 'actual_refund_time'" json:"actual_refund_time" form:"actual_refund_time" xml:"actual_refund_time"`
	// BookFeeDeduction defined
	BookFeeDeduction null.Float `xorm:"float(50,2) 'book_fee_deduction'" json:"book_fee_deduction" form:"book_fee_deduction" xml:"book_fee_deduction"`
	// GiftFeeDeduction defined
	GiftFeeDeduction null.Float `xorm:"float(50,2) 'gift_fee_deduction'" json:"gift_fee_deduction" form:"gift_fee_deduction" xml:"gift_fee_deduction"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// SbtId defined
	SbtId null.Int `xorm:"int(11) 'sbt_id'" json:"sbt_id" form:"sbt_id" xml:"sbt_id"`
	// ZdCheckState defined
	ZdCheckState null.Int `xorm:"int(11) 'zd_check_state'" json:"zd_check_state" form:"zd_check_state" xml:"zd_check_state"`
	// CheckDate defined
	CheckDate null.Time `xorm:"datetime 'check_date'" json:"check_date" form:"check_date" xml:"check_date"`
	// PkFlowSet defined
	PkFlowSet null.Int `xorm:"int(11) 'pk_flow_set'" json:"pk_flow_set" form:"pk_flow_set" xml:"pk_flow_set"`
	// NowCheckUser defined
	NowCheckUser null.String `xorm:"varchar(500) 'now_check_user'" json:"now_check_user" form:"now_check_user" xml:"now_check_user"`
	// NowFloor defined
	NowFloor null.Int `xorm:"int(11) 'now_floor'" json:"now_floor" form:"now_floor" xml:"now_floor"`
	// TurnFloor defined
	TurnFloor null.Int `xorm:"int(11) 'turn_floor'" json:"turn_floor" form:"turn_floor" xml:"turn_floor"`
	// HistoryCheckUser defined
	HistoryCheckUser null.String `xorm:"varchar(500) 'history_check_user'" json:"history_check_user" form:"history_check_user" xml:"history_check_user"`
	// PkCfpId defined
	PkCfpId null.Int `xorm:"int(11) 'pk_cfp_id'" json:"pk_cfp_id" form:"pk_cfp_id" xml:"pk_cfp_id"`
	// PkCheckUser defined
	PkCheckUser null.Int `xorm:"int(11) 'pk_check_user'" json:"pk_check_user" form:"pk_check_user" xml:"pk_check_user"`
	// RefApplyFile defined
	RefApplyFile null.Int `xorm:"int(11) 'ref_apply_file'" json:"ref_apply_file" form:"ref_apply_file" xml:"ref_apply_file"`
	// ParentFile defined
	ParentFile null.Int `xorm:"int(11) 'parent_file'" json:"parent_file" form:"parent_file" xml:"parent_file"`
	// RefReductionFile defined
	RefReductionFile null.Int `xorm:"int(11) 'ref_reduction_file'" json:"ref_reduction_file" form:"ref_reduction_file" xml:"ref_reduction_file"`
	// RefDisclaimerFile defined
	RefDisclaimerFile null.Int `xorm:"int(11) 'ref_disclaimer_file'" json:"ref_disclaimer_file" form:"ref_disclaimer_file" xml:"ref_disclaimer_file"`
	// RefHn defined
	RefHn null.String `xorm:"varchar(1000) 'ref_hn'" json:"ref_hn" form:"ref_hn" xml:"ref_hn"`
	// RefActualmoney defined
	RefActualmoney null.Float `xorm:"float(50,2) 'ref_actualmoney'" json:"ref_actualmoney" form:"ref_actualmoney" xml:"ref_actualmoney"`
	// RefParentbackFile defined
	RefParentbackFile null.Int `xorm:"int(11) 'ref_parentback_file'" json:"ref_parentback_file" form:"ref_parentback_file" xml:"ref_parentback_file"`
	// NowCfid defined
	NowCfid null.Int `xorm:"int(11) 'now_cfid'" json:"now_cfid" form:"now_cfid" xml:"now_cfid"`
	// NextCfid defined
	NextCfid null.Int `xorm:"int(11) 'next_cfid'" json:"next_cfid" form:"next_cfid" xml:"next_cfid"`
	// IfRelief defined
	IfRelief null.Int `xorm:"int(11) 'if_relief'" json:"if_relief" form:"if_relief" xml:"if_relief"`
	// OtherRefund defined
	OtherRefund null.Float `xorm:"float(50,2) 'other_refund'" json:"other_refund" form:"other_refund" xml:"other_refund"`
	// NowCheckid defined
	NowCheckid null.String `xorm:"varchar(1000) 'now_checkid'" json:"now_checkid" form:"now_checkid" xml:"now_checkid"`
	// NowCheckname defined
	NowCheckname null.String `xorm:"varchar(1000) 'now_checkname'" json:"now_checkname" form:"now_checkname" xml:"now_checkname"`
	// NextCheckid defined
	NextCheckid null.String `xorm:"varchar(1000) 'next_checkid'" json:"next_checkid" form:"next_checkid" xml:"next_checkid"`
	// NextCheckname defined
	NextCheckname null.String `xorm:"varchar(1000) 'next_checkname'" json:"next_checkname" form:"next_checkname" xml:"next_checkname"`
	// CheckServiceRate defined
	CheckServiceRate null.Float `xorm:"float(50,2) 'check_service_rate'" json:"check_service_rate" form:"check_service_rate" xml:"check_service_rate"`
	// CheckServiceMoney defined
	CheckServiceMoney null.Float `xorm:"float(50,2) 'check_service_money'" json:"check_service_money" form:"check_service_money" xml:"check_service_money"`
	// CheckInvoiceRate defined
	CheckInvoiceRate null.Float `xorm:"float(50,2) 'check_invoice_rate'" json:"check_invoice_rate" form:"check_invoice_rate" xml:"check_invoice_rate"`
	// CheckInvoiceMoney defined
	CheckInvoiceMoney null.Float `xorm:"float(50,2) 'check_invoice_money'" json:"check_invoice_money" form:"check_invoice_money" xml:"check_invoice_money"`
	// CheckIfkouBook defined
	CheckIfkouBook null.Int `xorm:"int(11) 'check_ifkou_book'" json:"check_ifkou_book" form:"check_ifkou_book" xml:"check_ifkou_book"`
	// CheckKouBookmoney defined
	CheckKouBookmoney null.Float `xorm:"float(50,2) 'check_kou_bookmoney'" json:"check_kou_bookmoney" form:"check_kou_bookmoney" xml:"check_kou_bookmoney"`
	// CheckIfkouGift defined
	CheckIfkouGift null.Int `xorm:"int(11) 'check_ifkou_gift'" json:"check_ifkou_gift" form:"check_ifkou_gift" xml:"check_ifkou_gift"`
	// CheckKouGiftmoney defined
	CheckKouGiftmoney null.Float `xorm:"float(50,2) 'check_kou_giftmoney'" json:"check_kou_giftmoney" form:"check_kou_giftmoney" xml:"check_kou_giftmoney"`
	// CheckKouOrtermoney defined
	CheckKouOrtermoney null.Float `xorm:"float(50,2) 'check_kou_ortermoney'" json:"check_kou_ortermoney" form:"check_kou_ortermoney" xml:"check_kou_ortermoney"`
	// CheckActualRefmoney defined
	CheckActualRefmoney null.Float `xorm:"float(50,2) 'check_actual_refmoney'" json:"check_actual_refmoney" form:"check_actual_refmoney" xml:"check_actual_refmoney"`
	// RefApplyFile2 defined
	RefApplyFile2 null.Int `xorm:"int(11) 'ref_apply_file2'" json:"ref_apply_file2" form:"ref_apply_file2" xml:"ref_apply_file2"`
	// RefHour defined
	RefHour null.Float `xorm:"float(50,2) 'ref_hour'" json:"ref_hour" form:"ref_hour" xml:"ref_hour"`
	// RefBeforeMoney defined
	RefBeforeMoney null.Float `xorm:"float(50,2) 'ref_before_money'" json:"ref_before_money" form:"ref_before_money" xml:"ref_before_money"`
	// RdrId defined
	RdrId null.Int `xorm:"int(11) 'rdr_id'" json:"rdr_id" form:"rdr_id" xml:"rdr_id"`
	// SxfCheckLevel defined
	SxfCheckLevel null.Int `xorm:"int(11) 'sxf_check_level'" json:"sxf_check_level" form:"sxf_check_level" xml:"sxf_check_level"`
	// MzCheckLevel defined
	MzCheckLevel null.Int `xorm:"int(11) 'mz_check_level'" json:"mz_check_level" form:"mz_check_level" xml:"mz_check_level"`
	// Version defined
	Version null.Int `xorm:"int(11) 'version'" json:"version" form:"version" xml:"version"`
}

// TableName table name of defined Refund
func (m *Refund) TableName() string {
	return "refund"
}

func (r *Refund) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRefund(data []byte) (Refund, error) {
	var r Refund
	err := json.Unmarshal(data, &r)
	return r, err
}
