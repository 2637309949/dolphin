select
    sys_area.id
from
	sys_area
where
	sys_area.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
