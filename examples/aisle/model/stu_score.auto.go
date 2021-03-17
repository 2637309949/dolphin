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

// StuScore defined
type StuScore struct {
	// StuScoreId defined
	StuScoreId null.Int `xorm:"int(11) pk notnull autoincr 'stu_score_id'" json:"stu_score_id" form:"stu_score_id" xml:"stu_score_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// TmId defined
	TmId null.Int `xorm:"int(11) 'tm_id'" json:"tm_id" form:"tm_id" xml:"tm_id"`
	// LevelId defined
	LevelId null.Int `xorm:"int(11) 'level_id'" json:"level_id" form:"level_id" xml:"level_id"`
	// UnitId defined
	UnitId null.Int `xorm:"int(11) 'unit_id'" json:"unit_id" form:"unit_id" xml:"unit_id"`
	// TestTime defined
	TestTime null.Time `xorm:"datetime 'test_time'" json:"test_time" form:"test_time" xml:"test_time"`
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
	// AllScore defined
	AllScore null.Float `xorm:"float(50,2) 'all_score'" json:"all_score" form:"all_score" xml:"all_score"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// TauserId defined
	TauserId null.Int `xorm:"int(11) 'tauser_id'" json:"tauser_id" form:"tauser_id" xml:"tauser_id"`
}

// Marshal defined
func (m *StuScore) With(s interface{}) (interface{}, error) {
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
func (m *StuScore) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StuScore) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StuScore) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StuScore) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StuScore) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuScore) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuScore
func (m *StuScore) TableName() string {
	return "stu_score"
}
