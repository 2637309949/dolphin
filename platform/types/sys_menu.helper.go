package types

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// InitSysData defined inital system data
func (m *SysMenu) InitSysData(s *xorm.Session) error {
	items := []SysMenu{
		{
			ID:         null.IntFrom(1),
			Name:       null.StringFrom("Application"),
			Code:       null.StringFrom("Application"),
			URL:        null.StringFrom("application"),
			Component:  null.StringFrom("application/index"),
			Type:       null.IntFrom(1),
			Icon:       null.StringFrom("application"),
			Order:      null.IntFrom(0),
			Hidden:     null.IntFrom(0),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
			Remark:     null.String{},
		},
		{
			ID:     null.IntFrom(2),
			Name:   null.StringFrom("Setting"),
			Code:   null.StringFrom("Setting"),
			URL:    null.StringFrom("setting"),
			Type:   null.IntFrom(0),
			Icon:   null.StringFrom("setting"),
			Order:  null.IntFrom(1),
			Hidden: null.IntFrom(0),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(3),
			Name:        null.StringFrom("User"),
			Code:        null.StringFrom("User"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("user"),
			Component:   null.StringFrom("user/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("user"),
			Order:       null.IntFrom(0),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|2|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(4),
			Name:        null.StringFrom("Org"),
			Code:        null.StringFrom("Org"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("org"),
			Component:   null.StringFrom("org/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("org"),
			Order:       null.IntFrom(1),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|3|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(5),
			Name:        null.StringFrom("Role"),
			Code:        null.StringFrom("Role"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("role"),
			Component:   null.StringFrom("role/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("role"),
			Order:       null.IntFrom(2),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|4|"),
			Creater:     null.IntFrom(1),
			CreateTime:  null.TimeFrom(time.Now()),
			Updater:     null.IntFrom(1),
			UpdateTime:  null.TimeFrom(time.Now()),
			IsDelete:    null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(6),
			Name:        null.StringFrom("OptionSet"),
			Code:        null.StringFrom("OptionSet"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("optionset"),
			Component:   null.StringFrom("optionset/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("optionset"),
			Order:       null.IntFrom(3),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|5|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(7),
			Name:        null.StringFrom("Menu"),
			Code:        null.StringFrom("Menu"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("menu"),
			Component:   null.StringFrom("menu/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("menu"),
			Order:       null.IntFrom(4),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|6|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(8),
			Name:        null.StringFrom("Scheduling"),
			Code:        null.StringFrom("Scheduling"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("scheduling"),
			Component:   null.StringFrom("scheduling/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("scheduling"),
			Order:       null.IntFrom(5),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|7|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(9),
			Name:        null.StringFrom("Attachment"),
			Code:        null.StringFrom("Attachment"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("attachment"),
			Component:   null.StringFrom("attachment/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("attachment"),
			Order:       null.IntFrom(6),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|8|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(10),
			Name:        null.StringFrom("Table"),
			Code:        null.StringFrom("Table"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("table"),
			Component:   null.StringFrom("table/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("table"),
			Order:       null.IntFrom(7),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|9|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:          null.IntFrom(11),
			Name:        null.StringFrom("Tracker"),
			Code:        null.StringFrom("Tracker"),
			Parent:      null.IntFrom(2),
			URL:         null.StringFrom("tracker"),
			Component:   null.StringFrom("tracker/index"),
			Type:        null.IntFrom(1),
			Icon:        null.StringFrom("tracker"),
			Order:       null.IntFrom(8),
			Hidden:      null.IntFrom(0),
			Inheritance: null.StringFrom("|1|10|"),

			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
	}
	for i := range items {
		if ct, err := s.Where("id=?", items[i].ID.Int64).Count(new(SysMenu)); ct == 0 || err != nil {
			if err != nil {
				return err
			}
			if _, err := s.Insert(&items[i]); err != nil {
				return err
			}
		}
	}
	return nil
}