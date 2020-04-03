select
    sys_menu.*
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
{{if ne .name ""}}
    and sys_menu.name = "{{.name}}"
{{end}}

order by order_num
