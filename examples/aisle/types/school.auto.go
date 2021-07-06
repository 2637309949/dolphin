// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// School defined
type School struct {
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// SchoolId defined
	SchoolId null.Int `xorm:"int(11) pk notnull autoincr 'school_id'" json:"school_id" form:"school_id" xml:"school_id"`
	// SchoolName defined
	SchoolName null.String `xorm:"varchar(500) 'school_name'" json:"school_name" form:"school_name" xml:"school_name"`
	// SchoolArea defined
	SchoolArea null.String `xorm:"varchar(2000) 'school_area'" json:"school_area" form:"school_area" xml:"school_area"`
	// SchoolRemark defined
	SchoolRemark null.String `xorm:"varchar(2000) 'school_remark'" json:"school_remark" form:"school_remark" xml:"school_remark"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// OrganId defined
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
}

// TableName table name of defined School
func (m *School) TableName() string {
	return "school"
}