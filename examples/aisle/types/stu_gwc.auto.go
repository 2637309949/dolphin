// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StuGwc defined
type StuGwc struct {
	// StuGwcId defined
	StuGwcId null.Int `xorm:"int(11) pk notnull autoincr 'stu_gwc_id'" json:"stu_gwc_id" form:"stu_gwc_id" xml:"stu_gwc_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// GwcName defined
	GwcName null.String `xorm:"varchar(100) 'gwc_name'" json:"gwc_name" form:"gwc_name" xml:"gwc_name"`
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
	// GwcPassword defined
	GwcPassword null.String `xorm:"varchar(1000) 'gwc_password'" json:"gwc_password" form:"gwc_password" xml:"gwc_password"`
}

// TableName table name of defined StuGwc
func (m *StuGwc) TableName() string {
	return "stu_gwc"
}

func (r *StuGwc) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStuGwc(data []byte) (StuGwc, error) {
	var r StuGwc
	err := json.Unmarshal(data, &r)
	return r, err
}
