// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// NetworkActivityOrgan defined
type NetworkActivityOrgan struct {
	// NAOId defined
	NAOId null.Int `xorm:"int(11) pk notnull autoincr 'n_a_o_id'" json:"n_a_o_id" form:"n_a_o_id" xml:"n_a_o_id"`
	// NaId defined
	NaId null.Int `xorm:"int(11) 'na_id'" json:"na_id" form:"na_id" xml:"na_id"`
	// OrganId defined
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
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

// TableName table name of defined NetworkActivityOrgan
func (m *NetworkActivityOrgan) TableName() string {
	return "network_activity_organ"
}

func (r *NetworkActivityOrgan) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNetworkActivityOrgan(data []byte) (NetworkActivityOrgan, error) {
	var r NetworkActivityOrgan
	err := json.Unmarshal(data, &r)
	return r, err
}
