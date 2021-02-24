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

// SystemMessage defined
type SystemMessage struct {
	// IfNotice defined
	IfNotice null.Int `xorm:"int(11) 'if_notice'" json:"if_notice" form:"if_notice" xml:"if_notice"`
	// IfNoticecn defined
	IfNoticecn null.String `xorm:"varchar 'if_noticecn'" json:"if_noticecn" form:"if_noticecn" xml:"if_noticecn"`
	// IfSee defined
	IfSee null.Int `xorm:"int(11) 'if_see'" json:"if_see" form:"if_see" xml:"if_see"`
	// IfSeecn defined
	IfSeecn null.String `xorm:"varchar 'if_seecn'" json:"if_seecn" form:"if_seecn" xml:"if_seecn"`
	// NoticeUser defined
	NoticeUser null.Int `xorm:"int(11) 'notice_user'" json:"notice_user" form:"notice_user" xml:"notice_user"`
	// NoticeName defined
	NoticeName null.String `xorm:"varchar(500) 'notice_name'" json:"notice_name" form:"notice_name" xml:"notice_name"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// T80 defined
	T80 null.Int `xorm:"int(11) pk notnull autoincr 't_8_0'" json:"t_8_0" form:"t_8_0" xml:"t_8_0"`
	// Title defined
	Title null.String `xorm:"varchar(500) 'title'" json:"title" form:"title" xml:"title"`
	// Content defined
	Content null.String `xorm:"varchar(2000) 'content'" json:"content" form:"content" xml:"content"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Parser defined
func (m *SystemMessage) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SystemMessage) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SystemMessage
func (m *SystemMessage) TableName() string {
	return "system_message"
}
