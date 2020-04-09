select
    sys_menu.id,
    sys_menu.parent,
    sys_menu.code,
    sys_menu.hidden,
    sys_menu.icon,
    sys_menu.name,
    sys_menu.order_num,
    sys_menu.url
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
	and
	sys_menu.del_flag {{.ne}} 0
LIMIT {{.size}} OFFSET {{.offset}}
