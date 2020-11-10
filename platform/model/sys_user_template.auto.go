// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysUserTemplate defined 用户信息扩展模板
type SysUserTemplate struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// 模板名称
	Name null.String `xorm:"varchar(200) notnull comment('模板名称') 'name'" json:"name" form:"name" xml:"name"`
	// 模板类型
	Type null.Int `xorm:"comment('模板类型') 'type'" json:"type" form:"type" xml:"type"`
	// 是否默认 1：是 0：否
	Default null.Int `xorm:"comment('是否默认 1：是 0：否') 'default'" json:"default" form:"default" xml:"default"`
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

// TableName table name of defined SysUserTemplate
func (m *SysUserTemplate) TableName() string {
	return "sys_user_template"
}
