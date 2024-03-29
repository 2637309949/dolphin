// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StuConfirmHour defined
type StuConfirmHour struct {
	// SCHId defined
	SCHId null.Int `xorm:"int(11) pk notnull autoincr 's_c_h_id'" json:"s_c_h_id" form:"s_c_h_id" xml:"s_c_h_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
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

// TableName table name of defined StuConfirmHour
func (m *StuConfirmHour) TableName() string {
	return "stu_confirm_hour"
}

func (r *StuConfirmHour) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStuConfirmHour(data []byte) (StuConfirmHour, error) {
	var r StuConfirmHour
	err := json.Unmarshal(data, &r)
	return r, err
}
