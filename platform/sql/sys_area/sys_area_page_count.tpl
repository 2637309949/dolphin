select
    count(*) records
from
	sys_area
where
	sys_area.is_delete {{.ne}} 1
