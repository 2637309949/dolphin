select
    sys_tag.id
from
	sys_tag
where
	sys_tag.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
