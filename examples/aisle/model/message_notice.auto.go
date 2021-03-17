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

// MessageNotice defined
type MessageNotice struct {
	// MNId defined
	MNId null.Int `xorm:"int(11) pk notnull autoincr 'm_n_id'" json:"m_n_id" form:"m_n_id" xml:"m_n_id"`
	// MessageContent defined
	MessageContent null.String `xorm:"varchar(10000) 'message_content'" json:"message_content" form:"message_content" xml:"message_content"`
	// NoticeDate defined
	NoticeDate null.Time `xorm:"datetime 'notice_date'" json:"notice_date" form:"notice_date" xml:"notice_date"`
	// SpId defined
	SpId null.Int `xorm:"int(11) 'sp_id'" json:"sp_id" form:"sp_id" xml:"sp_id"`
	// IfRead defined
	IfRead null.Int `xorm:"int(11) 'if_read'" json:"if_read" form:"if_read" xml:"if_read"`
	// ScsId defined
	ScsId null.Int `xorm:"int(11) 'scs_id'" json:"scs_id" form:"scs_id" xml:"scs_id"`
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
	// CsId defined
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
	// CctId defined
	CctId null.Int `xorm:"int(11) 'cct_id'" json:"cct_id" form:"cct_id" xml:"cct_id"`
	// ScfId defined
	ScfId null.Int `xorm:"int(11) 'scf_id'" json:"scf_id" form:"scf_id" xml:"scf_id"`
	// MessageType defined
	MessageType null.Int `xorm:"int(11) 'message_type'" json:"message_type" form:"message_type" xml:"message_type"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// AmnId defined
	AmnId null.Int `xorm:"int(11) 'amn_id'" json:"amn_id" form:"amn_id" xml:"amn_id"`
	// ParId defined
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" form:"par_id" xml:"par_id"`
	// IfSend defined
	IfSend null.Int `xorm:"int(11) 'if_send'" json:"if_send" form:"if_send" xml:"if_send"`
	// SaaId defined
	SaaId null.Int `xorm:"int(11) 'saa_id'" json:"saa_id" form:"saa_id" xml:"saa_id"`
}

// Marshal defined
func (m *MessageNotice) With(s interface{}) (interface{}, error) {
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
func (m *MessageNotice) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MessageNotice) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *MessageNotice) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *MessageNotice) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *MessageNotice) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MessageNotice) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MessageNotice
func (m *MessageNotice) TableName() string {
	return "message_notice"
}
