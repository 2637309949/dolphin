// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// GiftCity defined
type GiftCity struct {
	// GiftCityId defined
	GiftCityId null.Int `xorm:"int(11) pk notnull autoincr 'gift_city_id'" json:"gift_city_id" form:"gift_city_id" xml:"gift_city_id"`
	// GiftId defined
	GiftId null.Int `xorm:"int(11) 'gift_id'" json:"gift_id" form:"gift_id" xml:"gift_id"`
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

// TableName table name of defined GiftCity
func (m *GiftCity) TableName() string {
	return "gift_city"
}

func (r *GiftCity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalGiftCity(data []byte) (GiftCity, error) {
	var r GiftCity
	err := json.Unmarshal(data, &r)
	return r, err
}
