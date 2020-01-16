select
    count(*) records
from
	platform_sys_user
where
	platform_sys_user.id {{.ne}} ""
