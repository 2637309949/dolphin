select
    distinct
    sys_menu.id,
    sys_menu.parent,
    sys_menu.code,
    sys_menu.hidden,
    sys_menu.icon,
    sys_menu.name,
    sys_menu.order,
    sys_menu.url
from
	sys_menu{{if ne .is_admin true}}, sys_role_menu{{end}}
where
	sys_menu.id {{.ne}} ""
{{if ne .is_admin true}}
    and sys_menu.id = sys_role_menu.menu_id
    and sys_role_menu.role_id = "{{.role_id}}"
{{end}}
    and sys_menu.is_delete = 0

{{if ne .name ""}}
    and sys_menu.name = "{{.name}}"
{{end}}

order by `order`
