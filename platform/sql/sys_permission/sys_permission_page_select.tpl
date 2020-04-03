select
    sys_permission.id
from
	sys_permission
where
	sys_permission.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
