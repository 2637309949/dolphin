package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/null"
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
	ctx.engine.PlatformDB.ID(ctx.GetToken().GetUserID()).Get(&user)
	return user
}

// InRole defined
func (ctx *Context) InRole(role ...string) bool {
	return false
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	code := 200
	if len(status) > 0 {
		code = status[0]
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code: null.IntFrom(int64(code)),
		Data: data,
	})
}

// Fail defined failt result
func (ctx *Context) Fail(err error, status ...int) {
	code := 500
	msg := fmt.Sprintf("%v", errors.WithStack(err))
	if mErr, ok := err.(model.Error); ok {
		code = mErr.Code
	}
	if len(status) > 0 {
		code = status[0]
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code: null.IntFrom(int64(code)),
		Msg:  null.StringFrom(msg),
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
