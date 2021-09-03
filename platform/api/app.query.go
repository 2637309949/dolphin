// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/thoas/go-funk"
)

// Query defined TODO
type Query struct {
	ctx *Context
	sm  map[string]interface{}
}

// GetInt defined TODO
func (q *Query) GetInt(key string, init ...int) int {
	ret, _ := q.sm[key].(int)
	return ret
}

// SetInt defined TODO
func (q *Query) SetInt(key string, init ...int) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		q.sm[key] = util.SomeOne(init, 0).(int)
	} else {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		q.sm[key] = int(i)
	}
	return func() {
		q.sm[key] = util.SomeOne(init, 0).(int)
	}
}

// GetBool defined TODO
func (q *Query) GetBool(key string, init ...bool) bool {
	ret, _ := q.sm[key].(bool)
	return ret
}

// SetBool defined TODO
func (q *Query) SetBool(key string, init ...bool) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		q.sm[key] = util.SomeOne(init, false).(bool)
	} else {
		i, err := strconv.ParseBool(v)
		if err != nil {
			panic(err)
		}
		q.sm[key] = i
	}
	return func() {
		q.sm[key] = util.SomeOne(init, false).(bool)
	}
}

// GetString defined TODO
func (q *Query) GetString(key string, init ...string) string {
	ret, _ := q.sm[key].(string)
	return ret
}

// SetUnescaped defined TODO
func (q *Query) SetUnescaped(key string, value string) {
	q.sm[key] = template.HTML(value)
}

// SetString defined TODO
func (q *Query) SetString(key string, init ...string) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		q.sm[key] = util.SomeOne(init, "").(string)
	} else {
		q.sm[key] = v
	}
	return func() {
		q.sm[key] = util.SomeOne(init, "").(string)
	}
}

// GetRange defined TODO
func (q *Query) GetRange(key string, init ...string) (string, string) {
	start, _ := q.sm[fmt.Sprintf("%v_start", key)].(string)
	end, _ := q.sm[fmt.Sprintf("%v_end", key)].(string)
	return start, end
}

// SetRange defined TODO
func (q *Query) SetRange(key string, init ...[]string) func() {
	v := q.ctx.Query(key)
	q.sm[fmt.Sprintf("%v_start", key)] = ""
	q.sm[fmt.Sprintf("%v_end", key)] = ""
	if strings.TrimSpace(v) == "" {
		if len(init) >= 2 {
			q.sm[fmt.Sprintf("%v_start", key)] = init[0]
			q.sm[fmt.Sprintf("%v_end", key)] = init[1]
		}
	} else {
		values := strings.Split(v, ",")
		if len(values) >= 2 {
			q.sm[fmt.Sprintf("%v_start", key)] = values[0]
			q.sm[fmt.Sprintf("%v_end", key)] = values[1]
		}
	}
	return func() {
		if len(init) >= 2 {
			q.sm[fmt.Sprintf("%v_start", key)] = init[0]
			q.sm[fmt.Sprintf("%v_end", key)] = init[1]
		}
	}
}

// SetArrayString defined TODO
func (q *Query) SetArrayString(key string, init ...[]string) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		q.sm[key] = util.SomeOne(init, []string{}).([]string)
	} else {
		q.sm[key] = strings.Split(v, ",")
	}
	return func() {
		q.sm[key] = util.SomeOne(init, []string{}).([]string)
	}
}

// GetArrayString defined TODO
func (q *Query) GetArrayString(key string, init ...[]string) []string {
	ret, _ := q.sm[key].([]string)
	return ret
}

// SetArrayInt defined TODO
func (q *Query) SetArrayInt(key string, init ...[]int64) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		q.sm[key] = util.SomeOne(init, []int64{}).([]int64)
	} else {
		q.sm[key] = funk.Map(strings.Split(v, ","), func(i string) int64 {
			it, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				panic(err)
			}
			return it
		}).([]int64)
	}
	return func() {
		q.sm[key] = util.SomeOne(init, []int64{}).([]int64)
	}
}

