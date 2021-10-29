// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ReplaceGift defined
type ReplaceGift struct {
	// RGId defined
	RGId null.Int `xorm:"int(11) pk notnull autoincr 'r_g_id'" json:"r_g_id" form:"r_g_id" xml:"r_g_id"`
	// SgId defined
	SgId null.Int `xorm:"int(11) 'sg_id'" json:"sg_id" form:"sg_id" xml:"sg_id"`
	// Remark defined
	Remark null.String `xorm:"varchar(400) 'remark'" json:"remark" form:"remark" xml:"remark"`
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

// TableName table name of defined ReplaceGift
func (m *ReplaceGift) TableName() string {
	return "replace_gift"
}

func (r *ReplaceGift) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalReplaceGift(data []byte) (ReplaceGift, error) {
	var r ReplaceGift
	err := json.Unmarshal(data, &r)
	return r, err
}
