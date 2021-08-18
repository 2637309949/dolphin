select
    sys_tag.id
from
	sys_tag
where
	sys_tag.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
