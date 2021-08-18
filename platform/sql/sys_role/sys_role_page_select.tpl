select
    sys_role.id,
	sys_role.name,
	sys_role.code,
	sys_role.status,
	sys_role.app_index,
	sys_role.admin_index
from
	sys_role
where
	sys_role.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
