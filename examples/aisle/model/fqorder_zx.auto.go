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

// FqorderZx defined
type FqorderZx struct {
	// FqorderZxId defined
	FqorderZxId null.Int `xorm:"int(11) pk notnull autoincr 'fqorder_zx_id'" json:"fqorder_zx_id" form:"fqorder_zx_id" xml:"fqorder_zx_id"`
	// OrderId defined
	OrderId null.Int `xorm:"int(11) 'order_id'" json:"order_id" form:"order_id" xml:"order_id"`
	// ZcSchid defined
	ZcSchid null.Int `xorm:"int(11) 'zc_schid'" json:"zc_schid" form:"zc_schid" xml:"zc_schid"`
	// ZrSchid defined
	ZrSchid null.Int `xorm:"int(11) 'zr_schid'" json:"zr_schid" form:"zr_schid" xml:"zr_schid"`
	// ZxDesc defined
	ZxDesc null.String `xorm:"varchar(1000) 'zx_desc'" json:"zx_desc" form:"zx_desc" xml:"zx_desc"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// CheckUser defined
	CheckUser null.Int `xorm:"int(11) 'check_user'" json:"check_user" form:"check_user" xml:"check_user"`
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
	// AllotTa defined
	AllotTa null.Int `xorm:"int(11) 'allot_ta'" json:"allot_ta" form:"allot_ta" xml:"allot_ta"`
}

// Marshal defined
func (m *FqorderZx) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *FqorderZx) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *FqorderZx) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *FqorderZx) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *FqorderZx) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *FqorderZx) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined FqorderZx
func (m *FqorderZx) TableName() string {
	return "fqorder_zx"
}
