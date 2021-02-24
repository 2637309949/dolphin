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

// StudentClass defined
type StudentClass struct {
	// SCId defined
	SCId null.Int `xorm:"int(11) pk notnull autoincr 's_c_id'" json:"s_c_id" form:"s_c_id" xml:"s_c_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// CmId defined
	CmId null.Int `xorm:"int(11) 'cm_id'" json:"cm_id" form:"cm_id" xml:"cm_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// ScId defined
	ScId null.Int `xorm:"int(11) 'sc_id'" json:"sc_id" form:"sc_id" xml:"sc_id"`
	// IfCb defined
	IfCb null.Int `xorm:"int(11) 'if_cb'" json:"if_cb" form:"if_cb" xml:"if_cb"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// ScState defined
	ScState null.Int `xorm:"int(11) 'sc_state'" json:"sc_state" form:"sc_state" xml:"sc_state"`
	// ZcClass defined
	ZcClass null.Int `xorm:"int(11) 'zc_class'" json:"zc_class" form:"zc_class" xml:"zc_class"`
	// ScPrice defined
	ScPrice null.Int `xorm:"int(11) 'sc_price'" json:"sc_price" form:"sc_price" xml:"sc_price"`
	// JoinSource defined
	JoinSource null.Int `xorm:"int(11) 'join_source'" json:"join_source" form:"join_source" xml:"join_source"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// OsMoney defined
	OsMoney null.Float `xorm:"float(10,2) 'os_money'" json:"os_money" form:"os_money" xml:"os_money"`
	// FinalMoney defined
	FinalMoney null.Float `xorm:"float(10,2) 'final_money'" json:"final_money" form:"final_money" xml:"final_money"`
	// OnePrice defined
	OnePrice null.Float `xorm:"float(10,2) 'one_price'" json:"one_price" form:"one_price" xml:"one_price"`
	// RefundOnePrice defined
	RefundOnePrice null.Float `xorm:"float(10,2) 'refund_one_price'" json:"refund_one_price" form:"refund_one_price" xml:"refund_one_price"`
	// VisitId defined
	VisitId null.Int `xorm:"int(11) 'visit_id'" json:"visit_id" form:"visit_id" xml:"visit_id"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// OutDate defined
	OutDate null.Time `xorm:"datetime 'out_date'" json:"out_date" form:"out_date" xml:"out_date"`
	// Remark defined
	Remark null.String `xorm:"varchar(100) 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// Parser defined
func (m *StudentClass) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StudentClass) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentClass
func (m *StudentClass) TableName() string {
	return "student_class"
}
