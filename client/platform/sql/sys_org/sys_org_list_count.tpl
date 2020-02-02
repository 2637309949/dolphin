select
    count(*) records
from
	platform_sys_org
where
	platform_sys_org.id {{.ne}} ""
