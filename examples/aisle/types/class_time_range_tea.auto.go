// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ClassTimeRangeTea defined
type ClassTimeRangeTea struct {
	// CTRTId defined
	CTRTId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_r_t_id'" json:"c_t_r_t_id" form:"c_t_r_t_id" xml:"c_t_r_t_id"`
	// PkCtr defined
	PkCtr null.Int `xorm:"int(11) 'pk_ctr'" json:"pk_ctr" form:"pk_ctr" xml:"pk_ctr"`
	// PkTea defined
	PkTea null.Int `xorm:"int(11) 'pk_tea'" json:"pk_tea" form:"pk_tea" xml:"pk_tea"`
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

// TableName table name of defined ClassTimeRangeTea
func (m *ClassTimeRangeTea) TableName() string {
	return "class_time_range_tea"
}

func (r *ClassTimeRangeTea) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClassTimeRangeTea(data []byte) (ClassTimeRangeTea, error) {
	var r ClassTimeRangeTea
	err := json.Unmarshal(data, &r)
	return r, err
}
