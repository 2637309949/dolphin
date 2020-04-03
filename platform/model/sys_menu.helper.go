package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// InitSysData defined inital system data
func (m *SysMenu) InitSysData(s *xorm.Session) {
	items := []*SysMenu{
		&SysMenu{
			ID:       null.StringFrom("1729b295-d0f9-4754-9e6d-64c5ada23971"),
			Name:     null.StringFrom("Home"),
			Code:     null.StringFrom("1R56D38U2K"),
			URL:      null.StringFrom("home"),
			Type:     null.IntFrom(0),
			Icon:     null.StringFrom("menu_home_white"),
			OrderNum: null.IntFrom(0),
			Hidden:   null.IntFrom(0),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},

		&SysMenu{
			ID:       null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			Name:     null.StringFrom("Setting"),
			Code:     null.StringFrom("2J56B38U2m"),
			URL:      null.StringFrom("setting"),
			Type:     null.IntFrom(0),
			Icon:     null.StringFrom("menu_system_white"),
			OrderNum: null.IntFrom(0),
			Hidden:   null.IntFrom(1),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},

		&SysMenu{
			ID:       null.StringFrom("1da0a707-54c2-4289-97a3-5cdbdcf899e8"),
			Name:     null.StringFrom("Role"),
			Code:     null.StringFrom("1J76B38U3m"),
			Parent:   null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:      null.StringFrom("role"),
			Type:     null.IntFrom(0),
			OrderNum: null.IntFrom(0),
			Hidden:   null.IntFrom(0),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},

		&SysMenu{
			ID:       null.StringFrom("1da0a707-54c2-4289-97a3-5cdbdcf899e8"),
			Name:     null.StringFrom("User"),
			Code:     null.StringFrom("4mK6o38L3m"),
			Parent:   null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:      null.StringFrom("user"),
			Type:     null.IntFrom(0),
			OrderNum: null.IntFrom(1),
			Hidden:   null.IntFrom(0),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
	}
	for _, item := range items {
		if ct, err := s.Where("id=?", item.ID.String).Count(new(SysMenu)); ct == 0 || err != nil {
			if err != nil {
				s.Rollback()
				panic(err)
			}
			if _, err := s.InsertOne(item); err != nil {
				s.Rollback()
				panic(err)
			}
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
