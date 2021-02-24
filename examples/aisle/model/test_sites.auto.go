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

// TestSites defined
type TestSites struct {
	//
	TestSitesId null.Int `xorm:"int(11) pk notnull autoincr 'test_sites_id'" json:"test_sites_id" form:"test_sites_id" xml:"test_sites_id"`
	//
	TsName null.String `xorm:"varchar(100) 'ts_name'" json:"ts_name" form:"ts_name" xml:"ts_name"`
	//
	TsRemark null.String `xorm:"varchar(200) 'ts_remark'" json:"ts_remark" form:"ts_remark" xml:"ts_remark"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	SchoolId null.Int `xorm:"int(11) 'school_id'" json:"school_id" form:"school_id" xml:"school_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
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

// TableName table name of defined TestSites
func (m *TestSites) TableName() string {
	return "test_sites"
}
