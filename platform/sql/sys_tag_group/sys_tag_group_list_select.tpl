select
    sys_tag_group.id
from
	sys_tag_group
where
	sys_tag_group.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
