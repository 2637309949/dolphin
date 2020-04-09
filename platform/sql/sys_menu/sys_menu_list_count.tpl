select
    count(*) records
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