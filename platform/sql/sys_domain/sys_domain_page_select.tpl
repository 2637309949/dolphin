select
    sys_domain.id
from
	sys_domain
where
	sys_domain.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
