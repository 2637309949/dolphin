// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// OrganSchoolBussType defined
type OrganSchoolBussType struct {
	// OSBTId defined
	OSBTId null.Int `xorm:"int(11) pk notnull autoincr 'o_s_b_t_id'" json:"o_s_b_t_id" form:"o_s_b_t_id" xml:"o_s_b_t_id"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// OsBussType defined
	OsBussType null.Int `xorm:"int(11) 'os_buss_type'" json:"os_buss_type" form:"os_buss_type" xml:"os_buss_type"`
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

// TableName table name of defined OrganSchoolBussType
func (m *OrganSchoolBussType) TableName() string {
	return "organ_school_buss_type"
}

func (r *OrganSchoolBussType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOrganSchoolBussType(data []byte) (OrganSchoolBussType, error) {
	var r OrganSchoolBussType
	err := json.Unmarshal(data, &r)
	return r, err
}
