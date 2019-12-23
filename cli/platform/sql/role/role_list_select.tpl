select
    platform_role.id
from
	platform_role
where
	platform_role.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
