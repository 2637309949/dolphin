package types

import (
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/spf13/viper"
)

// DefaultClient default client
var DefaultClient = SysClient{
	ID:         null.IntFrom(1),
	Client:     null.StringFrom("Y76U9344RABF4"),
	Name:       null.StringFrom("default"),
	AppName:    null.StringFrom(viper.GetString("app.name")),
	Secret:     null.StringFrom("8UYO6FVB8UYO6FVB"),
	Domain:     null.StringFrom("localhost"),
	Creater:    null.IntFrom(1),
	CreateTime: null.TimeFromNow(),
	Updater:    null.IntFrom(1),
	UpdateTime: null.TimeFromNow(),
	IsDelete:   null.IntFrom(0),
}

// InitSysData defined inital system data
func (m *SysClient) InitSysData(s *xorm.Session) error {
	if ct, err := s.Where("id=?", DefaultClient.ID.Int64).Count(new(SysClient)); ct == 0 || err != nil {
		if err != nil {
			return err
		}
		if _, err := s.Insert(&DefaultClient); err != nil {
			return err
		}
	}
	return nil
}
