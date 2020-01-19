select
    count(*) records
from
	platform_sys_user_template
where
	platform_sys_user_template.id {{.ne}} ""
