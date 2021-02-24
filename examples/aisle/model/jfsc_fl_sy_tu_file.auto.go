// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// JfscFlSyTuFile defined
type JfscFlSyTuFile struct {
	//
	JFSTFId null.Int `xorm:"int(11) pk notnull autoincr 'j_f_s_t_f_id'" json:"j_f_s_t_f_id" form:"j_f_s_t_f_id" xml:"j_f_s_t_f_id"`
	//
	JfscFlSyTu null.Int `xorm:"int(11) 'jfsc_fl_sy_tu'" json:"jfsc_fl_sy_tu" form:"jfsc_fl_sy_tu" xml:"jfsc_fl_sy_tu"`
	//
	Fileid null.Int `xorm:"int(11) 'fileid'" json:"fileid" form:"fileid" xml:"fileid"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Parser defined
func (m *JfscFlSyTuFile) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *JfscFlSyTuFile) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined JfscFlSyTuFile
func (m *JfscFlSyTuFile) TableName() string {
	return "jfsc_fl_sy_tu_file"
}
