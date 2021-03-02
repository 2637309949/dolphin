// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// AppMessageNotification defined
type AppMessageNotification struct {
	// AMNId defined
	AMNId null.Int `xorm:"int(11) pk notnull autoincr 'a_m_n_id'" json:"a_m_n_id" form:"a_m_n_id" xml:"a_m_n_id"`
	// MessageContent defined
	MessageContent null.String `xorm:"varchar(10000) 'message_content'" json:"message_content" form:"message_content" xml:"message_content"`
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
	// AmnTitle defined
	AmnTitle null.String `xorm:"varchar(100) 'amn_title'" json:"amn_title" form:"amn_title" xml:"amn_title"`
}

// Marshal defined
func (m *AppMessageNotification) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *AppMessageNotification) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *AppMessageNotification) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *AppMessageNotification) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined AppMessageNotification
func (m *AppMessageNotification) TableName() string {
	return "app_message_notification"
}
