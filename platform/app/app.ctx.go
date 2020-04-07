package app

import (
	"container/list"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util"
)

type (
	// Context defined http handle hook context
	Context struct {
		AuthInfo
		*gin.Context
		DB         *xorm.Engine
		PlatformDB *xorm.Engine
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
	ctx.DB.Sql(fmt.Sprintf(`
		select count(distinct(role_id)) from sys_role_user where role_id in (select id from sys_role where code in (%v)) and user_id = "%v"
	`, strings.Join(funk.Map(role, func(r string) string {
		return fmt.Sprintf(`"%v"`, r)
	}).([]string), ","), ctx.GetToken().GetUserID())).Get(&cnt)
	return cnt == len(role)
}

// InAdmin defined
func (ctx *Context) InAdmin() bool {
	exit, _ := ctx.DB.Where(`role_id = ? and user_id = ?`, model.DefaultRole.ID.String, ctx.GetToken().GetUserID()).Exist(new(model.SysRoleUser))
	return exit
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	code := 200
	if len(status) > 0 {
		code = status[0]
	}
	ctx.JSON(http.StatusOK, model.Response{
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
	ctx.JSON(http.StatusOK, model.Response{
		Code: code,
		Msg:  msg,
	})
}

// TypeQuery defined failt result
func (ctx *Context) TypeQuery() *Query {
	return &Query{m: util.M{}, ctx: ctx}
}

// PageSearch defined
func (ctx *Context) PageSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error) {
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
	if result == nil {
		ret := util.M{}
		ret["page"] = page
		ret["size"] = size
		ret["data"] = []interface{}{}
		ret["totalrecords"] = 0
		ret["totalpages"] = 0
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
	ret := util.M{}
	ret["page"] = page
	ret["size"] = size
	ret["data"] = result
	ret["totalrecords"] = records
	ret["totalpages"] = totalpages
	return &ret, nil
}

func toString(v interface{}, defaultValue string) interface{} {
	if v == nil {
		return defaultValue
	}
	switch t := v.(type) {
	case int:
		return t
	case string:
		return t
	default:
		return t
	}
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
	textFiled := "name"
	parentValue := root

	treeNodeList := list.New()
	originNodeList := list.New()

	// 区分root节点和子节点数组
	for _, params := range paramsArr {
		value, parent, text := "", "", ""
		if params[valueFiled] != nil {
			value = fmt.Sprintf("%v", params[valueFiled])
		}
		if params[parentField] != nil {
			parent = fmt.Sprintf("%v", params[parentField])
		}
		if params[textFiled] != nil {
			text = fmt.Sprintf("%v", params[textFiled])
		}
		//json,_ := xorm.JSONString(params, true)
		// 如果根节点root为空，则从parent为空中获取root节点数组
		// 如果root不为空且treeSrcs中包含id为root的，则获取id等于的节点为root节点数组(相当于显示root节点的树)
		// 如果root不为空且treeSrcs中不包含id为root的，则获取parent等于的节点为root节点数组(相当于不显示root节点的树)
		if (parentValue == "" && parent == "") || value == parentValue || parent == parentValue {
			node := &model.TreeNode{
				ID:     value,
				Text:   text,
				Parent: parent,
				Tag:    params,
			}
			treeNodeList.PushBack(node)
			rootArr = append(rootArr, node)
		} else {
			originNodeList.PushBack(&model.TreeNode{
				ID:     value,
				Text:   text,
				Parent: parent,
				Tag:    params,
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

// GetInt defined
func (q *Query) GetInt(key string, init ...interface{}) int64 {
	ret, _ := q.m[key].(int64)
	return ret
}

// SetInt defined
func (q *Query) SetInt(key string, init ...interface{}) {
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
		q.m[key] = i
	}
}

// GetBool defined
func (q *Query) GetBool(key string, init ...interface{}) bool {
	ret, _ := q.m[key].(bool)
	return ret
}

// SetBool defined
func (q *Query) SetBool(key string, init ...interface{}) {
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
}

// GetString defined
func (q *Query) GetString(key string, init ...interface{}) string {
	ret, _ := q.m[key].(string)
	return ret
}

// SetString defined
func (q *Query) SetString(key string, init ...interface{}) {
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
}

// SetArrayString defined
func (q *Query) SetArrayString(key string, init ...[]string) {
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
}

// GetArrayString defined
func (q *Query) GetArrayString(key string, init ...[]string) []string {
	ret, _ := q.m[key].([]string)
	return ret
}

// SetArrayInt defined
func (q *Query) SetArrayInt(key string, init ...[]int64) {
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
}

// GetArrayInt defined
func (q *Query) GetArrayInt(key string, init ...[]int) []int64 {
	ret, _ := q.m[key].([]int64)
	return ret
}

// SetArrayFloat64 defined
func (q *Query) SetArrayFloat64(key string, init ...[]float64) {
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
}

// GetArrayFloat64 defined
func (q *Query) GetArrayFloat64(key string, init ...[]float64) []float64 {
	ret, _ := q.m[key].([]float64)
	return ret
}

// SetArrayBool defined
func (q *Query) SetArrayBool(key string, init ...[]bool) {
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
}

// GetArrayBool defined
func (q *Query) GetArrayBool(key string, init ...[]bool) []bool {
	ret, _ := q.m[key].([]bool)
	return ret
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

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) gin.HandlerFunc {
		return rg.engine.HandlerFunc(h)
	}).([]gin.HandlerFunc)...)
}