// GetArrayInt defined TODO
func (q *Query) GetArrayInt(key string, init ...[]int) []int64 {
	ret, _ := q.sm[key].([]int64)
	return ret
}

// SetArrayFloat64 defined TODO
func (q *Query) SetArrayFloat64(key string, init ...[]float64) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		q.sm[key] = util.SomeOne(init, []float64{}).([]float64)
	} else {
		q.sm[key] = funk.Map(strings.Split(v, ","), func(i string) float64 {
			it, err := strconv.ParseFloat(i, 64)
			if err != nil {
				panic(err)
			}
			return it
		}).([]float64)
	}
	return func() {
		q.sm[key] = util.SomeOne(init, []float64{}).([]float64)
	}
}

// GetArrayFloat64 defined TODO
func (q *Query) GetArrayFloat64(key string, init ...[]float64) []float64 {
	ret, _ := q.sm[key].([]float64)
	return ret
}

// SetArrayBool defined TODO
func (q *Query) SetArrayBool(key string, init ...[]bool) func() {
	v := q.ctx.Query(key)
	if strings.TrimSpace(v) == "" {
		q.sm[key] = util.SomeOne(init, []bool{}).([]bool)
	} else {
		q.sm[key] = funk.Map(strings.Split(v, ","), func(i string) bool {
			it, err := strconv.ParseBool(i)
			if err != nil {
				panic(err)
			}
			return it
		}).([]bool)
	}
	return func() {
		q.sm[key] = util.SomeOne(init, []bool{}).([]bool)
	}
}

// GetArrayBool defined TODO
func (q *Query) GetArrayBool(key string, init ...[]bool) []bool {
	ret, _ := q.sm[key].([]bool)
	return ret
}

// SetRule defined TODO, role_rule for data`permission
func (q *Query) SetRule(rule string) {
	var roleRules []types.SysDataPermissionDetail
	userID := q.ctx.GetToken().GetUserID()
	q.ctx.DB.SqlMapClient("get_user_rule_by_code", rule, userID).Find(&roleRules)
	for i := range roleRules {
		q.sm["is_role_rule_"+roleRules[i].Rule.String] = true
	}
}

// SetUser defined TODO
func (q *Query) SetUser(uid ...int64) {
	it, _ := strconv.ParseInt(q.ctx.GetToken().GetUserID(), 10, 64)
	q.sm["uid"] = util.SomeOne(uid, it).(int64)
}

// GetUser defined TODO
func (q *Query) GetUser() int64 {
	ret, _ := q.sm["uid"].(int64)
	return ret
}

// Value defined TODO
func (q *Query) Value() map[string]interface{} {
	return q.sm
}

// Unmarshal defined TODO
func (q *Query) Unmarshal(v interface{}) error {
	mbyte, e := json.Marshal(q.sm)
	if e != nil {
		return e
	}
	return json.Unmarshal(mbyte, v)
}

// SetTags defined TODO TODO TODO TODO TODO
func (q *Query) SetTags(params ...struct {
	Key   string
	Value string
}) {
	q.SetUnescaped("eq", "=")
	q.SetUnescaped("ne", "<>")
	q.SetUnescaped("lt", "<")
	q.SetUnescaped("lte", "<=")
	q.SetUnescaped("gt", ">")
	q.SetUnescaped("gte", ">=")
	for i := range params {
		q.SetString(params[i].Key, params[i].Value)
	}
}

// Remove defined TODO
func (q *Query) Remove(keys ...string) {
	for i := range keys {
		delete(q.sm, keys[i])
	}
}

// Reset defined TODO
func (q *Query) Reset() {
	q.sm = util.M{}
}

// NewQuery defined TODO
func NewQuery(ctx *Context) *Query {
	return &Query{sm: util.M{}, ctx: ctx}
}
