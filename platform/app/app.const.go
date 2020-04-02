package app

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
)

// DefaultDomain default domain
var DefaultDomain = model.SysDomain{
	ID:         null.StringFrom(util.AdminID),
	Name:       null.StringFrom("default"),
	FullName:   null.StringFrom("default"),
	Theme:      null.StringFrom("default"),
	DataSource: null.StringFrom(""),
	DriverName: null.StringFrom("mysql"),
	DomainUrl:  null.StringFrom("localhost"),
	LoginUrl:   null.StringFrom("localhost"),
	Type:       null.IntFrom(0),
	Status:     null.IntFrom(1),
	SyncFlag:   null.IntFrom(0),
	Domain:     null.StringFrom("localhost"),
	ApiUrl:     null.StringFrom("http://localhost:8086"),
	CreateBy:   null.StringFrom(util.AdminID),
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   null.StringFrom(util.AdminID),
	UpdateTime: null.TimeFrom(time.Now()),
}

// DefaultAdmin default admin
var DefaultAdmin = model.SysUser{
	ID:         null.StringFrom(util.AdminID),
	Password:   null.StringFrom("admin"),
	Name:       null.StringFrom("admin"),
	FullName:   null.StringFrom("admin"),
	Status:     null.IntFrom(1),
	Domain:     null.StringFrom("localhost"),
	CreateBy:   null.StringFrom(util.AdminID),
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   null.StringFrom(util.AdminID),
	UpdateTime: null.TimeFrom(time.Now()),
}
