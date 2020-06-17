// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"container/list"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
)

type (
	// ZeroType defined
	ZeroType interface {
		IsZero() bool
	}
	// Context defined http handle hook context
	Context struct {
		AuthInfo
		*gin.Context
		DB         *xorm.Engine
		PlatformDB *xorm.Engine
		OAuth2     *server.Server
		engine     *Engine
	}
	// HandlerFunc defines the handler used by gin middleware as return value.
	HandlerFunc func(*Context)
	// RouterGroup defines struct that extend from gin.RouterGroup
	RouterGroup struct {
		*gin.RouterGroup
		engine *Engine
	}
	// Query defined parse struct from query
	Query struct {
		ctx *Context
		m   map[string]interface{}
	}
)

// LoginInInfo defined
func (ctx *Context) LoginInInfo() model.SysUser {
	user := model.SysUser{}
	ctx.PlatformDB.ID(ctx.GetToken().GetUserID()).Get(&user)
	return user
}

// InRole defined
func (ctx *Context) InRole(role ...string) bool {
	var cnt int
	if _, err := ctx.DB.SqlTemplateClient("sys_role_in_role_cnt.tpl", &map[string]interface{}{
		"user_id": ctx.GetToken().GetUserID(),
		"roles": template.HTML(strings.Join(funk.Map(role, func(r string) string {
			return fmt.Sprintf(`"%v"`, r)
		}).([]string), ",")),
	}).Get(&cnt); err != nil {
		logrus.Error(err)
		return false
	}
	return cnt == len(role)
}

// InAdmin defined
func (ctx *Context) InAdmin() bool {
	exit, _ := ctx.DB.Where(`role_id = ? and user_id = ?`, model.AdminRole.ID.String, ctx.GetToken().GetUserID()).Exist(new(model.SysRoleUser))
	return exit
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	code := 200
	if len(status) > 0 {
		code = status[0]
	}
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
	if len(status) > 0 {
		code = status[0]
	}
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
func (ctx *Context) PageSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (*model.PageList, error) {
	page := q["page"].(int)
	size := q["size"].(int)
	q["offset"] = (page - 1) * size
	sqlTagName := fmt.Sprintf("%s_%s_select.tpl", controller, api)
	result, err := db.SqlTemplateClient(sqlTagName, &q).Query().List()
	if err != nil {
		return nil, err
	}
	sqlTagName = fmt.Sprintf("%s_%s_count.tpl", controller, api)
	cresult, err := db.SqlTemplateClient(sqlTagName, &q).Query().List()
	if err != nil {
		return nil, err
	}
	var ret model.PageList
	if result == nil || len(result) == 0 {
		ret.Data = []map[string]interface{}{}
		return &ret, nil
	}
	records := cresult[0]["records"].(int64)
	var totalpages int64 = 0
	if records < int64(size) {
		totalpages = 1
	} else if records%int64(size) == 0 {
		totalpages = records / int64(size)
	} else {
		totalpages = records/int64(size) + 1
	}
	ret.Page = page
	ret.Size = size
	ret.Data = result
	ret.TotalRecords = int(records)
	ret.TotalPages = int(totalpages)
	return &ret, nil
}

// TreeSearch defined
func (ctx *Context) TreeSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error) {
	stplkey := fmt.Sprintf("%s_%s.tpl", controller, api)
	result, err := db.SqlTemplateClient(stplkey, &q).Query().List()
	if err != nil {
		return nil, err
	}

	root := ""
	rootArr := []*model.TreeNode{}
	paramsArr := result
	valueFiled := "id"
	parentField := "parent"
	nameFiled := "name"
	parentValue := root

	treeNodeList := list.New()
	originNodeList := list.New()

	// 区分root节点和子节点数组
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
		//json,_ := xorm.JSONString(params, true)
		// 如果根节点root为空，则从parent为空中获取root节点数组
		// 如果root不为空且treeSrcs中包含id为root的，则获取id等于的节点为root节点数组(相当于显示root节点的树)
		// 如果root不为空且treeSrcs中不包含id为root的，则获取parent等于的节点为root节点数组(相当于不显示root节点的树)
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

	// 把子节点根据parent分配到对应的父节点上
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
	var optionset []model.SysOptionset
	err := db.Where("del_flag = 0").In("code", keys).Find(&optionset)
	if err != nil {
		return nil, err
	}

	optionMap := make(map[string]map[string]interface{})
	for _, v := range optionset {
		var options []Options
		if err = json.Unmarshal([]byte(v.Value.String), &options); err != nil {
			return nil, err
		}
		if optionMap[v.Code.String] == nil {
			optionMap[v.Code.String] = map[string]interface{}{}
		}
		for _, item := range options {
			optionMap[v.Code.String][item.Text] = item.Value
		}
	}
	return optionMap, nil
}

