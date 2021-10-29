// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// NetworkActivityHead defined
type NetworkActivityHead struct {
	// NAHId defined
	NAHId null.Int `xorm:"int(11) pk notnull autoincr 'n_a_h_id'" json:"n_a_h_id" form:"n_a_h_id" xml:"n_a_h_id"`
	// NaId defined
	NaId null.Int `xorm:"int(11) 'na_id'" json:"na_id" form:"na_id" xml:"na_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
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

// TableName table name of defined NetworkActivityHead
func (m *NetworkActivityHead) TableName() string {
	return "network_activity_head"
}

func (r *NetworkActivityHead) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNetworkActivityHead(data []byte) (NetworkActivityHead, error) {
	var r NetworkActivityHead
	err := json.Unmarshal(data, &r)
	return r, err
}
