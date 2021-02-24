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

// Organ defined
type Organ struct {
	// OrganId defined
	OrganId null.Int `xorm:"int(11) pk notnull autoincr 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
	// OrganName defined
	OrganName null.String `xorm:"varchar(50) 'organ_name'" json:"organ_name" form:"organ_name" xml:"organ_name"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// OrganPinyin defined
	OrganPinyin null.String `xorm:"varchar(100) 'organ_pinyin'" json:"organ_pinyin" form:"organ_pinyin" xml:"organ_pinyin"`
	// OrganNumber defined
	OrganNumber null.String `xorm:"varchar(50) 'organ_number'" json:"organ_number" form:"organ_number" xml:"organ_number"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// ParentId defined
	ParentId null.Int `xorm:"int(11) 'parent_id'" json:"parent_id" form:"parent_id" xml:"parent_id"`
	// IfBuyOnline defined
	IfBuyOnline null.Int `xorm:"int(11) 'if_buy_online'" json:"if_buy_online" form:"if_buy_online" xml:"if_buy_online"`
	// OrganTell defined
	OrganTell null.String `xorm:"varchar(20) 'organ_tell'" json:"organ_tell" form:"organ_tell" xml:"organ_tell"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// ShengId defined
	ShengId null.Int `xorm:"int(11) 'sheng_id'" json:"sheng_id" form:"sheng_id" xml:"sheng_id"`
}

// Parser defined
func (m *Organ) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *Organ) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined Organ
func (m *Organ) TableName() string {
	return "organ"
}
