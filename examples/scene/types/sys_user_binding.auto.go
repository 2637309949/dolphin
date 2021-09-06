// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysUserBinding defined
type SysUserBinding struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) pk notnull autoincr comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// UserId defined 用户ID
	UserId null.String `xorm:"varchar(36) notnull comment('用户ID') 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// OpenId defined open_id
	OpenId null.String `xorm:"varchar(36) notnull comment('open_id') 'open_id'" json:"open_id" form:"open_id" xml:"open_id"`
	// UnionId defined union_id
	UnionId null.String `xorm:"varchar(36) notnull comment('union_id') 'union_id'" json:"union_id" form:"union_id" xml:"union_id"`
	// Type defined 类型(0: 微信 1：叮叮)
	Type null.Int `xorm:"int(11) comment('类型(0: 微信 1：叮叮)') 'type'" json:"type" form:"type" xml:"type"`
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

// TableName table name of defined SysUserBinding
func (m *SysUserBinding) TableName() string {
	return "sys_user_binding"
}
