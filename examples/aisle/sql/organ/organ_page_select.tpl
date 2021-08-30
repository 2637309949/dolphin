select
	organ_id,
    organ_name,
    create_date,
    update_date,
    organ_pinyin,
    organ_number,
    creater,
    updater,
    parent_id,
    if_buy_online,
    organ_tell,
    sheng_id
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
order by organ.update_date desc
limit {{.size}} offset {{.offset}}
