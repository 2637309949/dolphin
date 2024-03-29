// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// OrderFromStuIntegral defined
type OrderFromStuIntegral struct {
	// OFSIId defined
	OFSIId null.Int `xorm:"int(11) pk notnull autoincr 'o_f_s_i_id'" json:"o_f_s_i_id" form:"o_f_s_i_id" xml:"o_f_s_i_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// StuIntegral defined
	StuIntegral null.Int `xorm:"int(11) 'stu_integral'" json:"stu_integral" form:"stu_integral" xml:"stu_integral"`
	// UseIntegral defined
	UseIntegral null.Float `xorm:"float(11,2) 'use_integral'" json:"use_integral" form:"use_integral" xml:"use_integral"`
	// DiscountMoney defined
	DiscountMoney null.Float `xorm:"float(11,2) 'discount_money'" json:"discount_money" form:"discount_money" xml:"discount_money"`
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

// TableName table name of defined OrderFromStuIntegral
func (m *OrderFromStuIntegral) TableName() string {
	return "order_from_stu_integral"
}

func (r *OrderFromStuIntegral) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOrderFromStuIntegral(data []byte) (OrderFromStuIntegral, error) {
	var r OrderFromStuIntegral
	err := json.Unmarshal(data, &r)
	return r, err
}
