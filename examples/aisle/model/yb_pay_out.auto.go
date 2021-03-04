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

// YbPayOut defined
type YbPayOut struct {
	// YbPayOutId defined
	YbPayOutId null.Int `xorm:"int(11) pk notnull autoincr 'yb_pay_out_id'" json:"yb_pay_out_id" form:"yb_pay_out_id" xml:"yb_pay_out_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// YbPayId defined
	YbPayId null.Int `xorm:"int(11) 'yb_pay_id'" json:"yb_pay_id" form:"yb_pay_id" xml:"yb_pay_id"`
	// OutMoney defined
	OutMoney null.Float `xorm:"float(11,2) 'out_money'" json:"out_money" form:"out_money" xml:"out_money"`
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
	// FeeId defined
	FeeId null.Int `xorm:"int(11) 'fee_id'" json:"fee_id" form:"fee_id" xml:"fee_id"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) default(54) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// YboutDesc defined
	YboutDesc null.String `xorm:"varchar(1000) 'ybout_desc'" json:"ybout_desc" form:"ybout_desc" xml:"ybout_desc"`
	// IfZdj defined
	IfZdj null.Int `xorm:"int(11) 'if_zdj'" json:"if_zdj" form:"if_zdj" xml:"if_zdj"`
	// DjId defined
	DjId null.Int `xorm:"int(11) 'dj_id'" json:"dj_id" form:"dj_id" xml:"dj_id"`
}

// Marshal defined
func (m *YbPayOut) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *YbPayOut) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *YbPayOut) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *YbPayOut) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *YbPayOut) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *YbPayOut) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined YbPayOut
func (m *YbPayOut) TableName() string {
	return "yb_pay_out"
}
