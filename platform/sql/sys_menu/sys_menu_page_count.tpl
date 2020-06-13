select
    count(*) records
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
	and
	sys_menu.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}