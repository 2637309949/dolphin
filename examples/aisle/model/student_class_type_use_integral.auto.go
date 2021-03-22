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

// StudentClassTypeUseIntegral defined
type StudentClassTypeUseIntegral struct {
	// SCTUIId defined
	SCTUIId null.Int `xorm:"int(11) pk notnull autoincr 's_c_t_u_i_id'" json:"s_c_t_u_i_id" form:"s_c_t_u_i_id" xml:"s_c_t_u_i_id"`
	// SctId defined
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" form:"sct_id" xml:"sct_id"`
	// UseIntegral defined
	UseIntegral null.Float `xorm:"float(11,2) 'use_integral'" json:"use_integral" form:"use_integral" xml:"use_integral"`
	// DyMoney defined
	DyMoney null.Float `xorm:"float(11,2) 'dy_money'" json:"dy_money" form:"dy_money" xml:"dy_money"`
	// StuIntegral defined
	StuIntegral null.Int `xorm:"int(11) 'stu_integral'" json:"stu_integral" form:"stu_integral" xml:"stu_integral"`
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
}

// With defined
func (m *StudentClassTypeUseIntegral) With(s interface{}) (interface{}, error) {
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
func (m *StudentClassTypeUseIntegral) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StudentClassTypeUseIntegral) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StudentClassTypeUseIntegral) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StudentClassTypeUseIntegral) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StudentClassTypeUseIntegral) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *StudentClassTypeUseIntegral) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentClassTypeUseIntegral
func (m *StudentClassTypeUseIntegral) TableName() string {
	return "student_class_type_use_integral"
}
