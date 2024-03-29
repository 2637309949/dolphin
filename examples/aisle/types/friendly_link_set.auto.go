// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// FriendlyLinkSet defined
type FriendlyLinkSet struct {
	// FLSId defined
	FLSId null.Int `xorm:"int(11) pk notnull autoincr 'f_l_s_id'" json:"f_l_s_id" form:"f_l_s_id" xml:"f_l_s_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// FlName defined
	FlName null.String `xorm:"varchar(50) 'fl_name'" json:"fl_name" form:"fl_name" xml:"fl_name"`
	// FlUrl defined
	FlUrl null.String `xorm:"varchar(50) 'fl_url'" json:"fl_url" form:"fl_url" xml:"fl_url"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined FriendlyLinkSet
func (m *FriendlyLinkSet) TableName() string {
	return "friendly_link_set"
}

func (r *FriendlyLinkSet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFriendlyLinkSet(data []byte) (FriendlyLinkSet, error) {
	var r FriendlyLinkSet
	err := json.Unmarshal(data, &r)
	return r, err
}
