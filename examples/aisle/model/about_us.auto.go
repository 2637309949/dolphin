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

// AboutUs defined
type AboutUs struct {
	// AboutUsId defined
	AboutUsId null.Int `xorm:"int(11) pk notnull autoincr 'about_us_id'" json:"about_us_id" form:"about_us_id" xml:"about_us_id"`
	// AboutUsPicture defined
	AboutUsPicture null.Int `xorm:"int(11) 'about_us_picture'" json:"about_us_picture" form:"about_us_picture" xml:"about_us_picture"`
	// AboutUsContent defined
	AboutUsContent null.String `xorm:"varchar(10000) 'about_us_content'" json:"about_us_content" form:"about_us_content" xml:"about_us_content"`
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
	// AboutUsVideo defined
	AboutUsVideo null.Int `xorm:"int(11) 'about_us_video'" json:"about_us_video" form:"about_us_video" xml:"about_us_video"`
	// AboutUsCntitle defined
	AboutUsCntitle null.String `xorm:"varchar(1000) 'about_us_cntitle'" json:"about_us_cntitle" form:"about_us_cntitle" xml:"about_us_cntitle"`
	// AboutUsEntitle defined
	AboutUsEntitle null.String `xorm:"varchar(1000) 'about_us_entitle'" json:"about_us_entitle" form:"about_us_entitle" xml:"about_us_entitle"`
}

// With defined
func (m *AboutUs) With(s interface{}) (interface{}, error) {
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
func (m *AboutUs) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *AboutUs) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *AboutUs) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *AboutUs) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *AboutUs) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *AboutUs) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined AboutUs
func (m *AboutUs) TableName() string {
	return "about_us"
}
