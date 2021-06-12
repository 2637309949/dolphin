// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/eriklott/mustache"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type (
	// Context defined http handle hook context
	Context struct {
		*gin.Context
		AuthInfo
		PlatformDB *xorm.Engine
		DB         *xorm.Engine
	}
	// HandlerFunc defines the handler used by gin middleware as return value
	HandlerFunc struct {
		Method       string
		RelativePath string
		Interceptor  []HandlerFunc
		Handler      func(ctx *Context)
	}
	// RouterGroup defines struct that extend from gin.RouterGroup
	RouterGroup struct {
		Handlers []HandlerFunc
		basePath string
		engine   *Engine
	}
)

// HF2Handler defined
func HF2Handler(h func(ctx *Context)) HandlerFunc {
	return HandlerFunc{Handler: h}
}

// Raw defined
func (ctx *Context) Raw() *gin.Context {
	return ctx.Context
}

// reset defined clean vars in ctx
func (ctx *Context) reset() {
	ctx.DB = nil
	ctx.Context = nil
}

// ShouldBindWith defined
func (ctx *Context) ShouldBindWith(ptr interface{}) error {
	req, err := ctx.Request, ctx.ShouldBindQuery(ptr)
	if err != nil {
		return err
	}
	if req.ContentLength > 0 &&
		strings.Contains(req.Header.Get("Content-Type"), "application/json") {
		return ctx.ShouldBindBodyWith(ptr, binding.JSON)
	}
	return nil
}

// LoginInInfo defined
func (ctx *Context) LoginInInfo() model.SysUser {
	user := model.SysUser{}
	ctx.PlatformDB.ID(ctx.GetToken().GetUserID()).Get(&user)
	return user
}

// InRole defined
func (ctx *Context) InRole(role ...string) bool {
	var cnt int
	role, _ = slice.RemoveStringDuplicates(role)
	if _, err := ctx.DB.SQL(
		fmt.Sprintf(`
		    select count(distinct(role_id)) cnt 
			from sys_role_user 
			left join sys_role on sys_role.id=sys_role_user.role_id 
			where sys_role.code in (%v) and sys_role_user.user_id = ? and sys_role_user.is_delete != 1
		`, strings.Join(funk.Map(role, func(id string) string {
			return fmt.Sprintf("'%v'", id)
		}).([]string), ",")),
		ctx.GetToken().GetUserID()).Get(&cnt); err != nil {
		logrus.Error(err)
		return false
	}
	return cnt == len(role)
}

// InAdmin defined
func (ctx *Context) InAdmin(role ...string) bool {
	var cnt int
	role, _ = slice.RemoveStringDuplicates(append(role, model.AdminRole.Code.String))
	if _, err := ctx.DB.SQL(
		fmt.Sprintf(`
			select count(sys_role_user.id) cnt 
			from sys_role_user 
			left join sys_role on sys_role.id=sys_role_user.role_id 
			where sys_role_user.user_id = ? and sys_role.code in (%v) and sys_role_user.is_delete != 1
		`, strings.Join(funk.Map(role, func(id string) string {
			return fmt.Sprintf("'%v'", id)
		}).([]string), ",")),
		ctx.GetToken().GetUserID()).Get(&cnt); err != nil {
		logrus.Error(err)
		return false
	}
	return cnt > 0
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	code := util.SomeOne(status, 200).(int)
	ctx.JSON(http.StatusOK, model.Success{
		Code: code,
		Data: data,
	})
}

// Fail defined failt result
func (ctx *Context) Fail(err error, status ...int) {
	code := 500
	msg := err.Error()
	if mErr, ok := err.(model.Error); ok {
		code = mErr.Code
	}
	code = util.SomeOne(status, code).(int)
	ctx.JSON(http.StatusOK, model.Fail{
		Code: code,
		Msg:  msg,
	})
}

// TypeQuery defined failt result
func (ctx *Context) TypeQuery() *Query {
	return &Query{m: util.M{}, ctx: ctx}
}

