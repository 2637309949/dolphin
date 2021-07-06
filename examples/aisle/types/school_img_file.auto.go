// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SchoolImgFile defined
type SchoolImgFile struct {
	// SIFId defined
	SIFId null.Int `xorm:"int(11) pk notnull autoincr 's_i_f_id'" json:"s_i_f_id" form:"s_i_f_id" xml:"s_i_f_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// FileId defined
	FileId null.Int `xorm:"int(11) 'file_id'" json:"file_id" form:"file_id" xml:"file_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined SchoolImgFile
func (m *SchoolImgFile) TableName() string {
	return "school_img_file"
}