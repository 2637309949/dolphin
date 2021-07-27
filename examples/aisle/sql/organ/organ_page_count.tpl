-- Code generated by dol build. Only Generate by tools if not existed.
-- 1. You must specify the appropriate field name, instead of the *, field can be referenced in sql/sqlmap/.xml.
-- 2. You must load user information from the code level if you need, because the user information is present in another table.
select
    count(*) records
from
	organ
where
	organ.isdelete={{.is_delete}}
{{if .creater}}
	and organ.creater='{{.creater}}'
{{end}}
{{if .updater}}
	and organ.updater='{{.updater}}'
{{end}}
{{if and .create_time_start .create_time_end}}
	and organ.create_date between '{{.create_time_start}}' and '{{.create_time_end}}'
{{end}}
{{if and .update_time_start .update_time_end}}
	and organ.update_date between '{{.update_time_start}}' and '{{.update_time_end}}'
{{end}}

