select
    platform_sys_tag_group.id
from
	platform_sys_tag_group
where
	platform_sys_tag_group.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
