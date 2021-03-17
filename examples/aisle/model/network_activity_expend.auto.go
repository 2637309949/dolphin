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

// NetworkActivityExpend defined
type NetworkActivityExpend struct {
	// NAEId defined
	NAEId null.Int `xorm:"int(11) pk notnull autoincr 'n_a_e_id'" json:"n_a_e_id" form:"n_a_e_id" xml:"n_a_e_id"`
	// NaId defined
	NaId null.Int `xorm:"int(11) 'na_id'" json:"na_id" form:"na_id" xml:"na_id"`
	// ExpendMoney defined
	ExpendMoney null.Float `xorm:"float(11,2) 'expend_money'" json:"expend_money" form:"expend_money" xml:"expend_money"`
	// Remark defined
	Remark null.String `xorm:"varchar(500) 'remark'" json:"remark" form:"remark" xml:"remark"`
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
	// FeeType defined
	FeeType null.Int `xorm:"int(11) 'fee_type'" json:"fee_type" form:"fee_type" xml:"fee_type"`
}

// Marshal defined
func (m *NetworkActivityExpend) With(s interface{}) (interface{}, error) {
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
func (m *NetworkActivityExpend) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *NetworkActivityExpend) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *NetworkActivityExpend) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *NetworkActivityExpend) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *NetworkActivityExpend) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *NetworkActivityExpend) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined NetworkActivityExpend
func (m *NetworkActivityExpend) TableName() string {
	return "network_activity_expend"
}
