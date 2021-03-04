// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// OrderFromStudyCard defined
type OrderFromStudyCard struct {
	// OFSCId defined
	OFSCId null.Int `xorm:"int(11) pk notnull autoincr 'o_f_s_c_id'" json:"o_f_s_c_id" form:"o_f_s_c_id" xml:"o_f_s_c_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// ScId defined
	ScId null.Int `xorm:"int(11) 'sc_id'" json:"sc_id" form:"sc_id" xml:"sc_id"`
	// DiscountMoney defined
	DiscountMoney null.Float `xorm:"float(11,2) 'discount_money'" json:"discount_money" form:"discount_money" xml:"discount_money"`
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
}

// Marshal defined
func (m *OrderFromStudyCard) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *OrderFromStudyCard) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *OrderFromStudyCard) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *OrderFromStudyCard) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *OrderFromStudyCard) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *OrderFromStudyCard) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OrderFromStudyCard
func (m *OrderFromStudyCard) TableName() string {
	return "order_from_study_card"
}
