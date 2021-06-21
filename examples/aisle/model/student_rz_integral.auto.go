// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentRzIntegral defined
type StudentRzIntegral struct {
	// SRIId defined
	SRIId null.Int `xorm:"int(11) pk notnull autoincr 's_r_i_id'" json:"s_r_i_id" form:"s_r_i_id" xml:"s_r_i_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// Integral defined
	Integral null.Float `xorm:"float(10,2) 'integral'" json:"integral" form:"integral" xml:"integral"`
	// IntegralSource defined
	IntegralSource null.Int `xorm:"int(11) 'integral_source'" json:"integral_source" form:"integral_source" xml:"integral_source"`
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
	// IntegralDesc defined
	IntegralDesc null.String `xorm:"varchar(500) 'integral_desc'" json:"integral_desc" form:"integral_desc" xml:"integral_desc"`
	// IntegralState defined
	IntegralState null.Int `xorm:"int(11) 'integral_state'" json:"integral_state" form:"integral_state" xml:"integral_state"`
	// TlsId defined
	TlsId null.Int `xorm:"int(11) 'tls_id'" json:"tls_id" form:"tls_id" xml:"tls_id"`
	// BuyOrderid defined
	BuyOrderid null.Int `xorm:"int(11) 'buy_orderid'" json:"buy_orderid" form:"buy_orderid" xml:"buy_orderid"`
	// LtorderId defined
	LtorderId null.Int `xorm:"int(11) 'ltorder_id'" json:"ltorder_id" form:"ltorder_id" xml:"ltorder_id"`
	// FeeId defined
	FeeId null.Int `xorm:"int(11) 'fee_id'" json:"fee_id" form:"fee_id" xml:"fee_id"`
	// PkOdd defined
	PkOdd null.Int `xorm:"int(11) 'pk_odd'" json:"pk_odd" form:"pk_odd" xml:"pk_odd"`
}

// TableName table name of defined StudentRzIntegral
func (m *StudentRzIntegral) TableName() string {
	return "student_rz_integral"
}
