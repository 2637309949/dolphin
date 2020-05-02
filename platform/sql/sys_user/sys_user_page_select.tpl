select
    sys_user.id
from
	sys_user
where
	sys_user.id {{.ne}} ""
	and
	sys_user.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
