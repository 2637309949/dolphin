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

// PhoneMessage defined
type PhoneMessage struct {
	//
	T90 null.Int `xorm:"int(11) pk notnull autoincr 't_9_0'" json:"t_9_0" form:"t_9_0" xml:"t_9_0"`
	//
	Title null.String `xorm:"varchar(500) 'title'" json:"title" form:"title" xml:"title"`
	//
	Content null.String `xorm:"varchar(2000) 'content'" json:"content" form:"content" xml:"content"`
	//
	IfNotice null.Int `xorm:"int(11) 'if_notice'" json:"if_notice" form:"if_notice" xml:"if_notice"`
	//
	IfNoticecn null.String `xorm:"varchar 'if_noticecn'" json:"if_noticecn" form:"if_noticecn" xml:"if_noticecn"`
	//
	IfSee null.Int `xorm:"int(11) 'if_see'" json:"if_see" form:"if_see" xml:"if_see"`
	//
	IfSeecn null.String `xorm:"varchar 'if_seecn'" json:"if_seecn" form:"if_seecn" xml:"if_seecn"`
	//
	NoticeUser null.Int `xorm:"int(11) 'notice_user'" json:"notice_user" form:"notice_user" xml:"notice_user"`
	//
	NoticeName null.String `xorm:"varchar(500) 'notice_name'" json:"notice_name" form:"notice_name" xml:"notice_name"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Parser defined
func (m *PhoneMessage) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *PhoneMessage) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined PhoneMessage
func (m *PhoneMessage) TableName() string {
	return "phone_message"
}
