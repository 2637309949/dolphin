select
    sys_role.id,
	sys_role.name,
	sys_role.code,
	sys_role.status
from
	sys_role
where
	sys_role.id {{.ne}} ""
	and
	sys_role.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
