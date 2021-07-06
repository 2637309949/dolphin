// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ChangeTeaHisTea defined
type ChangeTeaHisTea struct {
	// CTHTId defined
	CTHTId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_h_t_id'" json:"c_t_h_t_id" form:"c_t_h_t_id" xml:"c_t_h_t_id"`
	// PkChangeTea defined
	PkChangeTea null.Int `xorm:"int(11) 'pk_change_tea'" json:"pk_change_tea" form:"pk_change_tea" xml:"pk_change_tea"`
	// PkHisTea defined
	PkHisTea null.Int `xorm:"int(11) 'pk_his_tea'" json:"pk_his_tea" form:"pk_his_tea" xml:"pk_his_tea"`
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
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// KqKc defined
	KqKc null.Float `xorm:"float(50,2) 'kq_kc'" json:"kq_kc" form:"kq_kc" xml:"kq_kc"`
	// IfKouHour defined
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
}

// TableName table name of defined ChangeTeaHisTea
func (m *ChangeTeaHisTea) TableName() string {
	return "change_tea_his_tea"
}