// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysAreaTemplateDetail defined
type SysAreaTemplateDetail struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) pk notnull autoincr comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 区域名称
	Name null.String `xorm:"varchar(200) notnull comment('区域名称') 'name'" json:"name" form:"name" xml:"name"`
	// TempId defined 区域信息模板id
	TempId null.Int `xorm:"bigint(20) comment('区域信息模板id') 'temp_id'" json:"temp_id" form:"temp_id" xml:"temp_id"`
	// Value defined 值
	Value null.String `xorm:"varchar(36) comment('值') 'value'" json:"value" form:"value" xml:"value"`
	// Type defined 类型 0:数值项 1:单选项 2:文字项 3:列表项
	Type null.Int `xorm:"int(11) comment('类型 0:数值项 1:单选项 2:文字项 3:列表项') 'type'" json:"type" form:"type" xml:"type"`
	// Priority defined 优先级
	Priority null.Int `xorm:"int(11) comment('优先级') 'priority'" json:"priority" form:"priority" xml:"priority"`
	// Content defined 内容
	Content null.String `xorm:"text(0) comment('内容') 'content'" json:"content" form:"content" xml:"content"`
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

// TableName table name of defined SysAreaTemplateDetail
func (m *SysAreaTemplateDetail) TableName() string {
	return "sys_area_template_detail"
}
