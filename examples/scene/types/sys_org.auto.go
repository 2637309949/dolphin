// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// SysOrg defined
type SysOrg struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) pk notnull autoincr comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Parent defined 上级组织
	Parent null.String `xorm:"varchar(36) comment('上级组织') 'parent'" json:"parent" form:"parent" xml:"parent"`
	// Inheritance defined 继承关系
	Inheritance null.String `xorm:"varchar(500) comment('继承关系') 'inheritance'" json:"inheritance" form:"inheritance" xml:"inheritance"`
	// Name defined 名称
	Name null.String `xorm:"varchar(36) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// FullName defined 全名
	FullName null.String `xorm:"varchar(36) comment('全名') 'full_name'" json:"full_name" form:"full_name" xml:"full_name"`
	// Leader defined 领导人
	Leader null.String `xorm:"varchar(36) comment('领导人') 'leader'" json:"leader" form:"leader" xml:"leader"`
	// Code defined 编码
	Code null.String `xorm:"varchar(36) comment('编码') 'code'" json:"code" form:"code" xml:"code"`
	// Type defined 组织类型
	Type null.Int `xorm:"int(11) comment('组织类型') 'type'" json:"type" form:"type" xml:"type"`
	// Order defined 排序
	Order null.Int `xorm:"int(11) notnull comment('排序') 'order'" json:"order" form:"order" xml:"order"`
	// Status defined 状态 0：禁用 1：正常
	Status null.Int `xorm:"int(11) notnull comment('状态 0：禁用 1：正常') 'status'" json:"status" form:"status" xml:"status"`
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

// TableName table name of defined SysOrg
func (m *SysOrg) TableName() string {
	return "sys_org"
}

func (r *SysOrg) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSysOrg(data []byte) (SysOrg, error) {
	var r SysOrg
	err := json.Unmarshal(data, &r)
	return r, err
}
