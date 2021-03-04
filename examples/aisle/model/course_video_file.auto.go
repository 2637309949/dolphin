// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// CourseVideoFile defined
type CourseVideoFile struct {
	// CVFId defined
	CVFId null.Int `xorm:"int(11) pk notnull autoincr 'c_v_f_id'" json:"c_v_f_id" form:"c_v_f_id" xml:"c_v_f_id"`
	// CvId defined
	CvId null.Int `xorm:"int(11) 'cv_id'" json:"cv_id" form:"cv_id" xml:"cv_id"`
	// Filed defined
	Filed null.Int `xorm:"int(11) 'filed'" json:"filed" form:"filed" xml:"filed"`
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
}

// Marshal defined
func (m *CourseVideoFile) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *CourseVideoFile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *CourseVideoFile) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *CourseVideoFile) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *CourseVideoFile) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *CourseVideoFile) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CourseVideoFile
func (m *CourseVideoFile) TableName() string {
	return "course_video_file"
}
