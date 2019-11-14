package tempalte

// TmplSQLSel sql
var TmplSQLSel = `select
    {{.Application.Name}}_{{.Api.Table}}.id
from
	{{.Application.Name}}_{{.Api.Table}}
where
	{{.Application.Name}}_{{.Api.Table}}.id {{.Application.Unescaped "<>"}} ""
LIMIT {{.size}} OFFSET {{.offset}}
`

// TmplSQLCount sql
var TmplSQLCount = `select
    count(*) records
from
	{{.Application.Name}}_{{.Api.Table}}
where
	{{.Application.Name}}_{{.Api.Table}}.id {{.Application.Unescaped "<>"}} ""
`
