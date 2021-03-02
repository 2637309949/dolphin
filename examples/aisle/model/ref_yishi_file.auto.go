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

// RefYishiFile defined
type RefYishiFile struct {
	// RYFId defined
	RYFId null.Int `xorm:"int(11) pk notnull autoincr 'r_y_f_id'" json:"r_y_f_id" form:"r_y_f_id" xml:"r_y_f_id"`
	// RefId defined
	RefId null.Int `xorm:"int(11) 'ref_id'" json:"ref_id" form:"ref_id" xml:"ref_id"`
	// Fileid defined
	Fileid null.Int `xorm:"int(11) 'fileid'" json:"fileid" form:"fileid" xml:"fileid"`
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
func (m *RefYishiFile) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *RefYishiFile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *RefYishiFile) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *RefYishiFile) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined RefYishiFile
func (m *RefYishiFile) TableName() string {
	return "ref_yishi_file"
}
