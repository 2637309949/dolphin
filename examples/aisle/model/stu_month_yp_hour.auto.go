// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// StuMonthYpHour defined
type StuMonthYpHour struct {
	// SMYHId defined
	SMYHId null.Int `xorm:"int(11) pk notnull autoincr 's_m_y_h_id'" json:"s_m_y_h_id" form:"s_m_y_h_id" xml:"s_m_y_h_id"`
	// PkStu defined
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" form:"pk_stu" xml:"pk_stu"`
	// PkSch defined
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
	// MonthHour defined
	MonthHour null.Float `xorm:"float(11,2) 'month_hour'" json:"month_hour" form:"month_hour" xml:"month_hour"`
	// MonthMoney defined
	MonthMoney decimal.Decimal `xorm:"decimal(11,2) 'month_money'" json:"month_money" form:"month_money" xml:"month_money"`
	// MonthDate defined
	MonthDate null.Time `xorm:"datetime 'month_date'" json:"month_date" form:"month_date" xml:"month_date"`
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
func (m *StuMonthYpHour) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StuMonthYpHour) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StuMonthYpHour) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StuMonthYpHour) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StuMonthYpHour) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuMonthYpHour) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuMonthYpHour
func (m *StuMonthYpHour) TableName() string {
	return "stu_month_yp_hour"
}
