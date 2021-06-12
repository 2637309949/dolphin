package model

import (
	"fmt"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/util"
	"golang.org/x/crypto/scrypt"
)

// DefaultAdmin default admin
var DefaultAdmin = SysUser{
	ID:         null.IntFrom(1),
	Password:   null.StringFrom("admin"),
	Name:       null.StringFrom("admin"),
	Nickname:   null.StringFrom("admin"),
	Avatar:     null.StringFrom("http://pic.616pic.com/ys_bnew_img/00/06/27/TWk2P5YJ5k.jpg"),
	Status:     null.IntFrom(1),
	Domain:     null.StringFrom("localhost"),
	OrgId:      null.IntFrom(1),
	TempId:     null.IntFrom(1),
	Creater:    null.IntFrom(1),
	CreateTime: null.TimeFrom(time.Now()),
	Updater:    null.IntFrom(1),
	UpdateTime: null.TimeFrom(time.Now()),
	IsDelete:   null.IntFrom(0),
}

// SetPassword Method to set salt and hash the password for a user
func (m *SysUser) SetPassword(password string) {
	b := util.RandString(16, util.RandNumChar)
	m.Salt = null.StringFrom(b)
	dk, err := scrypt.Key([]byte(password), []byte(m.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	m.Password = null.StringFrom(fmt.Sprintf("%x", dk))
}

// ValidPassword Method to check the entered password is correct or not
func (m *SysUser) ValidPassword(password string) bool {
	dk, err := scrypt.Key([]byte(password), []byte(m.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	return m.Password.String == fmt.Sprintf("%x", dk)
}

// InitSysData defined inital system data
func (m *SysUser) InitSysData(s *xorm.Session) {
	if ct, err := s.Where("id=?", DefaultAdmin.ID.Int64).Count(new(SysUser)); ct == 0 || err != nil {
		if err != nil {
			s.Rollback()
			panic(err)
		}
		DefaultAdmin.SetPassword(DefaultAdmin.Password.String)
		if _, err := s.InsertOne(&DefaultAdmin); err != nil {
			s.Rollback()
			panic(err)
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
