// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/shopspring/decimal"
)

// ChannelSchoolPlan defined
type ChannelSchoolPlan struct {
	// CSPId defined
	CSPId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_p_id'" json:"c_s_p_id" form:"c_s_p_id" xml:"c_s_p_id"`
	// PkSch defined
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
	// InPlan defined
	InPlan null.Int `xorm:"int(11) 'in_plan'" json:"in_plan" form:"in_plan" xml:"in_plan"`
	// OutPlan defined
	OutPlan null.Int `xorm:"int(11) 'out_plan'" json:"out_plan" form:"out_plan" xml:"out_plan"`
	// DemoPlan defined
	DemoPlan null.Int `xorm:"int(11) 'demo_plan'" json:"demo_plan" form:"demo_plan" xml:"demo_plan"`
	// AchievementPlan defined
	AchievementPlan decimal.Decimal `xorm:"decimal(11,2) 'achievement_plan'" json:"achievement_plan" form:"achievement_plan" xml:"achievement_plan"`
	// PlanDate defined
	PlanDate null.Time `xorm:"datetime 'plan_date'" json:"plan_date" form:"plan_date" xml:"plan_date"`
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

// TableName table name of defined ChannelSchoolPlan
func (m *ChannelSchoolPlan) TableName() string {
	return "channel_school_plan"
}

func (r *ChannelSchoolPlan) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChannelSchoolPlan(data []byte) (ChannelSchoolPlan, error) {
	var r ChannelSchoolPlan
	err := json.Unmarshal(data, &r)
	return r, err
}
