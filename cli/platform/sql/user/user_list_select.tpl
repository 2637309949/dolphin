select
    platform_user.id
from
	platform_user
where
	platform_user.id {{.lt }}{{.gt}} ""
LIMIT {{.size}} OFFSET {{.offset}}
