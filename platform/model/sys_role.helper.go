package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/util"
)

// DefaultRole default admin
var DefaultRole = SysRole{
	ID:         null.StringFrom("4c18ee66-c5e6-40a7-b190-86d115bae3e5"),
	Name:       null.StringFrom("admin"),
	Code:       null.StringFrom("X8e6D3y60K"),
	Status:     null.IntFrom(1),
	CreateBy:   DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
}

// GenCode table name of defined SysRole
func (m *SysRole) GenCode(rewrite ...bool) {
	if len(rewrite) > 0 && rewrite[0] == true || m.Code.String == "" {
		m.Code = null.StringFrom(util.RandString(10, util.RandNumChar))
	}
}

// InitSysData defined inital system data
func (m *SysRole) InitSysData(s *xorm.Session) {
	if ct, err := s.Where("id=?", DefaultRole.ID.String).Count(new(SysRole)); ct == 0 || err != nil {
		if err != nil {
			s.Rollback()
			panic(err)
		}
		if _, err := s.InsertOne(&DefaultRole); err != nil {
			s.Rollback()
			panic(err)
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
