select
    count(*) records
from
	platform_sys_oauth2
where
	platform_sys_oauth2.id {{.ne}} ""
