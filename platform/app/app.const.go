package app

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/model"
)

// DefaultDomain default domain
var DefaultDomain = model.SysDomain{
	ID:         null.StringFrom("2ba7b810-9dad-11d1-80b4-00c04fd430c3"),
	Name:       null.StringFrom("Default"),
	FullName:   null.StringFrom("Default"),
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
	CreateBy:   DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
}

// DefaultAdmin default admin
var DefaultAdmin = model.SysUser{
	ID:         null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	Password:   null.StringFrom("admin"),
	Name:       null.StringFrom("admin"),
	FullName:   null.StringFrom("admin"),
	Status:     null.IntFrom(1),
	Domain:     null.StringFrom("localhost"),
	CreateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
}

// DefaultClient default client
var DefaultClient = model.SysClient{
	ID:         null.StringFrom("5ba2b110-9dad-11d1-80b4-00c04fd431c4"),
	Client:     null.StringFrom("Y76U9344RABF4"),
	Name:       null.StringFrom("Default"),
	Secret:     null.StringFrom("98UYO6FVB865"),
	Domain:     null.StringFrom("http://127.0.0.1:8082"),
	CreateBy:   DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
}
