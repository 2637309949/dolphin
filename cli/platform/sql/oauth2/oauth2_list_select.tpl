select
    platform_user.id
from
	platform_user
where
	platform_user.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
