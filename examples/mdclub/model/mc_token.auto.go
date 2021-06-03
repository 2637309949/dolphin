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

// McToken defined 用户TOKEN
type McToken struct {
	// Token defined token 字符串
	Token null.String `xorm:"varchar(50) pk notnull default('') comment('token 字符串') 'token'" json:"token" form:"token" xml:"token"`
	// UserId defined 用户ID
	UserId null.Int `xorm:"int(11) notnull comment('用户ID') 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// Device defined 登陆设备，浏览器 UA 等信息
	Device null.String `xorm:"varchar(600) notnull default('') comment('登陆设备，浏览器 UA 等信息') 'device'" json:"device" form:"device" xml:"device"`
	// CreateTime defined 创建时间
	CreateTime null.Int `xorm:"int(10) notnull default(0) comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// UpdateTime defined 更新时间
	UpdateTime null.Int `xorm:"int(10) notnull default(0) comment('更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// ExpireTime defined 过期时间
	ExpireTime null.Int `xorm:"int(10) notnull default(0) comment('过期时间') 'expire_time'" json:"expire_time" form:"expire_time" xml:"expire_time"`
}

// With defined
func (m *McToken) With(s interface{}) (interface{}, error) {
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
func (m *McToken) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *McToken) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *McToken) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *McToken) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *McToken) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *McToken) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined McToken
func (m *McToken) TableName() string {
	return "mc_token"
}