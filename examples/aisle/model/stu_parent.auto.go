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

// StuParent defined
type StuParent struct {
	// StuParentId defined
	StuParentId null.Int `xorm:"int(11) pk notnull autoincr 'stu_parent_id'" json:"stu_parent_id" form:"stu_parent_id" xml:"stu_parent_id"`
	// StuParentName defined
	StuParentName null.String `xorm:"varchar(1000) 'stu_parent_name'" json:"stu_parent_name" form:"stu_parent_name" xml:"stu_parent_name"`
	// StuParentPhone defined
	StuParentPhone null.String `xorm:"varchar(11) 'stu_parent_phone'" json:"stu_parent_phone" form:"stu_parent_phone" xml:"stu_parent_phone"`
	// StuParentPassword defined
	StuParentPassword null.String `xorm:"varchar(1000) 'stu_parent_password'" json:"stu_parent_password" form:"stu_parent_password" xml:"stu_parent_password"`
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
	// UserLoginState defined
	UserLoginState null.Int `xorm:"int(11) 'user_login_state'" json:"user_login_state" form:"user_login_state" xml:"user_login_state"`
	// TempPhone defined
	TempPhone null.String `xorm:"varchar(11) 'temp_phone'" json:"temp_phone" form:"temp_phone" xml:"temp_phone"`
	// JpushCodePar defined
	JpushCodePar null.String `xorm:"varchar(100) 'jpush_code_par'" json:"jpush_code_par" form:"jpush_code_par" xml:"jpush_code_par"`
	// ParAppversion defined
	ParAppversion null.String `xorm:"varchar(20) 'par_appversion'" json:"par_appversion" form:"par_appversion" xml:"par_appversion"`
	// ParAppneedupdate defined
	ParAppneedupdate null.Int `xorm:"int(11) 'par_appneedupdate'" json:"par_appneedupdate" form:"par_appneedupdate" xml:"par_appneedupdate"`
	// IdentificationNumber defined
	IdentificationNumber null.String `xorm:"varchar(1000) 'identification_number'" json:"identification_number" form:"identification_number" xml:"identification_number"`
}

// Marshal defined
func (m *StuParent) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StuParent) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StuParent) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StuParent) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StuParent) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuParent) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuParent
func (m *StuParent) TableName() string {
	return "stu_parent"
}
