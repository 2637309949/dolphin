select
    count(*) records
from
	platform_role
where
	platform_role.id {{.lt }}{{.gt}} ""
