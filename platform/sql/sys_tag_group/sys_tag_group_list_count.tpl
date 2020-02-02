select
    count(*) records
from
	platform_sys_tag_group
where
	platform_sys_tag_group.id {{.ne}} ""
