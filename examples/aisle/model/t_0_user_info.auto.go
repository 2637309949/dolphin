// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// T0UserInfo defined
type T0UserInfo struct {
	//
	UserId null.Int `xorm:"int(11) pk notnull autoincr 'user_id'" json:"user_id" xml:"user_id"`
	//
	UserName null.String `xorm:"varchar(500) 'user_name'" json:"user_name" xml:"user_name"`
	//
	LoginName null.String `xorm:"varchar(500) 'login_name'" json:"login_name" xml:"login_name"`
	//
	Password null.String `xorm:"varchar(500) 'password'" json:"password" xml:"password"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	ZkUser null.Int `xorm:"int(11) 'zk_user'" json:"zk_user" xml:"zk_user"`
	//
	ZkDate null.Time `xorm:"datetime 'zk_date'" json:"zk_date" xml:"zk_date"`
	//
	UserCardnum null.String `xorm:"varchar(20) 'user_cardnum'" json:"user_cardnum" xml:"user_cardnum"`
	//
	UserPhone null.String `xorm:"varchar(11) 'user_phone'" json:"user_phone" xml:"user_phone"`
	//
	Nationality null.Int `xorm:"int(11) 'nationality'" json:"nationality" xml:"nationality"`
	//
	Sex null.Int `xorm:"int(11) 'sex'" json:"sex" xml:"sex"`
	//
	Birthday null.Time `xorm:"datetime 'birthday'" json:"birthday" xml:"birthday"`
	//
	PassportNumber null.String `xorm:"varchar(100) 'passport_number'" json:"passport_number" xml:"passport_number"`
	//
	Mailbox null.String `xorm:"varchar(100) 'mailbox'" json:"mailbox" xml:"mailbox"`
	//
	UrgentContacts null.String `xorm:"varchar(100) 'urgent_contacts'" json:"urgent_contacts" xml:"urgent_contacts"`
	//
	UrgentcontactsPhone null.String `xorm:"varchar(11) 'urgentcontacts_phone'" json:"urgentcontacts_phone" xml:"urgentcontacts_phone"`
	//
	UrgentcontactsEmail null.String `xorm:"varchar(100) 'urgentcontacts_email'" json:"urgentcontacts_email" xml:"urgentcontacts_email"`
	//
	Education null.Int `xorm:"int(11) 'education'" json:"education" xml:"education"`
	//
	MaritalStatus null.Int `xorm:"int(11) 'marital_status'" json:"marital_status" xml:"marital_status"`
	//
	Position null.Int `xorm:"int(11) 'position'" json:"position" xml:"position"`
	//
	ContractPeriod null.Float `xorm:"float(10,2) 'contract_period'" json:"contract_period" xml:"contract_period"`
	//
	ContractsNumber null.Float `xorm:"float(10,2) 'contracts_number'" json:"contracts_number" xml:"contracts_number"`
	//
	Wages null.Float `xorm:"float(10,2) 'wages'" json:"wages" xml:"wages"`
	//
	Graduate null.String `xorm:"varchar(200) 'graduate'" json:"graduate" xml:"graduate"`
	//
	ContractStart null.Time `xorm:"datetime 'contract_start'" json:"contract_start" xml:"contract_start"`
	//
	ContractEnd null.Time `xorm:"datetime 'contract_end'" json:"contract_end" xml:"contract_end"`
	//
	ContractType null.Int `xorm:"int(11) 'contract_type'" json:"contract_type" xml:"contract_type"`
	//
	EnglishName null.String `xorm:"varchar(500) 'english_name'" json:"english_name" xml:"english_name"`
	//
	WorkDate null.Time `xorm:"datetime 'work_date'" json:"work_date" xml:"work_date"`
	//
	UserNationalityType null.Int `xorm:"int(11) 'user_nationality_type'" json:"user_nationality_type" xml:"user_nationality_type"`
	//
	Onoff null.Int `xorm:"int(11) 'onoff'" json:"onoff" xml:"onoff"`
	//
	UserLoginState null.Int `xorm:"int(11) 'user_login_state'" json:"user_login_state" xml:"user_login_state"`
	//
	LoginName1 null.String `xorm:"varchar(50) 'login_name1'" json:"login_name1" xml:"login_name1"`
	//
	UserIcon null.Int `xorm:"int(11) 'user_icon'" json:"user_icon" xml:"user_icon"`
	//
	LzDate null.Time `xorm:"datetime 'lz_date'" json:"lz_date" xml:"lz_date"`
	//
	PhoneNumber null.String `xorm:"varchar(11) 'phone_number'" json:"phone_number" xml:"phone_number"`
	//
	JpushCode null.String `xorm:"varchar(100) 'jpush_code'" json:"jpush_code" xml:"jpush_code"`
	//
	ZxNum null.String `xorm:"varchar(200) 'zx_num'" json:"zx_num" xml:"zx_num"`
	//
	ZxPwd null.String `xorm:"varchar(200) 'zx_pwd'" json:"zx_pwd" xml:"zx_pwd"`
	//
	AppVersion null.String `xorm:"varchar(20) 'app_version'" json:"app_version" xml:"app_version"`
	//
	AppNeedupdate null.Int `xorm:"int(11) 'app_needupdate'" json:"app_needupdate" xml:"app_needupdate"`
	//
	EasEmployeeId null.String `xorm:"varchar(50) 'eas_employee_id'" json:"eas_employee_id" xml:"eas_employee_id"`
	//
	OfHour null.Float `xorm:"float(50,2) 'of_hour'" json:"of_hour" xml:"of_hour"`
	//
	OfHourType null.Int `xorm:"int(11) 'of_hour_type'" json:"of_hour_type" xml:"of_hour_type"`
	//
	IfShow null.Int `xorm:"int(11) 'if_show'" json:"if_show" xml:"if_show"`
}

// TableName table name of defined T0UserInfo
func (m *T0UserInfo) TableName() string {
	return "t_0_user_info"
}