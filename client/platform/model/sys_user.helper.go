package model

import (
	"fmt"

	"github.com/2637309949/dolphin/client/platform/util"
	"github.com/2637309949/dolphin/packages/null"
	"golang.org/x/crypto/scrypt"
)

// Encode structs into URL query parameters
func (u *SysUser) Encode() (string, error) {
	v, err := util.Values(u)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

// SetPassword Method to set salt and hash the password for a user
func (u *SysUser) SetPassword(password string) {
	b, err := util.RandomBytes(16)
	if err != nil {
		panic(err)
	}
	u.Salt = null.StringFrom(fmt.Sprintf("%x", b))
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
	fmt.Println("fmt.Sprintf(dk) = ", fmt.Sprintf("%x", dk))
	return u.Password.String == fmt.Sprintf("%x", dk)
}
