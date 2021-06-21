// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// NetworkActivity defined
type NetworkActivity struct {
	// NAId defined
	NAId null.Int `xorm:"int(11) pk notnull autoincr 'n_a_id'" json:"n_a_id" form:"n_a_id" xml:"n_a_id"`
	// ActivityName defined
	ActivityName null.String `xorm:"varchar(50) 'activity_name'" json:"activity_name" form:"activity_name" xml:"activity_name"`
	// NdId defined
	NdId null.Int `xorm:"int(11) 'nd_id'" json:"nd_id" form:"nd_id" xml:"nd_id"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// MaUseMoney defined
	MaUseMoney null.Float `xorm:"float(11,2) 'ma_use_money'" json:"ma_use_money" form:"ma_use_money" xml:"ma_use_money"`
	// ActivityExposure defined
	ActivityExposure null.Float `xorm:"float(11,2) 'activity_exposure'" json:"activity_exposure" form:"activity_exposure" xml:"activity_exposure"`
	// ActivityContent defined
	ActivityContent null.String `xorm:"varchar(500) 'activity_content'" json:"activity_content" form:"activity_content" xml:"activity_content"`
	// ActivityDesc defined
	ActivityDesc null.String `xorm:"varchar(500) 'activity_desc'" json:"activity_desc" form:"activity_desc" xml:"activity_desc"`
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
	// ActivityEstimateIncome defined
	ActivityEstimateIncome null.Float `xorm:"float(10,2) 'activity_estimate_income'" json:"activity_estimate_income" form:"activity_estimate_income" xml:"activity_estimate_income"`
}

// TableName table name of defined NetworkActivity
func (m *NetworkActivity) TableName() string {
	return "network_activity"
}
