package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/util"
)

// AdminRole default admin
var AdminRole = SysRole{
	ID:         null.IntFrom(1),
	Name:       null.StringFrom("管理员"),
	Code:       null.StringFrom("admin"),
	Status:     null.IntFrom(1),
	Creater:    null.IntFrom(1),
	CreateTime: null.TimeFrom(time.Now()),
	Updater:    null.IntFrom(1),
	UpdateTime: null.TimeFrom(time.Now()),
	IsDelete:   null.IntFrom(0),
}

var roles = []SysRole{
	AdminRole,
	{
		Name:       null.StringFrom("客户关系"),
		Code:       null.StringFrom("crm"),
		Status:     null.IntFrom(1),
		Creater:    null.IntFrom(1),
		CreateTime: null.TimeFrom(time.Now()),
		Updater:    null.IntFrom(1),
		UpdateTime: null.TimeFrom(time.Now()),
		IsDelete:   null.IntFrom(0),
	},
}

// GenCode table name of defined SysRole
func (m *SysRole) GenCode(rewrite ...bool) {
	if len(rewrite) > 0 && rewrite[0] || m.Code.String == "" {
		m.Code = null.StringFrom(util.RandString(10, util.RandNumChar))
	}
}

// InitSysData defined inital system data
func (m *SysRole) InitSysData(s *xorm.Session) error {
	for i := range roles {
		if ct, err := s.Where("code=?", roles[i].Code.String).Count(new(SysRole)); ct == 0 || err != nil {
			if err != nil {
				return err
			}
			if _, err := s.InsertOne(&roles[i]); err != nil {
				return err
			}
		}
	}
	return nil
}
