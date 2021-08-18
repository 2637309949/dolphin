select
    sys_permission.id
from
	sys_permission
where
	sys_permission.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
