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

// StudentTypeJournal defined
type StudentTypeJournal struct {
	// STJId defined
	STJId null.Int `xorm:"int(11) pk notnull autoincr 's_t_j_id'" json:"s_t_j_id" form:"s_t_j_id" xml:"s_t_j_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// StuType defined
	StuType null.Int `xorm:"int(11) 'stu_type'" json:"stu_type" form:"stu_type" xml:"stu_type"`
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
	// StuTypeName defined
	StuTypeName null.String `xorm:"varchar(100) 'stu_type_name'" json:"stu_type_name" form:"stu_type_name" xml:"stu_type_name"`
	// Stutypedesc defined
	Stutypedesc null.String `xorm:"varchar(500) 'stutypedesc'" json:"stutypedesc" form:"stutypedesc" xml:"stutypedesc"`
	// OldStuType defined
	OldStuType null.Int `xorm:"int(11) 'old_stu_type'" json:"old_stu_type" form:"old_stu_type" xml:"old_stu_type"`
	// OldStuTypeName defined
	OldStuTypeName null.String `xorm:"varchar(1000) 'old_stu_type_name'" json:"old_stu_type_name" form:"old_stu_type_name" xml:"old_stu_type_name"`
}

// With defined
func (m *StudentTypeJournal) With(s interface{}) (interface{}, error) {
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
func (m *StudentTypeJournal) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StudentTypeJournal) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StudentTypeJournal) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StudentTypeJournal) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StudentTypeJournal) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *StudentTypeJournal) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentTypeJournal
func (m *StudentTypeJournal) TableName() string {
	return "student_type_journal"
}
