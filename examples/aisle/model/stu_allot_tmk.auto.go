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

// StuAllotTmk defined
type StuAllotTmk struct {
	// SATId defined
	SATId null.Int `xorm:"int(11) pk notnull autoincr 's_a_t_id'" json:"s_a_t_id" form:"s_a_t_id" xml:"s_a_t_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// TmkUserid defined
	TmkUserid null.Int `xorm:"int(11) 'tmk_userid'" json:"tmk_userid" form:"tmk_userid" xml:"tmk_userid"`
	// AllotState defined
	AllotState null.Int `xorm:"int(11) 'allot_state'" json:"allot_state" form:"allot_state" xml:"allot_state"`
	// AllotDate defined
	AllotDate null.Time `xorm:"datetime 'allot_date'" json:"allot_date" form:"allot_date" xml:"allot_date"`
	// QxAllotDate defined
	QxAllotDate null.Time `xorm:"datetime 'qx_allot_date'" json:"qx_allot_date" form:"qx_allot_date" xml:"qx_allot_date"`
	// HisTmkUserid defined
	HisTmkUserid null.Int `xorm:"int(11) 'his_tmk_userid'" json:"his_tmk_userid" form:"his_tmk_userid" xml:"his_tmk_userid"`
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
	// AllotDesc defined
	AllotDesc null.String `xorm:"varchar(5000) 'allot_desc'" json:"allot_desc" form:"allot_desc" xml:"allot_desc"`
	// IfPl defined
	IfPl null.Int `xorm:"int(11) 'if_pl'" json:"if_pl" form:"if_pl" xml:"if_pl"`
}

// With defined
func (m *StuAllotTmk) With(s interface{}) (interface{}, error) {
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
func (m *StuAllotTmk) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StuAllotTmk) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StuAllotTmk) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StuAllotTmk) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StuAllotTmk) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *StuAllotTmk) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuAllotTmk
func (m *StuAllotTmk) TableName() string {
	return "stu_allot_tmk"
}
