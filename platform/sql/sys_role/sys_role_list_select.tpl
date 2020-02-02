select
    platform_sys_role.id
from
	platform_sys_role
where
	platform_sys_role.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
