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

// StuKwJfx defined
type StuKwJfx struct {
	// StuKwJfxId defined
	StuKwJfxId null.Int `xorm:"int(11) pk notnull autoincr 'stu_kw_jfx_id'" json:"stu_kw_jfx_id" form:"stu_kw_jfx_id" xml:"stu_kw_jfx_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// JfxDate defined
	JfxDate null.Time `xorm:"datetime 'jfx_date'" json:"jfx_date" form:"jfx_date" xml:"jfx_date"`
	// Jfx defined
	Jfx null.Int `xorm:"int(11) 'jfx'" json:"jfx" form:"jfx" xml:"jfx"`
	// JfxNum defined
	JfxNum null.Int `xorm:"int(11) 'jfx_num'" json:"jfx_num" form:"jfx_num" xml:"jfx_num"`
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
func (m *StuKwJfx) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StuKwJfx) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StuKwJfx) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StuKwJfx) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StuKwJfx) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuKwJfx) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuKwJfx
func (m *StuKwJfx) TableName() string {
	return "stu_kw_jfx"
}
