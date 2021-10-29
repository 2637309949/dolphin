// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// HourAllotRatio defined
type HourAllotRatio struct {
	// HARId defined
	HARId null.Int `xorm:"int(11) pk notnull autoincr 'h_a_r_id'" json:"h_a_r_id" form:"h_a_r_id" xml:"h_a_r_id"`
	// Period defined
	Period null.Int `xorm:"int(11) 'period'" json:"period" form:"period" xml:"period"`
	// Ratio defined
	Ratio null.String `xorm:"varchar(10) 'ratio'" json:"ratio" form:"ratio" xml:"ratio"`
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
	// PkHamId defined
	PkHamId null.Int `xorm:"int(11) 'pk_ham_id'" json:"pk_ham_id" form:"pk_ham_id" xml:"pk_ham_id"`
	// EndMolecule defined
	EndMolecule null.Int `xorm:"int(11) 'end_molecule'" json:"end_molecule" form:"end_molecule" xml:"end_molecule"`
	// StartMolecule defined
	StartMolecule null.Int `xorm:"int(11) 'start_molecule'" json:"start_molecule" form:"start_molecule" xml:"start_molecule"`
}

// TableName table name of defined HourAllotRatio
func (m *HourAllotRatio) TableName() string {
	return "hour_allot_ratio"
}

func (r *HourAllotRatio) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalHourAllotRatio(data []byte) (HourAllotRatio, error) {
	var r HourAllotRatio
	err := json.Unmarshal(data, &r)
	return r, err
}
