select
    count(*) records
from
	sys_tag_group
where
	sys_tag_group.is_delete {{.ne}} 1
