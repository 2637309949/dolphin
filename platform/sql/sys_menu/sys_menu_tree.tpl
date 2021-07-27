select
    sys_menu.id,
    sys_menu.parent,
    sys_menu.code,
    sys_menu.hidden,
    sys_menu.icon,
    sys_menu.name,
    sys_menu.order,
    sys_menu.url
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
    and is_delete = 0

{{if ne .name ""}}
    and sys_menu.name = "{{.name}}"
{{end}}

order by `order`
