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

// StudentCzIntegral defined
type StudentCzIntegral struct {
	// SCIId defined
	SCIId null.Int `xorm:"int(11) pk notnull autoincr 's_c_i_id'" json:"s_c_i_id" form:"s_c_i_id" xml:"s_c_i_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// Integral defined
	Integral null.Float `xorm:"float(50,2) 'integral'" json:"integral" form:"integral" xml:"integral"`
	// IntegralDesc defined
	IntegralDesc null.String `xorm:"varchar(500) 'integral_desc'" json:"integral_desc" form:"integral_desc" xml:"integral_desc"`
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
	// IntegralState defined
	IntegralState null.Int `xorm:"int(11) 'integral_state'" json:"integral_state" form:"integral_state" xml:"integral_state"`
	// JfOrder defined
	JfOrder null.Int `xorm:"int(11) 'jf_order'" json:"jf_order" form:"jf_order" xml:"jf_order"`
	// YhjfOrderid defined
	YhjfOrderid null.Int `xorm:"int(11) 'yhjf_orderid'" json:"yhjf_orderid" form:"yhjf_orderid" xml:"yhjf_orderid"`
	// SctId defined
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" form:"sct_id" xml:"sct_id"`
}

// With defined
func (m *StudentCzIntegral) With(s interface{}) (interface{}, error) {
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
func (m *StudentCzIntegral) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StudentCzIntegral) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StudentCzIntegral) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StudentCzIntegral) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StudentCzIntegral) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *StudentCzIntegral) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentCzIntegral
func (m *StudentCzIntegral) TableName() string {
	return "student_cz_integral"
}
