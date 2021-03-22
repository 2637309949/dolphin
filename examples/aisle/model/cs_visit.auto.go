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

// CsVisit defined
type CsVisit struct {
	// CsVisitId defined
	CsVisitId null.Int `xorm:"int(11) pk notnull autoincr 'cs_visit_id'" json:"cs_visit_id" form:"cs_visit_id" xml:"cs_visit_id"`
	// VisitContent defined
	VisitContent null.String `xorm:"varchar(300) 'visit_content'" json:"visit_content" form:"visit_content" xml:"visit_content"`
	// VisitStartTime defined
	VisitStartTime null.Time `xorm:"datetime 'visit_start_time'" json:"visit_start_time" form:"visit_start_time" xml:"visit_start_time"`
	// VisitEndTime defined
	VisitEndTime null.Time `xorm:"datetime 'visit_end_time'" json:"visit_end_time" form:"visit_end_time" xml:"visit_end_time"`
	// NextVisitTime defined
	NextVisitTime null.Time `xorm:"datetime 'next_visit_time'" json:"next_visit_time" form:"next_visit_time" xml:"next_visit_time"`
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
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// VisituserId defined
	VisituserId null.Int `xorm:"int(11) 'visituser_id'" json:"visituser_id" form:"visituser_id" xml:"visituser_id"`
}

// With defined
func (m *CsVisit) With(s interface{}) (interface{}, error) {
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
func (m *CsVisit) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *CsVisit) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *CsVisit) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *CsVisit) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *CsVisit) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *CsVisit) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CsVisit
func (m *CsVisit) TableName() string {
	return "cs_visit"
}