// OmitByZero omit invalid fileds
func (ctx *Context) OmitByZero(source interface{}) (target interface{}) {
	sv := []reflect.StructField{}
	t := reflect.TypeOf(source)
	v := reflect.ValueOf(source)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if zeroType, ok := v.Field(i).Interface().(ZeroType); !ok || !zeroType.IsZero() {
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
// []Msi{ Msi{"prop": "os_name", "label": "校区", "code": "sch_id", "align": "center", "minWidth": 100, "maxWidth": 150}}
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
		json.Unmarshal([]byte(cstr), &columns)
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

// GetInt defined
func (q *Query) GetInt(key string, init ...interface{}) int64 {
	ret, _ := q.m[key].(int64)
	return ret
}

// SetInt defined
func (q *Query) SetInt(key string, init ...interface{}) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = 0
		}
	} else {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		q.m[key] = int(i)
	}
	return func() {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = 0
		}
	}
}

// GetBool defined
func (q *Query) GetBool(key string, init ...interface{}) bool {
	ret, _ := q.m[key].(bool)
	return ret
}

// SetBool defined
func (q *Query) SetBool(key string, init ...interface{}) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = false
		}
	} else {
		i, err := strconv.ParseBool(v)
		if err != nil {
			panic(err)
		}
		q.m[key] = i
	}
	return func() {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = false
		}
	}
}

// GetString defined
func (q *Query) GetString(key string, init ...interface{}) string {
	ret, _ := q.m[key].(string)
	return ret
}

// SetString defined
func (q *Query) SetString(key string, init ...interface{}) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = ""
		}
	} else {
		q.m[key] = v
	}
	return func() {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = ""
		}
	}
}

// GetRange defined
func (q *Query) GetRange(key string, init ...interface{}) (string, string) {
	start, _ := q.m[fmt.Sprintf("%v_start", key)].(string)
	end, _ := q.m[fmt.Sprintf("%v_end", key)].(string)
	return start, end
}

// SetRange defined
func (q *Query) SetRange(key string, init ...[]string) func() {
	v := q.ctx.Query(key)
	q.m[fmt.Sprintf("%v_start", key)] = ""
	q.m[fmt.Sprintf("%v_end", key)] = ""
	if strings.TrimSpace(v) == "" {
		if len(init) >= 2 {
			q.m[fmt.Sprintf("%v_start", key)] = init[0]
			q.m[fmt.Sprintf("%v_end", key)] = init[1]
		}
	} else {
		values := strings.Split(v, ",")
		if len(values) >= 2 {
			q.m[fmt.Sprintf("%v_start", key)] = values[0]
			q.m[fmt.Sprintf("%v_end", key)] = values[1]
		}
	}
	return func() {
		if len(init) >= 2 {
			q.m[fmt.Sprintf("%v_start", key)] = init[0]
			q.m[fmt.Sprintf("%v_end", key)] = init[1]
		}
	}
}

// SetArrayString defined
func (q *Query) SetArrayString(key string, init ...[]string) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []string{}
		}
	} else {
		q.m[key] = strings.Split(v, ",")
	}
	return func() {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []string{}
		}
	}
}

// GetArrayString defined
func (q *Query) GetArrayString(key string, init ...[]string) []string {
	ret, _ := q.m[key].([]string)
	return ret
}

