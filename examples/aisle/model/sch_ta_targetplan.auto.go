// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SchTaTargetplan defined
type SchTaTargetplan struct {
	// STTId defined
	STTId null.Int `xorm:"int(11) pk notnull autoincr 's_t_t_id'" json:"s_t_t_id" form:"s_t_t_id" xml:"s_t_t_id"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// PlanMonth defined
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" form:"plan_month" xml:"plan_month"`
	// YxTarget defined
	YxTarget null.Float `xorm:"float(11,2) 'yx_target'" json:"yx_target" form:"yx_target" xml:"yx_target"`
	// TqxTarget defined
	TqxTarget null.Float `xorm:"float(11,2) 'tqx_target'" json:"tqx_target" form:"tqx_target" xml:"tqx_target"`
	// ZbTarget defined
	ZbTarget null.Float `xorm:"float(11,2) 'zb_target'" json:"zb_target" form:"zb_target" xml:"zb_target"`
	// LtQdNum defined
	LtQdNum null.Float `xorm:"float(11,2) 'lt_qd_num'" json:"lt_qd_num" form:"lt_qd_num" xml:"lt_qd_num"`
	// LtZyNum defined
	LtZyNum null.Float `xorm:"float(11,2) 'lt_zy_num'" json:"lt_zy_num" form:"lt_zy_num" xml:"lt_zy_num"`
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
func (m *SchTaTargetplan) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SchTaTargetplan) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *SchTaTargetplan) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SchTaTargetplan) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SchTaTargetplan
func (m *SchTaTargetplan) TableName() string {
	return "sch_ta_targetplan"
}
