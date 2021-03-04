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

// SysUserTemplateDetail defined 用户信息扩展模板明细
type SysUserTemplateDetail struct {
	// ID defined 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// TempId defined 模板ID
	TempId null.String `xorm:"varchar(36) notnull comment('模板ID') 'temp_id'" json:"temp_id" form:"temp_id" xml:"temp_id"`
	// Name defined 名称
	Name null.String `xorm:"varchar(200) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// Value defined 值
	Value null.String `xorm:"varchar(200) notnull comment('值') 'value'" json:"value" form:"value" xml:"value"`
	// Type defined 模板类型 0:数值项 1:单选项 2:文字项 3:列表项
	Type null.Int `xorm:"comment('模板类型 0:数值项 1:单选项 2:文字项 3:列表项') 'type'" json:"type" form:"type" xml:"type"`
	// Content defined 名称
	Content null.String `xorm:"varchar(200) notnull comment('名称') 'content'" json:"content" form:"content" xml:"content"`
	// Priority defined 优先级
	Priority null.Int `xorm:"comment('优先级') 'priority'" json:"priority" form:"priority" xml:"priority"`
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
func (m *SysUserTemplateDetail) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysUserTemplateDetail) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SysUserTemplateDetail) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SysUserTemplateDetail) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SysUserTemplateDetail) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysUserTemplateDetail) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysUserTemplateDetail
func (m *SysUserTemplateDetail) TableName() string {
	return "sys_user_template_detail"
}
