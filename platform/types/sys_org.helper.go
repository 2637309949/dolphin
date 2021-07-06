package types

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// InitSysData defined
func (m *SysOrg) InitSysData(s *xorm.Session) error {
	orgs := []SysOrg{
		{
			ID:          null.IntFrom(1),
			Inheritance: null.StringFrom("|1"),
			Name:        null.StringFrom("Headquarters"),
			Code:        null.StringFrom("HQ"),
			Order:       null.IntFrom(0),
			Status:      null.IntFrom(1),
			Creater:     null.IntFrom(1),
			CreateTime:  null.TimeFrom(time.Now()),
			Updater:     null.IntFrom(1),
			UpdateTime:  null.TimeFrom(time.Now()),
			IsDelete:    null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(2),
			Parent:      null.StringFrom("b6a7bt50-7dad-31d1-81b5-10c34fd460e3"),
			Inheritance: null.StringFrom("|2"),
			Name:        null.StringFrom("Administrative"),
			Code:        null.StringFrom("AD"),
			Order:       null.IntFrom(0),
			Status:      null.IntFrom(1),
			Creater:     null.IntFrom(1),
			CreateTime:  null.TimeFrom(time.Now()),
			Updater:     null.IntFrom(1),
			UpdateTime:  null.TimeFrom(time.Now()),
			IsDelete:    null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(3),
			Parent:      null.StringFrom("b6a7bt50-7dad-31d1-81b5-10c34fd460e3"),
			Inheritance: null.StringFrom("|3"),
			Name:        null.StringFrom("Technology"),
			Code:        null.StringFrom("TD"),
			Order:       null.IntFrom(1),
			Status:      null.IntFrom(1),
			Creater:     null.IntFrom(1),
			CreateTime:  null.TimeFrom(time.Now()),
			Updater:     null.IntFrom(1),
			UpdateTime:  null.TimeFrom(time.Now()),
			IsDelete:    null.IntFrom(0),
		},
	}
	for i := range orgs {
		if ct, err := s.Where("id=?", orgs[i].ID.Int64).Count(new(SysOrg)); ct == 0 || err != nil {
			if err != nil {
				return err
			}
			if _, err := s.InsertOne(&orgs[i]); err != nil {
				return err
			}
		}
	}
	return nil
}
