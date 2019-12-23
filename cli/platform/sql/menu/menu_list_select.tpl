select
    platform_menu.id
from
	platform_menu
where
	platform_menu.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
