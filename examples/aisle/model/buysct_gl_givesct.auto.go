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

// BuysctGlGivesct defined
type BuysctGlGivesct struct {
	// BGGId defined
	BGGId null.Int `xorm:"int(11) pk notnull autoincr 'b_g_g_id'" json:"b_g_g_id" form:"b_g_g_id" xml:"b_g_g_id"`
	// BuysctId defined
	BuysctId null.Int `xorm:"int(11) 'buysct_id'" json:"buysct_id" form:"buysct_id" xml:"buysct_id"`
	// GivesctId defined
	GivesctId null.Int `xorm:"int(11) 'givesct_id'" json:"givesct_id" form:"givesct_id" xml:"givesct_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Parser defined
func (m *BuysctGlGivesct) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *BuysctGlGivesct) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined BuysctGlGivesct
func (m *BuysctGlGivesct) TableName() string {
	return "buysct_gl_givesct"
}
