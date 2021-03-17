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

// ActiveLesson defined
type ActiveLesson struct {
	// ALId defined
	ALId null.Int `xorm:"int(11) pk notnull autoincr 'a_l_id'" json:"a_l_id" form:"a_l_id" xml:"a_l_id"`
	// ActiveName defined
	ActiveName null.String `xorm:"varchar(100) 'active_name'" json:"active_name" form:"active_name" xml:"active_name"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// ActiveState defined
	ActiveState null.Int `xorm:"int(11) 'active_state'" json:"active_state" form:"active_state" xml:"active_state"`
	// MaxNum defined
	MaxNum null.Int `xorm:"int(11) 'max_num'" json:"max_num" form:"max_num" xml:"max_num"`
	// TeaId defined
	TeaId null.Int `xorm:"int(11) 'tea_id'" json:"tea_id" form:"tea_id" xml:"tea_id"`
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
	// TlDate defined
	TlDate null.Time `xorm:"datetime 'tl_date'" json:"tl_date" form:"tl_date" xml:"tl_date"`
	// RoomId defined
	RoomId null.Int `xorm:"int(11) 'room_id'" json:"room_id" form:"room_id" xml:"room_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// ActiveType defined
	ActiveType null.Int `xorm:"int(11) 'active_type'" json:"active_type" form:"active_type" xml:"active_type"`
	// ReportedNumber defined
	ReportedNumber null.Int `xorm:"int(11) 'reported_number'" json:"reported_number" form:"reported_number" xml:"reported_number"`
	// Note defined
	Note null.String `xorm:"varchar(200) 'note'" json:"note" form:"note" xml:"note"`
	// AgeGroup defined
	AgeGroup null.Int `xorm:"int(11) 'age_group'" json:"age_group" form:"age_group" xml:"age_group"`
	// BeforeId defined
	BeforeId null.Int `xorm:"int(11) 'before_id'" json:"before_id" form:"before_id" xml:"before_id"`
	// TkType defined
	TkType null.Int `xorm:"int(11) 'tk_type'" json:"tk_type" form:"tk_type" xml:"tk_type"`
}

// Marshal defined
func (m *ActiveLesson) With(s interface{}) (interface{}, error) {
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
func (m *ActiveLesson) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ActiveLesson) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ActiveLesson) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ActiveLesson) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ActiveLesson) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ActiveLesson) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ActiveLesson
func (m *ActiveLesson) TableName() string {
	return "active_lesson"
}
