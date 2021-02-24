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

// CssCsTaskFile defined
type CssCsTaskFile struct {
	//
	CCTFId null.Int `xorm:"int(11) pk notnull autoincr 'c_c_t_f_id'" json:"c_c_t_f_id" form:"c_c_t_f_id" xml:"c_c_t_f_id"`
	//
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
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
	//
	CssCsTId null.Int `xorm:"int(11) 'css_cs_t_id'" json:"css_cs_t_id" form:"css_cs_t_id" xml:"css_cs_t_id"`
	//
	CsTFiled null.Int `xorm:"int(11) 'cs_t_filed'" json:"cs_t_filed" form:"cs_t_filed" xml:"cs_t_filed"`
	//
	CssId null.Int `xorm:"int(11) 'css_id'" json:"css_id" form:"css_id" xml:"css_id"`
	//
	CstId null.Int `xorm:"int(11) 'cst_id'" json:"cst_id" form:"cst_id" xml:"cst_id"`
}

// Parser defined
func (m *CssCsTaskFile) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *CssCsTaskFile) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CssCsTaskFile
func (m *CssCsTaskFile) TableName() string {
	return "css_cs_task_file"
}
