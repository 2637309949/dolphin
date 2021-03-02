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

// OnSaleFiles defined
type OnSaleFiles struct {
	// OSFId defined
	OSFId null.Int `xorm:"int(11) pk notnull autoincr 'o_s_f_id'" json:"o_s_f_id" form:"o_s_f_id" xml:"o_s_f_id"`
	// OnSaleId defined
	OnSaleId null.Int `xorm:"int(11) 'on_sale_id'" json:"on_sale_id" form:"on_sale_id" xml:"on_sale_id"`
	// Files defined
	Files null.Int `xorm:"int(11) 'files'" json:"files" form:"files" xml:"files"`
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
func (m *OnSaleFiles) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *OnSaleFiles) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *OnSaleFiles) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *OnSaleFiles) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OnSaleFiles
func (m *OnSaleFiles) TableName() string {
	return "on_sale_files"
}
