// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
    "github.com/2637309949/dolphin/packages/null"
    "github.com/shopspring/decimal"
)

// SysMenu defined  
type SysMenu struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) pk notnull autoincr comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 名称
	Name null.String `xorm:"varchar(36) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// Code defined 编码
	Code null.String `xorm:"varchar(36) notnull comment('编码') 'code'" json:"code" form:"code" xml:"code"`
	// Parent defined 父菜单ID，一级菜单为null
	Parent null.Int `xorm:"bigint(20) comment('父菜单ID，一级菜单为null') 'parent'" json:"parent" form:"parent" xml:"parent"`
	// Inheritance defined 继承关系
	Inheritance null.String `xorm:"varchar(500) comment('继承关系') 'inheritance'" json:"inheritance" form:"inheritance" xml:"inheritance"`
	// URL defined 菜单URL
	URL null.String `xorm:"varchar(200) notnull comment('菜单URL') 'url'" json:"url" form:"url" xml:"url"`
	// Component defined 菜单组件
	Component null.String `xorm:"varchar(100) comment('菜单组件') 'component'" json:"component" form:"component" xml:"component"`
	// Perms defined 授权(多个用逗号分隔，如：sys:user:add
	Perms null.String `xorm:"varchar(500) comment('授权(多个用逗号分隔，如：sys:user:add') 'perms'" json:"perms" form:"perms" xml:"perms"`
	// Type defined 类型 0：目录 1：菜单 2：按钮
	Type null.Int `xorm:"int(11) notnull comment('类型 0：目录 1：菜单 2：按钮') 'type'" json:"type" form:"type" xml:"type"`
	// Icon defined 菜单图标
	Icon null.String `xorm:"varchar(50) comment('菜单图标') 'icon'" json:"icon" form:"icon" xml:"icon"`
	// Order defined 排序
	Order null.Int `xorm:"int(11) notnull comment('排序') 'order'" json:"order" form:"order" xml:"order"`
	// Hidden defined 是否隐藏
	Hidden null.Int `xorm:"int(11) notnull comment('是否隐藏') 'hidden'" json:"hidden" form:"hidden" xml:"hidden"`
	// Creater defined 创建人
	Creater null.Int `xorm:"bigint(20) notnull comment('创建人') 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateTime defined 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// Updater defined 最后更新人
	Updater null.Int `xorm:"bigint(20) notnull comment('最后更新人') 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateTime defined 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// IsDelete defined 删除标记
	IsDelete null.Int `xorm:"int(11) notnull comment('删除标记') 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// Remark defined 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// TableName table name of defined SysMenu
func (m *SysMenu) TableName() string {
	return "sys_menu"
}
