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

// JfscFlSytu defined
type JfscFlSytu struct {
	// JFSId defined
	JFSId null.Int `xorm:"int(11) pk notnull autoincr 'j_f_s_id'" json:"j_f_s_id" form:"j_f_s_id" xml:"j_f_s_id"`
	// JfscFlSy defined
	JfscFlSy null.String `xorm:"varchar(1000) 'jfsc_fl_sy'" json:"jfsc_fl_sy" form:"jfsc_fl_sy" xml:"jfsc_fl_sy"`
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
func (m *JfscFlSytu) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *JfscFlSytu) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *JfscFlSytu) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *JfscFlSytu) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined JfscFlSytu
func (m *JfscFlSytu) TableName() string {
	return "jfsc_fl_sytu"
}
