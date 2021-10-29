// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// LanguagePackCity defined
type LanguagePackCity struct {
	// LPCId defined
	LPCId null.Int `xorm:"int(11) pk notnull autoincr 'l_p_c_id'" json:"l_p_c_id" form:"l_p_c_id" xml:"l_p_c_id"`
	// LpId defined
	LpId null.Int `xorm:"int(11) 'lp_id'" json:"lp_id" form:"lp_id" xml:"lp_id"`
	// CityId defined
	CityId null.Int `xorm:"int(11) 'city_id'" json:"city_id" form:"city_id" xml:"city_id"`
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

// TableName table name of defined LanguagePackCity
func (m *LanguagePackCity) TableName() string {
	return "language_pack_city"
}

func (r *LanguagePackCity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalLanguagePackCity(data []byte) (LanguagePackCity, error) {
	var r LanguagePackCity
	err := json.Unmarshal(data, &r)
	return r, err
}
