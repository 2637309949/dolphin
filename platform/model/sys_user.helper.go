package model

import (
	"fmt"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/util"
	"golang.org/x/crypto/scrypt"
)

// SetPassword Method to set salt and hash the password for a user
func (u *SysUser) SetPassword(password string) {
	b := util.RandString(16)
	u.Salt = null.StringFrom(b)
	dk, err := scrypt.Key([]byte(password), []byte(u.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	u.Password = null.StringFrom(fmt.Sprintf("%x", dk))
}

// ValidPassword Method to check the entered password is correct or not
func (u *SysUser) ValidPassword(password string) bool {
	dk, err := scrypt.Key([]byte(password), []byte(u.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	return u.Password.String == fmt.Sprintf("%x", dk)
}
