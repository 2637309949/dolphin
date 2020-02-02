package tempalte

// TmplSQLSel sql
var TmplSQLSel = `select
    {{.Api.TableName .Application.Name .Api.Table}}.id
from
	{{.Api.TableName .Application.Name .Api.Table}}
where
	{{.Api.TableName .Application.Name .Api.Table}}.id {{.Application.Unescaped "{{.ne}}"}} ""
LIMIT {{.Application.Unescaped "{{.size}}"}} OFFSET {{.Application.Unescaped "{{.offset}}"}}
`

// TmplSQLCount sql
var TmplSQLCount = `select
    count(*) records
from
	{{.Api.TableName .Application.Name .Api.Table}}
where
	{{.Api.TableName .Application.Name .Api.Table}}.id {{.Application.Unescaped "{{.ne}}"}} ""
`
