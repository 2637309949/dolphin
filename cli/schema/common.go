// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package schema

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

// Common struct
type Common struct {
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
	if name == "id" {
		return "ID"
	}
	if name == "url" {
		return "URL"
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

// VersionPrefix version path
func (c *Common) VersionPrefix(v string) string {
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

// ToTitle title
func (c *Common) ToTitle(title string) string {
	return strings.Title(title)
}

// Ref defined model name
func (c *Common) Ref(m string) string {
	if strings.HasPrefix(m, "$") || strings.HasPrefix(m, "[]$") {
		if strings.HasPrefix(m, "[]$") {
			return fmt.Sprintf("[]model.%v", c.ToUpperCase(strings.ReplaceAll(m, "[]", "")))
		}
		return fmt.Sprintf("model.%v", c.ToUpperCase(m))
	}
	return m
}

// ISArray defined isarray
func (c *Common) ISArray(m string) bool {
	return strings.HasPrefix(m, "[]")
}

// TableName defined table
func (c *Common) TableName(app string, table string) string {
	return fmt.Sprintf("%v_%v", app, table)
}

// SplitExtends extends model,bean
func (c *Common) SplitExtends(parent string) []string {
	return funk.Map(strings.Split(parent, ","), func(p string) string {
		return strings.ReplaceAll(c.Ref(p), "model.", "")
	}).([]string)
}
