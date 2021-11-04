// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/packages/web/gin"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
)

// Context defined TODO
type Context struct {
	*gin.Context
	AuthProtocol
	DB         *xorm.Engine
	PlatformDB *xorm.Engine
}

// NewContext defined TODO
func NewContext(dol *Dolphin) *Context {
	return &Context{PlatformDB: dol.PlatformDB, AuthProtocol: &AuthOAuth2{oauth2: dol.OAuth2, jwt: dol.JWT}}
}

// LoginInInfo defined TODO
func (ctx *Context) LoginInInfo(user *types.SysUser) (bool, error) {
	tk := ctx.GetToken()
	if tk != nil {
		return ctx.PlatformDB.ID(tk.GetUserID()).Get(user)
	}
	return false, nil
}

// BusinessDB defined TODO
func (ctx *Context) BusinessDB(domain string) *xorm.Engine {
	return App.Manager.GetBusinessDB(domain)
}

// TypeQuery defined failt result
func (ctx *Context) TypeQuery() *Query {
	return NewQuery(ctx)
}

// TypeQuery defined failt result
func (ctx *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// OmitByZero omit invalid fileds
func (ctx *Context) OmitByZero(source interface{}) (target interface{}) {
	sv := []reflect.StructField{}
	t := reflect.TypeOf(source)
	v := reflect.ValueOf(source)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if zeroType, ok := v.Field(i).Interface().(interface {
				IsZero() bool
			}); !ok || !zeroType.IsZero() {
				sv = append(sv, reflect.StructField{
					Name: strings.Title(t.Field(i).Name),
					Type: t.Field(i).Type,
					Tag:  t.Field(i).Tag,
				})
			}
		}
	}
	target = reflect.New(reflect.StructOf(sv)).Interface()
	sbt, _ := json.Marshal(source)
	json.Unmarshal(sbt, &target)
	return
}

// QueryRange defined TODO
func (ctx *Context) QueryRange(key string, init ...string) (string, string) {
	v := ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) >= 2 {
			return init[0], init[1]
		}
	}
	values := strings.Split(v, ",")
	if len(values) >= 2 {
		return values[0], values[1]
	}
	return "", ""
}

// QueryInt defined TODO
func (ctx *Context) QueryInt(key string, init ...int) int {
	v := ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			return init[0]
		}
		return 0
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// QueryBool defined TODO
func (ctx *Context) QueryBool(key string, init ...bool) bool {
	v := ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			return init[0]
		}
		return false
	}
	i, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return i
}

// QueryString defined TODO
func (ctx *Context) QueryString(key string, init ...string) string {
	v := ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			return init[0]
		}
	}
	return v
}
