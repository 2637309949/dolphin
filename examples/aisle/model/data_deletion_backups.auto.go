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

// DataDeletionBackups defined
type DataDeletionBackups struct {
	//
	DDBId null.Int `xorm:"int(11) pk notnull autoincr 'd_d_b_id'" json:"d_d_b_id" form:"d_d_b_id" xml:"d_d_b_id"`
	//
	Tableid null.Int `xorm:"int(11) 'tableid'" json:"tableid" form:"tableid" xml:"tableid"`
	//
	Dataid null.Int `xorm:"int(11) 'dataid'" json:"dataid" form:"dataid" xml:"dataid"`
	//
	Userid null.Int `xorm:"int(11) 'userid'" json:"userid" form:"userid" xml:"userid"`
	//
	Username null.String `xorm:"varchar(100) 'username'" json:"username" form:"username" xml:"username"`
	//
	Deleteuserip null.String `xorm:"varchar(100) 'deleteuserip'" json:"deleteuserip" form:"deleteuserip" xml:"deleteuserip"`
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
func (m *SysCommentReply) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysCommentReply) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined DataDeletionBackups
func (m *DataDeletionBackups) TableName() string {
	return "data_deletion_backups"
}
