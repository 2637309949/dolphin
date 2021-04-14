select
    count(*) records
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
	and
	sys_menu.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}