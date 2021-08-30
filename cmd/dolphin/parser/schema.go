// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

type (
	// Common defined TODO
	Common struct {
		Name string `validate:"required"`
		Desc string `validate:"required"`
		Path string
	}
	// Success defined TODO
	Success struct {
		Common
		Type string
	}
	// Failure defined TODO
	Failure struct {
		Common
		Type string
	}
	// Return defined TODO
	Return struct {
		Common
		Success *Success
		Failure *Failure
	}
	// Param defined TODO
	Param struct {
		Common
		Type  string
		Value string
	}
	// API defined TODO
	API struct {
		Common
		Roles   []string
		Auth    []string
		Version string
		Path    string
		Func    string
		Table   string
		Method  string
		Cache   uint64
		Params  []*Param
		Return  *Return
	}
	// Controller defined TODO
	Controller struct {
		Common
		APIS   []*API
		Prefix string
	}

	// Service defined TODO
	Service struct {
		Common
		RPCS []*RPC
	}

	// RPC defined TODO
	RPC struct {
		Common
		Request *Request
		Reply   *Reply
	}

	// Request defined TODO
	Request struct {
		Common
		Type string
	}

	// Reply defined TODO
	Reply struct {
		Common
		Type string
	}

	// Prop defined TODO
	Prop struct {
		Common
		Type    string
		JSON    string
		Form    string
		Example string
	}

	// Bean defined TODO
	Bean struct {
		Common
		Packages string
		Props    []*Prop
		Extends  string
	}

	// Column defined TODO
	Column struct {
		Common
		Type    string
		Xorm    string
		JSON    string
		Form    string
		Example string
	}

	// Table defined TODO
	Table struct {
		Common
		Bind     string
		Packages string
		Columns  []*Column
		Extends  string
	}

	// Application defined TODO
	Application struct {
		Common
		PackageName string `validate:"required"`
		Controllers []*Controller
		Services    []*Service
		Beans       []*Bean
		Tables      []*Table
	}
)

// Import packages
func (c *Common) Import(pkgs ...string) template.HTML {
	pkgs = funk.FilterString(pkgs, func(s string) bool { return strings.TrimSpace(s) != "" })
	if pkg := strings.Join(pkgs, ","); len(pkg) > 0 {
		packages := strings.Split(pkg, ",")
		for i, v := range packages {
			packages[i] = fmt.Sprintf(`    "%v"`, v)
		}
		tmpl := "import (\n%s\n)"
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
	return `select ` + names + ` from ` + name + ` where id =?id`
}

// SQLDelOne remove one
func (c *Common) SQLDelOne(table Table, name string) string {
	return `delete from ` + name + ` where id =?id`
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
		where id =?id`
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

// Title first word
func (c *Common) Title(str string) string {
	return strings.Title(strings.ToLower(str))
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

// Contains defined
func (c *Common) Contains(args []string, s string) bool {
	for _, v := range args {
		if v == s {
			return true
		}
	}
	return false
}

// FormatString defined
func (c *Common) FormatString(args []string, segm string) template.HTML {
	strs := []string{}
	for _, v := range args {
		strs = append(strs, fmt.Sprintf(`"%v"`, v))
	}
	return template.HTML(strings.Join(strs, segm))
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
func (c *Common) APIPath(ctrName, apiName, apiPath string) string {
	if apiPath == "" {
		return "/" + strings.ReplaceAll(ctrName, "_", "/") + "/" + apiName
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
			return fmt.Sprintf("[]%v.%v", viper.GetString("dir.types"), c.ToUpperCase(strings.ReplaceAll(m, "[]$", "")))
		} else if strings.HasPrefix(m, "[]*$") {
			return fmt.Sprintf("[]*%v.%v", viper.GetString("dir.types"), c.ToUpperCase(strings.ReplaceAll(m, "[]*$", "")))
		}
		return fmt.Sprintf("%v.%v", viper.GetString("dir.types"), c.ToUpperCase(m))
	}
	return m
}

// TypeWithPointer defined model name
func (c *Common) TypeWithPointer(m string) string {
	if strings.Contains(m, "*") {
		return m
	}
	splites := strings.Split(m, "[]")
	if len(splites) == 1 {
		return fmt.Sprintf("%v%v", "*", m)
	}
	return fmt.Sprintf("%v%v%v", "[]", "*", splites[1])
}

// PRef defined model name
func (c *Common) PRef(m string) string {
	if strings.HasPrefix(m, "$") {
		return fmt.Sprintf("%v", c.ToUpperCase(strings.ReplaceAll(m, "$", "")))
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
		return strings.ReplaceAll(c.Ref(p), viper.GetString("dir.types")+".", "")
	}).([]string)
}

// TableNameOfType defined
func (c *Common) TableNameOfType(t string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(t, "[]", ""), "$", ""), viper.GetString("dir.types")+".", "")
}

// AutoIncr defined
func (c *Common) AutoIncr(tables []*Table, tableName string) (fields []string) {
	for i := range tables {
		if tables[i].Name == tableName {
			for c := range tables[i].Columns {
				if strings.Contains(tables[i].Columns[c].Xorm, "pk") && !strings.Contains(tables[i].Columns[c].Xorm, "autoincr") {
					fields = append(fields, tables[i].Columns[c].ToUpperCase(tables[i].Columns[c].Name))
				}
			}
		}
	}
	return fields
}
