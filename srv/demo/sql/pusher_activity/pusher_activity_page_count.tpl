select
    count(*) records
from
    pusher_activity
where
    pusher_activity.del_flag = {{.del_flag}}
{{if ne .title ""}}
    and pusher_activity.title like '%{{.title}}%'
{{end}}
{{if ne .hidden ""}}
    and pusher_activity.hidden = '{{.hidden}}'
{{end}}
