// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// Enterprises defined
type Enterprises struct {
	// EId defined
	EId null.Int `xorm:"int(11) pk notnull autoincr 'e_id'" json:"e_id" form:"e_id" xml:"e_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// EntName defined
	EntName null.String `xorm:"varchar(100) 'ent_name'" json:"ent_name" form:"ent_name" xml:"ent_name"`
	// EntRemark defined
	EntRemark null.String `xorm:"varchar(100) 'ent_remark'" json:"ent_remark" form:"ent_remark" xml:"ent_remark"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined Enterprises
func (m *Enterprises) TableName() string {
	return "enterprises"
}

func (r *Enterprises) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEnterprises(data []byte) (Enterprises, error) {
	var r Enterprises
	err := json.Unmarshal(data, &r)
	return r, err
}
