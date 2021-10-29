// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// TempShaoerShow defined
type TempShaoerShow struct {
	// UserName defined
	UserName null.String `xorm:"varchar(255) 'user_name'" json:"user_name" form:"user_name" xml:"user_name"`
	// UserPhoneNum defined
	UserPhoneNum null.String `xorm:"varchar(255) 'user_phone_num'" json:"user_phone_num" form:"user_phone_num" xml:"user_phone_num"`
	// PhoneType defined
	PhoneType null.String `xorm:"varchar(255) 'phone_type'" json:"phone_type" form:"phone_type" xml:"phone_type"`
}

// TableName table name of defined TempShaoerShow
func (m *TempShaoerShow) TableName() string {
	return "temp_shaoer_show"
}

func (r *TempShaoerShow) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTempShaoerShow(data []byte) (TempShaoerShow, error) {
	var r TempShaoerShow
	err := json.Unmarshal(data, &r)
	return r, err
}
