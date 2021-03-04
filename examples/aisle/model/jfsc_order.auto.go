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

// JfscOrder defined
type JfscOrder struct {
	// JfscOrderId defined
	JfscOrderId null.Int `xorm:"int(11) pk notnull autoincr 'jfsc_order_id'" json:"jfsc_order_id" form:"jfsc_order_id" xml:"jfsc_order_id"`
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
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// SgId defined
	SgId null.Int `xorm:"int(11) 'sg_id'" json:"sg_id" form:"sg_id" xml:"sg_id"`
	// GeId defined
	GeId null.Int `xorm:"int(11) 'ge_id'" json:"ge_id" form:"ge_id" xml:"ge_id"`
	// RecipientInfoId defined
	RecipientInfoId null.Int `xorm:"int(11) 'recipient_info_id'" json:"recipient_info_id" form:"recipient_info_id" xml:"recipient_info_id"`
	// DingdanState defined
	DingdanState null.Int `xorm:"int(11) 'dingdan_state'" json:"dingdan_state" form:"dingdan_state" xml:"dingdan_state"`
}

// Marshal defined
func (m *JfscOrder) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *JfscOrder) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *JfscOrder) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *JfscOrder) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *JfscOrder) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *JfscOrder) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined JfscOrder
func (m *JfscOrder) TableName() string {
	return "jfsc_order"
}