// PageSearch defined
func (ctx *Context) PageSearch(db *xorm.Engine, ctr, api, table string, params map[string]interface{}) (*model.PageList, error) {
	page, ok := params["page"].(int)
	if !ok {
		return nil, errors.New("not found page")
	}
	size, ok := params["size"].(int)
	if !ok {
		return nil, errors.New("not found size")
	}
	params["offset"] = (page - 1) * size
	rowsSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s_select.tpl", ctr, api), &params).Query().List()
	if err != nil {
		return nil, err
	}
	cntSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s_count.tpl", ctr, api), &params).Query().List()
	if err != nil {
		return nil, err
	}
	var plt model.PageList
	if len(rowsSet) == 0 {
		plt.Data = []map[string]interface{}{}
		return &plt, nil
	}
	records, _ := cntSet[0]["records"].(int64)
	var totalpages int64 = 0
	if records < int64(size) {
		totalpages = 1
	} else if records%int64(size) == 0 {
		totalpages = records / int64(size)
	} else {
		totalpages = records/int64(size) + 1
	}
	plt.Page = page
	plt.Size = size
	plt.Data = rowsSet
	plt.TotalRecords = int(records)
	plt.TotalPages = int(totalpages)
	return &plt, nil
}

// TreeSearch defined
func (ctx *Context) TreeSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error) {
	rowsSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s.tpl", controller, api), &q).Query().List()
	if err != nil {
		return nil, err
	}

	valueFiled, parentField, nameFiled := "id", "parent", "name"
	treeNodeList, originNodeList := list.New(), list.New()
	root, rootArr := "", []*model.TreeNode{}
	paramsArr, parentValue := rowsSet, root

	for _, params := range paramsArr {
		value, parent, name := "", "", ""
		if params[valueFiled] != nil {
			value = fmt.Sprintf("%v", params[valueFiled])
		}
		if params[parentField] != nil {
			parent = fmt.Sprintf("%v", params[parentField])
		}
		if params[nameFiled] != nil {
			name = fmt.Sprintf("%v", params[nameFiled])
		}
		if (parentValue == "" && parent == "") || value == parentValue || parent == parentValue {
			node := &model.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    params,
				Nodes:  make([]*model.TreeNode, 0),
			}
			treeNodeList.PushBack(node)
			rootArr = append(rootArr, node)
		} else {
			originNodeList.PushBack(&model.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    params,
				Nodes:  make([]*model.TreeNode, 0),
			})
		}
	}

	for ele := treeNodeList.Front(); ele != nil; ele = ele.Next() {
		treeNode := ele.Value.(*model.TreeNode)
		originEle := originNodeList.Front()
		if originEle == nil {
			break
		}
		for originEle != nil {
			originNextEle := originEle.Next()
			originNode := originEle.Value.(*model.TreeNode)
			if originNode.Parent == treeNode.ID {
				treeNodeList.InsertAfter(originNode, ele)
				treeNode.Nodes = append(treeNode.Nodes, originNode)
				originNodeList.Remove(originEle)
			}
			originEle = originNextEle
		}
	}
	return rootArr, nil
}

// GetOptions defined
func (ctx *Context) GetOptions(db *xorm.Engine, keys ...string) (map[string]map[string]interface{}, error) {
	type Options struct {
		Value interface{} `json:"value"`
		Text  string      `json:"text"`
	}
	var optSets []model.SysOptionset
	if err := db.Where("is_delete = 0").In("code", keys).Find(&optSets); err != nil {
		return nil, err
	}

	optMap := make(map[string]map[string]interface{})
	for _, v := range optSets {
		var options []Options
		if err := json.Unmarshal([]byte(v.Value.String), &options); err != nil {
			return nil, err
		}
		if optMap[v.Code.String] == nil {
			optMap[v.Code.String] = map[string]interface{}{}
		}
		for _, item := range options {
			optMap[v.Code.String][item.Text] = item.Value
		}
	}
	return optMap, nil
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

// QueryRange defined
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

// QueryInt defined
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

// QueryBool defined
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

// QueryString defined
func (ctx *Context) QueryString(key string, init ...string) string {
	v := ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			return init[0]
		}
	}
	return v
}

