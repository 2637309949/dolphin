// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CsVisit defined
type CsVisit struct {
	// CsVisitId defined
	CsVisitId null.Int `xorm:"int(11) pk notnull autoincr 'cs_visit_id'" json:"cs_visit_id" form:"cs_visit_id" xml:"cs_visit_id"`
	// VisitContent defined
	VisitContent null.String `xorm:"varchar(300) 'visit_content'" json:"visit_content" form:"visit_content" xml:"visit_content"`
	// VisitStartTime defined
	VisitStartTime null.Time `xorm:"datetime 'visit_start_time'" json:"visit_start_time" form:"visit_start_time" xml:"visit_start_time"`
	// VisitEndTime defined
	VisitEndTime null.Time `xorm:"datetime 'visit_end_time'" json:"visit_end_time" form:"visit_end_time" xml:"visit_end_time"`
	// NextVisitTime defined
	NextVisitTime null.Time `xorm:"datetime 'next_visit_time'" json:"next_visit_time" form:"next_visit_time" xml:"next_visit_time"`
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
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// VisituserId defined
	VisituserId null.Int `xorm:"int(11) 'visituser_id'" json:"visituser_id" form:"visituser_id" xml:"visituser_id"`
}

// TableName table name of defined CsVisit
func (m *CsVisit) TableName() string {
	return "cs_visit"
}
