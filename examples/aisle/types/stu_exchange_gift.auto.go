// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StuExchangeGift defined
type StuExchangeGift struct {
	// SEGId defined
	SEGId null.Int `xorm:"int(11) pk notnull autoincr 's_e_g_id'" json:"s_e_g_id" form:"s_e_g_id" xml:"s_e_g_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// SegRemark defined
	SegRemark null.String `xorm:"varchar(1000) 'seg_remark'" json:"seg_remark" form:"seg_remark" xml:"seg_remark"`
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
	// UseTotalPoint defined
	UseTotalPoint null.Float `xorm:"float(10,2) 'use_total_point'" json:"use_total_point" form:"use_total_point" xml:"use_total_point"`
}

// TableName table name of defined StuExchangeGift
func (m *StuExchangeGift) TableName() string {
	return "stu_exchange_gift"
}

func (r *StuExchangeGift) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStuExchangeGift(data []byte) (StuExchangeGift, error) {
	var r StuExchangeGift
	err := json.Unmarshal(data, &r)
	return r, err
}
