select
    platform_sys_user.id
from
	platform_sys_user
where
	platform_sys_user.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
