// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// OldGiveGift defined
type OldGiveGift struct {
	// OGGId defined
	OGGId null.Int `xorm:"int(11) pk notnull autoincr 'o_g_g_id'" json:"o_g_g_id" form:"o_g_g_id" xml:"o_g_g_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// Remark defined
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" form:"remark" xml:"remark"`
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

// TableName table name of defined OldGiveGift
func (m *OldGiveGift) TableName() string {
	return "old_give_gift"
}

func (r *OldGiveGift) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOldGiveGift(data []byte) (OldGiveGift, error) {
	var r OldGiveGift
	err := json.Unmarshal(data, &r)
	return r, err
}
