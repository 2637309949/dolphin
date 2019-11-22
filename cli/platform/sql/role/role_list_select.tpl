select
    platform_role.id
from
	platform_role
where
	platform_role.id {{.lt }}{{.gt}} ""
LIMIT {{.size}} OFFSET {{.offset}}
