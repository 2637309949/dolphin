// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/viper"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// DefaultDomain default domain
var DefaultDomain = SysDomain{
	ID:         null.StringFrom("2ba7b810-9dad-11d1-80b4-00c04fd430c3"),
	Name:       null.StringFrom("default"),
	FullName:   null.StringFrom("default"),
	Theme:      null.StringFrom("default"),
	DataSource: null.StringFrom(strings.Replace(viper.GetString("db.dataSource"), "?", "_localhost?", 1)),
	DriverName: null.StringFrom("mysql"),
	LoginUrl:   null.StringFrom("localhost"),
	Type:       null.IntFrom(0),
	Status:     null.IntFrom(1),
	SyncFlag:   null.IntFrom(0),
	AuthMode:   null.IntFrom(1),
	Domain:     null.StringFrom("localhost"),
	ApiUrl:     null.StringFrom("http://localhost:8086"),
	CreateBy:   DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
}

// InitSysData defined inital system data
func (m *SysDomain) InitSysData(s *xorm.Session) {
	if ct, err := s.Where("id=?", DefaultDomain.ID.String).Count(new(SysDomain)); ct == 0 || err != nil {
		if err != nil {
			s.Rollback()
			panic(err)
		}
		DefaultDomain.DataSource = null.StringFrom(strings.Replace(viper.GetString("db.dataSource"), "?", "localhost?", 1))
		if _, err := s.InsertOne(&DefaultDomain); err != nil {
			s.Rollback()
			panic(err)
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
