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

// TeachLog defined
type TeachLog struct {
	// TeachLogId defined
	TeachLogId null.Int `xorm:"int(11) pk notnull autoincr 'teach_log_id'" json:"teach_log_id" form:"teach_log_id" xml:"teach_log_id"`
	// TlContent defined
	TlContent null.String `xorm:"varchar(500) 'tl_content'" json:"tl_content" form:"tl_content" xml:"tl_content"`
	// TlStuShow defined
	TlStuShow null.String `xorm:"varchar(500) 'tl_stu_show'" json:"tl_stu_show" form:"tl_stu_show" xml:"tl_stu_show"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// TlType defined
	TlType null.Int `xorm:"int(11) 'tl_type'" json:"tl_type" form:"tl_type" xml:"tl_type"`
	// CsId defined
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
	// AppId defined
	AppId null.Int `xorm:"int(11) 'app_id'" json:"app_id" form:"app_id" xml:"app_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// With defined
func (m *TeachLog) With(s interface{}) (interface{}, error) {
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
func (m *TeachLog) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TeachLog) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *TeachLog) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *TeachLog) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *TeachLog) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *TeachLog) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TeachLog
func (m *TeachLog) TableName() string {
	return "teach_log"
}
