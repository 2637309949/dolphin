package app

import (
	"strconv"
	"strings"
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
