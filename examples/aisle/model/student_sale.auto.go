// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentSale defined
type StudentSale struct {
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// SSId defined
	SSId null.Int `xorm:"int(11) pk notnull autoincr 's_s_id'" json:"s_s_id" form:"s_s_id" xml:"s_s_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// SsState defined
	SsState null.Int `xorm:"int(11) 'ss_state'" json:"ss_state" form:"ss_state" xml:"ss_state"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// AllotSourse defined
	AllotSourse null.Int `xorm:"int(11) 'allot_sourse'" json:"allot_sourse" form:"allot_sourse" xml:"allot_sourse"`
	// AllotDate defined
	AllotDate null.Time `xorm:"datetime 'allot_date'" json:"allot_date" form:"allot_date" xml:"allot_date"`
	// CancelDate defined
	CancelDate null.Time `xorm:"datetime 'cancel_date'" json:"cancel_date" form:"cancel_date" xml:"cancel_date"`
	// Remark defined
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// YesornoBatchAllocation defined
	YesornoBatchAllocation null.Int `xorm:"int(11) 'yesorno_batch_allocation'" json:"yesorno_batch_allocation" form:"yesorno_batch_allocation" xml:"yesorno_batch_allocation"`
	// History defined
	History null.Int `xorm:"int(11) 'history'" json:"history" form:"history" xml:"history"`
}

// TableName table name of defined StudentSale
func (m *StudentSale) TableName() string {
	return "student_sale"
}
