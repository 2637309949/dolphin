// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuParent defined
type StuParent struct {
	// StuParentId defined
	StuParentId null.Int `xorm:"int(11) pk notnull autoincr 'stu_parent_id'" json:"stu_parent_id" form:"stu_parent_id" xml:"stu_parent_id"`
	// StuParentName defined
	StuParentName null.String `xorm:"varchar(1000) 'stu_parent_name'" json:"stu_parent_name" form:"stu_parent_name" xml:"stu_parent_name"`
	// StuParentPhone defined
	StuParentPhone null.String `xorm:"varchar(11) 'stu_parent_phone'" json:"stu_parent_phone" form:"stu_parent_phone" xml:"stu_parent_phone"`
	// StuParentPassword defined
	StuParentPassword null.String `xorm:"varchar(1000) 'stu_parent_password'" json:"stu_parent_password" form:"stu_parent_password" xml:"stu_parent_password"`
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
	// UserLoginState defined
	UserLoginState null.Int `xorm:"int(11) 'user_login_state'" json:"user_login_state" form:"user_login_state" xml:"user_login_state"`
	// TempPhone defined
	TempPhone null.String `xorm:"varchar(11) 'temp_phone'" json:"temp_phone" form:"temp_phone" xml:"temp_phone"`
	// JpushCodePar defined
	JpushCodePar null.String `xorm:"varchar(100) 'jpush_code_par'" json:"jpush_code_par" form:"jpush_code_par" xml:"jpush_code_par"`
	// ParAppversion defined
	ParAppversion null.String `xorm:"varchar(20) 'par_appversion'" json:"par_appversion" form:"par_appversion" xml:"par_appversion"`
	// ParAppneedupdate defined
	ParAppneedupdate null.Int `xorm:"int(11) 'par_appneedupdate'" json:"par_appneedupdate" form:"par_appneedupdate" xml:"par_appneedupdate"`
	// IdentificationNumber defined
	IdentificationNumber null.String `xorm:"varchar(1000) 'identification_number'" json:"identification_number" form:"identification_number" xml:"identification_number"`
}

// TableName table name of defined StuParent
func (m *StuParent) TableName() string {
	return "stu_parent"
}
