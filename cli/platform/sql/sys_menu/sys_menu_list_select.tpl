select
    platform_sys_menu.id
from
	platform_sys_menu
where
	platform_sys_menu.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