// ParseExcel defined
// []map[string]string{ map[string]string{"prop": "os_name", "label": "校区", "code": "sch_id", "align": "center", "minWidth": 100, "maxWidth": 150}}
func (ctx *Context) ParseExcel(cfg ExcelConfig) ([]map[string]string, error) {
	if ctx.QueryString("__columns__") != "" {
		cstr := ctx.QueryString("__columns__")
		columns := []map[string]interface{}{}
		json.Unmarshal([]byte(cstr), &columns)
		cfg.Header = columns
	}
	return ParseExcel(cfg)
}

// SuccessWithExcel defined
func (ctx *Context) SuccessWithExcel(cfg ExcelConfig) {
	if ctx.QueryString("__columns__") != "" {
		cstr := ctx.QueryString("__columns__")
		columns := []map[string]interface{}{}
		err := json.Unmarshal([]byte(cstr), &columns)
		if err != nil {
			logrus.Error(err)
		}
		cfg.Header = columns
	}
	excelInfo, err := BuildExcel(cfg)
	if err != nil {
		ctx.Fail(err)
		return
	}
	if ctx.QueryString("__name__") != "" {
		cfg.FileName = ctx.QueryString("__name__")
	}
	excelInfo.FileName = cfg.FileName
	ctx.Success(excelInfo)
}

// BusinessDB defined
func (ctx *Context) BusinessDB(domain string) *xorm.Engine {
	return App.Manager.GetBusinessDB(domain)
}

// Persist defined
func (ctx *Context) Persist(db *xorm.Session, ids ...string) (int64, error) {
	return new(model.SysAttachment).Persist(db, ids...)
}

// PersistFile defined
func (ctx *Context) PersistFile(db *xorm.Session, cb func([]model.SysAttachment) error, ids ...string) (int64, error) {
	return new(model.SysAttachment).PersistFile(db, cb, ids...)
}

// Remove defined
func (ctx *Context) Remove(db *xorm.Session, ids ...string) (int64, error) {
	return new(model.SysAttachment).Remove(db, ids...)
}

// RemoveFile defined
func (ctx *Context) RemoveFile(db *xorm.Session, cb func([]model.SysAttachment) error, ids ...string) (int64, error) {
	return new(model.SysAttachment).RemoveFile(db, cb, ids...)
}

// RenderString defined
func (ctx *Context) String(code int, data string, context ...interface{}) error {
	str, err := mustache.NewTemplate().Render(data, context...)
	if err != nil {
		return err
	}
	ctx.Context.String(code, str)
	return nil
}

// RenderFile defined
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

// RenderHTL defined
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

// RenderXML defined
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

// ShouldBindQuery defined
func (ctx *Context) ShouldBindQuery(ptr interface{}) error {
	if ptr == nil {
		return errors.New("first parameter must be an struct")
	}
	urls, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err != nil {
		return err
	}
	ptrVal := reflect.ValueOf(ptr)
	if ptrVal.Kind() == reflect.Ptr {
		ptrVal = ptrVal.Elem()
	}
	if ptrVal.Kind() == reflect.Map &&
		ptrVal.Type().Key().Kind() == reflect.String {
		return ctx.Context.ShouldBindWith(ptr, binding.Query)
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
			if (sf.PkgPath != "" && !sf.Anonymous) || sf.Tag.Get("form") == "-" { // unexported
				continue
			}
			tagValue, _ := head(sf.Tag.Get("form"), ",")
			if tagValue == "" { // default value is FieldName
				tagValue = sf.Name
				if tagValue == "" { // default value is FieldName
					continue
				}
			}
			switch sf.Type {
			case reflect.TypeOf(null.String{}), reflect.TypeOf(null.Time{}):
				if u, ok := urls[tagValue]; ok {
					urls[tagValue] = []string{}
					for i := range u {
						bte, err := json.Marshal(u[i])
						if err != nil {
							return err
						}
						urls[tagValue] = append(urls[tagValue], string(bte))
					}
				}
			}
		}
	}
	ctx.Request.URL.RawQuery = urls.Encode()
	return ctx.Context.ShouldBindWith(ptr, binding.Query)
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		method := methods[i]
		rg.engine.Http.Handle(method, path.Join(rg.basePath, relativePath), append(rg.Handlers, handlers...)...)
	}
}
