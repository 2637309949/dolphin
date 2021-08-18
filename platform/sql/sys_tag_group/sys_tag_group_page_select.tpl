select
    sys_tag_group.id
from
	sys_tag_group
where
	sys_tag_group.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
