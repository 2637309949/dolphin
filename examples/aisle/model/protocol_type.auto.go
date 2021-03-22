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

// ProtocolType defined
type ProtocolType struct {
	// PTId defined
	PTId null.Int `xorm:"int(11) pk notnull autoincr 'p_t_id'" json:"p_t_id" form:"p_t_id" xml:"p_t_id"`
	// CourseName defined
	CourseName null.String `xorm:"varchar(100) 'course_name'" json:"course_name" form:"course_name" xml:"course_name"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"notnull 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// PtState defined
	PtState null.Int `xorm:"int(11) 'pt_state'" json:"pt_state" form:"pt_state" xml:"pt_state"`
	// PtDescribe defined
	PtDescribe null.Int `xorm:"int(11) 'pt_describe'" json:"pt_describe" form:"pt_describe" xml:"pt_describe"`
	// PtLevel defined
	PtLevel null.Int `xorm:"int(11) 'pt_level'" json:"pt_level" form:"pt_level" xml:"pt_level"`
	// TypeName defined
	TypeName null.String `xorm:"varchar(50) 'type_name'" json:"type_name" form:"type_name" xml:"type_name"`
	// PtContent defined
	PtContent null.Int `xorm:"int(11) 'pt_content'" json:"pt_content" form:"pt_content" xml:"pt_content"`
}

// With defined
func (m *ProtocolType) With(s interface{}) (interface{}, error) {
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
func (m *ProtocolType) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ProtocolType) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ProtocolType) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ProtocolType) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ProtocolType) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *ProtocolType) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ProtocolType
func (m *ProtocolType) TableName() string {
	return "protocol_type"
}
