select
    platform_sys_org.id
from
	platform_sys_org
where
	platform_sys_org.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
