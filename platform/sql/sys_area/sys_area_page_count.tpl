select
    count(*) records
from
	sys_area
where
	sys_area.id {{.ne}} ""
