select
    sys_tracker.id
from
	sys_tracker
where
	sys_tracker.id {{.ne}} ""
	and
	sys_tracker.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
