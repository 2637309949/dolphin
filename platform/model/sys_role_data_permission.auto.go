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

// SysRoleDataPermission defined 角色对应数据权限
type SysRoleDataPermission struct {
	// ID defined 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// RoleId defined 角色ID
	RoleId null.String `xorm:"varchar(36) notnull comment('角色ID') 'role_id'" json:"role_id" form:"role_id" xml:"role_id"`
	// DataPermissionId defined 数据权限ID
	DataPermissionId null.String `xorm:"varchar(36) notnull comment('数据权限ID') 'data_permission_id'" json:"data_permission_id" form:"data_permission_id" xml:"data_permission_id"`
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
func (m *SysRoleDataPermission) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysRoleDataPermission) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SysRoleDataPermission) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SysRoleDataPermission) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SysRoleDataPermission) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysRoleDataPermission) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysRoleDataPermission
func (m *SysRoleDataPermission) TableName() string {
	return "sys_role_data_permission"
}
