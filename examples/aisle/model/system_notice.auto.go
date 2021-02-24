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

// SystemNotice defined
type SystemNotice struct {
	//
	SysnId null.Int `xorm:"int(11) pk notnull autoincr 'sysn_id'" json:"sysn_id" form:"sysn_id" xml:"sysn_id"`
	//
	SysnName null.String `xorm:"varchar(500) 'sysn_name'" json:"sysn_name" form:"sysn_name" xml:"sysn_name"`
	//
	SysnContent null.String `xorm:"varchar(4000) 'sysn_content'" json:"sysn_content" form:"sysn_content" xml:"sysn_content"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
}

// Parser defined
func (m *SystemNotice) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SystemNotice) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SystemNotice
func (m *SystemNotice) TableName() string {
	return "system_notice"
}
