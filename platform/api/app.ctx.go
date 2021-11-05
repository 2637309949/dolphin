// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/2637309949/dolphin/packages/web/gin"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/errors"
)

// Context defined TODO
type Context struct {
	*gin.Context
}

// LoginInInfo defined TODO
func (ctx *Context) LoginInInfo(user *types.SysUser) (bool, error) {
	return App.PlatformDB.ID(ctx.MustToken().GetUserID()).Get(user)
}

// BusinessDB defined TODO
func (ctx *Context) BusinessDB(domain string) *xorm.Engine {
	return App.Manager.GetBusinessDB(domain)
}

// DB defined TODO
func (ctx *Context) DB() (*xorm.Engine, bool) {
	if db := ctx.Get("DB"); db != nil {
		return db.(*xorm.Engine), true
	}
	return nil, false
}

// MustDB defined TODO
func (ctx *Context) MustDB() *xorm.Engine {
	db, ok := ctx.DB()
	if !ok {
		panic(errors.ErrNotFoundDB)
	}
	return db
}

// Token defined TODO
func (ctx *Context) Token() (TokenInfo, bool) {
	if info := ctx.Get("AuthInfo"); info != nil {
		return info.(TokenInfo), true
	}
	return nil, false
}

// MustToken defined TODO
func (ctx *Context) MustToken() TokenInfo {
	tk, ok := ctx.Token()
	if !ok {
		panic(errors.ErrAccessDenied)
	}
	return tk
}

// TypeQuery defined TODO
func (ctx *Context) TypeQuery() *Query {
	return NewQuery(ctx)
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
