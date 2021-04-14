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

// StudentOrganSchool defined
type StudentOrganSchool struct {
	// SOSId defined
	SOSId null.Int `xorm:"int(11) pk notnull autoincr 's_o_s_id'" json:"s_o_s_id" form:"s_o_s_id" xml:"s_o_s_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// StudentId defined
	StudentId null.Int `xorm:"int(11) 'student_id'" json:"student_id" form:"student_id" xml:"student_id"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// SyId defined
	SyId null.Int `xorm:"int(11) 'sy_id'" json:"sy_id" form:"sy_id" xml:"sy_id"`
	// StuGxYy defined
	StuGxYy null.Int `xorm:"int(11) 'stu_gx_yy'" json:"stu_gx_yy" form:"stu_gx_yy" xml:"stu_gx_yy"`
	// SfPlzx defined
	SfPlzx null.Int `xorm:"int(11) 'sf_plzx'" json:"sf_plzx" form:"sf_plzx" xml:"sf_plzx"`
}

// With defined
func (m *StudentOrganSchool) With(s interface{}) (interface{}, error) {
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
func (m *StudentOrganSchool) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StudentOrganSchool) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StudentOrganSchool) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StudentOrganSchool) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StudentOrganSchool) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *StudentOrganSchool) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentOrganSchool
func (m *StudentOrganSchool) TableName() string {
	return "student_organ_school"
}
