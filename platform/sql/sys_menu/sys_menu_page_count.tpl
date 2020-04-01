select
    count(*) records
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
