// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysDomain defined 域
type SysDomain struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" xml:"id"`
	// 名字
	Name null.String `xorm:"varchar(36) notnull comment('名字') 'name'" json:"name" xml:"name"`
	// 应用归属
	AppName null.String `xorm:"varchar(36) notnull comment('应用归属') 'app_name'" json:"app_name" xml:"app_name"`
	// 域
	Domain null.String `xorm:"varchar(36) notnull comment('域') 'domain'" json:"domain" xml:"domain"`
	// 域绑定
	DomainUrl null.String `xorm:"varchar(36) notnull comment('域绑定') 'domain_url'" json:"domain_url" xml:"domain_url"`
	// 全名
	FullName null.String `xorm:"varchar(36) comment('全名') 'full_name'" json:"full_name" xml:"full_name"`
	// 负责人
	ContactName null.String `xorm:"varchar(36) comment('负责人') 'contact_name'" json:"contact_name" xml:"contact_name"`
	// 负责人邮箱
	ContactEmail null.String `xorm:"varchar(50)  comment('负责人邮箱') 'contact_email'" json:"contact_email" xml:"contact_email"`
	// 负责人电话
	ContactMobile null.String `xorm:"varchar(50)  comment('负责人电话') 'contact_mobile'" json:"contact_mobile" xml:"contact_mobile"`
	// 数据库链接串
	DataSource null.String `xorm:"varchar(200) notnull comment('数据库链接串') 'data_source'" json:"data_source" xml:"data_source"`
	// 驱动名称
	DriverName null.String `xorm:"varchar(50) notnull comment('驱动名称') 'driver_name'" json:"driver_name" xml:"driver_name"`
	// 登录地址
	LoginUrl null.String `xorm:"varchar(200) comment('登录地址') 'login_url'" json:"login_url" xml:"login_url"`
	// 请求地址
	ApiUrl null.String `xorm:"varchar(200) comment('请求地址') 'api_url'" json:"api_url" xml:"api_url"`
	// 静态地址
	StaticUrl null.String `xorm:"varchar(200) comment('静态地址') 'static_url'" json:"static_url" xml:"static_url"`
	// 样式
	Theme null.String `xorm:"varchar(50)  comment('样式') 'theme'" json:"theme" xml:"theme"`
	// 域类型
	Type null.Int `xorm:"notnull comment('域类型') 'type'" json:"type" xml:"type"`
	// 状态 0：禁用 1：正常
	Status null.Int `xorm:"notnull comment('状态 0：禁用 1：正常') 'status'" json:"status" xml:"status"`
	// 认证模式 0：集成登录 1：单点登录
	AuthMode null.Int `xorm:"notnull comment('认证模式 0：集成登录 1：单点登录') 'auth_mode'" json:"auth_mode" xml:"auth_mode"`
	// 是否同步了数据库标志
	SyncFlag null.Int `xorm:"notnull comment('是否同步了数据库标志') 'sync_flag'" json:"sync_flag" xml:"sync_flag"`
	// 创建人
	CreateBy null.String `xorm:"varchar(36) notnull comment('创建人') 'create_by'" json:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull comment('最后更新人') 'update_by'" json:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull comment('删除标记') 'del_flag'" json:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" xml:"remark"`
}

// TableName table name of defined SysDomain
func (m *SysDomain) TableName() string {
	return "sys_domain"
}
