// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// Department defined
type Department struct {
	// DepartmentId defined
	DepartmentId null.Int `xorm:"int(11) pk notnull autoincr 'department_id'" json:"department_id" form:"department_id" xml:"department_id"`
	// ParentId defined
	ParentId null.Int `xorm:"int(11) 'parent_id'" json:"parent_id" form:"parent_id" xml:"parent_id"`
	// DepartmentName defined
	DepartmentName null.String `xorm:"varchar(100) 'department_name'" json:"department_name" form:"department_name" xml:"department_name"`
	// DepartmentNumber defined
	DepartmentNumber null.String `xorm:"varchar(100) 'department_number'" json:"department_number" form:"department_number" xml:"department_number"`
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
	// SchoolId defined
	SchoolId null.Int `xorm:"int(11) 'school_id'" json:"school_id" form:"school_id" xml:"school_id"`
	// DepCity defined
	DepCity null.Int `xorm:"int(11) 'dep_city'" json:"dep_city" form:"dep_city" xml:"dep_city"`
}

// TableName table name of defined Department
func (m *Department) TableName() string {
	return "department"
}

func (r *Department) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDepartment(data []byte) (Department, error) {
	var r Department
	err := json.Unmarshal(data, &r)
	return r, err
}
