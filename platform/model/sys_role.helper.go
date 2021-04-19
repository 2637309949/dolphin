package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/util"
)

// AdminRole default admin
var AdminRole = SysRole{
	ID:         null.StringFrom("4c18ee66-c5e6-40a7-b190-86d115bae3e5"),
	Name:       null.StringFrom("管理员"),
	Code:       null.StringFrom("admin"),
	Status:     null.IntFrom(1),
	Creater:    DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	Updater:    DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	IsDelete:   null.IntFrom(0),
}

var roles = []SysRole{
	AdminRole,
	{
		ID:         null.StringFrom("3b18ee66-a5e6-40a3-b190-86d115bae3e2"),
		Name:       null.StringFrom("客户关系"),
		Code:       null.StringFrom("crm"),
		Status:     null.IntFrom(1),
		Creater:    DefaultAdmin.ID,
		CreateTime: null.TimeFrom(time.Now()),
		Updater:    DefaultAdmin.ID,
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
func (m *SysRole) InitSysData(s *xorm.Session) {
	for _, role := range roles {
		if ct, err := s.Where("id=?", role.ID.String).Count(new(SysRole)); ct == 0 || err != nil {
			if err != nil {
				s.Rollback()
				panic(err)
			}
			if _, err := s.InsertOne(&role); err != nil {
				s.Rollback()
				panic(err)
			}
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
