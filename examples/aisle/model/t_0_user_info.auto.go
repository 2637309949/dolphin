// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// T0UserInfo defined
type T0UserInfo struct {
	// UserId defined
	UserId null.Int `xorm:"int(11) pk notnull autoincr 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// UserName defined
	UserName null.String `xorm:"varchar(500) 'user_name'" json:"user_name" form:"user_name" xml:"user_name"`
	// LoginName defined
	LoginName null.String `xorm:"varchar(500) 'login_name'" json:"login_name" form:"login_name" xml:"login_name"`
	// Password defined
	Password null.String `xorm:"varchar(500) 'password'" json:"password" form:"password" xml:"password"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// ZkUser defined
	ZkUser null.Int `xorm:"int(11) 'zk_user'" json:"zk_user" form:"zk_user" xml:"zk_user"`
	// ZkDate defined
	ZkDate null.Time `xorm:"datetime 'zk_date'" json:"zk_date" form:"zk_date" xml:"zk_date"`
	// UserCardnum defined
	UserCardnum null.String `xorm:"varchar(20) 'user_cardnum'" json:"user_cardnum" form:"user_cardnum" xml:"user_cardnum"`
	// UserPhone defined
	UserPhone null.String `xorm:"varchar(11) 'user_phone'" json:"user_phone" form:"user_phone" xml:"user_phone"`
	// Nationality defined
	Nationality null.Int `xorm:"int(11) 'nationality'" json:"nationality" form:"nationality" xml:"nationality"`
	// Sex defined
	Sex null.Int `xorm:"int(11) 'sex'" json:"sex" form:"sex" xml:"sex"`
	// Birthday defined
	Birthday null.Time `xorm:"datetime 'birthday'" json:"birthday" form:"birthday" xml:"birthday"`
	// PassportNumber defined
	PassportNumber null.String `xorm:"varchar(100) 'passport_number'" json:"passport_number" form:"passport_number" xml:"passport_number"`
	// Mailbox defined
	Mailbox null.String `xorm:"varchar(100) 'mailbox'" json:"mailbox" form:"mailbox" xml:"mailbox"`
	// UrgentContacts defined
	UrgentContacts null.String `xorm:"varchar(100) 'urgent_contacts'" json:"urgent_contacts" form:"urgent_contacts" xml:"urgent_contacts"`
	// UrgentcontactsPhone defined
	UrgentcontactsPhone null.String `xorm:"varchar(11) 'urgentcontacts_phone'" json:"urgentcontacts_phone" form:"urgentcontacts_phone" xml:"urgentcontacts_phone"`
	// UrgentcontactsEmail defined
	UrgentcontactsEmail null.String `xorm:"varchar(100) 'urgentcontacts_email'" json:"urgentcontacts_email" form:"urgentcontacts_email" xml:"urgentcontacts_email"`
	// Education defined
	Education null.Int `xorm:"int(11) 'education'" json:"education" form:"education" xml:"education"`
	// MaritalStatus defined
	MaritalStatus null.Int `xorm:"int(11) 'marital_status'" json:"marital_status" form:"marital_status" xml:"marital_status"`
	// Position defined
	Position null.Int `xorm:"int(11) 'position'" json:"position" form:"position" xml:"position"`
	// ContractPeriod defined
	ContractPeriod null.Float `xorm:"float(10,2) 'contract_period'" json:"contract_period" form:"contract_period" xml:"contract_period"`
	// ContractsNumber defined
	ContractsNumber null.Float `xorm:"float(10,2) 'contracts_number'" json:"contracts_number" form:"contracts_number" xml:"contracts_number"`
	// Wages defined
	Wages null.Float `xorm:"float(10,2) 'wages'" json:"wages" form:"wages" xml:"wages"`
	// Graduate defined
	Graduate null.String `xorm:"varchar(200) 'graduate'" json:"graduate" form:"graduate" xml:"graduate"`
	// ContractStart defined
	ContractStart null.Time `xorm:"datetime 'contract_start'" json:"contract_start" form:"contract_start" xml:"contract_start"`
	// ContractEnd defined
	ContractEnd null.Time `xorm:"datetime 'contract_end'" json:"contract_end" form:"contract_end" xml:"contract_end"`
	// ContractType defined
	ContractType null.Int `xorm:"int(11) 'contract_type'" json:"contract_type" form:"contract_type" xml:"contract_type"`
	// EnglishName defined
	EnglishName null.String `xorm:"varchar(500) 'english_name'" json:"english_name" form:"english_name" xml:"english_name"`
	// WorkDate defined
	WorkDate null.Time `xorm:"datetime 'work_date'" json:"work_date" form:"work_date" xml:"work_date"`
	// UserNationalityType defined
	UserNationalityType null.Int `xorm:"int(11) 'user_nationality_type'" json:"user_nationality_type" form:"user_nationality_type" xml:"user_nationality_type"`
	// Onoff defined
	Onoff null.Int `xorm:"int(11) 'onoff'" json:"onoff" form:"onoff" xml:"onoff"`
	// UserLoginState defined
	UserLoginState null.Int `xorm:"int(11) 'user_login_state'" json:"user_login_state" form:"user_login_state" xml:"user_login_state"`
	// LoginName1 defined
	LoginName1 null.String `xorm:"varchar(50) 'login_name1'" json:"login_name1" form:"login_name1" xml:"login_name1"`
	// UserIcon defined
	UserIcon null.Int `xorm:"int(11) 'user_icon'" json:"user_icon" form:"user_icon" xml:"user_icon"`
	// LzDate defined
	LzDate null.Time `xorm:"datetime 'lz_date'" json:"lz_date" form:"lz_date" xml:"lz_date"`
	// PhoneNumber defined
	PhoneNumber null.String `xorm:"varchar(11) 'phone_number'" json:"phone_number" form:"phone_number" xml:"phone_number"`
	// JpushCode defined
	JpushCode null.String `xorm:"varchar(100) 'jpush_code'" json:"jpush_code" form:"jpush_code" xml:"jpush_code"`
	// ZxNum defined
	ZxNum null.String `xorm:"varchar(200) 'zx_num'" json:"zx_num" form:"zx_num" xml:"zx_num"`
	// ZxPwd defined
	ZxPwd null.String `xorm:"varchar(200) 'zx_pwd'" json:"zx_pwd" form:"zx_pwd" xml:"zx_pwd"`
	// AppVersion defined
	AppVersion null.String `xorm:"varchar(20) 'app_version'" json:"app_version" form:"app_version" xml:"app_version"`
	// AppNeedupdate defined
	AppNeedupdate null.Int `xorm:"int(11) 'app_needupdate'" json:"app_needupdate" form:"app_needupdate" xml:"app_needupdate"`
	// EasEmployeeId defined
	EasEmployeeId null.String `xorm:"varchar(50) 'eas_employee_id'" json:"eas_employee_id" form:"eas_employee_id" xml:"eas_employee_id"`
	// OfHour defined
	OfHour null.Float `xorm:"float(50,2) 'of_hour'" json:"of_hour" form:"of_hour" xml:"of_hour"`
	// OfHourType defined
	OfHourType null.Int `xorm:"int(11) 'of_hour_type'" json:"of_hour_type" form:"of_hour_type" xml:"of_hour_type"`
	// IfShow defined
	IfShow null.Int `xorm:"int(11) 'if_show'" json:"if_show" form:"if_show" xml:"if_show"`
}

// TableName table name of defined T0UserInfo
func (m *T0UserInfo) TableName() string {
	return "t_0_user_info"
}
