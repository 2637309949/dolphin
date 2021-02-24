// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SysArea defined 区域管理
type SysArea struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// 区域名称
	Name null.String `xorm:"varchar(200) notnull comment('区域名称') 'name'" json:"name" form:"name" xml:"name"`
	// 父节点
	Parent null.String `xorm:"varchar(36) comment('父节点') 'parent'" json:"parent" form:"parent" xml:"parent"`
	// 继承关系
	Inheritance null.String `xorm:"varchar(500) comment('继承关系') 'inheritance'" json:"inheritance" form:"inheritance" xml:"inheritance"`
	// 组织ID
	OrgId null.String `xorm:"varchar(36) comment('组织ID') 'org_id'" json:"org_id" form:"org_id" xml:"org_id"`
	// 模板ID
	TempId null.String `xorm:"varchar(36) comment('模板ID') 'temp_id'" json:"temp_id" form:"temp_id" xml:"temp_id"`
	// 模板内容
	TempValue null.String `xorm:"text comment('模板内容') 'temp_value'" json:"temp_value" form:"temp_value" xml:"temp_value"`
	// 负责人
	Manager null.String `xorm:"varchar(36) comment('负责人') 'manager'" json:"manager" form:"manager" xml:"manager"`
	// 创建人
	CreateBy null.String `xorm:"varchar(36) notnull comment('创建人') 'create_by'" json:"create_by" form:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull comment('最后更新人') 'update_by'" json:"update_by" form:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull comment('删除标记') 'del_flag'" json:"del_flag" form:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// Parser defined
func (m *SysArea) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysArea) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysArea
func (m *SysArea) TableName() string {
	return "sys_area"
}
