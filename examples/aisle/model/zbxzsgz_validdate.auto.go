// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// ZbxzsgzValiddate defined
type ZbxzsgzValiddate struct {
	// ZVId defined
	ZVId null.Int `xorm:"int(11) pk notnull autoincr 'z_v_id'" json:"z_v_id" form:"z_v_id" xml:"z_v_id"`
	// ZbxzsgzId defined
	ZbxzsgzId null.Int `xorm:"int(11) 'zbxzsgz_id'" json:"zbxzsgz_id" form:"zbxzsgz_id" xml:"zbxzsgz_id"`
	// BeginDate defined
	BeginDate null.Time `xorm:"datetime 'begin_date'" json:"begin_date" form:"begin_date" xml:"begin_date"`
	// EndDate defined
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" form:"end_date" xml:"end_date"`
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
	// SyhourZsbl defined
	SyhourZsbl null.Float `xorm:"float(50,2) 'syhour_zsbl'" json:"syhour_zsbl" form:"syhour_zsbl" xml:"syhour_zsbl"`
}

// Marshal defined
func (m *ZbxzsgzValiddate) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *ZbxzsgzValiddate) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ZbxzsgzValiddate) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ZbxzsgzValiddate) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ZbxzsgzValiddate) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ZbxzsgzValiddate) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ZbxzsgzValiddate) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ZbxzsgzValiddate
func (m *ZbxzsgzValiddate) TableName() string {
	return "zbxzsgz_validdate"
}
