-- Code generated by dol build. Only Generate by tools if not existed.
-- 1. You must specify the appropriate field name, instead of the *, field can be referenced in sql/sqlmap/.xml.
-- 2. You must load user information from the code level if you need, because the user information is present in another table.
select
    article.id
from
	article
where
	article.is_delete={{.is_delete}}
{{if .creater}}
	and article.creater='{{.creater}}'
{{end}}
{{if .updater}}
	and article.updater='{{.updater}}'
{{end}}
{{if and .create_time_start .create_time_end}}
	and article.create_time between '{{.create_time_start}}' and '{{.create_time_end}}'
{{end}}
{{if and .update_time_start .update_time_end}}
	and article.update_time between '{{.update_time_start}}' and '{{.update_time_end}}'
{{end}}
{{if .role_rule}}
	and {{.role_rule}}
{{end}}
{{if .sort}}
	order by article.{{.sort}}
{{end}}
	limit {{.size}} offset {{.offset}}
