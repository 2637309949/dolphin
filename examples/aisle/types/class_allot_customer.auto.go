// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ClassAllotCustomer defined
type ClassAllotCustomer struct {
	// CACId defined
	CACId null.Int `xorm:"int(11) pk notnull autoincr 'c_a_c_id'" json:"c_a_c_id" form:"c_a_c_id" xml:"c_a_c_id"`
	// ClassId defined
	ClassId null.Int `xorm:"int(11) 'class_id'" json:"class_id" form:"class_id" xml:"class_id"`
	// CustomerId defined
	CustomerId null.Int `xorm:"int(11) 'customer_id'" json:"customer_id" form:"customer_id" xml:"customer_id"`
	// AllotState defined
	AllotState null.Int `xorm:"int(11) 'allot_state'" json:"allot_state" form:"allot_state" xml:"allot_state"`
	// AllotCancelTime defined
	AllotCancelTime null.Time `xorm:"datetime 'allot_cancel_time'" json:"allot_cancel_time" form:"allot_cancel_time" xml:"allot_cancel_time"`
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
	// HiscustomerId defined
	HiscustomerId null.Int `xorm:"int(11) 'hiscustomer_id'" json:"hiscustomer_id" form:"hiscustomer_id" xml:"hiscustomer_id"`
	// CacDesc defined
	CacDesc null.String `xorm:"varchar(500) 'cac_desc'" json:"cac_desc" form:"cac_desc" xml:"cac_desc"`
}

// TableName table name of defined ClassAllotCustomer
func (m *ClassAllotCustomer) TableName() string {
	return "class_allot_customer"
}

func (r *ClassAllotCustomer) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClassAllotCustomer(data []byte) (ClassAllotCustomer, error) {
	var r ClassAllotCustomer
	err := json.Unmarshal(data, &r)
	return r, err
}
