// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Deletestudentbackups defined
type Deletestudentbackups struct {
	// StuLinkPerson defined
	StuLinkPerson null.String `xorm:"varchar(100) 'stu_link_person'" json:"stu_link_person" form:"stu_link_person" xml:"stu_link_person"`
	// Age defined
	Age null.Int `xorm:"int(11) 'age'" json:"age" form:"age" xml:"age"`
	// StuAddress defined
	StuAddress null.String `xorm:"varchar(100) 'stu_address'" json:"stu_address" form:"stu_address" xml:"stu_address"`
	// StuRemark defined
	StuRemark null.String `xorm:"varchar(300) 'stu_remark'" json:"stu_remark" form:"stu_remark" xml:"stu_remark"`
	// RecommendStu defined
	RecommendStu null.Int `xorm:"int(11) 'recommend_stu'" json:"recommend_stu" form:"recommend_stu" xml:"recommend_stu"`
	// StudentId defined
	StudentId null.Int `xorm:"int(11) pk notnull autoincr 'student_id'" json:"student_id" form:"student_id" xml:"student_id"`
	// StuName defined
	StuName null.String `xorm:"varchar(100) 'stu_name'" json:"stu_name" form:"stu_name" xml:"stu_name"`
	// StuPhone defined
	StuPhone null.String `xorm:"varchar(11) 'stu_phone'" json:"stu_phone" form:"stu_phone" xml:"stu_phone"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// StuType defined
	StuType null.Int `xorm:"int(11) 'stu_type'" json:"stu_type" form:"stu_type" xml:"stu_type"`
	// MaType defined
	MaType null.Int `xorm:"int(11) 'ma_type'" json:"ma_type" form:"ma_type" xml:"ma_type"`
	// StuSysState defined
	StuSysState null.Int `xorm:"int(11) 'stu_sys_state'" json:"stu_sys_state" form:"stu_sys_state" xml:"stu_sys_state"`
	// StuEnName defined
	StuEnName null.String `xorm:"varchar(100) 'stu_en_name'" json:"stu_en_name" form:"stu_en_name" xml:"stu_en_name"`
	// StuBir defined
	StuBir null.Time `xorm:"datetime 'stu_bir'" json:"stu_bir" form:"stu_bir" xml:"stu_bir"`
	// StuLoadSchool defined
	StuLoadSchool null.Int `xorm:"int(11) 'stu_load_school'" json:"stu_load_school" form:"stu_load_school" xml:"stu_load_school"`
	// MaTypeRemark defined
	MaTypeRemark null.String `xorm:"varchar(200) 'ma_type_remark'" json:"ma_type_remark" form:"ma_type_remark" xml:"ma_type_remark"`
	// LoadUser defined
	LoadUser null.Int `xorm:"int(11) 'load_user'" json:"load_user" form:"load_user" xml:"load_user"`
	// StuPwd defined
	StuPwd null.String `xorm:"varchar(300) 'stu_pwd'" json:"stu_pwd" form:"stu_pwd" xml:"stu_pwd"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// StuSex defined
	StuSex null.Int `xorm:"int(11) 'stu_sex'" json:"stu_sex" form:"stu_sex" xml:"stu_sex"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// StuZcMoney defined
	StuZcMoney null.Float `xorm:"float(10,2) 'stu_zc_money'" json:"stu_zc_money" form:"stu_zc_money" xml:"stu_zc_money"`
	// CallType defined
	CallType null.Int `xorm:"int(11) 'call_type'" json:"call_type" form:"call_type" xml:"call_type"`
	// StuLinkPerson2 defined
	StuLinkPerson2 null.String `xorm:"varchar(50) 'stu_link_person2'" json:"stu_link_person2" form:"stu_link_person2" xml:"stu_link_person2"`
	// StuPhone2 defined
	StuPhone2 null.String `xorm:"varchar(11) 'stu_phone2'" json:"stu_phone2" form:"stu_phone2" xml:"stu_phone2"`
	// AttendSchool defined
	AttendSchool null.String `xorm:"varchar(200) 'attend_school'" json:"attend_school" form:"attend_school" xml:"attend_school"`
	// Wechatid defined
	Wechatid null.String `xorm:"varchar(50) 'wechatid'" json:"wechatid" form:"wechatid" xml:"wechatid"`
	// NetworkDetail defined
	NetworkDetail null.Int `xorm:"int(11) 'network_detail'" json:"network_detail" form:"network_detail" xml:"network_detail"`
	// KhType defined
	KhType null.Int `xorm:"int(11) 'kh_type'" json:"kh_type" form:"kh_type" xml:"kh_type"`
	// TimerVisitTime defined
	TimerVisitTime null.Time `xorm:"datetime 'timer_visit_time'" json:"timer_visit_time" form:"timer_visit_time" xml:"timer_visit_time"`
	// StuDjMoney defined
	StuDjMoney null.Float `xorm:"float(10,2) 'stu_dj_money'" json:"stu_dj_money" form:"stu_dj_money" xml:"stu_dj_money"`
	// Relation1 defined
	Relation1 null.Int `xorm:"int(11) 'relation1'" json:"relation1" form:"relation1" xml:"relation1"`
	// Relation2 defined
	Relation2 null.Int `xorm:"int(11) 'relation2'" json:"relation2" form:"relation2" xml:"relation2"`
	// TjUser defined
	TjUser null.Int `xorm:"int(11) 'tj_user'" json:"tj_user" form:"tj_user" xml:"tj_user"`
	// StuJifen defined
	StuJifen null.Int `xorm:"int(11) 'stu_jifen'" json:"stu_jifen" form:"stu_jifen" xml:"stu_jifen"`
	// StuNumber defined
	StuNumber null.String `xorm:"varchar(500) 'stu_number'" json:"stu_number" form:"stu_number" xml:"stu_number"`
	// StuLevel defined
	StuLevel null.Int `xorm:"int(11) 'stu_level'" json:"stu_level" form:"stu_level" xml:"stu_level"`
	// StuFollowUpMonthly defined
	StuFollowUpMonthly null.Int `xorm:"int(11) 'stu_follow_up_monthly'" json:"stu_follow_up_monthly" form:"stu_follow_up_monthly" xml:"stu_follow_up_monthly"`
	// BcId defined
	BcId null.Int `xorm:"int(11) 'bc_id'" json:"bc_id" form:"bc_id" xml:"bc_id"`
	// BigCustomerType defined
	BigCustomerType null.Int `xorm:"int(11) 'big_customer_type'" json:"big_customer_type" form:"big_customer_type" xml:"big_customer_type"`
	// OverseasStuFollowUpMonthly defined
	OverseasStuFollowUpMonthly null.Int `xorm:"int(11) 'overseas_stu_follow_up_monthly'" json:"overseas_stu_follow_up_monthly" form:"overseas_stu_follow_up_monthly" xml:"overseas_stu_follow_up_monthly"`
	// SportsStuFollowUpMonthly defined
	SportsStuFollowUpMonthly null.Int `xorm:"int(11) 'sports_stu_follow_up_monthly'" json:"sports_stu_follow_up_monthly" form:"sports_stu_follow_up_monthly" xml:"sports_stu_follow_up_monthly"`
	// ArtStuFollowUpMonthly defined
	ArtStuFollowUpMonthly null.Int `xorm:"int(11) 'art_stu_follow_up_monthly'" json:"art_stu_follow_up_monthly" form:"art_stu_follow_up_monthly" xml:"art_stu_follow_up_monthly"`
	// BigClientStuFollowUpMonthly defined
	BigClientStuFollowUpMonthly null.Int `xorm:"int(11) 'big_client_stu_follow_up_monthly'" json:"big_client_stu_follow_up_monthly" form:"big_client_stu_follow_up_monthly" xml:"big_client_stu_follow_up_monthly"`
	// PtuserId defined
	PtuserId null.Int `xorm:"int(11) 'ptuser_id'" json:"ptuser_id" form:"ptuser_id" xml:"ptuser_id"`
	// HwStuSysState defined
	HwStuSysState null.Int `xorm:"int(11) 'hw_stu_sys_state'" json:"hw_stu_sys_state" form:"hw_stu_sys_state" xml:"hw_stu_sys_state"`
	// TyStuSysState defined
	TyStuSysState null.Int `xorm:"int(11) 'ty_stu_sys_state'" json:"ty_stu_sys_state" form:"ty_stu_sys_state" xml:"ty_stu_sys_state"`
	// YsStuSysState defined
	YsStuSysState null.Int `xorm:"int(11) 'ys_stu_sys_state'" json:"ys_stu_sys_state" form:"ys_stu_sys_state" xml:"ys_stu_sys_state"`
	// DkhStuSysState defined
	DkhStuSysState null.Int `xorm:"int(11) 'dkh_stu_sys_state'" json:"dkh_stu_sys_state" form:"dkh_stu_sys_state" xml:"dkh_stu_sys_state"`
	// HwStuType defined
	HwStuType null.Int `xorm:"int(11) 'hw_stu_type'" json:"hw_stu_type" form:"hw_stu_type" xml:"hw_stu_type"`
	// TyStuType defined
	TyStuType null.Int `xorm:"int(11) 'ty_stu_type'" json:"ty_stu_type" form:"ty_stu_type" xml:"ty_stu_type"`
	// YsStuType defined
	YsStuType null.Int `xorm:"int(11) 'ys_stu_type'" json:"ys_stu_type" form:"ys_stu_type" xml:"ys_stu_type"`
	// DkhStuType defined
	DkhStuType null.Int `xorm:"int(11) 'dkh_stu_type'" json:"dkh_stu_type" form:"dkh_stu_type" xml:"dkh_stu_type"`
	// Plfpbz defined
	Plfpbz null.String `xorm:"varchar(2000) 'plfpbz'" json:"plfpbz" form:"plfpbz" xml:"plfpbz"`
	// UserLoginState defined
	UserLoginState null.Int `xorm:"int(11) 'user_login_state'" json:"user_login_state" form:"user_login_state" xml:"user_login_state"`
	// StuParentId defined
	StuParentId null.Int `xorm:"int(11) 'stu_parent_id'" json:"stu_parent_id" form:"stu_parent_id" xml:"stu_parent_id"`
	// AgeGroup defined
	AgeGroup null.Int `xorm:"int(11) 'age_group'" json:"age_group" form:"age_group" xml:"age_group"`
	// StuInvalidDate defined
	StuInvalidDate null.Time `xorm:"datetime 'stu_invalid_date'" json:"stu_invalid_date" form:"stu_invalid_date" xml:"stu_invalid_date"`
	// StuWxdqDate defined
	StuWxdqDate null.Time `xorm:"datetime 'stu_wxdq_date'" json:"stu_wxdq_date" form:"stu_wxdq_date" xml:"stu_wxdq_date"`
	// HeadImage defined
	HeadImage null.Int `xorm:"int(11) 'head_image'" json:"head_image" form:"head_image" xml:"head_image"`
	// StuWxksDate defined
	StuWxksDate null.Time `xorm:"datetime 'stu_wxks_date'" json:"stu_wxks_date" form:"stu_wxks_date" xml:"stu_wxks_date"`
}

// TableName table name of defined Deletestudentbackups
func (m *Deletestudentbackups) TableName() string {
	return "deletestudentbackups"
}
