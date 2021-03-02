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

// CcRewardUser defined
type CcRewardUser struct {
	// CRUId defined
	CRUId null.Int `xorm:"int(11) pk notnull autoincr 'c_r_u_id'" json:"c_r_u_id" form:"c_r_u_id" xml:"c_r_u_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// NdpdId defined
	NdpdId null.Int `xorm:"int(11) 'ndpd_id'" json:"ndpd_id" form:"ndpd_id" xml:"ndpd_id"`
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
	// CcMoney defined
	CcMoney null.Float `xorm:"float(10,2) 'cc_money'" json:"cc_money" form:"cc_money" xml:"cc_money"`
}

// Marshal defined
func (m *CcRewardUser) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *CcRewardUser) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *CcRewardUser) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *CcRewardUser) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CcRewardUser
func (m *CcRewardUser) TableName() string {
	return "cc_reward_user"
}
