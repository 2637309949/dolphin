select
    count(*) records
from
	platform_sys_role
where
	platform_sys_role.id {{.ne}} ""
