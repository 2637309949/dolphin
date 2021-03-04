// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// StudentClassYpkDkb defined
type StudentClassYpkDkb struct {
	// SCYId defined
	SCYId null.Int `xorm:"int(11) pk notnull autoincr 's_c_y_id'" json:"s_c_y_id" form:"s_c_y_id" xml:"s_c_y_id"`
	// PkYpk defined
	PkYpk null.Int `xorm:"int(11) 'pk_ypk'" json:"pk_ypk" form:"pk_ypk" xml:"pk_ypk"`
	// PkStu defined
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" form:"pk_stu" xml:"pk_stu"`
	// Kc defined
	Kc null.Float `xorm:"float(11,2) 'kc'" json:"kc" form:"kc" xml:"kc"`
	// OnePrice defined
	OnePrice decimal.Decimal `xorm:"decimal(11,2) 'one_price'" json:"one_price" form:"one_price" xml:"one_price"`
	// AllPrice defined
	AllPrice decimal.Decimal `xorm:"decimal(11,2) 'all_price'" json:"all_price" form:"all_price" xml:"all_price"`
	// PkSch defined
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
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
	// ClassDate defined
	ClassDate null.Time `xorm:"datetime 'class_date'" json:"class_date" form:"class_date" xml:"class_date"`
	// PkSct defined
	PkSct null.Int `xorm:"int(11) 'pk_sct'" json:"pk_sct" form:"pk_sct" xml:"pk_sct"`
	// PkCt defined
	PkCt null.Int `xorm:"int(11) 'pk_ct'" json:"pk_ct" form:"pk_ct" xml:"pk_ct"`
}

// Marshal defined
func (m *StudentClassYpkDkb) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StudentClassYpkDkb) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StudentClassYpkDkb) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StudentClassYpkDkb) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StudentClassYpkDkb) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StudentClassYpkDkb) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentClassYpkDkb
func (m *StudentClassYpkDkb) TableName() string {
	return "student_class_ypk_dkb"
}