// SetArrayInt defined
func (q *Query) SetArrayInt(key string, init ...[]int64) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []int64{}
		}
	} else {
		q.m[key] = funk.Map(strings.Split(v, ","), func(i string) int64 {
			it, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				panic(err)
			}
			return it
		}).([]int64)
	}
	return func() {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []int64{}
		}
	}
}

// GetArrayInt defined
func (q *Query) GetArrayInt(key string, init ...[]int) []int64 {
	ret, _ := q.m[key].([]int64)
	return ret
}

// SetArrayFloat64 defined
func (q *Query) SetArrayFloat64(key string, init ...[]float64) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []float64{}
		}
	} else {
		q.m[key] = funk.Map(strings.Split(v, ","), func(i string) float64 {
			it, err := strconv.ParseFloat(i, 64)
			if err != nil {
				panic(err)
			}
			return it
		}).([]float64)
	}
	return func() {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []float64{}
		}
	}
}

// GetArrayFloat64 defined
func (q *Query) GetArrayFloat64(key string, init ...[]float64) []float64 {
	ret, _ := q.m[key].([]float64)
	return ret
}

// SetArrayBool defined
func (q *Query) SetArrayBool(key string, init ...[]bool) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []bool{}
		}
	} else {
		q.m[key] = funk.Map(strings.Split(v, ","), func(i string) bool {
			it, err := strconv.ParseBool(i)
			if err != nil {
				panic(err)
			}
			return it
		}).([]bool)
	}
	return func() {
		if len(init) > 0 {
			q.m[key] = init[0]
		} else {
			q.m[key] = []bool{}
		}
	}
}

// GetArrayBool defined
func (q *Query) GetArrayBool(key string, init ...[]bool) []bool {
	ret, _ := q.m[key].([]bool)
	return ret
}

// SetRule defined
func (q *Query) SetRule(rules ...string) {
	rule := q.ctx.Query("rule_code")
	if len(rules) > 0 {
		rule = rules[0]
	}
	var roleRule interface{}
	var roleRules []model.SysDataPermissionDetail
	userID := q.ctx.GetToken().GetUserID()
	err := q.ctx.DB.SqlMapClient("get_user_rule_by_code", rule, userID).Find(&roleRules)
	if err != nil {
		roleRule = ""
	} else {
		roleRule = ParseRule(roleRules)
	}
	q.m["role_rule"] = roleRule
}

// SetUser defined
func (q *Query) SetUser() {
	q.m["uid"] = q.ctx.GetToken().GetUserID()
}

// GetUser defined
func (q *Query) GetUser() string {
	ret, _ := q.m["uid"].(string)
	return ret
}

// Value defined
func (q *Query) Value() map[string]interface{} {
	return q.m
}

// Unmarshal defined
func (q *Query) Unmarshal(v interface{}) error {
	mbyte, e := json.Marshal(q.m)
	if e != nil {
		return e
	}
	json.Unmarshal(mbyte, v)
	return nil
}

// Unescaped defined
func (q *Query) Unescaped(s string) template.HTML {
	return template.HTML(s)
}

// SetTags defined
func (q *Query) SetTags() {
	q.SetString("eq", q.Unescaped("="))
	q.SetString("ne", q.Unescaped("<>"))
	q.SetString("lt", q.Unescaped("<"))
	q.SetString("lte", q.Unescaped("<="))
	q.SetString("gt", q.Unescaped(">"))
	q.SetString("gte", q.Unescaped(">="))
}

// ParseRule defined
func ParseRule(roleRules []model.SysDataPermissionDetail) interface{} {
	roleRule := ""
	for i, rule := range roleRules {
		if i == 0 {
			roleRule = rule.Rule.String
		} else if i > 0 {
			if roleRule != "" {
				roleRule = roleRule + " or " + rule.Rule.String
			} else {
				roleRule = rule.Rule.String
			}
		}
	}
	if len(roleRules) > 1 {
		roleRule = "(" + roleRule + ")"
	}
	return template.HTML(roleRule)
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) gin.HandlerFunc {
		return rg.engine.HandlerFunc(h)
	}).([]gin.HandlerFunc)...)
}
