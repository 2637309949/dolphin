// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// SaleSchool defined
type SaleSchool struct {
	// SSId defined
	SSId null.Int `xorm:"int(11) pk notnull autoincr 's_s_id'" json:"s_s_id" form:"s_s_id" xml:"s_s_id"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
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

// TableName table name of defined SaleSchool
func (m *SaleSchool) TableName() string {
	return "sale_school"
}

func (r *SaleSchool) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSaleSchool(data []byte) (SaleSchool, error) {
	var r SaleSchool
	err := json.Unmarshal(data, &r)
	return r, err
}
