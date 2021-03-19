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

// MaTeacher defined
type MaTeacher struct {
	// MaTeacherId defined
	MaTeacherId null.Int `xorm:"int(11) pk notnull autoincr 'ma_teacher_id'" json:"ma_teacher_id" form:"ma_teacher_id" xml:"ma_teacher_id"`
	// MaId defined
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
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
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// TeaClassDate defined
	TeaClassDate null.Time `xorm:"datetime 'tea_class_date'" json:"tea_class_date" form:"tea_class_date" xml:"tea_class_date"`
	// TeaBegintime defined
	TeaBegintime null.Time `xorm:"datetime 'tea_begintime'" json:"tea_begintime" form:"tea_begintime" xml:"tea_begintime"`
	// TeaEndtime defined
	TeaEndtime null.Time `xorm:"datetime 'tea_endtime'" json:"tea_endtime" form:"tea_endtime" xml:"tea_endtime"`
	// IfKouHour defined
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
	// ClassHour defined
	ClassHour null.Float `xorm:"float(50,2) 'class_hour'" json:"class_hour" form:"class_hour" xml:"class_hour"`
}

// With defined
func (m *MaTeacher) With(s interface{}) (interface{}, error) {
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
func (m *MaTeacher) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *MaTeacher) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *MaTeacher) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *MaTeacher) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *MaTeacher) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *MaTeacher) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MaTeacher
func (m *MaTeacher) TableName() string {
	return "ma_teacher"
}
