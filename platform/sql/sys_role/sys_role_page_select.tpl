select
    sys_role.id
from
	sys_role
where
	sys_role.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
