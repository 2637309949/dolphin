select
    platform_sys_oauth2.id
from
	platform_sys_oauth2
where
	platform_sys_oauth2.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
