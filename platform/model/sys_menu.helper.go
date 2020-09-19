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
			ID:        null.StringFrom("7b8bf57f-76f6-49d9-8529-7690cf385219"),
			Name:      null.StringFrom("Application"),
			Code:      null.StringFrom("Application"),
			URL:       null.StringFrom("application"),
			Component: null.StringFrom("application/index"),
			Type:      null.IntFrom(1),
			Icon:      null.StringFrom("application"),
			Order:     null.IntFrom(0),
			Hidden:    null.IntFrom(0),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:     null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			Name:   null.StringFrom("Setting"),
			Code:   null.StringFrom("Setting"),
			URL:    null.StringFrom("setting"),
			Type:   null.IntFrom(0),
			Icon:   null.StringFrom("setting"),
			Order:  null.IntFrom(1),
			Hidden: null.IntFrom(0),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("33cec683-e7a9-4e29-98d2-e0151df32221"),
			Name:        null.StringFrom("User"),
			Code:        null.StringFrom("User"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("user"),
			Component:   null.StringFrom("user/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("user"),
			Order:       null.IntFrom(0),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|33cec683-e7a9-4e29-98d2-e0151df32221|"),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("53ceb683-e8a9-4e29-98d2-e0151df32277"),
			Name:        null.StringFrom("Org"),
			Code:        null.StringFrom("Org"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("org"),
			Component:   null.StringFrom("org/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("org"),
			Order:       null.IntFrom(1),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|53ceb683-e8a9-4e29-98d2-e0151df32277|"),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("1da0a707-54c2-4289-97a3-5cdbdcf899e8"),
			Name:        null.StringFrom("Role"),
			Code:        null.StringFrom("Role"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("role"),
			Component:   null.StringFrom("role/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("role"),
			Order:       null.IntFrom(2),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|1da0a707-54c2-4289-97a3-5cdbdcf899e8|"),
			CreateBy:    DefaultAdmin.ID,
			CreateTime:  null.TimeFrom(time.Now()),
			UpdateBy:    DefaultAdmin.ID,
			UpdateTime:  null.TimeFrom(time.Now()),
			DelFlag:     null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("33cec683-e7a9-4e29-98d2-e0151df32218"),
			Name:        null.StringFrom("OptionSet"),
			Code:        null.StringFrom("OptionSet"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("optionset"),
			Component:   null.StringFrom("optionset/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("optionset"),
			Order:       null.IntFrom(3),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|33cec683-e7a9-4e29-98d2-e0151df32218|"),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("33cec683-e7a9-4e29-98d2-e0151df32237"),
			Name:        null.StringFrom("Menu"),
			Code:        null.StringFrom("Menu"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("menu"),
			Component:   null.StringFrom("menu/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("menu"),
			Order:       null.IntFrom(4),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|33cec683-e7a9-4e29-98d2-e0151df32237|"),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("36bfc673-e7a9-4e77-18d2-e0121df3b220"),
			Name:        null.StringFrom("Scheduling"),
			Code:        null.StringFrom("Scheduling"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("scheduling"),
			Component:   null.StringFrom("scheduling/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("scheduling"),
			Order:       null.IntFrom(5),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|36bfc673-e7a9-4e77-18d2-e0121df3b220|"),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("26bfc673-ac79-4e77-18d2-e0121df3b220"),
			Name:        null.StringFrom("Attachment"),
			Code:        null.StringFrom("Attachment"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("attachment"),
			Component:   null.StringFrom("attachment/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("attachment"),
			Order:       null.IntFrom(6),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|26bfc673-ac79-4e77-18d2-e0121df3b220|"),

			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		&SysMenu{
			ID:          null.StringFrom("33cec683-e7a9-4e29-98d2-e0151df32220"),
			Name:        null.StringFrom("Tracker"),
			Code:        null.StringFrom("Tracker"),
			Parent:      null.StringFrom("0a8bc79f-76f6-49d9-8529-7690cf385212"),
			URL:         null.StringFrom("tracker"),
			Component:   null.StringFrom("tracker/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("tracker"),
			Order:       null.IntFrom(7),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|0a8bc79f-76f6-49d9-8529-7690cf385212|33cec683-e7a9-4e29-98d2-e0151df32220|"),

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
