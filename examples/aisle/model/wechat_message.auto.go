// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// WechatMessage defined
type WechatMessage struct {
	// T3230 defined
	T3230 null.Int `xorm:"int(11) pk notnull autoincr 't_323_0'" json:"t_323_0" form:"t_323_0" xml:"t_323_0"`
	// Title defined
	Title null.String `xorm:"varchar(500) 'title'" json:"title" form:"title" xml:"title"`
	// Content defined
	Content null.String `xorm:"varchar(2000) 'content'" json:"content" form:"content" xml:"content"`
	// IfNotice defined
	IfNotice null.Int `xorm:"int(11) 'if_notice'" json:"if_notice" form:"if_notice" xml:"if_notice"`
	// IfNoticecn defined
	IfNoticecn null.Int `xorm:"int(11) 'if_noticecn'" json:"if_noticecn" form:"if_noticecn" xml:"if_noticecn"`
	// IfSee defined
	IfSee null.Int `xorm:"int(11) 'if_see'" json:"if_see" form:"if_see" xml:"if_see"`
	// IfSeecn defined
	IfSeecn null.Int `xorm:"int(11) 'if_seecn'" json:"if_seecn" form:"if_seecn" xml:"if_seecn"`
	// NoticeUser defined
	NoticeUser null.Int `xorm:"int(11) 'notice_user'" json:"notice_user" form:"notice_user" xml:"notice_user"`
	// NoticeName defined
	NoticeName null.String `xorm:"varchar(500) 'notice_name'" json:"notice_name" form:"notice_name" xml:"notice_name"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// With defined
func (m *WechatMessage) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *WechatMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *WechatMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *WechatMessage) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *WechatMessage) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *WechatMessage) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *WechatMessage) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined WechatMessage
func (m *WechatMessage) TableName() string {
	return "wechat_message"
}
