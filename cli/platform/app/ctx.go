package app

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/2637309949/dolphin/cli/packages/gin"
	"github.com/2637309949/dolphin/cli/packages/go-funk"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/2637309949/dolphin/cli/platform/util"
)

type (
	// Context defined http handle hook context
	Context struct {
		*gin.Context
		AuthInfo
		engine *Engine
		DB     *xorm.Engine
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

// IsAdmin defined
func (ctx *Context) IsAdmin() bool {
	return false
}

// InRole defined
func (ctx *Context) InRole(role ...string) bool {
	return false
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	sise, code := http.StatusOK, http.StatusOK
	ctx.JSON(sise, model.Response{
		Code: null.IntFrom(int64(code)),
		Data: data,
	})
}

// Fail defined failt result
func (ctx *Context) Fail(err error, status ...int) {
	sise, code := http.StatusOK, http.StatusInternalServerError
	msg := fmt.Sprintf("%v", errors.WithStack(err))
	if cusErr, ok := err.(model.Error); ok {
		code = cusErr.Code
	}
	if len(status) > 0 {
		sise = status[0]
	}
	ctx.JSON(sise, model.Response{
		Code: null.IntFrom(int64(code)),
		Msg:  null.StringFrom(msg),
	})
}

// TypeQuery defined failt result
func (ctx *Context) TypeQuery() *Query {
	q := Query{m: util.M{}, ctx: ctx}
	q.SetTags()
	return &q
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

// Value defined
func (q *Query) Value() map[string]interface{} {
	return q.m
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

// HandlerFunc convert to gin.HandlerFunc
func (h HandlerFunc) HandlerFunc(e *Engine) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c := e.pool.Get().(*Context)
		c.Context = ctx
		// from upper middles dataset
		for _, v := range ctx.Keys {
			switch t := v.(type) {
			case *xorm.Engine:
				c.DB = t
			case AuthInfo:
				c.AuthInfo = t
			}
		}
		h(c)
		e.pool.Put(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) gin.HandlerFunc {
		return h.HandlerFunc(rg.engine)
	}).([]gin.HandlerFunc)...)
}
