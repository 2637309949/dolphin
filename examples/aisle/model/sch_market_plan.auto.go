// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SchMarketPlan defined
type SchMarketPlan struct {
	// SMPId defined
	SMPId null.Int `xorm:"int(11) pk notnull autoincr 's_m_p_id'" json:"s_m_p_id" form:"s_m_p_id" xml:"s_m_p_id"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// PlanMonth defined
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" form:"plan_month" xml:"plan_month"`
	// InPlan defined
	InPlan null.Float `xorm:"float(11,2) 'in_plan'" json:"in_plan" form:"in_plan" xml:"in_plan"`
	// OutPlan defined
	OutPlan null.Float `xorm:"float(11,2) 'out_plan'" json:"out_plan" form:"out_plan" xml:"out_plan"`
	// DemoPlan defined
	DemoPlan null.Float `xorm:"float(11,2) 'demo_plan'" json:"demo_plan" form:"demo_plan" xml:"demo_plan"`
	// AchievementPlan defined
	AchievementPlan null.Float `xorm:"float(11,2) 'achievement_plan'" json:"achievement_plan" form:"achievement_plan" xml:"achievement_plan"`
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
func (m *SchMarketPlan) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *SchMarketPlan) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SchMarketPlan) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SchMarketPlan) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SchMarketPlan) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SchMarketPlan) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SchMarketPlan) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SchMarketPlan
func (m *SchMarketPlan) TableName() string {
	return "sch_market_plan"
}
