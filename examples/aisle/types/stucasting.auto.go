// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// Stucasting defined
type Stucasting struct {
	// StucastingId defined
	StucastingId null.Int `xorm:"int(11) pk notnull autoincr 'stucasting_id'" json:"stucasting_id" form:"stucasting_id" xml:"stucasting_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// CastName defined
	CastName null.String `xorm:"varchar(50) 'cast_name'" json:"cast_name" form:"cast_name" xml:"cast_name"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined Stucasting
func (m *Stucasting) TableName() string {
	return "stucasting"
}

func (r *Stucasting) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStucasting(data []byte) (Stucasting, error) {
	var r Stucasting
	err := json.Unmarshal(data, &r)
	return r, err
}
