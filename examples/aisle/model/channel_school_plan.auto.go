// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
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
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *ChannelSchoolPlan) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ChannelSchoolPlan) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ChannelSchoolPlan) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ChannelSchoolPlan) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ChannelSchoolPlan) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ChannelSchoolPlan) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ChannelSchoolPlan
func (m *ChannelSchoolPlan) TableName() string {
	return "channel_school_plan"
}
