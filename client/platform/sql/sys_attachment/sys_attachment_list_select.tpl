select
    platform_sys_attachment.id
from
	platform_sys_attachment
where
	platform_sys_attachment.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
