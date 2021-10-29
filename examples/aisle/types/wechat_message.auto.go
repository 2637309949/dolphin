// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// WechatMessage defined
type WechatMessage struct {
	// T3230 defined
	T3230 null.Int `xorm:"int(11) pk notnull autoincr 't_323_0'" json:"t_323_0" form:"t_323_0" xml:"t_323_0"`
	// Title defined
	Title null.String `xorm:"varchar(500) 'title'" json:"title" form:"title" xml:"title"`
	// Content defined
	Content null.String `xorm:"varchar(2000) 'content'" json:"content" form:"content" xml:"content"`
	// IfNotice defined
	IfNotice null.Int `xorm:"int(11) 'if_notice'" json:"if_notice" form:"if_notice" xml:"if_notice"`
	// IfNoticecn defined
	IfNoticecn null.Int `xorm:"int(11) 'if_noticecn'" json:"if_noticecn" form:"if_noticecn" xml:"if_noticecn"`
	// IfSee defined
	IfSee null.Int `xorm:"int(11) 'if_see'" json:"if_see" form:"if_see" xml:"if_see"`
	// IfSeecn defined
	IfSeecn null.Int `xorm:"int(11) 'if_seecn'" json:"if_seecn" form:"if_seecn" xml:"if_seecn"`
	// NoticeUser defined
	NoticeUser null.Int `xorm:"int(11) 'notice_user'" json:"notice_user" form:"notice_user" xml:"notice_user"`
	// NoticeName defined
	NoticeName null.String `xorm:"varchar(500) 'notice_name'" json:"notice_name" form:"notice_name" xml:"notice_name"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined WechatMessage
func (m *WechatMessage) TableName() string {
	return "wechat_message"
}

func (r *WechatMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWechatMessage(data []byte) (WechatMessage, error) {
	var r WechatMessage
	err := json.Unmarshal(data, &r)
	return r, err
}
