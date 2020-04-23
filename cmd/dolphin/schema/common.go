// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package schema

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"unicode"

	"github.com/2637309949/dolphin/packages/go-funk"
)

// Common struct
type Common struct {
	Name string `validate:"required"`
	Desc string `validate:"required"`
	Path string
}

// Import packages
func (c *Common) Import(pkg string) template.HTML {
	if len(pkg) > 0 {
		packages := strings.Split(pkg, ",")
		for i, v := range packages {
			packages[i] = fmt.Sprintf(`    "%v"`, v)
		}
		tmpl := `import (
%s
)`
		return c.Unescaped(fmt.Sprintf(tmpl, strings.Join(packages, "\n")))
	}
	return c.Unescaped("")
}

// Unescaped unescaped
func (c *Common) Unescaped(x string) template.HTML {
	return template.HTML(x)
}

// SQLInsertOne insert one
func (c *Common) SQLInsertOne(table Table, name string) string {
	names := strings.Join(funk.Map(table.Columns, func(col *Column) string {
		return fmt.Sprintf("`%v`", col.Name)
	}).([]string), ",")
	values := strings.Join(funk.Map(table.Columns, func(col *Column) string {
		return fmt.Sprintf("?%v", col.Name)
	}).([]string), ",")
	return `insert into ` + name + `
		(` + names + `)
		values
		(` + values + `)`
}

// SQLSelectOne select one
func (c *Common) SQLSelectOne(table Table, name string) string {
	names := strings.Join(funk.Map(table.Columns, func(col *Column) string {
		return fmt.Sprintf("`%v`", col.Name)
	}).([]string), ",")
	return `select ` + names + ` from ` + name + `
		where  id =?id`
}

// SQLSelectAll select one
func (c *Common) SQLSelectAll(table Table, name string) string {
	names := strings.Join(funk.Map(table.Columns, func(col *Column) string {
		return fmt.Sprintf("`%v`", col.Name)
	}).([]string), ",")
	return `select ` + names + ` from ` + name
}

// SQLUpdateOne update one
func (c *Common) SQLUpdateOne(table Table, name string) string {
	names := strings.Join(funk.Map(table.Columns, func(col *Column) string {
		return fmt.Sprintf("`%v`=?%v", col.Name, col.Name)
	}).([]string), ",")
	return `update ` + name + ` set ` + names + `
		where  id =?id`
}

// ToUpper toUpper
func (c *Common) ToUpper(name string) string {
	return strings.ToUpper(name)
}

// ToTypeValue value
func (c *Common) ToTypeValue(t string, v string) template.HTML {
	if strings.TrimSpace(v) == "" {
		return c.Unescaped("")
	}
	switch t {
	case "int":
		i, e := strconv.ParseInt(v, 10, 64)
		if e != nil {
			panic(e)
		}
		return c.Unescaped(fmt.Sprintf("%v", i))
	case "string":
		return c.Unescaped(fmt.Sprintf(`"%v"`, v))
	case "bool":
		b, e := strconv.ParseBool(v)
		if e != nil {
			panic(e)
		}
		return c.Unescaped(fmt.Sprintf("%v", b))
	}
	return c.Unescaped("")
}

// ToUpperCase uppercase
func (c *Common) ToUpperCase(name string) string {
	var newName string
	var toupper bool
	var initialisms = []string{
		"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP",
		"HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA",
		"SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID",
		"URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS",
	}
	sms, ext := funk.FindString(initialisms, func(sms string) bool {
		return strings.ToUpper(name) == sms
	})
	if ext {
		return sms
	}
	for i, r := range name {
		if i == 0 {
			if fmt.Sprintf("%c", r) != "$" {
				newName = newName + strings.ToUpper(fmt.Sprintf("%c", r))
			} else {
				toupper = true
			}
		} else if fmt.Sprintf("%c", r) == "_" {
			toupper = true
		} else {
			if toupper {
				newName = newName + strings.ToUpper(fmt.Sprintf("%c", r))
				toupper = false
			} else {
				newName = newName + fmt.Sprintf("%c", r)
			}
		}
	}
	return newName
}

// LcFirst first word
func (c *Common) LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// UcFirst first word
func (c *Common) UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// APIPrefix version path
func (c *Common) APIPrefix(v string) string {
	if v == "" {
		return ""
	}
	vf, err := strconv.ParseFloat(v, 64)
	vi := int(vf)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("/v%v", vi)
}

// APIPath api path
func (c *Common) APIPath(apiName, apiPath string) string {
	if apiPath == "" {
		if apiName == "" {
			return ""
		}
		return strings.ReplaceAll(apiName, "_", "/")
	}
	return apiPath
}

// ToTitle title
func (c *Common) ToTitle(title string) string {
	return strings.Title(title)
}

// Ref defined model name
func (c *Common) Ref(m string) string {
	if strings.HasPrefix(m, "$") || strings.HasPrefix(m, "[]$") || strings.HasPrefix(m, "[]*$") {
		if strings.HasPrefix(m, "[]$") {
			return fmt.Sprintf("[]model.%v", c.ToUpperCase(strings.ReplaceAll(m, "[]$", "")))
		} else if strings.HasPrefix(m, "[]*$") {
			return fmt.Sprintf("[]*model.%v", c.ToUpperCase(strings.ReplaceAll(m, "[]*$", "")))
		}
		return fmt.Sprintf("model.%v", c.ToUpperCase(m))
	}
	return m
}

// SRef defined model name
func (c *Common) SRef(m string) string {
	if strings.HasPrefix(m, "$") || strings.HasPrefix(m, "[]$") || strings.HasPrefix(m, "[]*$") {
		if strings.HasPrefix(m, "[]$") {
			return fmt.Sprintf("[]%v", c.ToUpperCase(strings.ReplaceAll(m, "[]$", "")))
		} else if strings.HasPrefix(m, "[]*$") {
			return fmt.Sprintf("[]*%v", c.ToUpperCase(strings.ReplaceAll(m, "[]*$", "")))
		}
		return fmt.Sprintf("%v", c.ToUpperCase(m))
	}
	return m
}

// ORef defined model name
func (c *Common) ORef(m string) string {
	return strings.ReplaceAll(c.Ref(m), "[]", "")
}

// ISArray defined isarray
func (c *Common) ISArray(m string) bool {
	return strings.HasPrefix(m, "[]")
}

// TableName defined table
func (c *Common) TableName(app string, table string) string {
	return fmt.Sprintf("%v", table)
}

// SplitExtends extends model,bean
func (c *Common) SplitExtends(parent string) []string {
	return funk.Map(strings.Split(parent, ","), func(p string) string {
		return strings.ReplaceAll(c.Ref(p), "model.", "")
	}).([]string)
}
