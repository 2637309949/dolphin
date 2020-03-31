select
    count(*) records
from
	sys_tag
where
	sys_tag.id {{.ne}} ""
