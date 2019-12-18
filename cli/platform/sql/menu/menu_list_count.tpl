select
    count(*) records
from
	platform_menu
where
	platform_menu.id {{.lt }}{{.gt}} ""
