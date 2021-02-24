// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// CsKfHourDetail defined
type CsKfHourDetail struct {
	//
	CKHDId null.Int `xorm:"int(11) pk notnull autoincr 'c_k_h_d_id'" json:"c_k_h_d_id" form:"c_k_h_d_id" xml:"c_k_h_d_id"`
	//
	PkCs null.Int `xorm:"int(11) 'pk_cs'" json:"pk_cs" form:"pk_cs" xml:"pk_cs"`
	//
	PkKf null.Int `xorm:"int(11) 'pk_kf'" json:"pk_kf" form:"pk_kf" xml:"pk_kf"`
	//
	KqHour null.Float `xorm:"float(11,2) 'kq_hour'" json:"kq_hour" form:"kq_hour" xml:"kq_hour"`
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
	KqDate null.Time `xorm:"datetime 'kq_date'" json:"kq_date" form:"kq_date" xml:"kq_date"`
	//
	PkCss null.Int `xorm:"int(11) 'pk_css'" json:"pk_css" form:"pk_css" xml:"pk_css"`
	//
	PkCsckfId null.Int `xorm:"int(11) 'pk_csckf_id'" json:"pk_csckf_id" form:"pk_csckf_id" xml:"pk_csckf_id"`
	//
	PkCssckfId null.Int `xorm:"int(11) 'pk_cssckf_id'" json:"pk_cssckf_id" form:"pk_cssckf_id" xml:"pk_cssckf_id"`
	//
	ZdRsType null.Int `xorm:"int(11) 'zd_rs_type'" json:"zd_rs_type" form:"zd_rs_type" xml:"zd_rs_type"`
	//
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" form:"pk_stu" xml:"pk_stu"`
	//
	KxOnePrice decimal.Decimal `xorm:"decimal(50,3) 'kx_one_price'" json:"kx_one_price" form:"kx_one_price" xml:"kx_one_price"`
	//
	KxAllPrice decimal.Decimal `xorm:"decimal(50,3) 'kx_all_price'" json:"kx_all_price" form:"kx_all_price" xml:"kx_all_price"`
	//
	JfOnePrice decimal.Decimal `xorm:"decimal(50,3) 'jf_one_price'" json:"jf_one_price" form:"jf_one_price" xml:"jf_one_price"`
	//
	JfAllPrice decimal.Decimal `xorm:"decimal(50,3) 'jf_all_price'" json:"jf_all_price" form:"jf_all_price" xml:"jf_all_price"`
	//
	PkClass null.Int `xorm:"int(11) 'pk_class'" json:"pk_class" form:"pk_class" xml:"pk_class"`
	//
	KqKc null.Float `xorm:"float(50,2) 'kq_kc'" json:"kq_kc" form:"kq_kc" xml:"kq_kc"`
	//
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
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

// TableName table name of defined CsKfHourDetail
func (m *CsKfHourDetail) TableName() string {
	return "cs_kf_hour_detail"
}
