// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysUserTemplate defined 用户信息扩展模板
type SysUserTemplate struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) notnull autoincr unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 模板名称
	Name null.String `xorm:"varchar(200) notnull comment('模板名称') 'name'" json:"name" form:"name" xml:"name"`
	// Type defined 模板类型
	Type null.Int `xorm:"comment('模板类型') 'type'" json:"type" form:"type" xml:"type"`
	// Default defined 是否默认 1：是 0：否
	Default null.Int `xorm:"comment('是否默认 1：是 0：否') 'default'" json:"default" form:"default" xml:"default"`
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

// TableName table name of defined SysUserTemplate
func (m *SysUserTemplate) TableName() string {
	return "sys_user_template"
}