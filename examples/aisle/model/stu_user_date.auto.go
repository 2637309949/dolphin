// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuUserDate defined
type StuUserDate struct {
	// SUDId defined
	SUDId null.Int `xorm:"int(11) pk notnull autoincr 's_u_d_id'" json:"s_u_d_id" form:"s_u_d_id" xml:"s_u_d_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// ZxDate defined
	ZxDate null.Time `xorm:"datetime 'zx_date'" json:"zx_date" form:"zx_date" xml:"zx_date"`
	// UserType defined
	UserType null.String `xorm:"varchar(100) 'user_type'" json:"user_type" form:"user_type" xml:"user_type"`
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

// TableName table name of defined StuUserDate
func (m *StuUserDate) TableName() string {
	return "stu_user_date"
}
