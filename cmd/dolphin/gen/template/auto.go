// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

// TmplAuto defined template
var TmplAuto = `// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package app

import (
	"{{.PackageName}}/model"

	"github.com/2637309949/dolphin/packages/viper"
)

// Name project
{{- $pName := .Name }}
var Name = "{{.Name}}"

{{- range .Controllers}}
{{- $ctr := .}}
// {{.ToUpperCase .Name}} defined
type {{.ToUpperCase .Name}} struct {
	{{- $APIS := .APIS}}
	{{- range $index, $api := .APIS}}
	{{.ToUpperCase $api.Name}}{{- if last $index $APIS}} func(ctx *Context){{- else}},{{- end}}
	{{- end}}
}

// New{{.ToUpperCase .Name}} defined
func New{{.ToUpperCase .Name}}() *{{.ToUpperCase .Name}} {
	ctr := &{{.ToUpperCase .Name}}{}
	{{- range .APIS}}
	ctr.{{.ToUpperCase .Name}} = {{$ctr.ToUpperCase $ctr.Name}}{{.ToUpperCase .Name}}
	{{- end}}
	return ctr
}

// {{.ToUpperCase .Name}}Routes defined
{{- $ctrName := .Name}}
func {{.ToUpperCase .Name}}Routes(engine *Engine) {
	group := engine.Group({{- if ne .Prefix ""}}"{{.Prefix}}"{{- else}}viper.GetString("http.prefix"){{- end}})
	{{- range .APIS}}
	group.Handle("{{.ToUpper .Method}}", "{{.APIPrefix .Version}}/{{.APIPath $ctr.Name .Path}}/{{.Name}}"{{- if .Auth}}, engine.Auth(){{- end}}{{- if gt (len .Roles) 0}}, engine.Roles({{.FormatString .Roles ","}}){{- end}}, {{.ToUpperCase $ctrName}}Instance.{{.ToUpperCase .Name}})
	{{- end}}
}

// {{.ToUpperCase .Name}}Instance defined
var {{.ToUpperCase .Name}}Instance = New{{.ToUpperCase .Name}}()
{{end}}

func init() {
	// Sync models
	{{- range .Tables}}
	App.Manager.MSet().Add(new(model.{{.ToUpperCase .Name}}){{- if ne .Bind "" }}, "{{.Bind}}"{{- end}})
	{{- end}}
	// Async Ctr
	{{- range .Controllers}}
	{{.ToUpperCase .Name}}Routes(App)
	{{- end}}
}
`
