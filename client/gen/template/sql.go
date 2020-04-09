package template

// TmplSQLSel sql
var TmplSQLSel = `select
    {{.Api.Table}}.id
from
	{{.Api.Table}}
where
	{{.Api.Table}}.id {{.Application.Unescaped "{{.ne}}"}} ""
	and
	{{.Api.Table}}.del_flag {{.Application.Unescaped "{{.ne}}"}} 0
LIMIT {{.Application.Unescaped "{{.size}}"}} OFFSET {{.Application.Unescaped "{{.offset}}"}}
`

// TmplSQLCount sql
var TmplSQLCount = `select
    count(*) records
from
	{{.Api.Table}}
where
	{{.Api.Table}}.id {{.Application.Unescaped "{{.ne}}"}} ""
	and
	{{.Api.Table}}.del_flag {{.Application.Unescaped "{{.ne}}"}} 0
`
