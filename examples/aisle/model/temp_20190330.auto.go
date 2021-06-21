// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Temp20190330 defined
type Temp20190330 struct {
	// OrderFormId defined
	OrderFormId null.Int `xorm:"int(11) notnull default(0) 'order_form_id'" json:"order_form_id" form:"order_form_id" xml:"order_form_id"`
	// OfNumber defined
	OfNumber null.String `xorm:"varchar(100) 'of_number'" json:"of_number" form:"of_number" xml:"of_number"`
	// TeachingMateriaAmount defined
	TeachingMateriaAmount null.Float `xorm:"float(10,2) 'teaching_materia_amount'" json:"teaching_materia_amount" form:"teaching_materia_amount" xml:"teaching_materia_amount"`
	// TotalAch defined
	TotalAch null.Float `xorm:"double(19,2) 'total_ach'" json:"total_ach" form:"total_ach" xml:"total_ach"`
	// MustMoney defined
	MustMoney null.Float `xorm:"float(10,2) 'must_money'" json:"must_money" form:"must_money" xml:"must_money"`
	// GetMoney defined
	GetMoney null.Float `xorm:"double(19,2) 'get_money'" json:"get_money" form:"get_money" xml:"get_money"`
	// Danjia defined
	Danjia null.Float `xorm:"double(19,2) 'danjia'" json:"danjia" form:"danjia" xml:"danjia"`
	// SCTId defined
	SCTId null.Int `xorm:"int(11) notnull default(0) 's_c_t_id'" json:"s_c_t_id" form:"s_c_t_id" xml:"s_c_t_id"`
	// AllPrice defined
	AllPrice null.Float `xorm:"float(11,2) 'all_price'" json:"all_price" form:"all_price" xml:"all_price"`
	// OnePrice defined
	OnePrice null.Float `xorm:"float(10,2) 'one_price'" json:"one_price" form:"one_price" xml:"one_price"`
	// OsMoney defined
	OsMoney null.Float `xorm:"float(10,2) 'os_money'" json:"os_money" form:"os_money" xml:"os_money"`
	// FinalMoney defined
	FinalMoney null.Float `xorm:"float(10,2) 'final_money'" json:"final_money" form:"final_money" xml:"final_money"`
	// RefundOnePrice defined
	RefundOnePrice null.Float `xorm:"float(10,2) 'refund_one_price'" json:"refund_one_price" form:"refund_one_price" xml:"refund_one_price"`
	// BuyHour defined
	BuyHour null.Float `xorm:"float(11,2) 'buy_hour'" json:"buy_hour" form:"buy_hour" xml:"buy_hour"`
	// UseHour defined
	UseHour null.Float `xorm:"float(10,2) 'use_hour'" json:"use_hour" form:"use_hour" xml:"use_hour"`
	// SurplusHour defined
	SurplusHour null.Float `xorm:"float(10,2) 'surplus_hour'" json:"surplus_hour" form:"surplus_hour" xml:"surplus_hour"`
	// WxhMoney defined
	WxhMoney null.Float `xorm:"float(10,2) 'wxh_money'" json:"wxh_money" form:"wxh_money" xml:"wxh_money"`
	// XhMoney defined
	XhMoney null.Float `xorm:"float(10,2) 'xh_money'" json:"xh_money" form:"xh_money" xml:"xh_money"`
}

// TableName table name of defined Temp20190330
func (m *Temp20190330) TableName() string {
	return "temp_20190330"
}
