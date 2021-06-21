// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ClassYpSc defined
type ClassYpSc struct {
	// CYSId defined
	CYSId null.Int `xorm:"int(11) pk notnull autoincr 'c_y_s_id'" json:"c_y_s_id" form:"c_y_s_id" xml:"c_y_s_id"`
	// PkClass defined
	PkClass null.Int `xorm:"int(11) 'pk_class'" json:"pk_class" form:"pk_class" xml:"pk_class"`
	// Kc defined
	Kc null.Float `xorm:"float(11,2) 'kc'" json:"kc" form:"kc" xml:"kc"`
	// ClassDate defined
	ClassDate null.Time `xorm:"datetime 'class_date'" json:"class_date" form:"class_date" xml:"class_date"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// ClassWeek defined
	ClassWeek null.Int `xorm:"int(11) 'class_week'" json:"class_week" form:"class_week" xml:"class_week"`
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
	// PkSch defined
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
}

// TableName table name of defined ClassYpSc
func (m *ClassYpSc) TableName() string {
	return "class_yp_sc"
}
