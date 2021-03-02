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

// EcUserTargetplan defined
type EcUserTargetplan struct {
	// EUTId defined
	EUTId null.Int `xorm:"int(11) pk notnull autoincr 'e_u_t_id'" json:"e_u_t_id" form:"e_u_t_id" xml:"e_u_t_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// PlanMonth defined
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" form:"plan_month" xml:"plan_month"`
	// MbSignNum defined
	MbSignNum null.Float `xorm:"float(11,2) 'mb_sign_num'" json:"mb_sign_num" form:"mb_sign_num" xml:"mb_sign_num"`
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
	// UserName defined
	UserName null.String `xorm:"varchar(100) 'user_name'" json:"user_name" form:"user_name" xml:"user_name"`
	// UserPhone defined
	UserPhone null.String `xorm:"varchar(11) 'user_phone'" json:"user_phone" form:"user_phone" xml:"user_phone"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// MonthTask defined
	MonthTask null.Float `xorm:"float(50,2) 'month_task'" json:"month_task" form:"month_task" xml:"month_task"`
}

// Marshal defined
func (m *EcUserTargetplan) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *EcUserTargetplan) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *EcUserTargetplan) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *EcUserTargetplan) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined EcUserTargetplan
func (m *EcUserTargetplan) TableName() string {
	return "ec_user_targetplan"
}
