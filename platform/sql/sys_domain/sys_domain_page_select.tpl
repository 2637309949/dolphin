select
    sys_domain.id
from
	sys_domain
where
	sys_domain.id {{.ne}} ""
	and
	sys_domain.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
