// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/util"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// Ensure defined inital system data
func (m *SysDomain) Ensure(db *xorm.Engine) {
	if !util.EnsureLeft(db.IsTableExist(new(SysDomain))).(bool) {
		db.Sync2(new(SysDomain))
	}
}

// InitSysData defined inital system data
func (m *SysDomain) InitSysData(s *xorm.Session) {
	domains := []SysDomain{
		{
			ID:         null.StringFrom("5ba2b810-9dad-11d1-80b4-00c04fd430c1"),
			Name:       null.StringFrom("localhost"),
			FullName:   null.StringFrom("localhost"),
			AppName:    null.StringFrom(viper.GetString("app.name")),
			Theme:      null.StringFrom("default"),
			DriverName: null.StringFrom("mysql"),
			LoginUrl:   null.StringFrom("localhost"),
			Type:       null.IntFrom(0),
			Status:     null.IntFrom(1),
			SyncFlag:   null.IntFrom(0),
			AuthMode:   null.IntFrom(1),
			Domain:     null.StringFrom("localhost"),
			DomainUrl:  null.StringFrom("localhost"),
			ApiUrl:     null.StringFrom("http://localhost:8082"),
			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
	}
	for _, domain := range domains {
		if ct, err := s.Where("id=?", domain.ID.String).Count(new(SysDomain)); ct == 0 || err != nil {
			if err != nil {
				s.Rollback()
				panic(err)
			}
			domain.DataSource = null.StringFrom(strings.Replace(viper.GetString("db.dataSource"), "?", "_localhost?", 1))
			if _, err := s.InsertOne(&domain); err != nil {
				s.Rollback()
				panic(err)
			}
		}
	}

	if err := s.Commit(); err != nil {
		panic(err)
	}
}
