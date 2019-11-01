select
    pusher_activity.*,
    ifnull(GROUP_CONCAT(pusher_campus.name), '') campus_name
from
    pusher_activity
left join
    pusher_campus
on
    find_in_set(pusher_campus.id, pusher_activity.campus)
where
    pusher_activity.del_flag = 0
{{if ne .title ""}}
    and pusher_activity.title like '%{{.title}}%'
{{end}}
{{if ne .campus ""}}
    and pusher_campus.id = '{{.campus}}'
{{else if ne .city ""}}
    and pusher_campus.city = '{{.city}}'
{{end}}
{{if ne .hidden ""}}
    and pusher_activity.hidden = '{{.hidden}}'
{{end}}
group by pusher_activity.id
{{if ne .__sort__ ""}}
	order by {{.__sort__}}
{{else}}
	order by pusher_activity.priority desc
{{end}}
LIMIT {{.rows}} OFFSET {{.offset}}