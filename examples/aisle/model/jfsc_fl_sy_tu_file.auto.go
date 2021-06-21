// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// JfscFlSyTuFile defined
type JfscFlSyTuFile struct {
	// JFSTFId defined
	JFSTFId null.Int `xorm:"int(11) pk notnull autoincr 'j_f_s_t_f_id'" json:"j_f_s_t_f_id" form:"j_f_s_t_f_id" xml:"j_f_s_t_f_id"`
	// JfscFlSyTu defined
	JfscFlSyTu null.Int `xorm:"int(11) 'jfsc_fl_sy_tu'" json:"jfsc_fl_sy_tu" form:"jfsc_fl_sy_tu" xml:"jfsc_fl_sy_tu"`
	// Fileid defined
	Fileid null.Int `xorm:"int(11) 'fileid'" json:"fileid" form:"fileid" xml:"fileid"`
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

// TableName table name of defined JfscFlSyTuFile
func (m *JfscFlSyTuFile) TableName() string {
	return "jfsc_fl_sy_tu_file"
}
