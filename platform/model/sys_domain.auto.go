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

// SysDomain defined 域
type SysDomain struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) notnull autoincr unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Name defined 名字
	Name null.String `xorm:"varchar(36) notnull comment('名字') 'name'" json:"name" form:"name" xml:"name"`
	// AppName defined 应用归属
	AppName null.String `xorm:"varchar(36) notnull comment('应用归属') 'app_name'" json:"app_name" form:"app_name" xml:"app_name"`
	// Domain defined 域
	Domain null.String `xorm:"varchar(36) notnull comment('域') 'domain'" json:"domain" form:"domain" xml:"domain"`
	// DomainUrl defined 域绑定
	DomainUrl null.String `xorm:"varchar(36) notnull comment('域绑定') 'domain_url'" json:"domain_url" form:"domain_url" xml:"domain_url"`
	// FullName defined 全名
	FullName null.String `xorm:"varchar(36) comment('全名') 'full_name'" json:"full_name" form:"full_name" xml:"full_name"`
	// ContactName defined 负责人
	ContactName null.String `xorm:"varchar(36) comment('负责人') 'contact_name'" json:"contact_name" form:"contact_name" xml:"contact_name"`
	// ContactEmail defined 负责人邮箱
	ContactEmail null.String `xorm:"varchar(50)  comment('负责人邮箱') 'contact_email'" json:"contact_email" form:"contact_email" xml:"contact_email"`
	// ContactMobile defined 负责人电话
	ContactMobile null.String `xorm:"varchar(50)  comment('负责人电话') 'contact_mobile'" json:"contact_mobile" form:"contact_mobile" xml:"contact_mobile"`
	// DataSource defined 数据库链接串
	DataSource null.String `xorm:"varchar(200) notnull comment('数据库链接串') 'data_source'" json:"data_source" form:"data_source" xml:"data_source"`
	// DriverName defined 驱动名称
	DriverName null.String `xorm:"varchar(50) notnull comment('驱动名称') 'driver_name'" json:"driver_name" form:"driver_name" xml:"driver_name"`
	// LoginUrl defined 登录地址
	LoginUrl null.String `xorm:"varchar(200) comment('登录地址') 'login_url'" json:"login_url" form:"login_url" xml:"login_url"`
	// ApiUrl defined 请求地址
	ApiUrl null.String `xorm:"varchar(200) comment('请求地址') 'api_url'" json:"api_url" form:"api_url" xml:"api_url"`
	// StaticUrl defined 静态地址
	StaticUrl null.String `xorm:"varchar(200) comment('静态地址') 'static_url'" json:"static_url" form:"static_url" xml:"static_url"`
	// Theme defined 样式
	Theme null.String `xorm:"varchar(50)  comment('样式') 'theme'" json:"theme" form:"theme" xml:"theme"`
	// Type defined 域类型
	Type null.Int `xorm:"notnull comment('域类型') 'type'" json:"type" form:"type" xml:"type"`
	// Status defined 状态 0：禁用 1：正常
	Status null.Int `xorm:"notnull comment('状态 0：禁用 1：正常') 'status'" json:"status" form:"status" xml:"status"`
	// AuthMode defined 认证模式 0：集成登录 1：单点登录
	AuthMode null.Int `xorm:"notnull comment('认证模式 0：集成登录 1：单点登录') 'auth_mode'" json:"auth_mode" form:"auth_mode" xml:"auth_mode"`
	// IsSync defined 是否同步了数据库标志
	IsSync null.Int `xorm:"notnull comment('是否同步了数据库标志') 'is_sync'" json:"is_sync" form:"is_sync" xml:"is_sync"`
	// Creater defined 创建人
	Creater null.Int `xorm:"bigint(20) notnull comment('创建人') 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateTime defined 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// Updater defined 最后更新人
	Updater null.Int `xorm:"bigint(20) notnull comment('最后更新人') 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateTime defined 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// IsDelete defined 删除标记
	IsDelete null.Int `xorm:"notnull comment('删除标记') 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// Remark defined 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// With defined
func (m *SysDomain) With(s interface{}) (interface{}, error) {
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
func (m *SysDomain) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysDomain) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SysDomain) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SysDomain) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SysDomain) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *SysDomain) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysDomain
func (m *SysDomain) TableName() string {
	return "sys_domain"
}
