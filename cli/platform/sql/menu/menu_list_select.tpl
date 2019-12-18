select
    platform_menu.id
from
	platform_menu
where
	platform_menu.id {{.lt }}{{.gt}} ""
LIMIT {{.size}} OFFSET {{.offset}}
