// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// MaBusiness defined
type MaBusiness struct {
	// MBId defined
	MBId null.Int `xorm:"int(11) pk notnull autoincr 'm_b_id'" json:"m_b_id" form:"m_b_id" xml:"m_b_id"`
	// MaId defined
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
	// Business defined
	Business null.Int `xorm:"int(11) 'business'" json:"business" form:"business" xml:"business"`
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

// TableName table name of defined MaBusiness
func (m *MaBusiness) TableName() string {
	return "ma_business"
}

func (r *MaBusiness) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMaBusiness(data []byte) (MaBusiness, error) {
	var r MaBusiness
	err := json.Unmarshal(data, &r)
	return r, err
}
