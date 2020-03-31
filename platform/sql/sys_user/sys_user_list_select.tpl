select
    sys_user.id
from
	sys_user
where
	sys_user.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
