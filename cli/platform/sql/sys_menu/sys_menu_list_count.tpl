select
    count(*) records
from
	platform_sys_menu
where
	platform_sys_menu.id {{.ne}} ""
