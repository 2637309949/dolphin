// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// HomePhoneMessage defined
type HomePhoneMessage struct {
	// T2820 defined
	T2820 null.Int `xorm:"int(11) pk notnull autoincr 't_282_0'" json:"t_282_0" form:"t_282_0" xml:"t_282_0"`
	// Title defined
	Title null.String `xorm:"varchar(500) 'title'" json:"title" form:"title" xml:"title"`
	// Content defined
	Content null.String `xorm:"varchar(2000) 'content'" json:"content" form:"content" xml:"content"`
	// PhoneNumber defined
	PhoneNumber null.String `xorm:"varchar(11) 'phone_number'" json:"phone_number" form:"phone_number" xml:"phone_number"`
	// Sendtime defined
	Sendtime null.Time `xorm:"datetime 'sendtime'" json:"sendtime" form:"sendtime" xml:"sendtime"`
	// IfSend defined
	IfSend null.Int `xorm:"int(11) 'if_send'" json:"if_send" form:"if_send" xml:"if_send"`
	// IfSendcn defined
	IfSendcn null.String `xorm:"varchar(50) 'if_sendcn'" json:"if_sendcn" form:"if_sendcn" xml:"if_sendcn"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
}

// TableName table name of defined HomePhoneMessage
func (m *HomePhoneMessage) TableName() string {
	return "home_phone_message"
}
