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

// RecipientInfo defined
type RecipientInfo struct {
	// RIId defined
	RIId null.Int `xorm:"int(11) pk notnull autoincr 'r_i_id'" json:"r_i_id" form:"r_i_id" xml:"r_i_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// RecipientName defined
	RecipientName null.String `xorm:"varchar(100) 'recipient_name'" json:"recipient_name" form:"recipient_name" xml:"recipient_name"`
	// RecipientPhone defined
	RecipientPhone null.String `xorm:"varchar(100) 'recipient_phone'" json:"recipient_phone" form:"recipient_phone" xml:"recipient_phone"`
	// RecipientCity defined
	RecipientCity null.Int `xorm:"int(11) 'recipient_city'" json:"recipient_city" form:"recipient_city" xml:"recipient_city"`
	// RecipientCampus defined
	RecipientCampus null.Int `xorm:"int(11) 'recipient_campus'" json:"recipient_campus" form:"recipient_campus" xml:"recipient_campus"`
	// RecipientAdress defined
	RecipientAdress null.String `xorm:"varchar(2000) 'recipient_adress'" json:"recipient_adress" form:"recipient_adress" xml:"recipient_adress"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"notnull 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// RecipientRegion defined
	RecipientRegion null.String `xorm:"varchar(100) 'recipient_region'" json:"recipient_region" form:"recipient_region" xml:"recipient_region"`
	// IfDefault defined
	IfDefault null.Int `xorm:"int(11) 'if_default'" json:"if_default" form:"if_default" xml:"if_default"`
}

// With defined
func (m *RecipientInfo) With(s interface{}) (interface{}, error) {
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
func (m *RecipientInfo) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *RecipientInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *RecipientInfo) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *RecipientInfo) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *RecipientInfo) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *RecipientInfo) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined RecipientInfo
func (m *RecipientInfo) TableName() string {
	return "recipient_info"
}
