// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// SysDataPermissionDetail defined 数据权限表明细
type SysDataPermissionDetail struct {
	// ID defined 主键
	ID null.Int `xorm:"bigint(20) notnull autoincr unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// DataPermissionId defined 数据权限表ID
	DataPermissionId null.String `xorm:"varchar(36) comment('数据权限表ID') 'data_permission_id'" json:"data_permission_id" form:"data_permission_id" xml:"data_permission_id"`
	// RoleId defined 角色ID
	RoleId null.Int `xorm:"bigint(20) comment('角色ID') 'role_id'" json:"role_id" form:"role_id" xml:"role_id"`
	// Rule defined 权限规则
	Rule null.String `xorm:"varchar(1000) notnull comment('权限规则') 'rule'" json:"rule" form:"rule" xml:"rule"`
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

// TableName table name of defined SysDataPermissionDetail
func (m *SysDataPermissionDetail) TableName() string {
	return "sys_data_permission_detail"
}

func (r *SysDataPermissionDetail) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSysDataPermissionDetail(data []byte) (SysDataPermissionDetail, error) {
	var r SysDataPermissionDetail
	err := json.Unmarshal(data, &r)
	return r, err
}
