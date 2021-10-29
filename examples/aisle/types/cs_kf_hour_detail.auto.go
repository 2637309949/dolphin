// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/shopspring/decimal"
)

// CsKfHourDetail defined
type CsKfHourDetail struct {
	// CKHDId defined
	CKHDId null.Int `xorm:"int(11) pk notnull autoincr 'c_k_h_d_id'" json:"c_k_h_d_id" form:"c_k_h_d_id" xml:"c_k_h_d_id"`
	// PkCs defined
	PkCs null.Int `xorm:"int(11) 'pk_cs'" json:"pk_cs" form:"pk_cs" xml:"pk_cs"`
	// PkKf defined
	PkKf null.Int `xorm:"int(11) 'pk_kf'" json:"pk_kf" form:"pk_kf" xml:"pk_kf"`
	// KqHour defined
	KqHour null.Float `xorm:"float(11,2) 'kq_hour'" json:"kq_hour" form:"kq_hour" xml:"kq_hour"`
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
	// KqDate defined
	KqDate null.Time `xorm:"datetime 'kq_date'" json:"kq_date" form:"kq_date" xml:"kq_date"`
	// PkCss defined
	PkCss null.Int `xorm:"int(11) 'pk_css'" json:"pk_css" form:"pk_css" xml:"pk_css"`
	// PkCsckfId defined
	PkCsckfId null.Int `xorm:"int(11) 'pk_csckf_id'" json:"pk_csckf_id" form:"pk_csckf_id" xml:"pk_csckf_id"`
	// PkCssckfId defined
	PkCssckfId null.Int `xorm:"int(11) 'pk_cssckf_id'" json:"pk_cssckf_id" form:"pk_cssckf_id" xml:"pk_cssckf_id"`
	// ZdRsType defined
	ZdRsType null.Int `xorm:"int(11) 'zd_rs_type'" json:"zd_rs_type" form:"zd_rs_type" xml:"zd_rs_type"`
	// PkStu defined
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" form:"pk_stu" xml:"pk_stu"`
	// KxOnePrice defined
	KxOnePrice decimal.Decimal `xorm:"decimal(50,3) 'kx_one_price'" json:"kx_one_price" form:"kx_one_price" xml:"kx_one_price"`
	// KxAllPrice defined
	KxAllPrice decimal.Decimal `xorm:"decimal(50,3) 'kx_all_price'" json:"kx_all_price" form:"kx_all_price" xml:"kx_all_price"`
	// JfOnePrice defined
	JfOnePrice decimal.Decimal `xorm:"decimal(50,3) 'jf_one_price'" json:"jf_one_price" form:"jf_one_price" xml:"jf_one_price"`
	// JfAllPrice defined
	JfAllPrice decimal.Decimal `xorm:"decimal(50,3) 'jf_all_price'" json:"jf_all_price" form:"jf_all_price" xml:"jf_all_price"`
	// PkClass defined
	PkClass null.Int `xorm:"int(11) 'pk_class'" json:"pk_class" form:"pk_class" xml:"pk_class"`
	// KqKc defined
	KqKc null.Float `xorm:"float(50,2) 'kq_kc'" json:"kq_kc" form:"kq_kc" xml:"kq_kc"`
	// PkSch defined
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
}

// TableName table name of defined CsKfHourDetail
func (m *CsKfHourDetail) TableName() string {
	return "cs_kf_hour_detail"
}

func (r *CsKfHourDetail) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCsKfHourDetail(data []byte) (CsKfHourDetail, error) {
	var r CsKfHourDetail
	err := json.Unmarshal(data, &r)
	return r, err
}
