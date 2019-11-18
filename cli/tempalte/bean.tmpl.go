// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tempalte

// TmplBean defined template
var TmplBean = `// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

{{.Bean.Import .Bean.Packages}}

// {{.Bean.ToUpperCase .Bean.Name}} defined {{.Bean.Desc}} 
type {{.Bean.ToUpperCase .Bean.Name}} struct {
	{{- if ne .Bean.Extends ""}}
	{{- $ext:= .Bean.SplitExtends .Bean.Extends }}
	{{- range $ext }}
	*{{.}}
	{{- end}}
	{{- end}}
	{{- range .Bean.Props}}
	// {{.Desc}}
	{{.ToUpperCase .Name}} {{.Type}} ` + "`" + `json:"{{.Name}}" xml:"{{.Name}}"` + "`" + `
	{{- end}}
}
`
