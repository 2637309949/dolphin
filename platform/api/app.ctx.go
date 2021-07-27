// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/eriklott/mustache"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

// Context defined TODO
type Context struct {
	AuthInfo
	*gin.Context
	DB         *xorm.Engine
	PlatformDB *xorm.Engine
}

// NewContext defined TODO
func NewContext(dol *Dolphin) *Context {
	return &Context{PlatformDB: dol.PlatformDB, AuthInfo: &AuthOAuth2{oauth2: dol.OAuth2, jwt: dol.JWT}}
}

// reset defined clean vars in ctx
func (ctx *Context) reset() {
	ctx.DB = nil
	ctx.Context = nil
}

// LoginInInfo defined TODO
func (ctx *Context) LoginInInfo(user *types.SysUser) (bool, error) {
	tk := ctx.GetToken()
	return ctx.PlatformDB.ID(tk.GetUserID()).Get(user)
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	code := util.SomeOne(status, 200).(int)
	ctx.JSON(http.StatusOK, types.Success{
		Code: code,
		Data: data,
	})
}

// Fail defined failt result
func (ctx *Context) Fail(err error, status ...int) {
	code, msg := 500, err.Error()
	if mErr, ok := err.(types.Error); ok {
		code = mErr.Code
	}
	code = util.SomeOne(status, code).(int)
	ctx.JSON(http.StatusOK, types.Fail{
		Code: code,
		Msg:  msg,
	})
}

// TypeQuery defined failt result
func (ctx *Context) TypeQuery() *Query {
	return &Query{m: util.M{}, ctx: ctx}
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

// BusinessDB defined TODO
func (ctx *Context) BusinessDB(domain string) *xorm.Engine {
	return App.Manager.GetBusinessDB(domain)
}

// String defined TODO
func (ctx *Context) String(code int, data string, context ...interface{}) error {
	str, err := mustache.NewTemplate().Render(data, context...)
	if err != nil {
		return err
	}
	ctx.Context.String(code, str)
	return nil
}

// RenderFile defined TODO
func (ctx *Context) RenderFile(filepath string, filename string, context ...interface{}) {
	ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	file, err := ioutil.TempFile("", "*")
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	defer os.Remove(file.Name())
	bte, err := ioutil.ReadFile(filepath)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	str, err := mustache.NewTemplate().Render(string(bte), context...)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if _, err = file.WriteString(str); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	http.ServeFile(ctx.Writer, ctx.Request, file.Name())
}

// RenderHTL defined TODO
func (ctx *Context) RenderHTML(filepath string, context ...interface{}) {
	ctx.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	file, err := ioutil.TempFile("", "*")
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	defer os.Remove(file.Name())
	bte, err := ioutil.ReadFile(filepath)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	str, err := mustache.NewTemplate().Render(string(bte), context...)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if _, err = file.WriteString(str); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	http.ServeFile(ctx.Writer, ctx.Request, file.Name())
}

// RenderXML defined TODO
func (ctx *Context) RenderXML(filepath string, context ...interface{}) {
	ctx.Writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
	file, err := ioutil.TempFile("", "*")
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	defer os.Remove(file.Name())
	bte, err := ioutil.ReadFile(filepath)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	str, err := mustache.NewTemplate().Render(string(bte), context...)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if _, err = file.WriteString(str); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	http.ServeFile(ctx.Writer, ctx.Request, file.Name())
}

// ShouldBindWith defined TODO
func (ctx *Context) ShouldBindWith(ptr interface{}) error {
	if ptr == nil {
		return errors.New("first parameter must be an pointer")
	}
	err := ctx.ShouldBindQuery(ptr)
	if err != nil {
		return err
	}
	if util.JsonCheck(ctx.Request) {
		return ctx.ShouldBindBody(ptr)
	}
	return nil
}

// ShouldBindBody defined TODO
func (ctx *Context) ShouldBindBody(ptr interface{}) error {
	var err error
	var body []byte
	if cb, ok := ctx.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil {
		body, err = ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			return err
		}
		ctx.Set(gin.BodyBytesKey, body)
	}
	return binding.JSON.BindBody(body, ptr)
}

// ShouldBindQuery defined TODO, support for nullType
func (ctx *Context) ShouldBindQuery(ptr interface{}) error {
	if ptr == nil {
		return errors.New("first parameter must be an pointer")
	}
	rawQuery := ctx.Request.URL.RawQuery
	urls, err := url.ParseQuery(rawQuery)
	if err != nil {
		return err
	}
	ptrVal := reflect.ValueOf(ptr)
	if ptrVal.Kind() == reflect.Ptr {
		ptrVal = ptrVal.Elem()
	}
	if ptrVal.Kind() == reflect.Map &&
		ptrVal.Type().Key().Kind() == reflect.String {
		return binding.Query.Bind(ctx.Request, ptr)
	}
	var head = func(str, sep string) (head string, tail string) {
		idx := strings.Index(str, sep)
		if idx < 0 {
			return str, ""
		}
		return str[:idx], str[idx+len(sep):]
	}
	var vKind = ptrVal.Kind()
	if vKind == reflect.Struct {
		tValue := ptrVal.Type()
		for i := 0; i < ptrVal.NumField(); i++ {
			sf := tValue.Field(i)
			if (sf.PkgPath != "" && !sf.Anonymous) || sf.Tag.Get("form") == "-" {
				continue
			}
			tagValue, _ := head(sf.Tag.Get("form"), ",")
			if tagValue == "" {
				tagValue = sf.Name
				if tagValue == "" {
					continue
				}
			}
			switch sf.Type {
			case reflect.TypeOf(""), reflect.TypeOf(null.String{}), reflect.TypeOf(null.Time{}):
				if u, ok := urls[tagValue]; ok {
					urls[tagValue] = []string{}
					for i := range u {
						istr := u[i]
						if !json.Valid([]byte(istr)) {
							bte, err := json.Marshal(u[i])
							if err != nil {
								return err
							}
							urls[tagValue] = append(urls[tagValue], string(bte))
						} else {
							urls[tagValue] = append(urls[tagValue], istr)
						}
					}
				}
			}
		}
	}
	ctx.Request.URL.RawQuery = urls.Encode()
	err = binding.Query.Bind(ctx.Request, ptr)
	ctx.Request.URL.RawQuery = rawQuery
	return err
}
