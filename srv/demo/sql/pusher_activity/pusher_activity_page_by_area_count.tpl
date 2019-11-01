select
    count(distinct pusher_activity.id) records
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