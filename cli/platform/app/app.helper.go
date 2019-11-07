package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/cli/platform/util"
	"github.com/xormplus/xorm"
)

// Query defined parse struct from query
type Query struct {
	ctx *Context
	m   map[string]interface{}
}

// SetInt defined
func (q *Query) SetInt(key string, init ...int) {
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

// SetString defined
func (q *Query) SetString(key string, init ...string) {
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

// Query defined
func (e *Engine) Query(ctx *Context) *Query {
	return &Query{m: util.M{}, ctx: ctx}
}

// PageSearch defined
func (e *Engine) PageSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error) {
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
		ret[""] = 0
		return &ret, nil
	}

	records := cresult[0]["records"].(int64)
	var totalpages int64 = 0
	// if records < int64(size) {
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
