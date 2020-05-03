select
    sys_user.id,
	sys_user.name,
	sys_user.full_name,
	sys_user.email,
	sys_user.mobile
from
	sys_user
where
	sys_user.id {{.ne}} ""
	and
	sys_user.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
