select
    organ.id
from
	organ
where
	organ.is_delete={{.is_delete}}
{{if .creater}}
	and organ.creater={{.creater}}
{{end}}
{{if .updater}}
	and organ.updater={{.updater}}
{{end}}
{{if and .create_time_start .create_time_end}}
	and organ.create_time between '{{.create_time_start}}' and '{{.create_time_end}}'
{{end}}
{{if and .update_time_start .update_time_end}}
	and organ.update_time between '{{.update_time_start}}' and '{{.update_time_end}}'
{{end}}
{{if .sort}}
	order by organ.{{.sort}}
{{end}}
	limit {{.size}} offset {{.offset}}
