// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// UserDepartment defined
type UserDepartment struct {
	// UDId defined
	UDId null.Int `xorm:"int(11) pk notnull autoincr 'u_d_id'" json:"u_d_id" form:"u_d_id" xml:"u_d_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// DepartmentId defined
	DepartmentId null.Int `xorm:"int(11) 'department_id'" json:"department_id" form:"department_id" xml:"department_id"`
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

// TableName table name of defined UserDepartment
func (m *UserDepartment) TableName() string {
	return "user_department"
}

func (r *UserDepartment) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUserDepartment(data []byte) (UserDepartment, error) {
	var r UserDepartment
	err := json.Unmarshal(data, &r)
	return r, err
}
