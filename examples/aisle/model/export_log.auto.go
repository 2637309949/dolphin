// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ExportLog defined
type ExportLog struct {
	// ExportLogId defined
	ExportLogId null.Int `xorm:"int(11) pk notnull autoincr 'export_log_id'" json:"export_log_id" form:"export_log_id" xml:"export_log_id"`
	// ExportUser defined
	ExportUser null.Int `xorm:"int(11) 'export_user'" json:"export_user" form:"export_user" xml:"export_user"`
	// ExportIp defined
	ExportIp null.String `xorm:"varchar(100) 'export_ip'" json:"export_ip" form:"export_ip" xml:"export_ip"`
	// ExportTime defined
	ExportTime null.Time `xorm:"datetime 'export_time'" json:"export_time" form:"export_time" xml:"export_time"`
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

// TableName table name of defined ExportLog
func (m *ExportLog) TableName() string {
	return "export_log"
}
