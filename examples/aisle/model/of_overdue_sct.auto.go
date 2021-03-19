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

// OfOverdueSct defined
type OfOverdueSct struct {
	// OOSId defined
	OOSId null.Int `xorm:"int(11) pk notnull autoincr 'o_o_s_id'" json:"o_o_s_id" form:"o_o_s_id" xml:"o_o_s_id"`
	// OfOverdueId defined
	OfOverdueId null.Int `xorm:"int(11) 'of_overdue_id'" json:"of_overdue_id" form:"of_overdue_id" xml:"of_overdue_id"`
	// SctId defined
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" form:"sct_id" xml:"sct_id"`
	// OverdueHour defined
	OverdueHour null.Float `xorm:"float(11,2) 'overdue_hour'" json:"overdue_hour" form:"overdue_hour" xml:"overdue_hour"`
	// OverdueMoney defined
	OverdueMoney null.Float `xorm:"float(11,2) 'overdue_money'" json:"overdue_money" form:"overdue_money" xml:"overdue_money"`
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
	// AccountDate defined
	AccountDate null.Time `xorm:"datetime 'account_date'" json:"account_date" form:"account_date" xml:"account_date"`
	// AccountType defined
	AccountType null.Int `xorm:"int(11) 'account_type'" json:"account_type" form:"account_type" xml:"account_type"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// KfId defined
	KfId null.Int `xorm:"int(11) 'kf_id'" json:"kf_id" form:"kf_id" xml:"kf_id"`
}

// With defined
func (m *OfOverdueSct) With(s interface{}) (interface{}, error) {
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
func (m *OfOverdueSct) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *OfOverdueSct) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *OfOverdueSct) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *OfOverdueSct) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *OfOverdueSct) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *OfOverdueSct) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OfOverdueSct
func (m *OfOverdueSct) TableName() string {
	return "of_overdue_sct"
}
