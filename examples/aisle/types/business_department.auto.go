// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// BusinessDepartment defined
type BusinessDepartment struct {
	// BDId defined
	BDId null.Int `xorm:"int(11) pk notnull autoincr 'b_d_id'" json:"b_d_id" form:"b_d_id" xml:"b_d_id"`
	// BdName defined
	BdName null.String `xorm:"varchar(200) 'bd_name'" json:"bd_name" form:"bd_name" xml:"bd_name"`
	// BdDesc defined
	BdDesc null.String `xorm:"varchar(300) 'bd_desc'" json:"bd_desc" form:"bd_desc" xml:"bd_desc"`
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

// TableName table name of defined BusinessDepartment
func (m *BusinessDepartment) TableName() string {
	return "business_department"
}

func (r *BusinessDepartment) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBusinessDepartment(data []byte) (BusinessDepartment, error) {
	var r BusinessDepartment
	err := json.Unmarshal(data, &r)
	return r, err
}
