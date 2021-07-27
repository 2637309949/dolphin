select
    distinct sys_menu.id, sys_menu.parent, sys_menu.code, sys_menu.component, sys_menu.hidden, sys_menu.icon, sys_menu.name, sys_menu.order, sys_menu.url
from
	sys_menu
{{if ne .isAdmin true}}
left join sys_role_menu on sys_menu.id = sys_role_menu.menu_id
inner join sys_role_user on sys_role_menu.role_id = sys_role_user.role_id and sys_role_user.user_id = "{{.uid}}"
{{end}}
where
	sys_menu.id {{.ne}} ""
	and
    sys_menu.is_delete = 0

	and
	sys_menu.hidden = 0
order by `order`
