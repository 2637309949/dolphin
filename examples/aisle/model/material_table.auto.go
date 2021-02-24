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

// MaterialTable defined
type MaterialTable struct {
	// MTId defined
	MTId null.Int `xorm:"int(11) pk notnull autoincr 'm_t_id'" json:"m_t_id" form:"m_t_id" xml:"m_t_id"`
	// MtName defined
	MtName null.String `xorm:"varchar(100) 'mt_name'" json:"mt_name" form:"mt_name" xml:"mt_name"`
	// MtPrice defined
	MtPrice null.Float `xorm:"float(11,2) 'mt_price'" json:"mt_price" form:"mt_price" xml:"mt_price"`
	// MtRemark defined
	MtRemark null.String `xorm:"varchar(2000) 'mt_remark'" json:"mt_remark" form:"mt_remark" xml:"mt_remark"`
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
func (m *MaterialTable) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MaterialTable) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MaterialTable
func (m *MaterialTable) TableName() string {
	return "material_table"
}
