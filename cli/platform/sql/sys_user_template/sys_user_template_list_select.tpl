select
    platform_sys_user_template.id
from
	platform_sys_user_template
where
	platform_sys_user_template.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
