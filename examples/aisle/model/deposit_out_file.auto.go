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

// DepositOutFile defined
type DepositOutFile struct {
	// DOFId defined
	DOFId null.Int `xorm:"int(11) pk notnull autoincr 'd_o_f_id'" json:"d_o_f_id" form:"d_o_f_id" xml:"d_o_f_id"`
	// DoId defined
	DoId null.Int `xorm:"int(11) 'do_id'" json:"do_id" form:"do_id" xml:"do_id"`
	// FileId defined
	FileId null.Int `xorm:"int(11) 'file_id'" json:"file_id" form:"file_id" xml:"file_id"`
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
func (m *DepositOutFile) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *DepositOutFile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *DepositOutFile) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *DepositOutFile) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined DepositOutFile
func (m *DepositOutFile) TableName() string {
	return "deposit_out_file"
}
