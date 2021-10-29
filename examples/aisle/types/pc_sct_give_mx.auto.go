// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// PcSctGiveMx defined
type PcSctGiveMx struct {
	// PSGMId defined
	PSGMId null.Int `xorm:"int(11) pk notnull autoincr 'p_s_g_m_id'" json:"p_s_g_m_id" form:"p_s_g_m_id" xml:"p_s_g_m_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// SctId defined
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" form:"sct_id" xml:"sct_id"`
	// SfSurplusHour defined
	SfSurplusHour null.Float `xorm:"float(50,2) 'sf_surplus_hour'" json:"sf_surplus_hour" form:"sf_surplus_hour" xml:"sf_surplus_hour"`
	// YuanGiveHour defined
	YuanGiveHour null.Float `xorm:"float(11,2) 'yuan_give_hour'" json:"yuan_give_hour" form:"yuan_give_hour" xml:"yuan_give_hour"`
	// ZsBl defined
	ZsBl null.Float `xorm:"float(11,2) 'zs_bl'" json:"zs_bl" form:"zs_bl" xml:"zs_bl"`
	// BlGiveHour defined
	BlGiveHour null.Float `xorm:"float(11,2) 'bl_give_hour'" json:"bl_give_hour" form:"bl_give_hour" xml:"bl_give_hour"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// PcNumber defined
	PcNumber null.Float `xorm:"float(11,2) 'pc_number'" json:"pc_number" form:"pc_number" xml:"pc_number"`
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
	// BxType defined
	BxType null.Int `xorm:"int(11) 'bx_type'" json:"bx_type" form:"bx_type" xml:"bx_type"`
	// WxhMoney defined
	WxhMoney null.Float `xorm:"float(50,2) 'wxh_money'" json:"wxh_money" form:"wxh_money" xml:"wxh_money"`
	// OrganId defined
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// ClassCourseType defined
	ClassCourseType null.Int `xorm:"int(11) 'class_course_type'" json:"class_course_type" form:"class_course_type" xml:"class_course_type"`
}

// TableName table name of defined PcSctGiveMx
func (m *PcSctGiveMx) TableName() string {
	return "pc_sct_give_mx"
}

func (r *PcSctGiveMx) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPcSctGiveMx(data []byte) (PcSctGiveMx, error) {
	var r PcSctGiveMx
	err := json.Unmarshal(data, &r)
	return r, err
}
