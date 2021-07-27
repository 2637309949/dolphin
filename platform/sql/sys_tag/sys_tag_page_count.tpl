select
    count(*) records
from
	sys_tag
where
	sys_tag.id {{.ne}} ""
	and
	sys_tag.is_delete {{.ne}} 1
