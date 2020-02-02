select
    platform_sys_tag.id
from
	platform_sys_tag
where
	platform_sys_tag.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
