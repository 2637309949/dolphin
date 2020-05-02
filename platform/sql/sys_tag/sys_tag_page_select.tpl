select
    sys_tag.id
from
	sys_tag
where
	sys_tag.id {{.ne}} ""
	and
	sys_tag.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
