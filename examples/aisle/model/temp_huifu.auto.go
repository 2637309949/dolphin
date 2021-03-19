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

// TempHuifu defined
type TempHuifu struct {
	// Fid defined
	Fid null.Int `xorm:"int(11) pk notnull 'fid'" json:"fid" form:"fid" xml:"fid"`
	// StuName defined
	StuName null.String `xorm:"varchar(255) 'stu_name'" json:"stu_name" form:"stu_name" xml:"stu_name"`
	// StuType defined
	StuType null.String `xorm:"varchar(255) 'stu_type'" json:"stu_type" form:"stu_type" xml:"stu_type"`
	// StuCity defined
	StuCity null.String `xorm:"varchar(255) 'stu_city'" json:"stu_city" form:"stu_city" xml:"stu_city"`
	// OsName defined
	OsName null.String `xorm:"varchar(255) 'os_name'" json:"os_name" form:"os_name" xml:"os_name"`
	// StuPhone defined
	StuPhone null.String `xorm:"varchar(11) 'stu_phone'" json:"stu_phone" form:"stu_phone" xml:"stu_phone"`
	// OfNumber defined
	OfNumber null.String `xorm:"varchar(255) 'of_number'" json:"of_number" form:"of_number" xml:"of_number"`
}

// With defined
func (m *TempHuifu) With(s interface{}) (interface{}, error) {
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
func (m *TempHuifu) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TempHuifu) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *TempHuifu) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *TempHuifu) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *TempHuifu) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *TempHuifu) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TempHuifu
func (m *TempHuifu) TableName() string {
	return "temp_huifu"
}
