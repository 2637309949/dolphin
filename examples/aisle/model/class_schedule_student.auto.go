// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// ClassScheduleStudent defined
type ClassScheduleStudent struct {
	// CSSId defined
	CSSId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_s_id'" json:"c_s_s_id" form:"c_s_s_id" xml:"c_s_s_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// StuStartDate defined
	StuStartDate null.Time `xorm:"datetime 'stu_start_date'" json:"stu_start_date" form:"stu_start_date" xml:"stu_start_date"`
	// StuEndDate defined
	StuEndDate null.Time `xorm:"datetime 'stu_end_date'" json:"stu_end_date" form:"stu_end_date" xml:"stu_end_date"`
	// KqState defined
	KqState null.Int `xorm:"int(11) 'kq_state'" json:"kq_state" form:"kq_state" xml:"kq_state"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// CsId defined
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
	// KqType defined
	KqType null.Int `xorm:"int(11) 'kq_type'" json:"kq_type" form:"kq_type" xml:"kq_type"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// ScsId defined
	ScsId null.Int `xorm:"int(11) 'scs_id'" json:"scs_id" form:"scs_id" xml:"scs_id"`
	// IfKouHour defined
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
	// NotKouReason defined
	NotKouReason null.String `xorm:"varchar(1000) 'not_kou_reason'" json:"not_kou_reason" form:"not_kou_reason" xml:"not_kou_reason"`
	// KqKc defined
	KqKc null.Float `xorm:"float(10,2) 'kq_kc'" json:"kq_kc" form:"kq_kc" xml:"kq_kc"`
	// KqHour defined
	KqHour null.Float `xorm:"float(10,2) 'kq_hour'" json:"kq_hour" form:"kq_hour" xml:"kq_hour"`
	// ClassDate defined
	ClassDate null.Time `xorm:"datetime 'class_date'" json:"class_date" form:"class_date" xml:"class_date"`
	// ClassBeginTime defined
	ClassBeginTime null.Time `xorm:"datetime 'class_begin_time'" json:"class_begin_time" form:"class_begin_time" xml:"class_begin_time"`
	// ClassEndTime defined
	ClassEndTime null.Time `xorm:"datetime 'class_end_time'" json:"class_end_time" form:"class_end_time" xml:"class_end_time"`
	// KfId defined
	KfId null.Int `xorm:"int(11) 'kf_id'" json:"kf_id" form:"kf_id" xml:"kf_id"`
	// ClassId defined
	ClassId null.Int `xorm:"int(11) 'class_id'" json:"class_id" form:"class_id" xml:"class_id"`
	// KxPrice defined
	KxPrice null.Float `xorm:"float(10,2) 'kx_price'" json:"kx_price" form:"kx_price" xml:"kx_price"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// CheckUser defined
	CheckUser null.Int `xorm:"int(11) 'check_user'" json:"check_user" form:"check_user" xml:"check_user"`
	// CheckTime defined
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" form:"check_time" xml:"check_time"`
	// CheckRemark defined
	CheckRemark null.String `xorm:"varchar(2000) 'check_remark'" json:"check_remark" form:"check_remark" xml:"check_remark"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// SctId defined
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" form:"sct_id" xml:"sct_id"`
	// SingleAllKxmoney defined
	SingleAllKxmoney null.Float `xorm:"float(50,2) 'single_all_kxmoney'" json:"single_all_kxmoney" form:"single_all_kxmoney" xml:"single_all_kxmoney"`
	// SbtId defined
	SbtId null.Int `xorm:"int(11) 'sbt_id'" json:"sbt_id" form:"sbt_id" xml:"sbt_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// ParId defined
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" form:"par_id" xml:"par_id"`
	// IfUpdateKf defined
	IfUpdateKf null.Int `xorm:"int(11) 'if_update_kf'" json:"if_update_kf" form:"if_update_kf" xml:"if_update_kf"`
	// JfPrice defined
	JfPrice null.Float `xorm:"float(50,2) 'jf_price'" json:"jf_price" form:"jf_price" xml:"jf_price"`
	// JfAllprice defined
	JfAllprice null.Float `xorm:"float(50,2) 'jf_allprice'" json:"jf_allprice" form:"jf_allprice" xml:"jf_allprice"`
	// IfMz defined
	IfMz null.Int `xorm:"int(11) 'if_mz'" json:"if_mz" form:"if_mz" xml:"if_mz"`
	// DetailHour defined
	DetailHour null.Float `xorm:"float(50,2) 'detail_hour'" json:"detail_hour" form:"detail_hour" xml:"detail_hour"`
	// UseMz defined
	UseMz null.Float `xorm:"float(50,2) 'use_mz'" json:"use_mz" form:"use_mz" xml:"use_mz"`
	// MzReason defined
	MzReason null.String `xorm:"varchar(20) 'mz_reason'" json:"mz_reason" form:"mz_reason" xml:"mz_reason"`
	// ZdReason defined
	ZdReason null.Int `xorm:"int(11) 'zd_reason'" json:"zd_reason" form:"zd_reason" xml:"zd_reason"`
	// ZdRsType defined
	ZdRsType null.Int `xorm:"int(11) 'zd_rs_type'" json:"zd_rs_type" form:"zd_rs_type" xml:"zd_rs_type"`
	// PkOf defined
	PkOf null.Int `xorm:"int(11) 'pk_of'" json:"pk_of" form:"pk_of" xml:"pk_of"`
	// GdkbIfUse defined
	GdkbIfUse null.Int `xorm:"int(11) 'gdkb_if_use'" json:"gdkb_if_use" form:"gdkb_if_use" xml:"gdkb_if_use"`
	// GdkbUseDate defined
	GdkbUseDate null.Time `xorm:"datetime 'gdkb_use_date'" json:"gdkb_use_date" form:"gdkb_use_date" xml:"gdkb_use_date"`
	// TurnInSch defined
	TurnInSch null.Int `xorm:"int(11) 'turn_in_sch'" json:"turn_in_sch" form:"turn_in_sch" xml:"turn_in_sch"`
	// NextCssId defined
	NextCssId null.Int `xorm:"int(11) 'next_css_id'" json:"next_css_id" form:"next_css_id" xml:"next_css_id"`
}

// Marshal defined
func (m *ClassScheduleStudent) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *ClassScheduleStudent) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ClassScheduleStudent) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ClassScheduleStudent) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ClassScheduleStudent) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ClassScheduleStudent) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ClassScheduleStudent) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ClassScheduleStudent
func (m *ClassScheduleStudent) TableName() string {
	return "class_schedule_student"
}
