// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysUser defined 用户
type SysUser struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" xml:"id"`
	// 密码
	Password null.String `xorm:"varchar(150) notnull comment('密码') 'password'" json:"password" xml:"password"`
	// 盐噪点
	Salt null.String `xorm:"varchar(36) notnull comment('盐噪点') 'salt'" json:"salt" xml:"salt"`
	// 名字
	Name null.String `xorm:"varchar(36) notnull comment('名字') 'name'" json:"name" xml:"name"`
	// 全名
	Nickname null.String `xorm:"varchar(36) comment('全名') 'nickname'" json:"nickname" xml:"nickname"`
	// 简介
	Intro null.String `xorm:"text comment('简介') 'intro'" json:"intro" xml:"intro"`
	// 电话
	Mobile null.String `xorm:"varchar(50) comment('电话') 'mobile'" json:"mobile" xml:"mobile"`
	// 邮箱
	Email null.String `xorm:"varchar(50) comment('邮箱') 'email'" json:"email" xml:"email"`
	// 组织ID
	OrgId null.String `xorm:"varchar(36) comment('组织ID') 'org_id'" json:"org_id" xml:"org_id"`
	// 性别(0：女,1：男)
	Gender null.Int `xorm:"comment('性别(0：女,1：男)') 'gender'" json:"gender" xml:"gender"`
	// 用户类型
	Type null.Int `xorm:"comment('用户类型') 'type'" json:"type" xml:"type"`
	// 状态(0：禁用,1：正常)
	Status null.Int `xorm:"notnull comment('状态(0：禁用,1：正常)') 'status'" json:"status" xml:"status"`
	// 头像
	Avatar null.String `xorm:"varchar(255) comment('头像') 'avatar'" json:"avatar" xml:"avatar"`
	// 域
	Domain null.String `xorm:"varchar(50) notnull comment('域') 'domain'" json:"domain" xml:"domain"`
	// 创建人
	CreateBy null.String `xorm:"varchar(36) notnull comment('创建人') 'create_by'" json:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull comment('最后更新人') 'update_by'" json:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull comment('删除标记') 'del_flag'" json:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" xml:"remark"`
}

// TableName table name of defined SysUser
func (m *SysUser) TableName() string {
	return "sys_user"
}
