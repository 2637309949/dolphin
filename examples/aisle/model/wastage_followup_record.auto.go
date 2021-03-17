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

// WastageFollowupRecord defined
type WastageFollowupRecord struct {
	// WFRId defined
	WFRId null.Int `xorm:"int(11) pk notnull autoincr 'w_f_r_id'" json:"w_f_r_id" form:"w_f_r_id" xml:"w_f_r_id"`
	// EnglishLevel defined
	EnglishLevel null.String `xorm:"varchar(20) 'english_level'" json:"english_level" form:"english_level" xml:"english_level"`
	// EnglishSchool defined
	EnglishSchool null.String `xorm:"varchar(50) 'english_school'" json:"english_school" form:"english_school" xml:"english_school"`
	// IfStudy defined
	IfStudy null.Int `xorm:"int(11) 'if_study'" json:"if_study" form:"if_study" xml:"if_study"`
	// TrackingProgram defined
	TrackingProgram null.String `xorm:"varchar(200) 'tracking_program'" json:"tracking_program" form:"tracking_program" xml:"tracking_program"`
	// WastageMonth defined
	WastageMonth null.Float `xorm:"float(10,2) 'wastage_month'" json:"wastage_month" form:"wastage_month" xml:"wastage_month"`
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
}

// Marshal defined
func (m *WastageFollowupRecord) With(s interface{}) (interface{}, error) {
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
func (m *WastageFollowupRecord) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *WastageFollowupRecord) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *WastageFollowupRecord) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *WastageFollowupRecord) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *WastageFollowupRecord) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *WastageFollowupRecord) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined WastageFollowupRecord
func (m *WastageFollowupRecord) TableName() string {
	return "wastage_followup_record"
}
