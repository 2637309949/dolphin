// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// TempHuifu defined
type TempHuifu struct {
	// Fid defined
	Fid null.Int `xorm:"int(11) pk notnull 'fid'" json:"fid" form:"fid" xml:"fid"`
	// StuName defined
	StuName null.String `xorm:"varchar(255) 'stu_name'" json:"stu_name" form:"stu_name" xml:"stu_name"`
	// StuType defined
	StuType null.String `xorm:"varchar(255) 'stu_type'" json:"stu_type" form:"stu_type" xml:"stu_type"`
	// StuCity defined
	StuCity null.String `xorm:"varchar(255) 'stu_city'" json:"stu_city" form:"stu_city" xml:"stu_city"`
	// OsName defined
	OsName null.String `xorm:"varchar(255) 'os_name'" json:"os_name" form:"os_name" xml:"os_name"`
	// StuPhone defined
	StuPhone null.String `xorm:"varchar(11) 'stu_phone'" json:"stu_phone" form:"stu_phone" xml:"stu_phone"`
	// OfNumber defined
	OfNumber null.String `xorm:"varchar(255) 'of_number'" json:"of_number" form:"of_number" xml:"of_number"`
}

// TableName table name of defined TempHuifu
func (m *TempHuifu) TableName() string {
	return "temp_huifu"
}
