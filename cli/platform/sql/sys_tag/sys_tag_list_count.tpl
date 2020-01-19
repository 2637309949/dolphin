select
    count(*) records
from
	platform_sys_tag
where
	platform_sys_tag.id {{.ne}} ""
