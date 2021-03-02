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

// SysTableColumn defined 系统字段
type SysTableColumn struct {
	// ID defined 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// TbId defined 表ID
	TbId null.String `xorm:"varchar(36) notnull comment('表ID') 'tb_id'" json:"tb_id" form:"tb_id" xml:"tb_id"`
	// Name defined 表名
	Name null.String `xorm:"varchar(150) notnull comment('表名') 'name'" json:"name" form:"name" xml:"name"`
	// Desc defined 描述
	Desc null.String `xorm:"varchar(300) notnull comment('描述') 'desc'" json:"desc" form:"desc" xml:"desc"`
	// IsPrimaryKey defined 是否主键
	IsPrimaryKey null.Bool `xorm:"notnull comment('是否主键') 'is_primary_key'" json:"is_primary_key" form:"is_primary_key" xml:"is_primary_key"`
	// Type defined 类型
	Type null.String `xorm:"varchar(100) notnull comment('类型') 'type'" json:"type" form:"type" xml:"type"`
	// Nullable defined 是否为空
	Nullable null.Bool `xorm:"notnull comment('是否为空') 'nullable'" json:"nullable" form:"nullable" xml:"nullable"`
	// Default defined 默认值
	Default null.String `xorm:"varchar(100) notnull comment('默认值') 'default'" json:"default" form:"default" xml:"default"`
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
func (m *SysTableColumn) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysTableColumn) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *SysTableColumn) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysTableColumn) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysTableColumn
func (m *SysTableColumn) TableName() string {
	return "sys_table_column"
}
