// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tempalte

// TmplCtr defined template
var TmplCtr = `// Code generated by dol build. Only Generate by tools if not existed.
// source: {{.Controller.Name}}.go

package app

import (
	"{{.PackageName}}/model"
	"{{.PackageName}}/srv"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
)
{{range .Controller.APIS}}
// {{$.Controller.ToUpperCase $.Controller.Name}}{{.ToUpperCase .Name}} api implementation
// @Summary {{.Desc}} 
// @Tags {{$.Controller.Desc}}
{{- if ne .Version "" }}
// @version {{.Version}}
{{- end}}
{{- if ne .Method "get"}}
// @Accept application/json
{{- end}}
{{- $api := .}}
{{- if .Auth}}
// @Param Authorization header string false "认证令牌"
{{- end}}
{{- range .Params}}
// @Param {{.Name}} {{- if eq $api.Method "get"}} query {{- else }} body {{- end}} {{.Ref .Type}} false "{{.Desc}}"
{{- end}}
// @Failure 403 {object} model.Response
{{- if ne .Return.Success.Type ""}}
// @Success 200 {object} model.Response
{{- end}}
{{- if ne .Return.Failure.Type ""}}
// @Failure 500 {object} model.Response
{{- end}}
// @Router /api{{.APIPrefix .Version}}/{{.APIPath $.Controller.Name .Path}}/{{.Name}} [{{.Method}}]
func {{$.Controller.ToUpperCase $.Controller.Name}}{{.ToUpperCase .Name}}(ctx *Context) {
{{- if eq .Function "list"}}
	q := ctx.TypeQuery()
	{{- range .Params}}
	{{- $tv := .ToTypeValue .Type .Value}}
	q.Set{{.ToTitle .Type}}("{{.Name}}"{{- if ne "" $tv}}, {{$tv}}{{- end}})
	{{- end}}
	ret, err := ctx.PageSearch(ctx.DB, "{{$.Controller.Name}}", "{{.Name}}", "{{.Table}}", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
{{- else if eq .Function "add"}}
	{{- $bp := index .Params 0}}
	var form {{$bp.Ref $bp.Type}}
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	{{- if .ISArray $bp.Type}}
	for _, f := range form {
		f.ID = null.StringFromUUID()
	}
	{{- else}}
	form.ID = null.StringFromUUID()
	{{- end}}
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
{{- else if eq .Function "delete"}}
	{{- $bp := index .Params 0}}
	{{- $isArr := $bp.ISArray $bp.Type}}
	{{- if $isArr}}
	var form {{$bp.Ref $bp.Type}}
	var err error
	var ret []int64
	var r int64
	if err = ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for _, f := range form {
		r, err = s.ID(f.ID).Delete(&f)
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
	{{- else }}
	var form {{$bp.Ref $bp.Type}}
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.ID(form.ID).Delete(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
	{{- end}}
{{- else if eq .Function "update"}}
	{{- $bp := index .Params 0}}
	{{- $isArr := $bp.ISArray $bp.Type}}
	{{- if $isArr}}
	var form {{$bp.Ref $bp.Type}}
	var err error
	var ret []int64
	var r int64
	if err = ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for _, f := range form {
		r, err = s.ID(f.ID).Update(&f)
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
	{{- else }}
	var form {{$bp.Ref $bp.Type}}
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.ID(form.ID).Update(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
	{{- end}}
{{- else if eq .Function "one"}}
	{{- $bp := index .Params 0}}
	var entity model.{{.ToUpperCase .Table}}
	id := ctx.Query("{{$bp.Name}}")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
{{- else}}
	{{- if eq .Method "get"}}
	q := ctx.TypeQuery()
	{{- range .Params}}
	{{- $tv := .ToTypeValue .Type .Value}}
	q.Set{{.ToTitle .Type}}("{{.Name}}"{{- if ne "" $tv}}, {{$tv}}{{- end}})
	{{- end}}
	ret, err := srv.{{$.Controller.ToUpperCase $.Controller.Name}}Action(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
	{{- else}}
	{{- if ne (len .Params) 0}}
	{{- $bp := index .Params 0}}
	var form {{$bp.Ref $bp.Type}}
	{{- else}}
	var form struct{}
	{{- end}}
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.{{$.Controller.ToUpperCase $.Controller.Name}}Action(form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
	{{- end}}
{{- end}}
}
{{end}}
`
