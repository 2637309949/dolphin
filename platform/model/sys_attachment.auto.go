// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SysAttachment defined 附件表
type SysAttachment struct {
	// ID defined 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 名称
	Name null.String `xorm:"varchar(200) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// Icon defined 图标
	Icon null.String `xorm:"varchar(300) comment('图标') 'icon'" json:"icon" form:"icon" xml:"icon"`
	// UUID defined uuid
	UUID null.String `xorm:"varchar(36) notnull comment('uuid') 'uuid'" json:"uuid" form:"uuid" xml:"uuid"`
	// Size defined 大小
	Size null.Int `xorm:"comment('大小') 'size'" json:"size" form:"size" xml:"size"`
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
	Durable null.Int `xorm:"notnull comment('是否持久化 0：否 1：是') 'durable'" json:"durable" form:"durable" xml:"durable"`
	// CreateBy defined 创建人
	CreateBy null.String `xorm:"varchar(36) notnull comment('创建人') 'create_by'" json:"create_by" form:"create_by" xml:"create_by"`
	// CreateTime defined 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// UpdateBy defined 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull comment('最后更新人') 'update_by'" json:"update_by" form:"update_by" xml:"update_by"`
	// UpdateTime defined 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// DelFlag defined 删除标记
	DelFlag null.Int `xorm:"notnull comment('删除标记') 'del_flag'" json:"del_flag" form:"del_flag" xml:"del_flag"`
	// Remark defined 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// Marshal defined
func (m *SysAttachment) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysAttachment) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SysAttachment) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SysAttachment) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SysAttachment) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysAttachment) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysAttachment
func (m *SysAttachment) TableName() string {
	return "sys_attachment"
}
