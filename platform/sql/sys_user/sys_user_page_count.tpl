select
    count(*) records
from
	sys_user
where
	sys_user.id {{.ne}} ""
	and
	sys_user.del_flag {{.ne}} 1
