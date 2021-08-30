select
    count(*) records
from
	article
where
	article.is_delete={{.is_delete}}
{{if .creater}}
	and article.creater={{.creater}}
{{end}}
{{if .updater}}
	and article.updater={{.updater}}
{{end}}
{{if and .create_time_start .create_time_end}}
	and article.create_time between '{{.create_time_start}}' and '{{.create_time_end}}'
{{end}}
{{if and .update_time_start .update_time_end}}
	and article.update_time between '{{.update_time_start}}' and '{{.update_time_end}}'
{{end}}
