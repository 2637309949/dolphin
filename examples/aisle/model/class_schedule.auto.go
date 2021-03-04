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

// ClassSchedule defined
type ClassSchedule struct {
	// CSId defined
	CSId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_id'" json:"c_s_id" form:"c_s_id" xml:"c_s_id"`
	// ClassId defined
	ClassId null.Int `xorm:"int(11) 'class_id'" json:"class_id" form:"class_id" xml:"class_id"`
	// AttendDate defined
	AttendDate null.Time `xorm:"datetime 'attend_date'" json:"attend_date" form:"attend_date" xml:"attend_date"`
	// AppStartTime defined
	AppStartTime null.Time `xorm:"datetime 'app_start_time'" json:"app_start_time" form:"app_start_time" xml:"app_start_time"`
	// AppEndTime defined
	AppEndTime null.Time `xorm:"datetime 'app_end_time'" json:"app_end_time" form:"app_end_time" xml:"app_end_time"`
	// StudyNumber defined
	StudyNumber null.Float `xorm:"float(11,2) 'study_number'" json:"study_number" form:"study_number" xml:"study_number"`
	// ProId defined
	ProId null.Int `xorm:"int(11) 'pro_id'" json:"pro_id" form:"pro_id" xml:"pro_id"`
	// OcpId defined
	OcpId null.Int `xorm:"int(11) 'ocp_id'" json:"ocp_id" form:"ocp_id" xml:"ocp_id"`
	// CsKc defined
	CsKc null.Float `xorm:"float(50,2) 'cs_kc'" json:"cs_kc" form:"cs_kc" xml:"cs_kc"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// StudyState defined
	StudyState null.Int `xorm:"int(11) 'study_state'" json:"study_state" form:"study_state" xml:"study_state"`
	// ClassroomId defined
	ClassroomId null.Int `xorm:"int(11) 'classroom_id'" json:"classroom_id" form:"classroom_id" xml:"classroom_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// KfId defined
	KfId null.Int `xorm:"int(11) 'kf_id'" json:"kf_id" form:"kf_id" xml:"kf_id"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// CancelReason defined
	CancelReason null.String `xorm:"varchar(500) 'cancel_reason'" json:"cancel_reason" form:"cancel_reason" xml:"cancel_reason"`
	// CspId defined
	CspId null.Int `xorm:"int(11) 'csp_id'" json:"csp_id" form:"csp_id" xml:"csp_id"`
	// TspId defined
	TspId null.Int `xorm:"int(11) 'tsp_id'" json:"tsp_id" form:"tsp_id" xml:"tsp_id"`
	// ClassWhenBout defined
	ClassWhenBout null.Int `xorm:"int(11) 'class_when_bout'" json:"class_when_bout" form:"class_when_bout" xml:"class_when_bout"`
	// NowCustomer defined
	NowCustomer null.String `xorm:"varchar(100) 'now_customer'" json:"now_customer" form:"now_customer" xml:"now_customer"`
	// HisCustomer defined
	HisCustomer null.String `xorm:"varchar(100) 'his_customer'" json:"his_customer" form:"his_customer" xml:"his_customer"`
	// IfUpdateKf defined
	IfUpdateKf null.Int `xorm:"int(11) 'if_update_kf'" json:"if_update_kf" form:"if_update_kf" xml:"if_update_kf"`
	// KtnrNum defined
	KtnrNum null.Float `xorm:"float(50,2) 'ktnr_num'" json:"ktnr_num" form:"ktnr_num" xml:"ktnr_num"`
	// KtslNum defined
	KtslNum null.Float `xorm:"float(50,2) 'ktsl_num'" json:"ktsl_num" form:"ktsl_num" xml:"ktsl_num"`
	// KtfkNum defined
	KtfkNum null.Float `xorm:"float(50,2) 'ktfk_num'" json:"ktfk_num" form:"ktfk_num" xml:"ktfk_num"`
	// KtzyNum defined
	KtzyNum null.Float `xorm:"float(50,2) 'ktzy_num'" json:"ktzy_num" form:"ktzy_num" xml:"ktzy_num"`
	// KtfkAllnum defined
	KtfkAllnum null.Float `xorm:"float(50,2) 'ktfk_allnum'" json:"ktfk_allnum" form:"ktfk_allnum" xml:"ktfk_allnum"`
	// KtzyAllnum defined
	KtzyAllnum null.Float `xorm:"float(50,2) 'ktzy_allnum'" json:"ktzy_allnum" form:"ktzy_allnum" xml:"ktzy_allnum"`
	// StuZydNum defined
	StuZydNum null.Float `xorm:"float(50,2) 'stu_zyd_num'" json:"stu_zyd_num" form:"stu_zyd_num" xml:"stu_zyd_num"`
	// TeaId defined
	TeaId null.String `xorm:"varchar(1000) 'tea_id'" json:"tea_id" form:"tea_id" xml:"tea_id"`
	// TeaName defined
	TeaName null.String `xorm:"varchar(1000) 'tea_name'" json:"tea_name" form:"tea_name" xml:"tea_name"`
	// NowCustomername defined
	NowCustomername null.String `xorm:"varchar(100) 'now_customername'" json:"now_customername" form:"now_customername" xml:"now_customername"`
	// KfEgName defined
	KfEgName null.String `xorm:"varchar(1000) 'kf_eg_name'" json:"kf_eg_name" form:"kf_eg_name" xml:"kf_eg_name"`
	// TeaEgName defined
	TeaEgName null.String `xorm:"varchar(1000) 'tea_eg_name'" json:"tea_eg_name" form:"tea_eg_name" xml:"tea_eg_name"`
	// IfSendTea defined
	IfSendTea null.Int `xorm:"int(11) 'if_send_tea'" json:"if_send_tea" form:"if_send_tea" xml:"if_send_tea"`
	// Remark defined
	Remark null.String `xorm:"varchar(50) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// PkCtr defined
	PkCtr null.Int `xorm:"int(11) 'pk_ctr'" json:"pk_ctr" form:"pk_ctr" xml:"pk_ctr"`
}

// Marshal defined
func (m *ClassSchedule) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ClassSchedule) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ClassSchedule) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ClassSchedule) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ClassSchedule) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ClassSchedule) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ClassSchedule
func (m *ClassSchedule) TableName() string {
	return "class_schedule"
}
