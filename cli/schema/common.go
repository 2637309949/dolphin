// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package schema

import (
	"fmt"
	"html/template"
	"strings"
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
	return template.HTML("")
}

// Unescaped unescaped
func (c *Common) Unescaped(x string) template.HTML {
	return template.HTML(x)
}

// ToUpper toUpper
func (c *Common) ToUpper(name string) string {
	return strings.ToUpper(name)
}

// ToUpperCase uppercase
func (c *Common) ToUpperCase(name string) string {
	var newName string
	var toupper bool
	if name == "id" {
		return "ID"
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
