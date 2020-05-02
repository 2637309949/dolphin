select
    sys_area.id
from
	sys_area
where
	sys_area.id {{.ne}} ""
	and
	sys_area.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
