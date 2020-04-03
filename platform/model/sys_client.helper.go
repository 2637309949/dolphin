package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// DefaultClient default client
var DefaultClient = SysClient{
	ID:         null.StringFrom("5ba2b110-9dad-11d1-80b4-00c04fd431c4"),
	Client:     null.StringFrom("Y76U9344RABF4"),
	Name:       null.StringFrom("default"),
	Secret:     null.StringFrom("98UYO6FVB865"),
	Domain:     null.StringFrom("http://127.0.0.1:8082"),
	CreateBy:   DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
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
