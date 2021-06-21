// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuScoreItem defined
type StuScoreItem struct {
	// SSIId defined
	SSIId null.Int `xorm:"int(11) pk notnull autoincr 's_s_i_id'" json:"s_s_i_id" form:"s_s_i_id" xml:"s_s_i_id"`
	// SsId defined
	SsId null.Int `xorm:"int(11) 'ss_id'" json:"ss_id" form:"ss_id" xml:"ss_id"`
	// SiId defined
	SiId null.Int `xorm:"int(11) 'si_id'" json:"si_id" form:"si_id" xml:"si_id"`
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
	// Score defined
	Score null.Float `xorm:"float(50,2) 'score'" json:"score" form:"score" xml:"score"`
}

// TableName table name of defined StuScoreItem
func (m *StuScoreItem) TableName() string {
	return "stu_score_item"
}
