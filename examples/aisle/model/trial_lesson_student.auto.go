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

// TrialLessonStudent defined
type TrialLessonStudent struct {
	// TLSId defined
	TLSId null.Int `xorm:"int(11) pk notnull autoincr 't_l_s_id'" json:"t_l_s_id" form:"t_l_s_id" xml:"t_l_s_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// TlId defined
	TlId null.Int `xorm:"int(11) 'tl_id'" json:"tl_id" form:"tl_id" xml:"tl_id"`
	// TlsState defined
	TlsState null.Int `xorm:"int(11) 'tls_state'" json:"tls_state" form:"tls_state" xml:"tls_state"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// SvId defined
	SvId null.Int `xorm:"int(11) 'sv_id'" json:"sv_id" form:"sv_id" xml:"sv_id"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// TlDate defined
	TlDate null.Time `xorm:"datetime 'tl_date'" json:"tl_date" form:"tl_date" xml:"tl_date"`
	// IvId defined
	IvId null.Int `xorm:"int(11) 'iv_id'" json:"iv_id" form:"iv_id" xml:"iv_id"`
	// YjQdDate defined
	YjQdDate null.Time `xorm:"datetime 'yj_qd_date'" json:"yj_qd_date" form:"yj_qd_date" xml:"yj_qd_date"`
}

// With defined
func (m *TrialLessonStudent) With(s interface{}) (interface{}, error) {
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
func (m *TrialLessonStudent) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TrialLessonStudent) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *TrialLessonStudent) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *TrialLessonStudent) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *TrialLessonStudent) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *TrialLessonStudent) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TrialLessonStudent
func (m *TrialLessonStudent) TableName() string {
	return "trial_lesson_student"
}
