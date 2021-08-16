select
    count(*) records
from
	sys_menu
where
	sys_menu.is_delete {{.ne}} 1
{{if ne .name ""}}
    and sys_menu.name like "%{{.name}}%"
{{end}}
{{if ne .code ""}}
    and sys_menu.code like "%{{.code}}%"
{{end}}