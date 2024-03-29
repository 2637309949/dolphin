// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// LevelUnit defined
type LevelUnit struct {
	// LevelUnitId defined
	LevelUnitId null.Int `xorm:"int(11) pk notnull autoincr 'level_unit_id'" json:"level_unit_id" form:"level_unit_id" xml:"level_unit_id"`
	// LevelId defined
	LevelId null.Int `xorm:"int(11) 'level_id'" json:"level_id" form:"level_id" xml:"level_id"`
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
	// UnitTotalScore defined
	UnitTotalScore null.Float `xorm:"float(50,2) 'unit_total_score'" json:"unit_total_score" form:"unit_total_score" xml:"unit_total_score"`
}

// TableName table name of defined LevelUnit
func (m *LevelUnit) TableName() string {
	return "level_unit"
}

func (r *LevelUnit) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalLevelUnit(data []byte) (LevelUnit, error) {
	var r LevelUnit
	err := json.Unmarshal(data, &r)
	return r, err
}
