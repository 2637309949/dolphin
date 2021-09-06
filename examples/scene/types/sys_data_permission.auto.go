// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
    "github.com/2637309949/dolphin/packages/null"
    "github.com/shopspring/decimal"
)

// SysDataPermission defined  
type SysDataPermission struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) pk notnull autoincr comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 名称
	Name null.String `xorm:"varchar(36) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// Code defined 编码
	Code null.String `xorm:"varchar(36) comment('编码') 'code'" json:"code" form:"code" xml:"code"`
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

// TableName table name of defined SysDataPermission
func (m *SysDataPermission) TableName() string {
	return "sys_data_permission"
}
