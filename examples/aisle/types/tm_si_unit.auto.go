// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// TmSiUnit defined
type TmSiUnit struct {
	// TmSiUnitId defined
	TmSiUnitId null.Int `xorm:"int(11) pk notnull autoincr 'tm_si_unit_id'" json:"tm_si_unit_id" form:"tm_si_unit_id" xml:"tm_si_unit_id"`
	// TsId defined
	TsId null.Int `xorm:"int(11) 'ts_id'" json:"ts_id" form:"ts_id" xml:"ts_id"`
	// UnitId defined
	UnitId null.Int `xorm:"int(11) 'unit_id'" json:"unit_id" form:"unit_id" xml:"unit_id"`
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

// TableName table name of defined TmSiUnit
func (m *TmSiUnit) TableName() string {
	return "tm_si_unit"
}

func (r *TmSiUnit) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTmSiUnit(data []byte) (TmSiUnit, error) {
	var r TmSiUnit
	err := json.Unmarshal(data, &r)
	return r, err
}
