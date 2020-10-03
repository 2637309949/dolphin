// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ParFile defined
type ParFile struct {
	//
	ParFileId null.Int `xorm:"int(11) pk notnull autoincr 'par_file_id'" json:"par_file_id" xml:"par_file_id"`
	//
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" xml:"par_id"`
	//
	FileId null.Int `xorm:"int(11) 'file_id'" json:"file_id" xml:"file_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined ParFile
func (m *ParFile) TableName() string {
	return "par_file"
}