// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

// TmplModel defined template
var TmplModel = `// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

{{.Table.Import .Table.Packages}}

// {{.Table.ToUpperCase .Table.Name}} defined {{.Table.Desc}} 
type {{.Table.ToUpperCase .Table.Name}} struct {
	{{- range .Table.Columns}}
	// {{.Desc}}
	{{.ToUpperCase .Name}} {{.Type}} ` + "`" + `xorm:"{{.Xorm}} '{{.Name}}'" json:"{{.Name}}" xml:"{{.Name}}"` + "`" + `
	{{- end}}
}

// TableName table name of defined {{.Table.ToUpperCase .Table.Name}}
func (m {{.Table.ToUpperCase .Table.Name}}) TableName() string {
	return "{{.Table.Name}}"
}
`