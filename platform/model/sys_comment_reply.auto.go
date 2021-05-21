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

// SysCommentReply defined 评论回复表
type SysCommentReply struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) notnull autoincr unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// CommentId defined 评论id
	CommentId null.Int `xorm:"bigint(20) notnull comment('评论id') 'comment_id'" json:"comment_id" form:"comment_id" xml:"comment_id"`
	// ReplyId defined 回复目标id
	ReplyId null.Int `xorm:"bigint(20) notnull comment('回复目标id') 'reply_id'" json:"reply_id" form:"reply_id" xml:"reply_id"`
	// ReplyType defined 回复类型
	ReplyType null.Int `xorm:"notnull comment('回复类型') 'reply_type'" json:"reply_type" form:"reply_type" xml:"reply_type"`
	// Content defined 回复内容
	Content null.String `xorm:"varchar(36) comment('回复内容') 'content'" json:"content" form:"content" xml:"content"`
	// FromUid defined 评论用户id
	FromUid null.Int `xorm:"varchar(36) notnull comment('评论用户id') 'from_uid'" json:"from_uid" form:"from_uid" xml:"from_uid"`
	// ToUid defined 目标用户id
	ToUid null.Int `xorm:"varchar(36) notnull comment('目标用户id') 'to_uid'" json:"to_uid" form:"to_uid" xml:"to_uid"`
	// Creater defined 创建人
	Creater null.Int `xorm:"bigint(20) notnull comment('创建人') 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateTime defined 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// Updater defined 最后更新人
	Updater null.Int `xorm:"bigint(20) notnull comment('最后更新人') 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateTime defined 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// IsDelete defined 删除标记
	IsDelete null.Int `xorm:"notnull comment('删除标记') 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// Remark defined 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// With defined
func (m *SysCommentReply) With(s interface{}) (interface{}, error) {
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
func (m *SysCommentReply) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysCommentReply) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SysCommentReply) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SysCommentReply) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SysCommentReply) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *SysCommentReply) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysCommentReply
func (m *SysCommentReply) TableName() string {
	return "sys_comment_reply"
}
