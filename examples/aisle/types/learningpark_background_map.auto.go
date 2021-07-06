// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// LearningparkBackgroundMap defined
type LearningparkBackgroundMap struct {
	// LBMId defined
	LBMId null.Int `xorm:"int(11) pk notnull autoincr 'l_b_m_id'" json:"l_b_m_id" form:"l_b_m_id" xml:"l_b_m_id"`
	// DescName defined
	DescName null.String `xorm:"varchar(1000) 'desc_name'" json:"desc_name" form:"desc_name" xml:"desc_name"`
	// Lpbmfileid defined
	Lpbmfileid null.Int `xorm:"int(11) 'lpbmfileid'" json:"lpbmfileid" form:"lpbmfileid" xml:"lpbmfileid"`
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

// TableName table name of defined LearningparkBackgroundMap
func (m *LearningparkBackgroundMap) TableName() string {
	return "learningpark_background_map"
}