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
	"github.com/shopspring/decimal"
)

// OfDepositDetail defined
type OfDepositDetail struct {
	// ODDId defined
	ODDId null.Int `xorm:"int(11) pk notnull autoincr 'o_d_d_id'" json:"o_d_d_id" form:"o_d_d_id" xml:"o_d_d_id"`
	// PkOf defined
	PkOf null.Int `xorm:"int(11) 'pk_of'" json:"pk_of" form:"pk_of" xml:"pk_of"`
	// AllMoney defined
	AllMoney decimal.Decimal `xorm:"decimal(11,2) 'all_money'" json:"all_money" form:"all_money" xml:"all_money"`
	// StudyMoney defined
	StudyMoney decimal.Decimal `xorm:"decimal(11,2) 'study_money'" json:"study_money" form:"study_money" xml:"study_money"`
	// HostStartHour defined
	HostStartHour null.Float `xorm:"float(11,2) 'host_start_hour'" json:"host_start_hour" form:"host_start_hour" xml:"host_start_hour"`
	// HostEndHour defined
	HostEndHour null.Float `xorm:"float(11,2) 'host_end_hour'" json:"host_end_hour" form:"host_end_hour" xml:"host_end_hour"`
	// MinorStartHour defined
	MinorStartHour null.Float `xorm:"float(11,2) 'minor_start_hour'" json:"minor_start_hour" form:"minor_start_hour" xml:"minor_start_hour"`
	// MinorEndHour defined
	MinorEndHour null.Float `xorm:"float(11,2) 'minor_end_hour'" json:"minor_end_hour" form:"minor_end_hour" xml:"minor_end_hour"`
	// PeriodStartDate defined
	PeriodStartDate null.Time `xorm:"datetime 'period_start_date'" json:"period_start_date" form:"period_start_date" xml:"period_start_date"`
	// PeriodEndDate defined
	PeriodEndDate null.Time `xorm:"datetime 'period_end_date'" json:"period_end_date" form:"period_end_date" xml:"period_end_date"`
	// OverdueDays defined
	OverdueDays null.Float `xorm:"float(50,2) 'overdue_days'" json:"overdue_days" form:"overdue_days" xml:"overdue_days"`
	// OverdueMoney defined
	OverdueMoney decimal.Decimal `xorm:"decimal(50,3) 'overdue_money'" json:"overdue_money" form:"overdue_money" xml:"overdue_money"`
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
	// ZdPeriodNum defined
	ZdPeriodNum null.Int `xorm:"int(11) 'zd_period_num'" json:"zd_period_num" form:"zd_period_num" xml:"zd_period_num"`
	// PkPvaaId defined
	PkPvaaId null.Int `xorm:"int(11) 'pk_pvaa_id'" json:"pk_pvaa_id" form:"pk_pvaa_id" xml:"pk_pvaa_id"`
	// RealOverdueMoney defined
	RealOverdueMoney decimal.Decimal `xorm:"decimal(50,3) 'real_overdue_money'" json:"real_overdue_money" form:"real_overdue_money" xml:"real_overdue_money"`
	// GetMoney defined
	GetMoney decimal.Decimal `xorm:"decimal(50,3) 'get_money'" json:"get_money" form:"get_money" xml:"get_money"`
}

// With defined
func (m *OfDepositDetail) With(s interface{}) (interface{}, error) {
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
func (m *OfDepositDetail) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *OfDepositDetail) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *OfDepositDetail) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *OfDepositDetail) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *OfDepositDetail) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *OfDepositDetail) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OfDepositDetail
func (m *OfDepositDetail) TableName() string {
	return "of_deposit_detail"
}
