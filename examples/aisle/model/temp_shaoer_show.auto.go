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

// TempShaoerShow defined
type TempShaoerShow struct {
	// UserName defined
	UserName null.String `xorm:"varchar(255) 'user_name'" json:"user_name" form:"user_name" xml:"user_name"`
	// UserPhoneNum defined
	UserPhoneNum null.String `xorm:"varchar(255) 'user_phone_num'" json:"user_phone_num" form:"user_phone_num" xml:"user_phone_num"`
	// PhoneType defined
	PhoneType null.String `xorm:"varchar(255) 'phone_type'" json:"phone_type" form:"phone_type" xml:"phone_type"`
}

// With defined
func (m *TempShaoerShow) With(s interface{}) (interface{}, error) {
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
func (m *TempShaoerShow) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TempShaoerShow) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *TempShaoerShow) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *TempShaoerShow) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *TempShaoerShow) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *TempShaoerShow) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TempShaoerShow
func (m *TempShaoerShow) TableName() string {
	return "temp_shaoer_show"
}
