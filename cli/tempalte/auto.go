// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tempalte

// TmplAuto defined template
var TmplAuto = `// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package app

import (
	"{{.PackageName}}/model"

	{{if ne .Name "platform"}}pApp "github.com/2637309949/dolphin/cli/platform/app"{{- end}}
	"github.com/2637309949/dolphin/cli/packages/viper"
	"github.com/2637309949/dolphin/srv/cli"
)

// Name project
var Name = "{{.Name}}"

// Sync models
var _ = cli.Invoke(BuildEngine(func(e *Engine) {
	{{- range .Tables}}
	e.MSet.Add(Name, new(model.{{.ToUpperCase .Name}}))
	{{- end}}
}))
{{- range .Controllers}}

// Build{{.ToUpperCase .Name}}
var _ = cli.Invoke(Build{{.ToUpperCase .Name}}(func(ctr *{{.ToUpperCase .Name}}) {
	group := ctr.Group(viper.GetString("http.prefix"))
	{{- $ctr := .}}
	{{- range .APIS}}
	group.Handle("{{.ToUpper .Method}}", "{{.VersionPrefix .Version}}/{{$ctr.Name}}/{{.Name}}",{{- if .Auth}} ctr.Auth({{if ne $.Name "platform"}}pApp.{{- end}}OAuth2),{{- end}} ctr.{{.ToUpperCase .Name}})
	{{- end}}
}))
{{- end}}

// Run booting system
var _ = cli.Invoke(BuildEngine(func(e *Engine) {
	e.Run()
}))
`