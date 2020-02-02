select
    count(*) records
from
	platform_sys_attachment
where
	platform_sys_attachment.id {{.ne}} ""
