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
	m   map[string]interface{}
}

// GetInt defined TODO
func (q *Query) GetInt(key string, init ...interface{}) int64 {
	ret, _ := q.m[key].(int64)
	return ret
}

// SetInt defined TODO
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

// GetBool defined TODO
func (q *Query) GetBool(key string, init ...interface{}) bool {
	ret, _ := q.m[key].(bool)
	return ret
}

// SetBool defined TODO
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

// GetString defined TODO
func (q *Query) GetString(key string, init ...interface{}) string {
	ret, _ := q.m[key].(string)
	return ret
}

// SetString defined TODO
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

// GetRange defined TODO
func (q *Query) GetRange(key string, init ...interface{}) (string, string) {
	start, _ := q.m[fmt.Sprintf("%v_start", key)].(string)
	end, _ := q.m[fmt.Sprintf("%v_end", key)].(string)
	return start, end
}

// SetRange defined TODO
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

// SetArrayString defined TODO
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

// GetArrayString defined TODO
func (q *Query) GetArrayString(key string, init ...[]string) []string {
	ret, _ := q.m[key].([]string)
	return ret
}

// SetArrayInt defined TODO
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

// GetArrayInt defined TODO
func (q *Query) GetArrayInt(key string, init ...[]int) []int64 {
	ret, _ := q.m[key].([]int64)
	return ret
}

// SetArrayFloat64 defined TODO
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

// GetArrayFloat64 defined TODO
func (q *Query) GetArrayFloat64(key string, init ...[]float64) []float64 {
	ret, _ := q.m[key].([]float64)
	return ret
}

// SetArrayBool defined TODO
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

// GetArrayBool defined TODO
func (q *Query) GetArrayBool(key string, init ...[]bool) []bool {
	ret, _ := q.m[key].([]bool)
	return ret
}

// // SetRule defined TODO
// func (q *Query) SetRule(rules ...string) {
// 	rule := q.ctx.Query("rule_code")
// 	if len(rules) > 0 {
// 		rule = rules[0]
// 	}
// 	var roleRule interface{}
// 	var roleRules []types.SysDataPermissionDetail
// 	userID := q.ctx.GetToken().GetUserID()
// 	err := q.ctx.DB.SqlMapClient("get_user_rule_by_code", rule, userID).Find(&roleRules)
// 	if err != nil {
// 		roleRule = ""
// 	} else {
// 		roleRule = ParseRule(roleRules)
// 	}
// 	q.m["role_rule"] = roleRule
// }

// SetRule defined TODO, role_rule for data`permission
func (q *Query) SetRule(rule string) {
	var roleRules []types.SysDataPermissionDetail
	userID := q.ctx.GetToken().GetUserID()
	q.ctx.DB.SqlMapClient("get_user_rule_by_code", rule, userID).Find(&roleRules)
	for i := range roleRules {
		q.m["is_role_rule_"+roleRules[i].Rule.String] = true
	}
}

// SetUser defined TODO
func (q *Query) SetUser(uid ...string) {
	if len(uid) > 0 {
		q.m["uid"] = uid[0]
	} else {
		q.m["uid"] = q.ctx.GetToken().GetUserID()
	}
}

// GetUser defined TODO
func (q *Query) GetUser() string {
	ret, _ := q.m["uid"].(string)
	return ret
}

// Value defined TODO
func (q *Query) Value() map[string]interface{} {
	return q.m
}

// Unmarshal defined TODO
func (q *Query) Unmarshal(v interface{}) error {
	mbyte, e := json.Marshal(q.m)
	if e != nil {
		return e
	}
	return json.Unmarshal(mbyte, v)
}

// Unescaped defined TODO
func (q *Query) Unescaped(s string) template.HTML {
	return template.HTML(s)
}

// SetTags defined TODO TODO TODO TODO TODO
func (q *Query) SetTags(params ...struct {
	Key   string
	Value string
}) {
	q.SetString("eq", q.Unescaped("="))
	q.SetString("ne", q.Unescaped("<>"))
	q.SetString("lt", q.Unescaped("<"))
	q.SetString("lte", q.Unescaped("<="))
	q.SetString("gt", q.Unescaped(">"))
	q.SetString("gte", q.Unescaped(">="))
	for i := range params {
		q.SetString(params[i].Key, params[i].Value)
	}
}

// Remove defined TODO
func (q *Query) Remove(keys ...string) {
	for _, k := range keys {
		delete(q.m, k)
	}
}

// Reset defined TODO
func (q *Query) Reset() {
	q.m = util.M{}
}

// ParseRule defined TODO
func ParseRule(roleRules []types.SysDataPermissionDetail) interface{} {
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
