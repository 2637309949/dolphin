// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/shopspring/decimal"
)

// StudentUseHourDetail defined
type StudentUseHourDetail struct {
	// SUHDId defined
	SUHDId null.Int `xorm:"int(11) pk notnull autoincr 's_u_h_d_id'" json:"s_u_h_d_id" form:"s_u_h_d_id" xml:"s_u_h_d_id"`
	// PkCsId defined
	PkCsId null.Int `xorm:"int(11) 'pk_cs_id'" json:"pk_cs_id" form:"pk_cs_id" xml:"pk_cs_id"`
	// PkCssId defined
	PkCssId null.Int `xorm:"int(11) 'pk_css_id'" json:"pk_css_id" form:"pk_css_id" xml:"pk_css_id"`
	// PkStuId defined
	PkStuId null.Int `xorm:"int(11) 'pk_stu_id'" json:"pk_stu_id" form:"pk_stu_id" xml:"pk_stu_id"`
	// UseHour defined
	UseHour null.Float `xorm:"float(11,2) 'use_hour'" json:"use_hour" form:"use_hour" xml:"use_hour"`
	// OnePrice defined
	OnePrice decimal.Decimal `xorm:"decimal(11,2) 'one_price'" json:"one_price" form:"one_price" xml:"one_price"`
	// AllPirce defined
	AllPirce decimal.Decimal `xorm:"decimal(11,2) 'all_pirce'" json:"all_pirce" form:"all_pirce" xml:"all_pirce"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"notnull 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// PkOf defined
	PkOf null.Int `xorm:"int(11) 'pk_of'" json:"pk_of" form:"pk_of" xml:"pk_of"`
	// DicHourType defined
	DicHourType null.Int `xorm:"int(11) 'dic_hour_type'" json:"dic_hour_type" form:"dic_hour_type" xml:"dic_hour_type"`
	// PkSchId defined
	PkSchId null.Int `xorm:"int(11) 'pk_sch_id'" json:"pk_sch_id" form:"pk_sch_id" xml:"pk_sch_id"`
	// JfOnePrice defined
	JfOnePrice decimal.Decimal `xorm:"decimal(50,3) 'jf_one_price'" json:"jf_one_price" form:"jf_one_price" xml:"jf_one_price"`
	// JfAllPrice defined
	JfAllPrice decimal.Decimal `xorm:"decimal(50,3) 'jf_all_price'" json:"jf_all_price" form:"jf_all_price" xml:"jf_all_price"`
	// KqKc defined
	KqKc null.Float `xorm:"float(50,2) 'kq_kc'" json:"kq_kc" form:"kq_kc" xml:"kq_kc"`
	// ChangStuId defined
	ChangStuId null.Int `xorm:"int(11) 'chang_stu_id'" json:"chang_stu_id" form:"chang_stu_id" xml:"chang_stu_id"`
	// ZdRsType defined
	ZdRsType null.Int `xorm:"int(11) 'zd_rs_type'" json:"zd_rs_type" form:"zd_rs_type" xml:"zd_rs_type"`
	// KqDate defined
	KqDate null.Time `xorm:"datetime 'kq_date'" json:"kq_date" form:"kq_date" xml:"kq_date"`
	// Remark defined
	Remark null.String `xorm:"varchar(20) 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// TableName table name of defined StudentUseHourDetail
func (m *StudentUseHourDetail) TableName() string {
	return "student_use_hour_detail"
}
