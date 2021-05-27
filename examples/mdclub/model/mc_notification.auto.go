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

// McNotification defined 通知表
type McNotification struct {
	// NotificationId defined 通知ID
	NotificationId null.Int `xorm:"int(11) pk notnull autoincr comment('通知ID') 'notification_id'" json:"notification_id" form:"notification_id" xml:"notification_id"`
	// ReceiverId defined 接收者ID
	ReceiverId null.Int `xorm:"int(11) notnull comment('接收者ID') 'receiver_id'" json:"receiver_id" form:"receiver_id" xml:"receiver_id"`
	// SenderId defined 发送者ID
	SenderId null.Int `xorm:"int(11) notnull comment('发送者ID') 'sender_id'" json:"sender_id" form:"sender_id" xml:"sender_id"`
	// Type defined 消息类型： question_answered, question_commented, question_deleted, article_commented, article_deleted, answer_commented, answer_deleted, comment_replied, comment_deleted
	Type null.String `xorm:"varchar(40) notnull comment('消息类型： question_answered, question_commented, question_deleted, article_commented, article_deleted, answer_commented, answer_deleted, comment_replied, comment_deleted') 'type'" json:"type" form:"type" xml:"type"`
	// ArticleId defined 文章ID
	ArticleId null.Int `xorm:"int(11) notnull comment('文章ID') 'article_id'" json:"article_id" form:"article_id" xml:"article_id"`
	// QuestionId defined 提问ID
	QuestionId null.Int `xorm:"int(11) notnull comment('提问ID') 'question_id'" json:"question_id" form:"question_id" xml:"question_id"`
	// AnswerId defined 回答ID
	AnswerId null.Int `xorm:"int(11) notnull comment('回答ID') 'answer_id'" json:"answer_id" form:"answer_id" xml:"answer_id"`
	// CommentId defined 评论ID
	CommentId null.Int `xorm:"int(11) notnull comment('评论ID') 'comment_id'" json:"comment_id" form:"comment_id" xml:"comment_id"`
	// ReplyId defined 回复ID
	ReplyId null.Int `xorm:"int(11) notnull comment('回复ID') 'reply_id'" json:"reply_id" form:"reply_id" xml:"reply_id"`
	// ContentDeleted defined 被删除的内容的备份
	ContentDeleted null.String `xorm:"text(0) notnull comment('被删除的内容的备份') 'content_deleted'" json:"content_deleted" form:"content_deleted" xml:"content_deleted"`
	// CreateTime defined 发送时间
	CreateTime null.Int `xorm:"int(10) notnull default(0) comment('发送时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// ReadTime defined 阅读时间
	ReadTime null.Int `xorm:"int(10) notnull default(0) comment('阅读时间') 'read_time'" json:"read_time" form:"read_time" xml:"read_time"`
}

// With defined
func (m *McNotification) With(s interface{}) (interface{}, error) {
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
func (m *McNotification) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *McNotification) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *McNotification) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *McNotification) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *McNotification) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *McNotification) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined McNotification
func (m *McNotification) TableName() string {
	return "mc_notification"
}
