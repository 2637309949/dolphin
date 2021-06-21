// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SaleLanguageModel defined
type SaleLanguageModel struct {
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// BusinessType defined
	BusinessType null.Int `xorm:"int(11) 'business_type'" json:"business_type" form:"business_type" xml:"business_type"`
	// SLMId defined
	SLMId null.Int `xorm:"int(11) pk notnull autoincr 's_l_m_id'" json:"s_l_m_id" form:"s_l_m_id" xml:"s_l_m_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// PtId defined
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" form:"pt_id" xml:"pt_id"`
	// SlmName defined
	SlmName null.String `xorm:"varchar(50) 'slm_name'" json:"slm_name" form:"slm_name" xml:"slm_name"`
	// SlmContent defined
	SlmContent null.String `xorm:"varchar(500) 'slm_content'" json:"slm_content" form:"slm_content" xml:"slm_content"`
	// VisitStage defined
	VisitStage null.Int `xorm:"int(11) 'visit_stage'" json:"visit_stage" form:"visit_stage" xml:"visit_stage"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// SlmType defined
	SlmType null.Int `xorm:"int(11) 'slm_type'" json:"slm_type" form:"slm_type" xml:"slm_type"`
	// StuSystemSta defined
	StuSystemSta null.Int `xorm:"int(11) 'stu_system_sta'" json:"stu_system_sta" form:"stu_system_sta" xml:"stu_system_sta"`
}

// TableName table name of defined SaleLanguageModel
func (m *SaleLanguageModel) TableName() string {
	return "sale_language_model"
}
