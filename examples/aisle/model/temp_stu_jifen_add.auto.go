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

// TempStuJifenAdd defined
type TempStuJifenAdd struct {
	// Fid defined
	Fid null.Int `xorm:"int(255) pk notnull 'fid'" json:"fid" form:"fid" xml:"fid"`
	// StuName defined
	StuName null.String `xorm:"varchar(255) 'stu_name'" json:"stu_name" form:"stu_name" xml:"stu_name"`
	// StuPhone defined
	StuPhone null.String `xorm:"varchar(255) 'stu_phone'" json:"stu_phone" form:"stu_phone" xml:"stu_phone"`
	// ZjJifen defined
	ZjJifen null.Int `xorm:"int(11) 'zj_jifen'" json:"zj_jifen" form:"zj_jifen" xml:"zj_jifen"`
	// Remark defined
	Remark null.String `xorm:"varchar(255) 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// Marshal defined
func (m *TempStuJifenAdd) With(s interface{}) (interface{}, error) {
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
func (m *TempStuJifenAdd) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TempStuJifenAdd) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *TempStuJifenAdd) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *TempStuJifenAdd) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *TempStuJifenAdd) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TempStuJifenAdd) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TempStuJifenAdd
func (m *TempStuJifenAdd) TableName() string {
	return "temp_stu_jifen_add"
}
