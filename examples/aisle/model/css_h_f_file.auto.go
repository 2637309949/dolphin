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

// CssHFFile defined
type CssHFFile struct {
	// CHFFId defined
	CHFFId null.Int `xorm:"int(11) pk notnull autoincr 'c_h_f_f_id'" json:"c_h_f_f_id" form:"c_h_f_f_id" xml:"c_h_f_f_id"`
	// Fileid defined
	Fileid null.Int `xorm:"int(11) 'fileid'" json:"fileid" form:"fileid" xml:"fileid"`
	// CctId defined
	CctId null.Int `xorm:"int(11) 'cct_id'" json:"cct_id" form:"cct_id" xml:"cct_id"`
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
	// ScsId defined
	ScsId null.Int `xorm:"int(11) 'scs_id'" json:"scs_id" form:"scs_id" xml:"scs_id"`
}

// Marshal defined
func (m *CssHFFile) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *CssHFFile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *CssHFFile) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *CssHFFile) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *CssHFFile) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *CssHFFile) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CssHFFile
func (m *CssHFFile) TableName() string {
	return "css_h_f_file"
}
