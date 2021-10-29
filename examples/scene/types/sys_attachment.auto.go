// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// SysAttachment defined
type SysAttachment struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) pk notnull autoincr comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 名称
	Name null.String `xorm:"varchar(200) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// Icon defined 图标
	Icon null.String `xorm:"varchar(300) comment('图标') 'icon'" json:"icon" form:"icon" xml:"icon"`
	// UUID defined uuid
	UUID null.String `xorm:"varchar(36) notnull comment('uuid') 'uuid'" json:"uuid" form:"uuid" xml:"uuid"`
	// Size defined 大小
	Size null.Int `xorm:"int(11) comment('大小') 'size'" json:"size" form:"size" xml:"size"`
	// Type defined 类型
	Type null.String `xorm:"varchar(36) comment('类型') 'type'" json:"type" form:"type" xml:"type"`
	// Ext defined ext
	Ext null.String `xorm:"varchar(36) comment('ext') 'ext'" json:"ext" form:"ext" xml:"ext"`
	// Hash defined hash
	Hash null.String `xorm:"varchar(500) comment('hash') 'hash'" json:"hash" form:"hash" xml:"hash"`
	// Path defined path
	Path null.String `xorm:"varchar(300) comment('path') 'path'" json:"path" form:"path" xml:"path"`
	// URL defined url
	URL null.String `xorm:"varchar(300) comment('url') 'url'" json:"url" form:"url" xml:"url"`
	// Durable defined 是否持久化 0：否 1：是
	Durable null.Int `xorm:"int(11) notnull comment('是否持久化 0：否 1：是') 'durable'" json:"durable" form:"durable" xml:"durable"`
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

// TableName table name of defined SysAttachment
func (m *SysAttachment) TableName() string {
	return "sys_attachment"
}

func (r *SysAttachment) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSysAttachment(data []byte) (SysAttachment, error) {
	var r SysAttachment
	err := json.Unmarshal(data, &r)
	return r, err
}
