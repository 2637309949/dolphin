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

// ActiveLessonProType defined
type ActiveLessonProType struct {
	// ALPTId defined
	ALPTId null.Int `xorm:"int(11) pk notnull autoincr 'a_l_p_t_id'" json:"a_l_p_t_id" form:"a_l_p_t_id" xml:"a_l_p_t_id"`
	// ActiveId defined
	ActiveId null.Int `xorm:"int(11) 'active_id'" json:"active_id" form:"active_id" xml:"active_id"`
	// ProTypeId defined
	ProTypeId null.Int `xorm:"int(11) 'pro_type_id'" json:"pro_type_id" form:"pro_type_id" xml:"pro_type_id"`
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
func (m *ActiveLessonProType) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ActiveLessonProType) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ActiveLessonProType
func (m *ActiveLessonProType) TableName() string {
	return "active_lesson_pro_type"
}
