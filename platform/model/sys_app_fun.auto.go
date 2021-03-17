// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SysAppFun defined APP功能
type SysAppFun struct {
	// ID defined 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 名称
	Name null.String `xorm:"varchar(36) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// Code defined 编码
	Code null.String `xorm:"varchar(36) notnull comment('编码') 'code'" json:"code" form:"code" xml:"code"`
	// Parent defined 父菜单ID，一级菜单为null
	Parent null.String `xorm:"varchar(36) comment('父菜单ID，一级菜单为null') 'parent'" json:"parent" form:"parent" xml:"parent"`
	// Inheritance defined 继承关系
	Inheritance null.String `xorm:"varchar(500) comment('继承关系') 'inheritance'" json:"inheritance" form:"inheritance" xml:"inheritance"`
	// URL defined 菜单URL,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀&#43;目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	URL null.String `xorm:"varchar(200) notnull comment('菜单URL,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀&#43;目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)') 'url'" json:"url" form:"url" xml:"url"`
	// Perms defined 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
	Perms null.String `xorm:"varchar(500) comment('授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)') 'perms'" json:"perms" form:"perms" xml:"perms"`
	// Type defined 类型 0：目录 1：菜单 2：按钮
	Type null.Int `xorm:"notnull comment('类型 0：目录 1：菜单 2：按钮') 'type'" json:"type" form:"type" xml:"type"`
	// Image defined 图片
	Image null.String `xorm:"varchar(200) comment('图片') 'image'" json:"image" form:"image" xml:"image"`
	// Icon defined 菜单图标
	Icon null.String `xorm:"varchar(50) comment('菜单图标') 'icon'" json:"icon" form:"icon" xml:"icon"`
	// Order defined 排序
	Order null.Int `xorm:"notnull comment('排序') 'order'" json:"order" form:"order" xml:"order"`
	// Hidden defined 是否隐藏
	Hidden null.Int `xorm:"notnull comment('是否隐藏') 'hidden'" json:"hidden" form:"hidden" xml:"hidden"`
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
func (m *SysAppFun) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *SysAppFun) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysAppFun) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SysAppFun) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SysAppFun) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SysAppFun) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysAppFun) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysAppFun
func (m *SysAppFun) TableName() string {
	return "sys_app_fun"
}
