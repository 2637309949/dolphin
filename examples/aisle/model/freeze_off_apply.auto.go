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

// FreezeOffApply defined
type FreezeOffApply struct {
	// FOAId defined
	FOAId null.Int `xorm:"int(11) pk notnull autoincr 'f_o_a_id'" json:"f_o_a_id" form:"f_o_a_id" xml:"f_o_a_id"`
	// FreezeOffDesc defined
	FreezeOffDesc null.String `xorm:"varchar(500) 'freeze_off_desc'" json:"freeze_off_desc" form:"freeze_off_desc" xml:"freeze_off_desc"`
	// FreezeOffTime defined
	FreezeOffTime null.Time `xorm:"datetime 'freeze_off_time'" json:"freeze_off_time" form:"freeze_off_time" xml:"freeze_off_time"`
	// FreezeOffAuditing defined
	FreezeOffAuditing null.Int `xorm:"int(11) 'freeze_off_auditing'" json:"freeze_off_auditing" form:"freeze_off_auditing" xml:"freeze_off_auditing"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
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
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// ShenhePerson defined
	ShenhePerson null.Int `xorm:"int(11) 'shenhe_person'" json:"shenhe_person" form:"shenhe_person" xml:"shenhe_person"`
}

// With defined
func (m *FreezeOffApply) With(s interface{}) (interface{}, error) {
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
func (m *FreezeOffApply) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *FreezeOffApply) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *FreezeOffApply) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *FreezeOffApply) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *FreezeOffApply) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *FreezeOffApply) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined FreezeOffApply
func (m *FreezeOffApply) TableName() string {
	return "freeze_off_apply"
}
