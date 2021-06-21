// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CsChangeTeacher defined
type CsChangeTeacher struct {
	// CCTId defined
	CCTId null.Int `xorm:"int(11) pk notnull autoincr 'c_c_t_id'" json:"c_c_t_id" form:"c_c_t_id" xml:"c_c_t_id"`
	// PkCs defined
	PkCs null.Int `xorm:"int(11) 'pk_cs'" json:"pk_cs" form:"pk_cs" xml:"pk_cs"`
	// Remark defined
	Remark null.String `xorm:"varchar(500) 'remark'" json:"remark" form:"remark" xml:"remark"`
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
	// KqKc defined
	KqKc null.Float `xorm:"float(50,2) 'kq_kc'" json:"kq_kc" form:"kq_kc" xml:"kq_kc"`
}

// TableName table name of defined CsChangeTeacher
func (m *CsChangeTeacher) TableName() string {
	return "cs_change_teacher"
}
