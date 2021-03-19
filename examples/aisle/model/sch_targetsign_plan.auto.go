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

// SchTargetsignPlan defined
type SchTargetsignPlan struct {
	// STPId defined
	STPId null.Int `xorm:"int(11) pk notnull autoincr 's_t_p_id'" json:"s_t_p_id" form:"s_t_p_id" xml:"s_t_p_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// MarkInNum defined
	MarkInNum null.Float `xorm:"float(11,2) 'mark_in_num'" json:"mark_in_num" form:"mark_in_num" xml:"mark_in_num"`
	// MarkOutNum defined
	MarkOutNum null.Float `xorm:"float(50,2) 'mark_out_num'" json:"mark_out_num" form:"mark_out_num" xml:"mark_out_num"`
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
	// NetworkTargetsignNum defined
	NetworkTargetsignNum null.Float `xorm:"float(50,2) 'network_targetsign_num'" json:"network_targetsign_num" form:"network_targetsign_num" xml:"network_targetsign_num"`
	// LtTargetsignNum defined
	LtTargetsignNum null.Float `xorm:"float(50,2) 'lt_targetsign_num'" json:"lt_targetsign_num" form:"lt_targetsign_num" xml:"lt_targetsign_num"`
	// PpTargetsignNum defined
	PpTargetsignNum null.Float `xorm:"float(50,2) 'pp_targetsign_num'" json:"pp_targetsign_num" form:"pp_targetsign_num" xml:"pp_targetsign_num"`
	// QdTargetsignNum defined
	QdTargetsignNum null.Float `xorm:"float(50,2) 'qd_targetsign_num'" json:"qd_targetsign_num" form:"qd_targetsign_num" xml:"qd_targetsign_num"`
	// QtTargetsignNum defined
	QtTargetsignNum null.Float `xorm:"float(50,2) 'qt_targetsign_num'" json:"qt_targetsign_num" form:"qt_targetsign_num" xml:"qt_targetsign_num"`
	// PlanMonth defined
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" form:"plan_month" xml:"plan_month"`
	// AllQdNum defined
	AllQdNum null.Float `xorm:"float(50,2) 'all_qd_num'" json:"all_qd_num" form:"all_qd_num" xml:"all_qd_num"`
	// AllQdMoney defined
	AllQdMoney null.Float `xorm:"float(50,2) 'all_qd_money'" json:"all_qd_money" form:"all_qd_money" xml:"all_qd_money"`
	// AllTzMoney defined
	AllTzMoney null.Float `xorm:"float(50,2) 'all_tz_money'" json:"all_tz_money" form:"all_tz_money" xml:"all_tz_money"`
}

// With defined
func (m *SchTargetsignPlan) With(s interface{}) (interface{}, error) {
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
func (m *SchTargetsignPlan) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SchTargetsignPlan) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SchTargetsignPlan) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SchTargetsignPlan) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SchTargetsignPlan) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *SchTargetsignPlan) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SchTargetsignPlan
func (m *SchTargetsignPlan) TableName() string {
	return "sch_targetsign_plan"
}
