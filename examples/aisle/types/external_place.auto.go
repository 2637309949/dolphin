// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ExternalPlace defined
type ExternalPlace struct {
	// EPId defined
	EPId null.Int `xorm:"int(11) pk notnull autoincr 'e_p_id'" json:"e_p_id" form:"e_p_id" xml:"e_p_id"`
	// PlaceName defined
	PlaceName null.String `xorm:"varchar(500) 'place_name'" json:"place_name" form:"place_name" xml:"place_name"`
	// EpCity defined
	EpCity null.Int `xorm:"int(11) 'ep_city'" json:"ep_city" form:"ep_city" xml:"ep_city"`
	// EpRemark defined
	EpRemark null.String `xorm:"varchar(2000) 'ep_remark'" json:"ep_remark" form:"ep_remark" xml:"ep_remark"`
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

// TableName table name of defined ExternalPlace
func (m *ExternalPlace) TableName() string {
	return "external_place"
}

func (r *ExternalPlace) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExternalPlace(data []byte) (ExternalPlace, error) {
	var r ExternalPlace
	err := json.Unmarshal(data, &r)
	return r, err
}
