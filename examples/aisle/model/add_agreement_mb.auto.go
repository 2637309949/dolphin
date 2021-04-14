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

// AddAgreementMb defined
type AddAgreementMb struct {
	// AAMId defined
	AAMId null.Int `xorm:"int(11) pk notnull autoincr 'a_a_m_id'" json:"a_a_m_id" form:"a_a_m_id" xml:"a_a_m_id"`
	// AamName defined
	AamName null.String `xorm:"varchar(20) 'aam_name'" json:"aam_name" form:"aam_name" xml:"aam_name"`
	// AamHead defined
	AamHead null.Int `xorm:"int(11) 'aam_head'" json:"aam_head" form:"aam_head" xml:"aam_head"`
	// AamMidd defined
	AamMidd null.Int `xorm:"int(11) 'aam_midd'" json:"aam_midd" form:"aam_midd" xml:"aam_midd"`
	// AamLast defined
	AamLast null.Int `xorm:"int(11) 'aam_last'" json:"aam_last" form:"aam_last" xml:"aam_last"`
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
	// AamTitle defined
	AamTitle null.String `xorm:"varchar(500) 'aam_title'" json:"aam_title" form:"aam_title" xml:"aam_title"`
	// OpenOrClose defined
	OpenOrClose null.Int `xorm:"int(11) 'open_or_close'" json:"open_or_close" form:"open_or_close" xml:"open_or_close"`
}

// With defined
func (m *AddAgreementMb) With(s interface{}) (interface{}, error) {
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
func (m *AddAgreementMb) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *AddAgreementMb) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *AddAgreementMb) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *AddAgreementMb) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *AddAgreementMb) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *AddAgreementMb) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined AddAgreementMb
func (m *AddAgreementMb) TableName() string {
	return "add_agreement_mb"
}
