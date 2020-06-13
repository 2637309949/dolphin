select
    distinct sys_menu.id, sys_menu.parent, sys_menu.code, sys_menu.component, sys_menu.hidden, sys_menu.icon, sys_menu.name, sys_menu.order, sys_menu.url
from
	sys_menu
{{if ne .isAdmin true}}
left join sys_role_menu
on
    sys_menu.id = sys_role_menu.menu_id
left join sys_user_role
on 
    sys_role_menu.role_id = sys_user_role.role_id
left join sys_user
on
    sys_user_role.user_id = sys_user.id
{{end}}
where
	sys_menu.id {{.ne}} ""
	and
    sys_menu.del_flag = 0
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
	and
	sys_menu.hidden = 0
	order by `order`