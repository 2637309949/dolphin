// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/shopspring/decimal"
)

// NetUserPlan defined
type NetUserPlan struct {
	// NUPId defined
	NUPId null.Int `xorm:"int(11) pk notnull autoincr 'n_u_p_id'" json:"n_u_p_id" form:"n_u_p_id" xml:"n_u_p_id"`
	// PkSch defined
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
	// PkPlaner defined
	PkPlaner null.Int `xorm:"int(11) 'pk_planer'" json:"pk_planer" form:"pk_planer" xml:"pk_planer"`
	// InPlan defined
	InPlan null.Int `xorm:"int(11) 'in_plan'" json:"in_plan" form:"in_plan" xml:"in_plan"`
	// OutPlan defined
	OutPlan null.Int `xorm:"int(11) 'out_plan'" json:"out_plan" form:"out_plan" xml:"out_plan"`
	// DemoPlan defined
	DemoPlan null.Int `xorm:"int(11) 'demo_plan'" json:"demo_plan" form:"demo_plan" xml:"demo_plan"`
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
	// AchievementPlan defined
	AchievementPlan decimal.Decimal `xorm:"decimal(50,3) 'achievement_plan'" json:"achievement_plan" form:"achievement_plan" xml:"achievement_plan"`
	// PlanMonth defined
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" form:"plan_month" xml:"plan_month"`
}

// TableName table name of defined NetUserPlan
func (m *NetUserPlan) TableName() string {
	return "net_user_plan"
}