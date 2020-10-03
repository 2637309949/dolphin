// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ChangeTeaHisTea defined
type ChangeTeaHisTea struct {
	//
	CTHTId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_h_t_id'" json:"c_t_h_t_id" xml:"c_t_h_t_id"`
	//
	PkChangeTea null.Int `xorm:"int(11) 'pk_change_tea'" json:"pk_change_tea" xml:"pk_change_tea"`
	//
	PkHisTea null.Int `xorm:"int(11) 'pk_his_tea'" json:"pk_his_tea" xml:"pk_his_tea"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" xml:"start_time"`
	//
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" xml:"end_time"`
	//
	KqKc null.Float `xorm:"float(50,2) 'kq_kc'" json:"kq_kc" xml:"kq_kc"`
	//
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" xml:"if_kou_hour"`
}

// TableName table name of defined ChangeTeaHisTea
func (m *ChangeTeaHisTea) TableName() string {
	return "change_tea_his_tea"
}