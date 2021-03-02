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

// ClassType defined
type ClassType struct {
	// ClassTypeId defined
	ClassTypeId null.Int `xorm:"int(11) pk notnull autoincr 'class_type_id'" json:"class_type_id" form:"class_type_id" xml:"class_type_id"`
	// CtName defined
	CtName null.String `xorm:"varchar(50) 'ct_name'" json:"ct_name" form:"ct_name" xml:"ct_name"`
	// CtRemark defined
	CtRemark null.String `xorm:"varchar(100) 'ct_remark'" json:"ct_remark" form:"ct_remark" xml:"ct_remark"`
	// CtNumber defined
	CtNumber null.Int `xorm:"int(11) 'ct_number'" json:"ct_number" form:"ct_number" xml:"ct_number"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// CtMoney defined
	CtMoney null.Float `xorm:"float(50,2) 'ct_money'" json:"ct_money" form:"ct_money" xml:"ct_money"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// CtType defined
	CtType null.Int `xorm:"int(11) 'ct_type'" json:"ct_type" form:"ct_type" xml:"ct_type"`
	// CtOnePrice defined
	CtOnePrice null.Float `xorm:"float(50,2) 'ct_one_price'" json:"ct_one_price" form:"ct_one_price" xml:"ct_one_price"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// CtSort defined
	CtSort null.Int `xorm:"int(11) 'ct_sort'" json:"ct_sort" form:"ct_sort" xml:"ct_sort"`
	// CtMinit defined
	CtMinit null.Float `xorm:"float(10,2) 'ct_minit'" json:"ct_minit" form:"ct_minit" xml:"ct_minit"`
	// OrgType defined
	OrgType null.Int `xorm:"int(11) 'org_type'" json:"org_type" form:"org_type" xml:"org_type"`
	// Ifmaincourse defined
	Ifmaincourse null.Int `xorm:"int(11) default(0) 'ifmaincourse'" json:"ifmaincourse" form:"ifmaincourse" xml:"ifmaincourse"`
	// CtCourseType defined
	CtCourseType null.Int `xorm:"int(11) 'ct_course_type'" json:"ct_course_type" form:"ct_course_type" xml:"ct_course_type"`
	// EnglishName defined
	EnglishName null.String `xorm:"varchar(100) 'english_name'" json:"english_name" form:"english_name" xml:"english_name"`
	// AchievementNum defined
	AchievementNum null.Float `xorm:"float(50,2) 'achievement_num'" json:"achievement_num" form:"achievement_num" xml:"achievement_num"`
	// OrderEdition defined
	OrderEdition null.Int `xorm:"int(11) 'order_edition'" json:"order_edition" form:"order_edition" xml:"order_edition"`
	// Mzks defined
	Mzks null.Float `xorm:"float(50,2) 'mzks'" json:"mzks" form:"mzks" xml:"mzks"`
}

// Marshal defined
func (m *ClassType) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ClassType) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *ClassType) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ClassType) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ClassType
func (m *ClassType) TableName() string {
	return "class_type"
}
