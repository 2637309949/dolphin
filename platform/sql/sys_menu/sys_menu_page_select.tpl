select
    sys_menu.id
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
