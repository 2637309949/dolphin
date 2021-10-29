// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StuWxdqDate defined
type StuWxdqDate struct {
	// SWDId defined
	SWDId null.Int `xorm:"int(11) pk notnull autoincr 's_w_d_id'" json:"s_w_d_id" form:"s_w_d_id" xml:"s_w_d_id"`
	// DayNum defined
	DayNum null.Float `xorm:"float(11,2) 'day_num'" json:"day_num" form:"day_num" xml:"day_num"`
	// WxdqDesc defined
	WxdqDesc null.String `xorm:"varchar(500) 'wxdq_desc'" json:"wxdq_desc" form:"wxdq_desc" xml:"wxdq_desc"`
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

// TableName table name of defined StuWxdqDate
func (m *StuWxdqDate) TableName() string {
	return "stu_wxdq_date"
}

func (r *StuWxdqDate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStuWxdqDate(data []byte) (StuWxdqDate, error) {
	var r StuWxdqDate
	err := json.Unmarshal(data, &r)
	return r, err
}
