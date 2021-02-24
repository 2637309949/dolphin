// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SysUser defined 用户
type SysUser struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// 密码
	Password null.String `xorm:"varchar(150) notnull comment('密码') 'password'" json:"password" form:"password" xml:"password"`
	// 盐噪点
	Salt null.String `xorm:"varchar(36) notnull comment('盐噪点') 'salt'" json:"salt" form:"salt" xml:"salt"`
	// 名字
	Name null.String `xorm:"varchar(36) notnull comment('名字') 'name'" json:"name" form:"name" xml:"name"`
	// 全名
	FullName null.String `xorm:"varchar(36) comment('全名') 'full_name'" json:"full_name" form:"full_name" xml:"full_name"`
	// 昵称
	Nickname null.String `xorm:"varchar(36) comment('昵称') 'nickname'" json:"nickname" form:"nickname" xml:"nickname"`
	// 简介
	Intro null.String `xorm:"text comment('简介') 'intro'" json:"intro" form:"intro" xml:"intro"`
	// 地址
	Address null.String `xorm:"varchar(150) comment('地址') 'address'" json:"address" form:"address" xml:"address"`
	// 电话
	Mobile null.String `xorm:"varchar(50) comment('电话') 'mobile'" json:"mobile" form:"mobile" xml:"mobile"`
	// 邮箱
	Email null.String `xorm:"varchar(50) comment('邮箱') 'email'" json:"email" form:"email" xml:"email"`
	// 是否邮件确认(0：尚未,1：确认)
	IsEmailConfirmed null.Int `xorm:"comment('是否邮件确认(0：尚未,1：确认)') 'is_email_confirmed'" json:"is_email_confirmed" form:"is_email_confirmed" xml:"is_email_confirmed"`
	// 组织ID
	OrgId null.String `xorm:"varchar(36) comment('组织ID') 'org_id'" json:"org_id" form:"org_id" xml:"org_id"`
	// 性别(0：女,1：男)
	Gender null.Int `xorm:"comment('性别(0：女,1：男)') 'gender'" json:"gender" form:"gender" xml:"gender"`
	// 用户类型
	Type null.Int `xorm:"comment('用户类型') 'type'" json:"type" form:"type" xml:"type"`
	// 状态(0：禁用,1：正常)
	Status null.Int `xorm:"notnull comment('状态(0：禁用,1：正常)') 'status'" json:"status" form:"status" xml:"status"`
	// 头像
	Avatar null.String `xorm:"varchar(255) comment('头像') 'avatar'" json:"avatar" form:"avatar" xml:"avatar"`
	// 标记所有已读时间
	MarkedAllAsReadAt null.Time `xorm:"datetime comment('标记所有已读时间') 'marked_all_as_read_at'" json:"marked_all_as_read_at" form:"marked_all_as_read_at" xml:"marked_all_as_read_at"`
	// 标记读取消息时间
	ReadNotificationsAt null.Time `xorm:"datetime comment('标记读取消息时间') 'read_notifications_at'" json:"read_notifications_at" form:"read_notifications_at" xml:"read_notifications_at"`
	// 模板ID
	TempId null.String `xorm:"varchar(36) comment('模板ID') 'temp_id'" json:"temp_id" form:"temp_id" xml:"temp_id"`
	// 模板值
	TempValue null.String `xorm:"text comment('模板值') 'temp_value'" json:"temp_value" form:"temp_value" xml:"temp_value"`
	// 域
	Domain null.String `xorm:"varchar(50) notnull comment('域') 'domain'" json:"domain" form:"domain" xml:"domain"`
	// 创建人
	CreateBy null.String `xorm:"varchar(36) notnull comment('创建人') 'create_by'" json:"create_by" form:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull comment('最后更新人') 'update_by'" json:"update_by" form:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull comment('删除标记') 'del_flag'" json:"del_flag" form:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// Parser defined
func (m *SysUser) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysUser) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysUser
func (m *SysUser) TableName() string {
	return "sys_user"
}
