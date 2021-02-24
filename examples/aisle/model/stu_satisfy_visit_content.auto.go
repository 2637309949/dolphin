// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// StuSatisfyVisitContent defined
type StuSatisfyVisitContent struct {
	//
	SSVCId null.Int `xorm:"int(11) pk notnull autoincr 's_s_v_c_id'" json:"s_s_v_c_id" form:"s_s_v_c_id" xml:"s_s_v_c_id"`
	//
	SsvId null.Int `xorm:"int(11) 'ssv_id'" json:"ssv_id" form:"ssv_id" xml:"ssv_id"`
	//
	SvmcId null.Int `xorm:"int(11) 'svmc_id'" json:"svmc_id" form:"svmc_id" xml:"svmc_id"`
	//
	Fraction null.Int `xorm:"int(11) 'fraction'" json:"fraction" form:"fraction" xml:"fraction"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	CustomerId null.Int `xorm:"int(11) 'customer_id'" json:"customer_id" form:"customer_id" xml:"customer_id"`
	//
	SaleId null.Int `xorm:"int(11) 'sale_id'" json:"sale_id" form:"sale_id" xml:"sale_id"`
	//
	TeacherId null.Int `xorm:"int(11) 'teacher_id'" json:"teacher_id" form:"teacher_id" xml:"teacher_id"`
}

// Parser defined
func (m *StuSatisfyVisitContent) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StuSatisfyVisitContent) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuSatisfyVisitContent
func (m *StuSatisfyVisitContent) TableName() string {
	return "stu_satisfy_visit_content"
}
