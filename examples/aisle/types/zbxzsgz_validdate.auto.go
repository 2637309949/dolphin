// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ZbxzsgzValiddate defined
type ZbxzsgzValiddate struct {
	// ZVId defined
	ZVId null.Int `xorm:"int(11) pk notnull autoincr 'z_v_id'" json:"z_v_id" form:"z_v_id" xml:"z_v_id"`
	// ZbxzsgzId defined
	ZbxzsgzId null.Int `xorm:"int(11) 'zbxzsgz_id'" json:"zbxzsgz_id" form:"zbxzsgz_id" xml:"zbxzsgz_id"`
	// BeginDate defined
	BeginDate null.Time `xorm:"datetime 'begin_date'" json:"begin_date" form:"begin_date" xml:"begin_date"`
	// EndDate defined
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" form:"end_date" xml:"end_date"`
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
	// SyhourZsbl defined
	SyhourZsbl null.Float `xorm:"float(50,2) 'syhour_zsbl'" json:"syhour_zsbl" form:"syhour_zsbl" xml:"syhour_zsbl"`
}

// TableName table name of defined ZbxzsgzValiddate
func (m *ZbxzsgzValiddate) TableName() string {
	return "zbxzsgz_validdate"
}

func (r *ZbxzsgzValiddate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalZbxzsgzValiddate(data []byte) (ZbxzsgzValiddate, error) {
	var r ZbxzsgzValiddate
	err := json.Unmarshal(data, &r)
	return r, err
}
