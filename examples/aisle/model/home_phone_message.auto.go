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

// HomePhoneMessage defined
type HomePhoneMessage struct {
	// T2820 defined
	T2820 null.Int `xorm:"int(11) pk notnull autoincr 't_282_0'" json:"t_282_0" form:"t_282_0" xml:"t_282_0"`
	// Title defined
	Title null.String `xorm:"varchar(500) 'title'" json:"title" form:"title" xml:"title"`
	// Content defined
	Content null.String `xorm:"varchar(2000) 'content'" json:"content" form:"content" xml:"content"`
	// PhoneNumber defined
	PhoneNumber null.String `xorm:"varchar(11) 'phone_number'" json:"phone_number" form:"phone_number" xml:"phone_number"`
	// Sendtime defined
	Sendtime null.Time `xorm:"datetime 'sendtime'" json:"sendtime" form:"sendtime" xml:"sendtime"`
	// IfSend defined
	IfSend null.Int `xorm:"int(11) 'if_send'" json:"if_send" form:"if_send" xml:"if_send"`
	// IfSendcn defined
	IfSendcn null.String `xorm:"varchar(50) 'if_sendcn'" json:"if_sendcn" form:"if_sendcn" xml:"if_sendcn"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
}

// Marshal defined
func (m *HomePhoneMessage) With(s interface{}) (interface{}, error) {
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
func (m *HomePhoneMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *HomePhoneMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *HomePhoneMessage) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *HomePhoneMessage) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *HomePhoneMessage) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *HomePhoneMessage) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined HomePhoneMessage
func (m *HomePhoneMessage) TableName() string {
	return "home_phone_message"
}
