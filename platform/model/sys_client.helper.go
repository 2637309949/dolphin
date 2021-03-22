package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// DefaultClient default client
var DefaultClient = SysClient{
	ID:         null.StringFrom("5ba2b110-9dad-11d1-80b4-00c04fd431c4"),
	Client:     null.StringFrom("Y76U9344RABF4"),
	Name:       null.StringFrom("default"),
	AppName:    null.StringFrom(viper.GetString("app.name")),
	Secret:     null.StringFrom("98UYO6FVB865"),
	Domain:     null.StringFrom("localhost"),
	Creater:    DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	Updater:    DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	IsDelete:   null.IntFrom(0),
}

// InitSysData defined inital system data
func (m *SysClient) InitSysData(s *xorm.Session) {
	if ct, err := s.Where("id=?", DefaultClient.ID.String).Count(new(SysClient)); ct == 0 || err != nil {
		if err != nil {
			s.Rollback()
			panic(err)
		}
		if _, err := s.InsertOne(&DefaultClient); err != nil {
			s.Rollback()
			panic(err)
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
