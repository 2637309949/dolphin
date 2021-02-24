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

// StudentCzIntegral defined
type StudentCzIntegral struct {
	//
	SCIId null.Int `xorm:"int(11) pk notnull autoincr 's_c_i_id'" json:"s_c_i_id" form:"s_c_i_id" xml:"s_c_i_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	//
	Integral null.Float `xorm:"float(50,2) 'integral'" json:"integral" form:"integral" xml:"integral"`
	//
	IntegralDesc null.String `xorm:"varchar(500) 'integral_desc'" json:"integral_desc" form:"integral_desc" xml:"integral_desc"`
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
	IntegralState null.Int `xorm:"int(11) 'integral_state'" json:"integral_state" form:"integral_state" xml:"integral_state"`
	//
	JfOrder null.Int `xorm:"int(11) 'jf_order'" json:"jf_order" form:"jf_order" xml:"jf_order"`
	//
	YhjfOrderid null.Int `xorm:"int(11) 'yhjf_orderid'" json:"yhjf_orderid" form:"yhjf_orderid" xml:"yhjf_orderid"`
	//
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" form:"sct_id" xml:"sct_id"`
}

// Parser defined
func (m *SysCommentReply) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysCommentReply) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentCzIntegral
func (m *StudentCzIntegral) TableName() string {
	return "student_cz_integral"
}
