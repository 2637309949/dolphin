// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// OnSaleCity defined
type OnSaleCity struct {
	// OSCId defined
	OSCId null.Int `xorm:"int(11) pk notnull autoincr 'o_s_c_id'" json:"o_s_c_id" form:"o_s_c_id" xml:"o_s_c_id"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
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

// TableName table name of defined OnSaleCity
func (m *OnSaleCity) TableName() string {
	return "on_sale_city"
}

func (r *OnSaleCity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOnSaleCity(data []byte) (OnSaleCity, error) {
	var r OnSaleCity
	err := json.Unmarshal(data, &r)
	return r, err
}
