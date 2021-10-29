// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// Partner defined
type Partner struct {
	// PartnerId defined
	PartnerId null.Int `xorm:"int(11) pk notnull autoincr 'partner_id'" json:"partner_id" form:"partner_id" xml:"partner_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// PrName defined
	PrName null.String `xorm:"varchar(100) 'pr_name'" json:"pr_name" form:"pr_name" xml:"pr_name"`
	// PrRemark defined
	PrRemark null.String `xorm:"varchar(500) 'pr_remark'" json:"pr_remark" form:"pr_remark" xml:"pr_remark"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// PartnerType defined
	PartnerType null.Int `xorm:"int(11) 'partner_type'" json:"partner_type" form:"partner_type" xml:"partner_type"`
}

// TableName table name of defined Partner
func (m *Partner) TableName() string {
	return "partner"
}

func (r *Partner) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartner(data []byte) (Partner, error) {
	var r Partner
	err := json.Unmarshal(data, &r)
	return r, err
}
