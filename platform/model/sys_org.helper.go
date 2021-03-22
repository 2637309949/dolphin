package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// InitSysData defined
func (m *SysOrg) InitSysData(s *xorm.Session) {
	orgs := []SysOrg{
		{
			ID:          null.StringFrom("b6a7bt50-7dad-31d1-81b5-10c34fd460e3"),
			Inheritance: null.StringFrom("|b6a7bt50-7dad-31d1-81b5-10c34fd460e3"),
			Name:        null.StringFrom("Headquarters"),
			Code:        null.StringFrom("HQ"),
			Order:       null.IntFrom(0),
			Status:      null.IntFrom(1),
			CreateBy:    DefaultAdmin.ID,
			CreateTime:  null.TimeFrom(time.Now()),
			UpdateBy:    DefaultAdmin.ID,
			UpdateTime:  null.TimeFrom(time.Now()),
			IsDelete:    null.IntFrom(0),
		},
		{
			ID:          null.StringFrom("c637bt50-7dad-31d1-81b5-10c34fd460e1"),
			Parent:      null.StringFrom("b6a7bt50-7dad-31d1-81b5-10c34fd460e3"),
			Inheritance: null.StringFrom("|b6a7bt50-7dad-31d1-81b5-10c34fd460e3|c637bt50-7dad-31d1-81b5-10c34fd460e1"),
			Name:        null.StringFrom("Administrative"),
			Code:        null.StringFrom("AD"),
			Order:       null.IntFrom(0),
			Status:      null.IntFrom(1),
			CreateBy:    DefaultAdmin.ID,
			CreateTime:  null.TimeFrom(time.Now()),
			UpdateBy:    DefaultAdmin.ID,
			UpdateTime:  null.TimeFrom(time.Now()),
			IsDelete:    null.IntFrom(0),
		},
		{
			ID:          null.StringFrom("5637bt50-7dad-31d1-81b5-10c34fd460e1"),
			Parent:      null.StringFrom("b6a7bt50-7dad-31d1-81b5-10c34fd460e3"),
			Inheritance: null.StringFrom("|b6a7bt50-7dad-31d1-81b5-10c34fd460e3|5637bt50-7dad-31d1-81b5-10c34fd460e1"),
			Name:        null.StringFrom("Technology"),
			Code:        null.StringFrom("TD"),
			Order:       null.IntFrom(1),
			Status:      null.IntFrom(1),
			CreateBy:    DefaultAdmin.ID,
			CreateTime:  null.TimeFrom(time.Now()),
			UpdateBy:    DefaultAdmin.ID,
			UpdateTime:  null.TimeFrom(time.Now()),
			IsDelete:    null.IntFrom(0),
		},
	}
	for _, org := range orgs {
		if ct, err := s.Where("id=?", org.ID.String).Count(new(SysOrg)); ct == 0 || err != nil {
			if err != nil {
				s.Rollback()
				panic(err)
			}
			if _, err := s.InsertOne(&org); err != nil {
				s.Rollback()
				panic(err)
			}
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
