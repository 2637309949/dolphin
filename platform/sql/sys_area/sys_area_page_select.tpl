select
    sys_area.id
from
	sys_area
where
	sys_area.id {{.ne}} ""
	and
	sys_area.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
